package rpcserver

import (
	"github.com/internet-cash/prototype/blockchain"
	"sync/atomic"
	"net/http"
	"errors"
	"time"
	"log"
	"net"
	"io/ioutil"
	"fmt"
	"strconv"
	"github.com/internet-cash/prototype/jsonrpc"
	"encoding/json"
	"strings"
	"reflect"
	"sync"
	"io"
)

const (
	rpcAuthTimeoutSeconds = 10
)

// timeZeroVal is simply the zero value for a time.Time and is used to avoid
// creating multiple instances.
var timeZeroVal time.Time

// parsedRPCCmd represents a JSON-RPC request object that has been parsed into
// a known concrete command along with any error that might have happened while
// parsing it.
type parsedRPCCmd struct {
	id     interface{}
	method string
	cmd    interface{}
	err    *jsonrpc.RPCError
}

// rpcServer provides a concurrent safe RPC server to a chain server.
type RpcServer struct {
	started    int32
	shutdown   int32
	numClients int32

	Config      RpcServerConfig
	HttpServer  *http.Server
	statusLock  sync.RWMutex
	statusLines map[int]string

	requestProcessShutdown chan struct{}
	quit                   chan int
}

type RpcServerConfig struct {
	Listenters    []net.Listener
	ChainParams   *blockchain.Params
	RPCMaxClients int
	RPCQuirks     bool
}

func (self RpcServer) Init(config *RpcServerConfig) (*RpcServer, error) {
	self.Config = *config
	return &self, nil
}

// RequestedProcessShutdown returns a channel that is sent to when an authorized
// RPC client requests the process to shutdown.  If the request can not be read
// immediately, it is dropped.
func (self RpcServer) RequestedProcessShutdown() <-chan struct{} {
	return self.requestProcessShutdown
}

// limitConnections responds with a 503 service unavailable and returns true if
// adding another client would exceed the maximum allow RPC clients.
//
// This function is safe for concurrent access.
func (self RpcServer) limitConnections(w http.ResponseWriter, remoteAddr string) bool {
	if int(atomic.LoadInt32(&self.numClients)+1) > self.Config.RPCMaxClients {
		log.Printf("Max RPC clients exceeded [%d] - "+
			"disconnecting client %s", self.Config.RPCMaxClients,
			remoteAddr)
		http.Error(w, "503 Too busy.  Try again later.",
			http.StatusServiceUnavailable)
		return true
	}
	return false
}

// genCertPair generates a key/cert pair to the paths provided.
func genCertPair(certFile, keyFile string) error {
	// TODO for using TCL
	/*log.Println("Generating TLS certificates...")

	org := "btcd autogenerated cert"
	validUntil := time.Now().Add(10 * 365 * 24 * time.Hour)
	cert, key, err := btcutil.NewTLSCertPair(org, validUntil, nil)
	if err != nil {
		return err
	}

	// Write cert and key files.
	if err = ioutil.WriteFile(certFile, cert, 0666); err != nil {
		return err
	}
	if err = ioutil.WriteFile(keyFile, key, 0600); err != nil {
		os.Remove(certFile)
		return err
	}

	rpcsLog.Infof("Done generating TLS certificates")*/
	return nil
}

func (self RpcServer) Start() (error) {
	if atomic.AddInt32(&self.started, 1) != 1 {
		return errors.New("RPC server is already started")
	}
	rpcServeMux := http.NewServeMux()
	self.HttpServer = &http.Server{
		Handler: rpcServeMux,

		// Timeout connections which don't complete the initial
		// handshake within the allowed timeframe.
		ReadTimeout: time.Second * rpcAuthTimeoutSeconds,
	}

	rpcServeMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		self.RpcHandleRequest(w, r)
	})
	for _, listen := range self.Config.Listenters {
		go func(listen net.Listener) {
			log.Printf("RPC server listening on %s", listen.Addr())
			go self.HttpServer.Serve(listen)
			log.Printf("RPC listener done for %s", listen.Addr())
		}(listen)
	}
	self.started = 1
	return nil
}

