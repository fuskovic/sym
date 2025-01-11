package main

import (
	"github.com/spf13/cobra"
	"github.com/fuskovic/sym"
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
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var decryptCmd = &cobra.Command{
	Use:     "decrypt",
	Short:   "decrypt text or files",
	Example: `TODO`,
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var keyGenCmd = &cobra.Command{
	Use:     "keygen",
	Short:   "generate keys",
	Example: `TODO`,
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
