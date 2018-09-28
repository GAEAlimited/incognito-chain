if [ -z "$1" ]
then
    echo "Please enter How many node(s) to start"
    exit 0
fi

TOTAL=$1
SRC=$(pwd)

KEY1="2pmcvmL2rYC6D4ydq5Ppv4xQvMR1KC3xennM1aL8jkyv1y9Az27tFUutBZPKgsLTyaDSp1MhPPVfxqcLA35uLTPWAqJW1389jN2wXrukupdKgBZcYmvmLsgJzXyW2sB7rLs9LW5XDMcXGbcvN8g7fnTRQEM5NhaXkE9M9TTXhb8ZDpQ2Du1uY2wVwDeHtZM6BeAbcUCaYWgnm2cUjMSYNGactfSnEeKprcRrvAqd47SFrgX4X4NSBco2FWoKNRhn4ocS75BrwsWkHK13SAEhjuTEp1cNsavm135yFwXCd6ZS5nHsQzpLyLkUUNXFCzXTkTurSTEGLn9ppAPLubhtGJYN6wwFePSi4Udcd3nRx8tJcohytxBYhsf4bbzkb63FB1h3cZuH6c6Gt4W6dDJFe4RCFpSBpDYSHGTjzyLmRmvFgaikMnwdK8BiPST3csuCBVfxTopZCAX9D2UNPHngdSmpY9KeG1r3RsKwjUHWN4AZY2399WmQaevs2v2CfUhfkMob96Jq3CJ4dgkso3qeMyZCnRHTRYqkgNTqABiMqt2wfzJBWm5EnYUy886eycJ1Cabt6VXRDKLaHfCeN6jxuHFKwjY6THgBH7i2C8meTkpVntUiGJ4GKtMX4jhT989UW4iLbAC8eYjasapp994zjnrS6iQaBFoBB2r98f4NFEDtSEoaeUP5WeUmkaLeKQVjQxbjKg21QarSs3o3cKvqE9GV1zRFa3egEm55rUDuKG6vvoSA4nAnA2HWcZnbyJMdeFALnjW"
KEY2="2ELu3Qi2UhYTK7sNHJq5rjRDbB8sZp6e7bHuhj17NVwzP9VBsRicTLEnHNvfTR5C2Xn8iGKE5oas2mQNJqEx992Nv4VwRkTfzPbo8tMF3tQV5WPgjsxnnJzXUrtnXzdf8kBHWuPqucs2gs7mGA6CNgqsJosSiJQYJLTK8wgf4TKb5yiahikpRi7jLNCCQns8UunvJLMCNdMzgNeLDeDZo9adeFzDvzmz4D8uXKWn94hPiJ4XqKNgwNmo9fMh4EXjZRHz3XWq4kQj8JdsCG1zKAuDfQw9sPuB1jacHt3zTxqa4aN6V4xNf482FrU5zvxNBxLJr6mpP8Px4rax3AZB9znnQuiUzEMA5QsKJKVabhPNs5e1euQz4xGv7SG3fZAeogVUUWMKF3EQeKZggWPMLyrCAqrGZcDjMarMJSFtaFvaRxjZu8ugHpapvnZm2P4C5SqzyKUMoFsaxBr2KaW74Qxv5CbkNDvWgsaxJLURUHdxvCQkxhNkrPLwLfY5kSZKpTM6ZBfWDqJGGB4oWpruTTh1UShm7TFvzPe2J5ZGHMkpvq2r4FEjXWJH6HPAxNQt99X7Hzs2RMYxy1CK14aszA23uZ99GXA6YKfdF7wPzYW4fd87BTWrBjpfRUQgEb3dHu3AgbtJ6QGH15QCw5vMegjCfAZ8YA3AzgaY9GT4RzvVEyF6GmMXPvRCXrLWgo5FSrg6CKKEzqdQXgFwQKfu4A6RToFS8NwXLw7gy7v9HDBbPwSDek9fdaoPJ1xkSf7W5GY"
KEY3="cXZ1szvjp1Guhme8d7HQwcnVyTtVutR7K6AHWyFpT9EY318kWtm8yya9ZEKW2UT35CR3zuDmrn2x1rmhKeZpwWsPJccK5SwbSLAUtGMbZhkg75kSiiPYeU4xucRzuuwJNnGJx5r2w9crP55NEcBWtjWVTdUFNy3Vk2R879dRXqMsSbtvM6TVeDcNjaXvxHpLqncfdYuMPHMtVdNF4rSoSRSpzXj3G13hY8f5RcmEGnpHsi9RXUQiW8jzwcViYbaeNpjbKSL23gdcC1wsGVBh7FC9vJCaC1KmQmEpN5MM4pCWdH7yzzFyvYEQmYuWhoTN4FJdD1Xn7WCM192oYzUEvCdjd3npfeWCV8y4Po69WngHdkVXjk9kA1fvA21ynxX8JKZvdrkpFMYc94y8KhXBEAwt8p37QyzdYcV8fQhLwuv1ipxV8fL5Ui6uCHsyryyNpVM3cJZ9Zj8xtETtkgEwY5WNZHHqqcyPFX5DKmjLeU2iJ9oCVcWw6VPrAovkGhWTrgMnDeYLjDHJ6zPQF5qiNGNc7n24ZnouCcMyfXoK67kriPUC6PEaZHVrrXzimXofJD37YFfLSnGsnD8sHtPpZix5yJQVgBCcPiH2y5HakePz2dbUT5PXXtyhkabi7JC7YP9nrpJcnRbKcwcwVzcS5ikvegE5Q8bEjvT3WcXcALq1AuFHxgkSqd39uXtxTB94CWjrY9N6BcUsu9JbfAZLg2nCNCVPS8oDYrK6jWdYMtUYS7NPswpgWaueugEtJid2jRzgGLepU"
KEY4="9oxyvDBML3ewXfG5CZKShZj7VeB24sT6W1mXtKyFTHAyGRMpgXBzS9tJr9ajzY99YiUojDTutTHVjX4Mu8o48Ryhym4A5XajxHMcMQk7SXJFXFx2yrTuKwgvX5qLmfxZJqH3dqYvHBvjht85a41RbfwwEgjkZAyShyLBKsjXUJon8Wtg3K3T5Y2Dg26tBNGKvFXHHHapMkLXn7vwhHbfmG6k91BmmHfEmNjcUA92uaAwm5XwMNJoyddr8yerePwtUMBgB2JvU1zMKLRNrNwzSC2v5SPmUz2ed6xG6FWcoSoqeYEp2kHqeqK6auMCZLh1sXvMUUy9MDogQzuQVNxTsPowwhu6EwEJR4dxKm5HDqTdfN4HambJtYr8mYmvCsGn4wSYFkWa6hiBSgpfQarQeWqg6NsiYAmMSdd8tLXZR9bbYzWqYS1WNTR8in4NAhStM82yqwzV3nnveGQ5eBR2fYYnGWt2omCLXXoJpUPYCHFYVLXQT8tSTyifpkPBmtCrYDFYZtxXKg1AqWEtaagdqGUUAKtEei2Gi3k3hTf2iAfCiuHnUTn78SZuuvepmT9wSMTxAdimqWoemaFLE2499NFm1G8ELzBUPccwDYLRTce2htb1wrcQxrDrSdGSSiqWKowNS5JcYMMcuUQYHpbsvMTU3Z1FikAwXVGfqiaEnqFCykusiN1gv5xiUYrwMy17ze9bcbkLoyD1nEBe55QdErTZFhuXWyZEhPMbkUExBgPGjTBZByZEr4khH9HUANBcEhGuoEW8wyW8PqUECSDxNi8"
KEY5="Jh6VbVg3NYSY8ybqgS1XA6h855WiSz4QoEq6okjGwnxQHhnPWiwH96NtC2bUFocd1bMBdNp6KvEWPSUc1uHNFAZyu4s7RL46j4b4bz7oXfk1yoPTBahxk7L7d2FmtiMuvLWF6ze3NG1BVp6LVqY9CpTEhYNx5N9VBxQupp32fqphsYzmM83AzMK1hvMLT6a23LZhqY4UtuSjAnD659zWhsRbffnkADfocDMEwauAdNNuX2Fq4w8izowQRTjtxq4sB915kjsGLs6smKegeSb4Z3S5k5dQcRydmh9C9gigSMABnt3n6RSNFwuAPMjs9NnUY7dKoTCFCaExLqJwddfQXCMp5WFZEZ3c3hm3dD1ex6oYZi8AwbUrKqD6FRD6G1xobRKm3LsKRWtyYrmgc6gtvK82iu5HqTMSukPRhXZrQBUup3ihmNX8Akr6tEn9Eh5AWxrcP1RUVptCVdMVpPE4wBA8Zt3mnNsAHhERCw4hwuKo7daLX8EeqdCmyUP4C7aPuhYute5rJoZdh4evkScb1sgbc3bJbVHB4PE9cH2pXfBdo6r2ZcxDMxLfrL9MGeFJgbdTBS2W6XiaHPZXXf95YMaZPgB2pMC7kYNjmgkENsXmVvAJ12CvsQfwj7GEZRh3tB8CumsPWMzteQLUuSnR13f6sNhged8GPTkzMyUofQcH8T8pwHjMuuFD5qDXPJBbMMLCd2tStHZiHU4aHJYpdbTuDt26rbCByZZNPSuhEFvALNRRJQEwHVwxkpjLExSPz4gncBEV3d1QMpdi"
KEY6="Cw7ndvY3yMq6vhfuLwxvLKH1B1KVAyrGaZ1R1ftbGHk3ysWBuVmFuTQwj33WVjKhVZmKkr2Lq5Qje5wJqgD9TSZ3NZ3oaVipkyKcyDcC8tEe7TFjRb3myVd2g4xULZo61zrcWo266me8FzWwVoskGA4tfMSoLoGEWc2i7GwydgUBU1FGuimFp2uY9gM36NvCmADkx5EsMnMNT2sEfpPU7Z4mKFzkGK2pBmU5q4Cqd2hNACZawqUwTjtXjfj8L4ZRxRS3RoWX2VDzCNyexHzm8YgwUYhJbjoKmYdjXXVSRgDcD6icG7XejFBCvyqB2bEcwuNsuAX2ATNUViE4GXpYmYaveGHZBGa1EKmCnnRyCD9A9PiE9ySpnZ8ZLtLe9mEcAwipKm4r3cAeA1ocKanGEY97oWrHct4sysPqkqFbvKD6Kb2Yq2nsEhJXnDVQvt7tgsbqpAYvpD9KLXdwHciVSFP4GZ6K4A1MmcuCevTMyDrxsDpnCcZfgyiCbUREBmz8NHYTHaEw3qD3VvVR6PNTcE5zzvwivgth6mpQtBfySebPQgNpM82PBfcUXtmkg72PDDuz94PsFBiH5PVT7MPREguPfKED42pi9vZAEUGGZfQnzfesjzdxMP9WPDppcu3j8yKMKKwPFdieAT7thrdiF5tPX9yX1MFjMkGHe8Dig9JJUqDhmi2kkZaTaJ6ZDnVNWCoeCcPUcsJRjXNWwFAKTX5guqqv9NpD2iSVGJELNrXDZyXzNxXxzSqi31qHMaDZw5mVmx8XhEG4"
KEY7="3horJt6gBxUDDcx2teNjSg4ncDs14TfBy5qXpCyemv9VAMU2qWUEDEoFN3iL2mWZ7ah8gg7rF9Jiw6qF5Tho7ZQiG31CYCHeoBYG4ibZsyGTYS83iQ1S1F3VSrxVSeEWaRmook4N1N8gYkb4Mf3S3drbMnyVWjtv1YUUkkeU7bvgcncg1jPgWc4pBPm3TYXnUpzXBf4Tc5UEMzaKip8YaGFdeahUucSBH29MKKmWs3umHh7vHSn69XVbV6r2uAnyW5ahnFMjUMdVfb7sUMexcZ1ZTh49YiAFeRXNR2uq3GuhPkQGD772x9D4pBydVxxrcEiT6XeMmPVF2QBtfpAue1g3f4r2Hdhg4Y3fWRvj28bqkhZYsYvihs2pkDiqBNiLsLp5WR5UiS95LwkYfabnBvNW8d1AAFa1Mxm62wdGTZaAriGm2VNND2jLVVL4kxP9xm7aLcmJkgxSukgaZDGqGdRoPWWHZgUtmLQguxw9vzL1edJYkKEfXWWr36SoUeDn4tPXWfApWFX5FS4ki11EuPNxD1ZXAFTxBUtD5s8LxFEBQwoAGabS61cBP3hUPxUw7rhZg7DWMVKRWvaqHSvXxZBKaGEz7iaFfe1ZYfimRD8sph4dXrQn4G1RoewwUP2KZZgJAxUgaDFYGtWDRKfoHwZTsQxiKVtqSFEvuyWV32AVrY6QpSP3Brh3RZc69ywEQzhFRac4CN9oL81GbbvxQv3xi5HitgtcyrCf2coATPpFPepiiE7dvqSrEGxiFzEYraHNTuqbA7E"
KEY8="2pmcvmL2rYC6D4ydq5Ppv4yB2P9YpnPeeTVaCCZi7gbEwbGGasWvYZQgd9i76RULcu1wzUrz9r1VqHR3BxGNYxuE7znd8yeUW6pUDyEiNCxYy2kqEBCPkHxjPubui3jgBK9gMoZ58ceVneQyjz1cm1VQRVsFefeohQpkgUSZAMUp2w9zgN9zJ2p3nWztqS9E7SbM6x3o2E6um2Z4CFkSp9LXShRN1V7jPhLc2bLdnjgDipNMvJDvTv2PvJe4EWhzT3Hu9aGcXAcSYLUnNfmrhGCPK89FojtmNLPNPRXcoTLwAoL1Dm9a2rwPYaxpKJAYVUJXnZPDr88aevYw4heKEHubyXQXQ2BxQi1oV2ZaoEoUAe7yzuXUZVGVGgkHRcjDi3rM7j5apmikMGjCZVcLTsiJV2EL6rXLGTRhTCjj8WrXmbMF8QcxjuiM2n1yoQMeEgKEzcvN1dn84NwigLKu5F7KbTTeqxibA51wB1SwsitUtg4uWgH8Lc3jHfKSvX4Do8qhSmd4zTiXaCJs3UKyu1TZXXurBEK59ZXkbAXKjL2NYtRmfi4B9eZMkbt747aHqVL4hMJmmz3PEEXzmHRZY56P2UJFzKEsWgo3dVH76BZATusZaiiCMzpgDHo3K8unHuYTQe12KwwqF8RDnCyKeDjP8GE8p7r6LFp7XkizBo93ie9tvioSMhw7nrCD6VpQhphjEkxVMCkBGiHrEbLCNim1fUudj8xZuqkWtah3pimpS7KL7hUXkiNxwc3RNZVcmDTEDAU"
KEY9="3UGX3XbAPtDKQ9VzMcsHiz6TA4UQVP2jeSsXyrBrLZuFMMjePfE12wQNVpZ3F2FxxttRkTffDifdJeH3JuRvsCFWBrnVu4teeYv9X9GFhmiAoL6tSGm8wgMbqBnesg5rEG19jTGecJzxHmRejeFMReCaWJDpqJLyM4tzJtGZnoFwpvQC4ZkH5C3By2AZ3hJWZfBsdzBSyz3QAA7RzktrHN8t1sGrXTXh5hBVVJXszZBEamDQLYAdNCtawFxhrmfw5jwenARjLkLJvWCbFLTZNNrsTeruhdMcDkdzeVs5oaKADpNTi8Rk9T67rajHDbafjmYn7tp4j1seVcSN5SiDRXWXjeUsppPMw3Yu6Dwqgfh1cjpBzrm3r6MR7WArrK7s72sTuzkFd2Nj4W7qERyqjSxgurZkNePFgaXcp88dzByBMFhcYBCo8BB5kB2aTvbWyA5S3bsXkCCrC31rWvqjVCQ7qUkFFXwAMii3GJPzFKMvczSk89Nk2b7v7ukuQfU2Gx55waxyfRynqeavzCL5LEQJGt4DmNLhC13opEtzoDrhT9rgPHrPGAtkFVwh8cd7YUpyv3QA8HS3T7mUJQxWRWJPPG6AvfPdVkSPre23hmNAk8KVumJ5yAhJeSXGJeuRoJubHGjv7Kw6YPsaiEnb6nqsVLMqU56m5xrJQFDWjnGZRr3Um91MqAtxmfqYxc53P6v59hHCJLL8si6uaxJ41DdVcm5dapC3sAKu63yR9iFs35ncXGHHqt5piwde"
KEY10="4fQEeuxBejuw64XZsWeQwCuKcn8j4XCXPFu9kQDBWs2QVLo5UAsEokPM7Pfu7pmeSCMP3joV7ZQ5ocN97VgnRgJbkbDvNbYSJvPLQVABh5BewnTuuUs7zCd42JDUCJ4rxULAcoCzDPCU523CihZVBRY6kiP5LUA8ueZpBkn7U7VZ8t3oNS1UjXGWLMrK7SfTd4PkrwuRzdWVXyEsL8rpxgMPEPpYrEY1Tt8h6c9Z2UYs7evj9PKCwJ6RzQqT1ZKkWS2D2PqT5sV3QdpNFNk5Na14hqTVwTwfuSZ7Tw1YDKy8T6PsmPS4T8uCfvzMjTvAksSFQFckX2YT73JgVJpKJdXUbSBWvPgGyh7KSyP4CFR6RnoTzLcgVySfPbxRV6rACZbb9nHHP9MfSLgCQRCn8u3btBHXziDJUnChzLKXXjEG6ikFy8tH56VpUgnBKB5sA5PwbgZo8TVCnbKz1NG7sZzhbgXKnGLo3nRfyWr8uR3SAkUFn8e1YUs5uyaMy888t9qW8tBHHebTF6r5oUzXPUz4eEtEJNJkE9oAkWaZDV4vsbb51rmQqR94jhmQ7DJsc3TvG3b8pdeAjhAHsUY1pq7KcnpApahAHz4QxrX3tJNvfgmavKiuf46wvvqycG5vGYdK7S9S5CukQuHfx9BCzd3V34v464Dp6MwWbyqhtGP89v2Y4EPEX6dPYDC5yAFaY1xsuD3rz9RC3Y1bptVWUPUer9F531R8MmGWc58YASoM4tRTZkmCym3icrQ5T4jW"
KEY11="93qwaJf1BefwzGYyZCPkVEYNw6hUg6ZzxGTSLo3KQfpWNbJ88HEKey85HeaHmDexB34YFLN3B31oKzjTDhjBafGPjcsTugHo2bD2dtTMscLjfWNNFQTLvqkSwSdPtmLu18VVtrJQHHJn7PSFx8TRUnNLc3w4ydycqLXVAgDSGCkqaNCUkcEGJWX5z4aR34qMY94Bk4W3JMshn9nSdLwsQTM5UmWJTERycZcPM74hd9jo7XqtEmfBv1BiXwYQEd2bWQB2bAme73ZCWwQ3rYfDxbdaw8bYmCYPWg9bdeiZikoLCcXgGHy5cQM3Hg7bdZDDyfEWWtuQvyoLnDdhjg2XPVNthActGuZeCNmy7F3y8RsK598PLTobVRd4fosHb5ykigu5L6L8Usa2QvghB59zkKQFjuQeexV2g3frzwdvEaif16ncixn5jNzcH12qNck8L8XG6ED1bPVer1V35UDafwpo8DFxkmYfrptDqsJkC8PzjDZrBw2wvhKkNiSjXRoRhGELQC5SJzkXJAFwSntE6Dz2zWGRpwukir5jHi5JfzMRRAeqnPKo945JZfLyScUF7XPDWMKFmT1rwfqYyfnNGxSk6j33L73Z9gYm6motYW9ToyM6abdWSdFquGs7sF93FG6nHaDLuY8FPWyDTa5Kow9bqTdLHEiGpoR4ntyS1iR2B9wwPfXUyTpDmqksxVWVThmGNvuzsZrn4yDYDFLpFg2TMqD7A7dV7Fs5MmEsLkKevDy4kMWjHQ9LM1f2a6AerotDWLMa"
KEY12="QxsjYLiexYy1UBZp9FEaKZuKjyMdNsMzGfQ1JrXgMy1nJrd7R7TnFB812bCrLMe3LHyNvFFHv7AqisnXXu4cTf8GA4uovNuXnUikVpLaTNVyDQEeEHfLoAgrifB36jXdAmPGt8uUS4L6FYqZe7k6mKyBpEtq9mFSA55oiUQx3F9jDZLcxQg1oRVgvVCerTjZ7sL2YJkD886SSPaa2bvp1rFzN3awyHNCNNyvg5APWGUd1Qb4SjjNkHX98TQPNiApG38F66TXT2vUUaGjztmB6vnS3s28w3MdMiFzd8qJ4mTdd1WF1KkEWEFpxXpNYuRGSPQ1iyc8ZynXmnMYJdmmaeQXSLhEx2FTCdfUkfHRtkoS7dqCU4KkN8JjmXtoEQXQw9tcXhXvHMh9FAPTrFibVCsReiBLFjtrKEeZ3Dbkz5HhryQanVRcLgf9MG9j8rUmzLBaYwNXQEmaMForsY61yGd3ZF94Xx59gzEMznY8DCPPjJ1bhsrPCQ2ut71BamPwZbMi5XfyJq52F4nTgrpvA3PD6uuTS6d9A7NQGEfyCfvppffMsWEEcem9RcuLSfLuep4NED2dGdoi3aQN3SBdFQSrSwL4u1xaj4xfRPNXwBfkttzJP8uTyF1ZoyYTnezKERi5VtQ7L8GVdVAhrsrYdqZf3diuAtdY4zYuKumRGioUyY2gT42ukfbCSv6Pdn4N2BoPw9jXtgZ1pmn4oQZgccjNNbzajKBGMRQhnepSBeXoiDy2mLqtzhNTLfBUkhisYQFeU"
KEY13="ufrxRyJ47kwKA2Wi33bpQFduexSqoGku9jJnkLt7i2tGoyanviFi6pqHj5exMiFyTN5qjr8EKbyYM7jCaC3SzyqK3grdaenVmQrGBpxPqgVqjCtLJDYVPdXYAGCemeL4DTxETxKeSNQ3nKWvZfLyHNpGTmBYdkdwJ64yQi3BEvkCDhxkn2qTN1U4jK4oKhq3BTVhGGUyDA5aP1FVoSp1VoM4Se4VPDNVymLX2TwRwDLeZVuXn7FrEk8AgzBhSyh8EoNzJy1SDXWER3KG1jXw3HEBDcgGhDGjYyeuhfhjy8hTF5d38F6Lfbt7jBCgbJFFSZKqwhkQZ5N5W66mub647u3A9r2zXG1ZhNRttP4rGMu97Pyj1PPm7GrGHeTWphi8B1orpWGfxY4xNhfjdzKx4Nx1RMoCykdVhFgMoF9V2EGL9HKro4feKTTKXT8akncGGUsvbgnmsEp49RDZXM7RFHA7LF7oMJhr4eYqfXVefFLqYUyohQGMPvnek5FX7QHQXvLJZ5h1qcCCSfhBc1HJTyesDdhPBfKuNQ56SfyFVhshsopx55cGxVonG1ZREbDgaVhcSdFc7E8dK1FFwNZjx7opySxZsGVCcMbuF3WuK3eEAux4RGBj95XPVaAeMpu5YrKzcv4G36D9jUPGSniCpKGwj1Jw7Bk4nZbxVdwLAcY3D9oqPxnKHTe49h48833zfh9uk2j4KVv9zoQ25L2SmDmBa38KFQavrK3VbJco6HCLzUTc2U2qCG6ym7fneMLgYKb8n9EVT1TVW"
KEY14="TEESTzhmeZN5qBBcwkr1CjmcCtTRm87Mctj9tSpMkCLhKZ6N3JU48mF7uMLXnfn9fkwvs9oqpvCHpNg3XkYWJASmSchW74ApFRnjEjAkQdGjanTS6URn2fr5vSmMkXVWz4sf9GRvFvJ3qxPJwuu85QroHrwMJxSgSFGSC6Qhd2gLzyKoDLHqRuZaPbPAMSwu6NTrr7grHnaqfAWCbdj9oETHQ8wJJYEpF6CWTDPesgsx4kAuYAZ1FyaCfzypuxueAxLU247rrxvxAxEBq1hFFokBW3pHuLrEUaqaqxe8mwqGnFjc5f1i2KRt8kkrGr9rDjj5GMXNeJsT66rAEFLgDw9md4U7uEQiJxhpU6JfZJUUMFdhHW33j4EVMWfBkMBDPZ5sjmEoW4RkTDWP8X3LHiZxLEj4chgkxSwYGMjuwDxs9cwwqofnRCxd2jccuZDvnqGt4j97mR3sfHu4NuBbar5MpCU4hpgi7qVeYX8M1uMfEGdLSw9rZz12sdUkik8MVzUMDyddUbwwihVXRsmHtoKyNnNcKMnaNJXv8BBFXwcVFjTcum5v7f2drra9EhSM4XV9MBi99bUXbdVevfSSXtHWVDp3Raa5Z26dZ9kUXyZ9cApbta6h1RiJ54sM5gjScAQhyHqb6bEvSctbgfimRZo5hurEV3QByKZvNB28ym6gS7EpE4qJTqf2A3Pb4Qdi6yTysMeM3LUj12Y8nC3R2z98NTNk2KMrLXjaJbGhyNAg484t46NMXvx8bgnMPu9MVHXz7a9YVESVEpMWc6Mz"
KEY15="ZTdG5DGitUiKqbfNSM1AWhxycofUNUvEkRU4y3LPT8DwFEB1Ba4M55reVbXsTQN919FYB3UKbY1BvYGECjrMh9sqvogJEGTy4xJaSeoNwqeDkn43SBq6oBE9NkaX7NWeB717TVAMxadxizA3841QkYsFzEYCBbSgBQCCNUj8hRZ8ZrMCx7zimvfB5BNcAZdP5KeuWJBVCvuTakYJ3MxkH5KjGhW2KKiXHiGZmRqGvF1brLH7aTpT4aDJ7PW5NyEt7qbxgpR2iEqCAb79svUehYKadqrxbAK2dtR4KfKm6nHYinTVeBszfFzRwA1nf2AUGHR9N675cyVcwWzAYAq3BcXTzCGUcG1EY4MZsfviTU5FzPdDhdXR94V6HG3CUhekLUdAAMhDRU3MN6jXPxvfdDT99bxbzro2TqY9qYNMgw21i74MGF5w446W1ZtF7hXxq1dSAWBippPCD1xXshk5E3ywEJ3SyyrmmzspDg7VdJr5Qezk5ko7QmgM4j23oqPNXoU4LcC6VKAnHaNL1CEuCCwSZTGUdxsPkqGEs9yw2ZZpRVBbY3snLvE2PMFE5sEZQELZ4up7WL1o9cAG5qyths9AqECG35TkwUi4G9GT644iL5N5hHRZvAHLHQs9LJTSHwFdsKMZCfcjUNKYLLk68xWPvKPEnq8ZDFQxc5KBH4aUaYgdELFnVfriy4reyP2NY4rHWwPzdYT5N5kcWXqRgq3s421mLtieGjQCyaMQyiW7kUk7J8ywhZq9sA"
KEY16="QxsjYLiexYy1UBZp9FEaKZvhRrNRU487DkvNLUe7hinf5cjejcced6QbWdN8W9PoB8RqcvvUPtFh2dqyZ7Kf52ZLMJn1moxKpEZGg3cMZuAguVPtYWLKVPBKJLMCGQfKNcXm4Tcvyuu8cx7NFf6TcobvYRg2Jkaka3jrQUrXwGipS2dXZNrUHTJipmE1dGqzF4wu7sPdiWFCNiAbW5AJacwrSvvTFSA5JbrRyUttAKxDenGHmiTg74WnmLzm3npwXSdWxKT7oE1vTB9sJnZ9Kg1cSYGTsK6Ez4g5QZHA6FB4uSNpdBynNmHpgvW9CNQDmPoDbAvjiWX9FvMEWJCvvUVHe8TtyRfSuXrHGzJVZYw9h9xuokitrcrmCcuQmZEoApAm3yBMyabdQFo4Wyr5WmZgbxF7pZNf6XFt5hsQYSawFVj3N7iqFf3xgatpcXtS1oC8AWT25kUKUwQx7WPjJ3uH3LtaeCtjDReeVYxKfqY6oretakc5v54e7VXnh1KKoCzKmJaQhELSgbHpmxqdU48f75ia9Jm9D8DaACjb2B2cz3tEUUai1mLN3tm9SUe3AguszZ8UL8fognE8sd5HB3STNvu1JtjehKYRGuErvAbbvQ2cRJC9xh1aPRLYkd8tDf5tFH3A5VQh25hbkt8fjTcfFGenyZfgbGjfXW68ZrBuMVmu8mkt8PQZk7pj3wA6o8i9tdLqeRCmG7Tt3zw2YDJsy45KVWpkkYjUvvvqZCPkScGY1kcWh7DfbjRRdvgVi8z1J"
KEY17="6Rtphde5XFwq8yKJvyhESWAjzzZUD4FhfKHn5j6sfMaFGywWWKm4y6A4Sk15VHRP61gWyLMKbtBih2qTtgw4rCE1ZZ6TCHf3PNsmuXQ4pQPyac8YYrpiL4u7z4JP3SPQz1RAV19WR7fm6qP2LcBZAT8qVJeLXR6KGSE6GVwc4AxgR9Bv5ZUVAjW3kryP4NXvYZ5WDBE2F3emSwYKhH6b7ScYACcyGzWteZpeRrGYLfGoNmVfuuY7PJ3E1pwgcZDjshj3pUuA2xzsxfc9jfKvZGQJzpQfHkVes2qrnLxuHg5PtHd9dtvXrtwRb8btuZHb3wb2w9U5VWrsvhaHAALqq7bUGsLHG4EPbbF2v5MfALcEdjGA2qvSac8Ha3u12JvrVsJ5fNYZ8Kqe4KeRpd2nMppxYaMzRVTV5vshkc6LdwpdBZV4ZymM9G3zwGGbGQMJm86EgWYgbeFLuYVtomcshH3kDdcDke7N6NzxsFn3rJmgRjKGbNTDkLnjgkMrBH8tXpmYntxWHKpNbzN3ezLGgcf1PWtLgcKx96WzHKjPt99MRXW6PBCbY3xA8Nrt5EcB7dni3PrSqk2Ebsj1PANgpaZ6SX63ByiQyLzMEoB3S1mheb5Kz92Y88MSydbvCyqHruxoH9oEhs8c3wtbqNGcGPkP5VTWYJfx1Mjgcg53h7NfZasjHwmyARDFsB3n8d69VY1M7vs2H7WVyEdUXZTeMajhqYb2tAQPksNz1NrMi7J9jL8Ndc18K8tbUj3GpAGfAHBJ"
KEY18="q85TGAnTgBBSch1CSkNHMpwxZRfQPaYwMaKCCczGmgHBtSMP96FcBcUfqGY9Pgy7ZNw4PXRMk9AmSMsVN9YGsHNHLn6SR5Zih64U5MH1KYLGLxps6sJMjzzSFwg7Bjv2TAn3s2qcj3xqY4v6gsAyAKUAdYonQNQi5GCwNxCtuqEvJMtB8ffzRe67HZf3Fk99f3xhn2QhjFfpWcGkd8f88KqQqfigjNW2bFj3AU8rj6x3nCtiCmBvrYZGLEX9vDVwj6X69HUmk9gZmzijmNS4NNprAuZxKR8jzb65FY4ge972NjnxDkTxiaxBN9HNQjHK1NAMxHhTcE2Xva1aJnZK121JhhqYPXEU2ezw81mV3Xe2PeNfn92ygAjD4yezdLXE3CHAYsqa49uR7VkJEvpxJ2Mq6DqHtNUNzV9at71mtJ3RAujXQxmrcRk3F8o3iRSQWhc66wVPDrEmQrsRyfvXa5w9xBq3WKS7xh9g5ovgCfnnbh3DqJCe6X9hHtpwwuBmkKFkYkCutQbdFGVtXThYHjbdnMFCUWrKLSp4GqxcyYVM5Zr17nMHVdFLEvxub9bnEQNTs67ED4w3mj3mbFawRWaqPv33GxfPiULPoKkYNbEBwgNEC79MrGRW5P22qjHYewGiJ4TzQd7NysXD89auC7PWGeJy97vn32AZgbPpYD4To6yjAhBQM7c4CRNDcUPKZR2fZCgd4KQAaCph4VZrZ3Vn4Vdd5acxwm2sgU3WQaJSGoAAn5tq6ZZBsBMpLL"
KEY19="FmYN1DFvXmNVwNGVg3cQ5erRGgrwwh96HA7S1AJZDjxYVVwo9uHA6sTETwszqKJmpd1gwAhFfF4bT3gvLtHuHvit7ViHtCK8yQ96JRLsNkbaLXjvUSWE6hTZXcHxkBYfkWfNdLKnwfPxtYTSFcp2WahdsHdeo4AJkoJwAemmiHhB3ZuczYEz7tk4yEMAo8rTXjDgrG8FP6x3EHFZ3x9ioKeqCjasdhMmUc8az95Y6EhT356kNCiKYYXGXCFFc84JKpKcaUC3HTBSHFHB43e3xM6sS2fg2L3i4EoqbgPcFvwABvLgYquPUaBgMcZRFP3q9EFp1GiXWAhDWD4jT4YvBsv5Q2fNGE6qoECbBBT8q5Ugtf6L1GV7aRuAPVsSYp3QdD1nqZZPA4QMsnLhsrMSRC2aYRusMtChfkeaukFp7bK3u9yu6yUN6AM4qBJE7kDFngjhtHEGPfV6MyLHSpbxw7q85qeBf2q1gCeQWvcvnhDNbh9wGLYtmfELCeVmMZiAu6fZrw9HJp8UKfrVYTF7G6mdxQNAF7CAAhKFXNHVR1EKPYmBfNRwbYHSTrzuWjAgw1RDjR1VAoLYLKrXhvrVvSb5J4XicZSrcXKsaajzuKXKcfTjH9XuiyToUVn26QbXVtGcTMhKMSzufNY7CMriL9iHTYaq65EzT95RB59cyd7u3LjVTRmBWX2NiJqd1RYZ1wXxnLq1JqJNChhiNu4XwEsTpkLMDnKCGSdzMLYUPxTxBUQHgY"
KEY20="5xhkx3fntthLoJDbQZjAUE1YVW4XPef9WanZHgMvuXCCj1MnVQBqKvZBBLRS3ycBKrfBxV4eYthb9uFUrHvgT6Tb1J88Du7wAhGT5UbToaTXCyQcDL3F48y1Mqh3xZiTdvtHZ6SRr6TXTortUrw8d4H6qmQSQdKSjwA5AYwnzKzQzDnytE21fS5eu4BCgaMcP4nxVQ894arsEu6wuenKeDYHsonTSqjuwG1mfaDGPx67xYp9P89CGfA23tE3WgdMSdJbbHPukF9dpyuq7cF3YEgzWPTouAzoStwWy2MZsw9sBLoG5u35qV77Q1vVRyLac3x8rhM2LTWwPe5LUm1Qj74VtiKc6CCwrUyuM7sUhZ4wSdJwgAxozAiSRVQb41RCkn6P7oqsF9F48TsJ156Ne2X9MCfpKBhCGjoYJGn3yN2oNJ31B33wv8X2Jf5QNKuFh9addTcT9agA3RDxw6fb65uaeTKE3AtYVpQbqszyxPrUoLD6L26q4c3ZcNf4DHfyYqLytZgCSNnjX9qaDeh3LmTBAGaWwtKBqnFG78JaWhxas4CFfsp2NXXhXvKr7wp7zNkoS4sgxSRsagWVBCdtjKRyw6HHRYb5S1j82bgZjfpuPRaVHufUgc8xhBcpfqPWJKt8vYdKDq2CcuY16Dzy4PNnapDJCw6PsDmDx1moqBSLfgZ9gzofhJyQFxxHrUcvyxTuBBqgafzjw11N4eLjv2PjYcyT3Q1hVTLtnVWbhpkeNXCs1jW9A"

