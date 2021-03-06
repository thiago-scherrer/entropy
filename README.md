# entropy

Software to generate a cryptographically secure random seed.

[![Build](https://github.com/thiago-scherrer/entropy/actions/workflows/build.yml/badge.svg)](https://github.com/thiago-scherrer/entropy/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/thiago-scherrer/entropy)](https://goreportcard.com/report/github.com/thiago-scherrer/entropy)

- Using /dev/urandom on Linux.
- On OpenBSD getentropy(2).
- On other Unix-like systems /dev/urandom.
- On Windows systems, Reader uses the CryptGenRandom API. On Wasm uses the Web Crypto API.

## reqs

- golang >=1.16
- docker (only with use docker)

## run with docker

Build:

```sh
docker build -t entropy .
```

Running and create a seed with 20 bytes

```sh
docker run entropy -size 20
```

## test

```sh
go test
```

## build without docker

Build:

```sh
go build entropy.go
```

Running and create a seed with 20 bytes

```sh
./entropy -size 20
```

Help:

```bash
./entropy -help
```

## refs

- [crypto/rang](https://golang.org/pkg/crypto/rand/)
- [getrandom](http://man7.org/linux/man-pages/man2/getrandom.2.html)
- [enconding/base64](https://golang.org/pkg/encoding/base64)
- [Cryptographically secure pseudorandom](https://en.wikipedia.org/wiki/Cryptographically_secure_pseudorandom_number_generator)
