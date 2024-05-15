package cmd

import (
	"fmt"
	"os"

	"github.com/karthikkalarikal/password_generator/generator"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "password-generator",
		Short: "A CLI password generator",
		Long: `passwor-generator is a CLI generaotor written in go
		that would generate password according to users requirements.`,
		Run: generate,
	}
	length          uint
	upperCase       bool
	numbers         bool
	symbols         bool
	similarChars    bool
	symbolAmbiguous bool
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().UintVarP(&length, "length", "l", 6, "length of password")
	rootCmd.PersistentFlags().BoolVarP(&upperCase, "uppercase", "u", false, "include uppercase letters in password")
	rootCmd.PersistentFlags().BoolVarP(&numbers, "numbers", "n", false, "include numbers in password")
	rootCmd.PersistentFlags().BoolVarP(&symbols, "symbols", "s", false, "include symbols in password")
	rootCmd.PersistentFlags().BoolVarP(&similarChars, "similar-characters", "", true, "exclude similar characters in password(il1Lo0O)")
	rootCmd.PersistentFlags().Lookup("similar-characters").Shorthand = "c"

	rootCmd.PersistentFlags().BoolVarP(&symbolAmbiguous, "ambiguous-symbols", "", true, "exclude ambiguous symbols in password(<>[](){}:;'/|\\,)")
	rootCmd.PersistentFlags().Lookup("ambiguous-symbols").Shorthand = "a"
}

func generate(_ *cobra.Command, args []string) {
	if length < 6 {
		fmt.Println("error: the length must be at least 6 to generate a valid password.")
		length = 6
	}
	config := generator.Config{
		Length:          length,
		UpperCase:       upperCase,
		Numbers:         numbers,
		Symbols:         symbols,
		SimilarChars:    similarChars,
		SymbolAmbigious: symbolAmbiguous,
	}

	g := generator.New(&config)
	pwd := g.GeneratePassword()
	fmt.Println("password: ", pwd)
}
