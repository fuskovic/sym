# Sym

[![Go Reference](https://pkg.go.dev/badge/github.com/fuskovic/sym.svg)](https://pkg.go.dev/github.com/fuskovic/sym)
[![Go Report Card](https://goreportcard.com/badge/github.com/fuskovic/sym)](https://goreportcard.com/report/github.com/fuskovic/sym)
<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-98%25-brightgreen.svg?longCache=true&style=flat)</a>
![CI](https://github.com/fuskovic/sym/actions/workflows/ci.yaml/badge.svg)

A light wrapper around how you would use the standard library for symmetric encryption/decryption anyway.

# Install

    go get -u github.com/fuskovic/sym

# Examples

## Encrypt/Decrypt Strings

```go
    // The key needs to be either 16, 24, or 32 characters in length
    key := os.Getenv("SYMMETRIC_KEY")

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
    // The key needs to be either 16, 24, or 32 characters in length
    key := os.Getenv("SYMMETRIC_KEY")

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
    // The key needs to be either 16, 24, or 32 characters in length
    key := os.Getenv("SYMMETRIC_KEY")

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
