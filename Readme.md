# このリポジトリについて

このリポジトリはmarnysanがGoのちょっとした動作確認用に作ったリポジトリです。


# Usage

```sh
$ make run c={cmd配下のディレクトリ名}
```

例
```sh
$ make run c=rand

go build -o bin/rand cmd/rand/main.go
./bin/rand
6
7
3
0
9
8
4
0
9
9
```