// Stop is used by server.go to stop the rpc listener.
func (self RpcServer) Stop() error {
	if atomic.AddInt32(&self.shutdown, 1) != 1 {
		log.Println("RPC server is already in the process of shutting down")
		return nil
	}
	log.Println("RPC server shutting down")
	self.HttpServer.Close()
	for _, listen := range self.Config.Listenters {
		listen.Close()
	}
	close(self.quit)
	log.Println("RPC server shutdown complete")
	self.started = 0
	self.shutdown = 1
	return nil
}

func (self RpcServer) RpcHandleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Connection", "close")
	w.Header().Set("Content-Type", "application/json")
	r.Close = true

	// Limit the number of connections to max allowed.
	if self.limitConnections(w, r.RemoteAddr) {
		return
	}

	// Keep track of the number of connected clients.
	self.incrementClients()
	defer self.decrementClients()
	// TODO
	_, isAdmin, err := self.checkAuth(r, true)
	if err != nil {
		self.AuthFail(w)
		return
	}

	self.ProcessRpcRequest(w, r, isAdmin)
}

// checkAuth checks the HTTP Basic authentication supplied by a wallet
// or RPC client in the HTTP request r.  If the supplied authentication
// does not match the username and password expected, a non-nil error is
// returned.
//
// This check is time-constant.
//
// The first bool return value signifies auth success (true if successful) and
// the second bool return value specifies whether the user can change the state
// of the server (true) or whether the user is limited (false). The second is
// always false if the first is.
func (self RpcServer) checkAuth(r *http.Request, require bool) (bool, bool, error) {
	// TODO
	return true, true, nil
}

// incrementClients adds one to the number of connected RPC clients.  Note
// this only applies to standard clients.  Websocket clients have their own
// limits and are tracked separately.
//
// This function is safe for concurrent access.
func (self *RpcServer) incrementClients() {
	atomic.AddInt32(&self.numClients, 1)
}

// decrementClients subtracts one from the number of connected RPC clients.
// Note this only applies to standard clients.  Websocket clients have their own
// limits and are tracked separately.
//
// This function is safe for concurrent access.
func (self *RpcServer) decrementClients() {
	atomic.AddInt32(&self.numClients, -1)
}

// AuthFail sends a message back to the client if the http auth is rejected.
func (self RpcServer) AuthFail(w http.ResponseWriter) {
	w.Header().Add("WWW-Authenticate", `Basic realm="RPC"`)
	http.Error(w, "401 Unauthorized.", http.StatusUnauthorized)
}

/**
handles reading and responding to RPC messages.
 */
