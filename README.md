# entropy

Software to generate a cryptographically secure random seed.

Using /dev/urandom on Linux. On OpenBSD getentropy(2). On other Unix-like systems /dev/urandom. On Windows systems, Reader uses the CryptGenRandom API. On Wasm uses the Web Crypto API.

## reqs

- golang >=1.11.2
- crypto/rand golang module
- encoding/base64 golang module
- flag golang module
- fmt golang module
- os golang module

## build

Run:

```bash
go build entropy.go
```

## refs

- [crypto/rang](https://golang.org/pkg/crypto/rand/)
- [getrandom](http://man7.org/linux/man-pages/man2/getrandom.2.html)
- [enconding/base64](https://golang.org/pkg/encoding/base64)
- [Cryptographically secure pseudorandom](https://en.wikipedia.org/wiki/Cryptographically_secure_pseudorandom_number_generator)
