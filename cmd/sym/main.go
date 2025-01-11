package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/fuskovic/sym"
	"github.com/spf13/cobra"
)

var (
	shouldOutputVersion bool
	key                 string
	keyFile             string
	inFile              string
	outFile             string
	size                int
)

func init() {
	encryptCmd.Flags().StringVarP(&key, "key", "k", "", "symetric key to use for encryption.")
	encryptCmd.Flags().StringVarP(&keyFile, "key-file", "f", "", "path to file containing the symetric key.")
	encryptCmd.Flags().StringVarP(&inFile, "in", "i", "", "path to target file for encryption.")
	encryptCmd.Flags().StringVarP(&outFile, "out", "o", "", "filepath to write encrypted ciphertext to.")
	root.AddCommand(encryptCmd)

	decryptCmd.Flags().StringVarP(&key, "key", "k", "", "symetric key to use for decryption.")
	decryptCmd.Flags().StringVarP(&keyFile, "key-file", "f", "", "path to file containing the symetric key.")
	decryptCmd.Flags().StringVarP(&inFile, "in", "i", "", "path to target file for decryption.")
	decryptCmd.Flags().StringVarP(&outFile, "out", "o", "", "filepath to write decrypted plaintext to.")
	root.AddCommand(decryptCmd)

	keyGenCmd.Flags().IntVar(&size, "size", 16, "size of key to generate (valid-sizes: 16, 24, 32)")
	keyGenCmd.Flags().StringVarP(&outFile, "out", "o", "", "filepath to write generated key to.")
	root.AddCommand(keyGenCmd)

	root.Flags().BoolVarP(&shouldOutputVersion, "version", "v", false, "Print installed version.")
}

var root = &cobra.Command{
	Use:   "sym",
	Short: "A utility for encrypting/decrypting text and files.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if shouldOutputVersion {
			println("v1.1.0")
			return
		}
		cmd.Usage()
	},
}

func main() {
	root.Execute()
}

var encryptCmd = &cobra.Command{
	Use:     "encrypt",
	Short:   "encrypt text or files",
	Example: `TODO`,
	Args:    cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		k := getKey(cmd)
		w := getTargetWriter(cmd)
		if inFile != "" {
			if outFile == "" {
				outFile = getDefaultOutFile(cmd)
				fmt.Printf("defaulting to out-file: %s", outFile)
			}
			if err := sym.EncryptFile(k, inFile, outFile); err != nil {
				usageError(cmd, err)
			}
			return
		}
		if len(args) > 0 {
			plainText, err := sym.EncryptString(k, args[0])
			if err != nil {
				usageError(cmd, err)
			}
			w.Write([]byte(plainText))
		}
	},
}

var decryptCmd = &cobra.Command{
	Use:     "decrypt",
	Short:   "decrypt text or files",
	Example: `TODO`,
	Args:    cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		k := getKey(cmd)
		w := getTargetWriter(cmd)
		if inFile != "" {
			if outFile == "" {
				outFile = getDefaultOutFile(cmd)
				fmt.Printf("defaulting to out-file: %s", outFile)
			}
			if err := sym.DecryptFile(k, inFile, outFile); err != nil {
				usageError(cmd, err)
			}
			return
		}
		if len(args) > 0 {
			plainText, err := sym.DecryptString(k, args[0])
			if err != nil {
				usageError(cmd, err)
			}
			w.Write([]byte(plainText))
		}
	},
}

var keyGenCmd = &cobra.Command{
	Use:     "keygen",
	Short:   "generate keys",
	Example: `TODO`,
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		key, err := sym.KeyGen(size)
		if err != nil {
			usageErrorf(cmd, "failed to generate key: %w", err)
		}
		getTargetWriter(cmd).Write([]byte(key))
	},
}

func usageError(cmd *cobra.Command, errMsg any) {
	var msg string
	switch t := errMsg.(type) {
	case string:
		msg = t
	case error:
		msg = t.Error()
	default:
		msg = fmt.Sprintf("%v", msg)
	}
	cmd.Usage()
	fmt.Println(msg)
	os.Exit(1)
}

func usageErrorf(cmd *cobra.Command, format string, args ...any) {
	usageError(cmd, fmt.Errorf(format, args...))
}

func getTargetWriter(cmd *cobra.Command) io.Writer {
	var w io.Writer = os.Stdout
	var err error
	if outFile != "" {
		w, err = os.Create(outFile)
		if err != nil {
			usageErrorf(cmd, "failed to open %s: %w", outFile, err)
		}
	}
	return w
}

func getKey(cmd *cobra.Command) string {
	if key == "" && keyFile == "" {
		usageError(cmd, "must provide key or key-file")
	}
	if key != "" && keyFile != "" {
		usageError(cmd, "key and keyfile provided; must be either or")
	}

	var err error
	if keyFile != "" {
		key, err = sym.KeyFromFilePath(keyFile)
		if err != nil {
			usageErrorf(cmd, "failed to get key from file: %w", err)
		}
	}
	return key
}

func getDefaultOutFile(cmd *cobra.Command) string {
	base := strings.TrimSuffix(inFile, path.Ext(inFile))
	if cmd.Use == "encrypt" {
		return base + ".enc"
	}
	return base + ".txt"
}