func (self RpcServer) ProcessRpcRequest(w http.ResponseWriter, r *http.Request, isAdmin bool) {
	if atomic.LoadInt32(&self.shutdown) != 0 {
		return
	}
	// Read and close the JSON-RPC request body from the caller.
	body, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		errCode := http.StatusBadRequest
		http.Error(w, fmt.Sprintf("%d error reading JSON message: %v",
			errCode, err), errCode)
		return
	}

	// Unfortunately, the http server doesn't provide the ability to
	// change the read deadline for the new connection and having one breaks
	// long polling.  However, not having a read deadline on the initial
	// connection would mean clients can connect and idle forever.  Thus,
	// hijack the connecton from the HTTP server, clear the read deadline,
	// and handle writing the response manually.
	hj, ok := w.(http.Hijacker)
	if !ok {
		errMsg := "webserver doesn't support hijacking"
		log.Print(errMsg)
		errCode := http.StatusInternalServerError
		http.Error(w, strconv.Itoa(errCode)+" "+errMsg, errCode)
		return
	}
	conn, buf, err := hj.Hijack()
	if err != nil {
		log.Printf("Failed to hijack HTTP connection: %v", err)
		errCode := http.StatusInternalServerError
		http.Error(w, strconv.Itoa(errCode)+" "+err.Error(), errCode)
		return
	}
	defer conn.Close()
	defer buf.Flush()
	conn.SetReadDeadline(timeZeroVal)

	// Attempt to parse the raw body into a JSON-RPC request.
	var responseID interface{}
	var jsonErr error
	var result interface{}
	var request jsonrpc.Request
	if err := json.Unmarshal(body, &request); err != nil {
		jsonErr = &jsonrpc.RPCError{
			Code:    jsonrpc.ErrRPCParse.Code,
			Message: "Failed to parse request: " + err.Error(),
		}
	}

	if jsonErr == nil {
		// The JSON-RPC 1.0 spec defines that notifications must have their "id"
		// set to null and states that notifications do not have a response.
		//
		// A JSON-RPC 2.0 notification is a request with "json-rpc":"2.0", and
		// without an "id" member. The specification states that notifications
		// must not be responded to. JSON-RPC 2.0 permits the null value as a
		// valid request id, therefore such requests are not notifications.
		//
		// Bitcoin Core serves requests with "id":null or even an absent "id",
		// and responds to such requests with "id":null in the response.
		//
		// Btcd does not respond to any request without and "id" or "id":null,
		// regardless the indicated JSON-RPC protocol version unless RPC quirks
		// are enabled. With RPC quirks enabled, such requests will be responded
		// to if the reqeust does not indicate JSON-RPC version.
		//
		// RPC quirks can be enabled by the user to avoid compatibility issues
		// with software relying on Core's behavior.
		if request.ID == nil && !(self.Config.RPCQuirks && request.Jsonrpc == "") {
			return
		}

		// The parse was at least successful enough to have an ID so
		// set it for the response.
		responseID = request.ID

		// Setup a close notifier.  Since the connection is hijacked,
		// the CloseNotifer on the ResponseWriter is not available.
		closeChan := make(chan struct{}, 1)
		go func() {
			_, err := conn.Read(make([]byte, 1))
			if err != nil {
				close(closeChan)
			}
		}()

		// Check if the user is limited and set error if method unauthorized
		if !isAdmin {
			if _, ok := RpcLimited[request.Method]; !ok {
				jsonErr = &jsonrpc.RPCError{
					Code:    jsonrpc.ErrRPCInvalidParams.Code,
					Message: "limited user not authorized for this method",
				}
			}
		}
		if jsonErr == nil {
			// Attempt to parse the JSON-RPC request into a known concrete
			// command.
			parsedCmd := parseCmd(&request)
			if parsedCmd.err != nil {
				jsonErr = parsedCmd.err
			} else {
				result, jsonErr = self.standardCmdResult(parsedCmd, closeChan)
			}
		}
	}
	// Marshal the response.
	msg, err := createMarshalledReply(responseID, result, jsonErr)
	if err != nil {
		log.Printf("Failed to marshal reply: %v", err)
		return
	}

	// Write the response.
	err = self.writeHTTPResponseHeaders(r, w.Header(), http.StatusOK, buf)
	if err != nil {
		log.Println(err)
		return
	}
	if _, err := buf.Write(msg); err != nil {
		log.Printf("Failed to write marshalled reply: %v", err)
	}

	// Terminate with newline to maintain compatibility with Bitcoin Core.
	if err := buf.WriteByte('\n'); err != nil {
		log.Printf("Failed to append terminating newline to reply: %v", err)
	}
}

// createMarshalledReply returns a new marshalled JSON-RPC response given the
// passed parameters.  It will automatically convert errors that are not of
// the type *btcjson.RPCError to the appropriate type as needed.
func createMarshalledReply(id, result interface{}, replyErr error) ([]byte, error) {
	var jsonErr *jsonrpc.RPCError
	if replyErr != nil {
		if jErr, ok := replyErr.(*jsonrpc.RPCError); ok {
			jsonErr = jErr
		} else {
			jsonErr = internalRPCError(replyErr.Error(), "")
		}
	}

	return jsonrpc.MarshalResponse(id, result, jsonErr)
}

