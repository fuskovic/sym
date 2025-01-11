package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fuskovic/sym"
)

var (
	key         string
	keyFilePath string
	keySize     int
	in          string
	out         string
)

func init() {
	flag.StringVar(&key, "key", "", "symetric key used for encryption/decryption")
	flag.StringVar(&keyFilePath, "key-file", "", "path to the symetric key file")
	flag.StringVar(&in, "in", "", "target input file")
	flag.StringVar(&out, "out", "", "target output file")
	flag.IntVar(&keySize, "size", 0, "size of key to generate")
}

func usageExit(errMsg any) {
	var msg string
	switch t := errMsg.(type) {
	case string:
		msg = t
	case error:
		msg = t.Error()
	default:
		msg = fmt.Sprintf("%v", t)
	}
	println(msg)
	flag.Usage()
	os.Exit(1)
}

func usageExitf(format string, args ...any) {
	usageExit(fmt.Errorf(format, args...))
}

func main() {
	if key == "" && keyFilePath == "" {
		usageExit("must provide either -key or -keyfile flag")
	}

	if key != "" && keyFilePath != "" {
		usageExit("provided both a -key and -keyfile flag; must be one or the other")
	}

	if keyFilePath != "" {
		var err error
		key, err = sym.KeyFromFilePath(keyFilePath)
		if err != nil {
			usageExitf("failed to read key from %q: %w", keyFilePath, err)
		}
	}

	args := os.Args[1:]
	if len(args) > 0 {
		subCmd := args[0]
		var subCmdArgs []string
		if len(args) > 1 {
			subCmdArgs = args[1:]
		}

		switch subCmd {
		case "encrypt":
			if len(subCmdArgs) > 0 {
				sym.EncryptString(key, subCmdArgs[0])
			}
		case "decrypt":
			if len(subCmdArgs) > 0 {

			}
		case "key":
		}
	}
}
