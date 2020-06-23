# goimport
The test for import on golang.

## Environment
```shell-session
$ go version
go version go1.14.2 darwin/amd64

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

### Command logs
cf. https://github.com/golang/go/wiki/Modules#quick-start-example
```shell-session
$ go mod init github.com/kotaoue/goimport
$ go build -o main
# When build complete comments "//indirect" are atached on "go.mod".
```