# goimport
The test for import on golang.

## env
```shell-session
$ echo $GO111MODULE
on
```

## Tree
```
$ tree $GOPATH -L 5
/Users/kota.oue/go
├── bin
├── pkg
│   └── mod
│       ├── github.com
│       └── golang.org
└── src
    └── github.com
        └── kotaoue
            ├── goimport
            │   ├── LICENSE
            │   ├── README.md
            │   ├── main.go
            │   └── packages
            │       └── car.go
            ├── gotour
            └── the-go-programming-language
```

### go mod
```shell-session
$ go mod init github.com/kotaoue/goimport
```