# dir-tree-print
Print a crc32 hash of every file in a directory


Usage:

```shell
git clone https://github.com/reddaly/dir-tree-print.git
cd dir-tree-print
go build
./dir-tree-print --dir /tmp
```

Example output:

```
$ ./dir-tree-print --dir $PWD/
LICENSE: 11357 bytes; crc32 2069693628
README.md: 65 bytes; crc32 335623904
dir-print: 2468551 bytes; crc32 3067024907
dir-tree-print: 2516087 bytes; crc32 2552466017
go.mod: 91 bytes; crc32 2475372900
go.sum: 330 bytes; crc32 879836845
main.go: 1158 bytes; crc32 28460343
main.go~: 1053 bytes; crc32 1110480648
```