// internalRPCError is a convenience function to convert an internal error to
// an RPC error with the appropriate code set.  It also logs the error to the
// RPC server subsystem since internal errors really should not occur.  The
// context parameter is only used in the log message and may be empty if it's
// not needed.
func internalRPCError(errStr, context string) *jsonrpc.RPCError {
	logStr := errStr
	if context != "" {
		logStr = context + ": " + errStr
	}
	log.Println(logStr)
	return jsonrpc.NewRPCError(jsonrpc.ErrRPCInternal.Code, errStr)
}

// httpStatusLine returns a response Status-Line (RFC 2616 Section 6.1)
// for the given request and response status code.  This function was lifted and
// adapted from the standard library HTTP server code since it's not exported.
func (self RpcServer) httpStatusLine(req *http.Request, code int) string {
	// Fast path:
	key := code
	proto11 := req.ProtoAtLeast(1, 1)
	if !proto11 {
		key = -key
	}
	self.statusLock.RLock()
	line, ok := self.statusLines[key]
	self.statusLock.RUnlock()
	if ok {
		return line
	}

	// Slow path:
	proto := "HTTP/1.0"
	if proto11 {
		proto = "HTTP/1.1"
	}
	codeStr := strconv.Itoa(code)
	text := http.StatusText(code)
	if text != "" {
		line = proto + " " + codeStr + " " + text + "\r\n"
		self.statusLock.Lock()
		self.statusLines[key] = line
		self.statusLock.Unlock()
	} else {
		text = "status code " + codeStr
		line = proto + " " + codeStr + " " + text + "\r\n"
	}

	return line
}

// writeHTTPResponseHeaders writes the necessary response headers prior to
// writing an HTTP body given a request to use for protocol negotiation, headers
// to write, a status code, and a writer.
func (self RpcServer) writeHTTPResponseHeaders(req *http.Request, headers http.Header, code int, w io.Writer) error {
	_, err := io.WriteString(w, self.httpStatusLine(req, code))
	if err != nil {
		return err
	}

	err = headers.Write(w)
	if err != nil {
		return err
	}

	_, err = io.WriteString(w, "\r\n")
	return err
}

// standardCmdResult checks that a parsed command is a standard Bitcoin JSON-RPC
// command and runs the appropriate handler to reply to the command.  Any
// commands which are not recognized or not implemented will return an error
// suitable for use in replies.
func (self RpcServer) standardCmdResult(cmd *parsedRPCCmd, closeChan <-chan struct{}) (interface{}, error) {
	handler, ok := RpcHandler[cmd.method]
	if ok {
		goto handled
	}
	return nil, jsonrpc.ErrRPCMethodNotFound
handled:

	return handler(&self, cmd.cmd, closeChan)
}

// parseCmd parses a JSON-RPC request object into known concrete command.  The
// err field of the returned parsedRPCCmd struct will contain an RPC error that
// is suitable for use in replies if the command is invalid in some way such as
// an unregistered command or invalid parameters.
func parseCmd(request *jsonrpc.Request) *parsedRPCCmd {
	var parsedCmd parsedRPCCmd
	parsedCmd.id = request.ID
	parsedCmd.method = request.Method

	cmd, err := UnmarshalCmd(request)
	if err != nil {
		// When the error is because the method is not registered,
		// produce a method not found RPC error.
		if jerr, ok := err.(jsonrpc.Error); ok &&
			jerr.ErrorCode == jsonrpc.ErrUnregisteredMethod {

			parsedCmd.err = jsonrpc.ErrRPCMethodNotFound
			return &parsedCmd
		}

		// Otherwise, some type of invalid parameters is the
		// cause, so produce the equivalent RPC error.
		parsedCmd.err = jsonrpc.NewRPCError(
			jsonrpc.ErrRPCInvalidParams.Code, err.Error())
		return &parsedCmd
	}

	parsedCmd.cmd = cmd
	return &parsedCmd
}

var (
	// These fields are used to map the registered types to method names.
	registerLock         sync.RWMutex
	methodToConcreteType = make(map[string]reflect.Type)
	methodToInfo         = make(map[string]methodInfo)
	concreteTypeToMethod = make(map[reflect.Type]string)
)

// UsageFlag define flags that specify additional properties about the
// circumstances under which a command can be used.
type UsageFlag uint32