if [ ! -f ./cash-prototype ]
then
    go build
    echo "Build cash-prototype success!"
fi

if [ ! -f ./bootnode/bootnode ]
then
    cd ./bootnode
    go build
    cd ../
    echo "Build bootnode success!"
fi

if [ ! -f ./benchmark/benchmark ]
then
    cd ./benchmark
    go build
    cd ../
    echo "Build benchmark success!"
fi



tmux new -d -s cash-prototype
tmux new-window -d -n bootnode

tmux send-keys -t cash-prototype:0.0 "cd $SRC && cd bootnode && ./bootnode" ENTER

for ((i=1;i<=$TOTAL;i++));
do
    PORT=$((2333 + $i))
    eval KEY=\${KEY$i}
    
    # open new window in tmux
    tmux new-window -d -n node$1 

    # remove data folder
    rm -rf data$i

    # build options to start node
    opts="--listen 127.0.0.1:$PORT --discoverpeers --datadir data$i --generate --sealerkeyset $KEY --wallet --walletpassphrase '12345678'"
    if [ $i != 1 ]
    then
        opts="--norpc --listen 127.0.0.1:$PORT --discoverpeers --datadir data$i --generate --sealerkeyset $KEY "
    fi
    # send command to node window
    tmux send-keys -t cash-prototype:$i.0 "cd $SRC && ./cash-prototype $opts" ENTER
    echo "Start node with port $PORT, key $KEY and options $opts success"
done

