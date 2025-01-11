# Sym

[![Go Reference](https://pkg.go.dev/badge/github.com/fuskovic/sym.svg)](https://pkg.go.dev/github.com/fuskovic/sym)
[![Go Report Card](https://goreportcard.com/badge/github.com/fuskovic/sym)](https://goreportcard.com/report/github.com/fuskovic/sym)
<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-98%25-brightgreen.svg?longCache=true&style=flat)</a>
![CI](https://github.com/fuskovic/sym/actions/workflows/ci.yaml/badge.svg)

Can be used as an importable pkg or as a CLI.

# Use as a CLI

    go install github.com/fuskovic/sym/cmd/sym@latest

# Examples

## Generate a Key

    // default key size is 16
    sym keygen

    // other valid sizes include 24, and 32
    sym keygen --size 32

    // output to a file
    sym keygen -o key.txt

## Encrypt a file

    sym encrypt \
        -f key.txt \
        -i inputfile.txt \
        -o outputfile.enc

## Decrypt a file

    sym decrypt \
        -f key.txt \
        -i ciphertext.enc \
        -o plaintext.txt


# Use as a library

    go get -u github.com/fuskovic/sym

# Examples

## Generate a Key

```go

    // use KeyGen if you want a specific size
    key, err := sym.KeyGen(32)
    if err != nil {
        // handle err
    }

    // or use MustKeyGen to generate a default size (16) key
    key := sym.MustKeyGen()

```

## Use key from file

```go

    key, err := sym.KeyFromFilePath("/path/to/my.key")
    if err != nil {
        // handle err
    }

```

## Encrypt/Decrypt Strings

```go

    key := sym.MustKeyGen()

    ciphertext, err := sym.EncryptString(key, "hello world")
    if err != nil {
        // handle error
    }

    plaintext, err := sym.DecryptString(key, ciphertext)
    if err != nil {
        // handle error
    }

```

## Encrypt/Decrypt Bytes

```go

    key := sym.MustKeyGen()

    ciphertext, err := sym.EncryptBytes(key, []byte("hello world"))
    if err != nil {
        // handle error
    }

    plaintext, err := sym.DecryptBytes(key, ciphertext)
    if err != nil {
        // handle error
    }

```

## Encrypt/Decrypt Files

```go

    key := sym.MustKeyGen()

    err := sym.EncryptFile(key,
        "/path/to/existing/plaintext/file.txt",
        "/path/to/write/the/encrypted/file/to.enc",
    )

    if err != nil {
        // handle error
    }

    err = sym.DecryptFile(key,
        "/path/to/encrypted/file.enc",
        "/path/to/write/decrypted/file/to.txt",
    )

    if err != nil {
        // handle error
    }
    
```
