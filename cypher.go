package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
)

var sliceArray []int32

func shift(r rune, shift int) rune {
	// Shift character by specified number of places.
	// ... If beyond range, shift backward or forward.
	s := int(r) - shift
	if unicode.IsSpace(rune(r)) || !unicode.IsLetter(rune(r)) {
		return rune(r)
	} else if s > 'z' {
		return rune(s - 26)
	} else if s < 'a' {
		return rune(s + 26)
	}
	return rune(s)
}

func main() {
	var cmdDecrypt = &cobra.Command{
		Use:   "decrypt [string to decypt]",
		Short: "decrypt a given string",
		Long: `descript will take an encrypted string,
		then shift the string by some given amount.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			encryptedValue := strings.ToLower(strings.Join(args, " "))
			shiftFlag := cmd.PersistentFlags().Lookup("shift")
			shiftVal, err := strconv.Atoi(shiftFlag.Value.String())
			if err != nil {
				panic(err)
			}

			decryptedValue := strings.Map(func(r rune) rune {
				return shift(r, shiftVal)
			}, encryptedValue)
			fmt.Println("Decrypted: " + decryptedValue)
		},
	}

	var cmdRepeat = &cobra.Command{
		Use:   "times [string to echo]",
		Short: "decrypt a given string using a repeating list.",
		Long: `decrypt the string using the required shift list.
		This list will be used iterativly over the string to shift while decrypting.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			encryptedValue := strings.ToLower(strings.Join(args, " "))
			for _, v := range sliceArray {
				fmt.Print(v)
			}
			index := 0
			decryptedValue := strings.Map(func(r rune) rune {
				if unicode.IsSpace(rune(r)) || !unicode.IsLetter(rune(r)) {
					return r
				}
				s := shift(r, int(sliceArray[index]))
				index += 1
				if index >= len(sliceArray) {
					index = 0
				}
				return s
			}, encryptedValue)
			fmt.Println("Decrypted: " + decryptedValue)
		},
	}

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdDecrypt)
	cmdDecrypt.AddCommand(cmdRepeat)
	cmdDecrypt.PersistentFlags().Int32P("shift", "s", 1, "Value to shift the given input by")
	// TODO: fix nil as default
	cmdRepeat.PersistentFlags().Int32SliceVarP(&sliceArray, "shift", "s", nil, "A repeating array value to shift by")

	rootCmd.Execute()
}
