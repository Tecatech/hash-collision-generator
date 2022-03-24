# Tool for performing hash collision attacks on SHA-XX cryptographic algorithms

```sh
$ go build
$ go run .
Birthday paradox attack:
2022/03/23 15:10:22 [SHA15] alloc: 900 KiB
2022/03/23 15:10:22 [SHA15] time: 19.0189ms
2022/03/23 15:10:22 [SHA16] alloc: 1404 KiB
2022/03/23 15:10:22 [SHA16] time: 27.5181ms
2022/03/23 15:10:22 [SHA17] alloc: 2251 KiB
2022/03/23 15:10:22 [SHA17] time: 41.6503ms
2022/03/23 15:10:24 [SHA18] alloc: 3574 KiB
2022/03/23 15:10:24 [SHA18] time: 62.1146ms
2022/03/23 15:10:24 [SHA19] alloc: 4549 KiB
2022/03/23 15:10:24 [SHA19] time: 85.1164ms
2022/03/23 15:10:25 [SHA20] alloc: 5205 KiB
2022/03/23 15:10:25 [SHA20] time: 121.5366ms
Pollard rho algorithm:
2022/03/23 15:10:54 [SHA15] alloc: 681 KiB
2022/03/23 15:10:54 [SHA15] time: 1m40.0695101s
2022/03/23 15:12:37 [SHA16] alloc: 1095 KiB
2022/03/23 15:12:37 [SHA16] time: 3m10.7641922s
2022/03/23 15:15:47 [SHA17] alloc: 1542 KiB
2022/03/23 15:15:47 [SHA17] time: 4m46.3459585s
2022/03/23 15:20:29 [SHA18] alloc: 2373 KiB
2022/03/23 15:20:29 [SHA18] time: 7m12.4940705s
2022/03/23 15:27:40 [SHA19] alloc: 3655 KiB
2022/03/23 15:27:40 [SHA19] time: 9m48.9932266s
2022/03/23 15:37:28 [SHA20] alloc: 4427 KiB
2022/03/23 15:37:28 [SHA20] time: 14m25.681136s
$ go clean
```
