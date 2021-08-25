# Sym

[![Go Reference](https://pkg.go.dev/badge/github.com/fuskovic/sym.svg)](https://pkg.go.dev/github.com/fuskovic/sym)
[![Go Report Card](https://goreportcard.com/badge/github.com/fuskovic/sym)](https://goreportcard.com/report/github.com/fuskovic/sym)
<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-94%25-brightgreen.svg?longCache=true&style=flat)</a>
![CI](https://github.com/fuskovic/sym/actions/workflows/ci.yaml/badge.svg)


A small and lightweight symmetric encryption/decryption package.

Useful for encrypting and decrypting strings, bytes, and files.

Only external dependency is [stretchr/testify](https://github.com/stretchr/testify) for the unit tests.

# Install 

    go get -u github.com/fuskovic/sym

# Examples

## Encrypt/Decrypt Strings

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



## Encrypt/Decrypt Bytes

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

## Encrypt/Decrypt Files

    var (
        // The key needs to be either 16, 24, or 32 characters in length
        key                 = os.Getenv("SYMMETRIC_KEY")
        plaintextFilePath   = "/path/to/existing/file.txt"
        // New files get created and existing files get overwritten.
        encryptedFilePath   = "/path/to/new/or/existing/file.txt"
    )

    if err := sym.EncryptFile(key, plaintextFilePath, encryptedFilePath); err != nil {
        // handle error
    }

    in := encryptedFilePath
    out := "/path/to/new/or/existing/file.txt"

    if err := sym.DecryptFile(key, in, out); err != nil {
        // handle error
    }