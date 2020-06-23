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
```shell-session
$ tree $GOPATH/src/github.com/kotaoue/goimport -L 3
/$GOPATH/src/github.com/kotaoue/goimport
├── LICENSE
├── README.md
├── go.mod
├── go.sum
├── main.go
└── packages
    └── car
        ├── car.go
        └── go.mod
```

## cf.
* https://github.com/golang/go/wiki/Modules#quick-start-example
* https://github.com/golang/go/wiki/Modules#when-should-i-use-the-replace-directive

## Command logs
```shell-session
$ cd $GOPATH/src/github.com/kotaoue/goimport/packages/car
$ go mod init

$ cd $GOPATH/src/github.com/kotaoue/goimport/
$ go mod init github.com/kotaoue/goimport
$ go mod edit -replace github.com/kotaoue/goimport/packages/car=./packages/car

# When build completed a require for "packages" are attached on "go.mod".
$ go build main.go

$ go list -m all
github.com/kotaoue/goimport
github.com/kotaoue/goimport/packages/car v0.0.0-00010101000000-000000000000 => ./packages/car
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
rsc.io/quote v1.5.2
rsc.io/sampler v1.3.0
```

## Problrem
```shell-session
$ go vet main.go 
# command-line-arguments
vet: ./main.go:14:18: ignition not declared by package car

$ go run main.go 
# command-line-arguments
./main.go:14:14: cannot refer to unexported name car.ignition
./main.go:14:14: undefined: car.ignition
```