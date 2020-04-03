# entropy

Software to generate a cryptographically secure random seed.

[![Go Report Card](https://goreportcard.com/badge/github.com/thiago-scherrer/entropy)](https://goreportcard.com/report/github.com/thiago-scherrer/entropy)

- Using /dev/urandom on Linux.
- On OpenBSD getentropy(2).
- On other Unix-like systems /dev/urandom.
- On Windows systems, Reader uses the CryptGenRandom API. On Wasm uses the Web Crypto API.

## reqs

- golang >=1.11.2

## test

```bash
go test
```

## build

```bash
go build entropy.go
```

## run

Create a seed with 20 bytes

```bash
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