// methodInfo keeps track of information about each registered method such as
// the parameter information.
type methodInfo struct {
	maxParams    int
	numReqParams int
	numOptParams int
	defaults     map[int]reflect.Value
	flags        UsageFlag
	usage        string
}

// makeError creates an Error given a set of arguments.
func makeError(c jsonrpc.ErrorCode, desc string) jsonrpc.Error {
	return jsonrpc.Error{ErrorCode: c, Description: desc}
}

// checkNumParams ensures the supplied number of params is at least the minimum
// required number for the command and less than the maximum allowed.
func checkNumParams(numParams int, info *methodInfo) error {
	if numParams < info.numReqParams || numParams > info.maxParams {
		if info.numReqParams == info.maxParams {
			str := fmt.Sprintf("wrong number of params (expected "+
				"%d, received %d)", info.numReqParams,
				numParams)
			return makeError(jsonrpc.ErrNumParams, str)
		}

		str := fmt.Sprintf("wrong number of params (expected "+
			"between %d and %d, received %d)", info.numReqParams,
			info.maxParams, numParams)
		return makeError(jsonrpc.ErrNumParams, str)
	}

	return nil
}

// UnmarshalCmd unmarshals a JSON-RPC request into a suitable concrete command
// so long as the method type contained within the marshalled request is
// registered.
func UnmarshalCmd(r *jsonrpc.Request) (interface{}, error) {
	registerLock.RLock()
	rtp, ok := methodToConcreteType[r.Method]
	info := methodToInfo[r.Method]
	registerLock.RUnlock()
	if !ok {
		str := fmt.Sprintf("%q is not registered", r.Method)
		return nil, makeError(jsonrpc.ErrUnregisteredMethod, str)
	}
	rt := rtp.Elem()
	rvp := reflect.New(rt)
	rv := rvp.Elem()

	// Ensure the number of parameters are correct.
	numParams := len(r.Params)
	if err := checkNumParams(numParams, &info); err != nil {
		return nil, err
	}

	// Loop through each of the struct fields and unmarshal the associated
	// parameter into them.
	for i := 0; i < numParams; i++ {
		rvf := rv.Field(i)
		// Unmarshal the parameter into the struct field.
		concreteVal := rvf.Addr().Interface()
		if err := json.Unmarshal(r.Params[i], &concreteVal); err != nil {
			// The most common error is the wrong type, so
			// explicitly detect that error and make it nicer.
			fieldName := strings.ToLower(rt.Field(i).Name)
			if jerr, ok := err.(*json.UnmarshalTypeError); ok {
				str := fmt.Sprintf("parameter #%d '%s' must "+
					"be type %v (got %v)", i+1, fieldName,
					jerr.Type, jerr.Value)
				return nil, makeError(jsonrpc.ErrInvalidType, str)
			}

			// Fallback to showing the underlying error.
			str := fmt.Sprintf("parameter #%d '%s' failed to "+
				"unmarshal: %v", i+1, fieldName, err)
			return nil, makeError(jsonrpc.ErrInvalidType, str)
		}
	}

	// When there are less supplied parameters than the total number of
	// params, any remaining struct fields must be optional.  Thus, populate
	// them with their associated default value as needed.
	if numParams < info.maxParams {
		populateDefaults(numParams, &info, rv)
	}

	return rvp.Interface(), nil
}

// populateDefaults populates default values into any remaining optional struct
// fields that did not have parameters explicitly provided.  The caller should
// have previously checked that the number of parameters being passed is at
// least the required number of parameters to avoid unnecessary work in this
// function, but since required fields never have default values, it will work
// properly even without the check.
func populateDefaults(numParams int, info *methodInfo, rv reflect.Value) {
	// When there are no more parameters left in the supplied parameters,
	// any remaining struct fields must be optional.  Thus, populate them
	// with their associated default value as needed.
	for i := numParams; i < info.maxParams; i++ {
		rvf := rv.Field(i)
		if defaultVal, ok := info.defaults[i]; ok {
			rvf.Set(defaultVal)
		}
	}
}
