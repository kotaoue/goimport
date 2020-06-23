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
$ cd $GOPATH/src/github.com/kotaoue/goimport/packages
$ go mod init

$ cd $GOPATH/src/github.com/kotaoue/goimport/packages/car
$ go mod init

$ go mod init github.com/kotaoue/goimport
$ echo "replace github.com/kotaoue/goimport/packages/car => ./packages/car" >> go.mod

# When build completed a require for "packages" are attached on "go.mod".
$ go build main.go

$ go list -m all
github.com/kotaoue/goimport
github.com/kotaoue/goimport/packages/car v0.0.0-00010101000000-000000000000 => ./packages/car
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
rsc.io/quote v1.5.2
rsc.io/sampler v1.3.0
```