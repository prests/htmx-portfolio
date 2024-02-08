# Portfolio

## Installation

### Setup Golang Deps

```sh
go install github.com/a-h/templ/cmd/templ@latest
go install github.com/cosmtrek/air@latest
```

### Setup TLS

You'll need to use the [generate_cert](https://go.dev/src/crypto/tls/generate_cert.go) tool

```sh
mkdir .cert
cd .cert
go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```
