package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pwdfy",
	Short: "A tiny CLI that makes a word into a functioning password",
	Long: `A tiny CLI that makes a given word into a function password,
          with numbers, special characters, and uppercase / lowercase letters`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(passwordify(args[0]))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pwdfy.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func passwordify(baseWord string) string {
	pwd := baseWord
	s := strings.Split(pwd, "")
	var pass []string

	passlen := len(baseWord)
	rem := 12 - (passlen + 3)

	if rem > 0 {
		for i := 0; i < rem; i++ {
			pass = append(pass, RandStringBytes(1))
		}
	}

	pass = append(pass, "!")

	for i := 0; i < len(s); i++ {
		s[i] = checkChar(s[i])

		pass = append(pass, s[i])
	}

	pass = append(pass, "##")

	pwd = strings.Join(pass, "")

	return pwd
}

func checkChar(s string) string {
	if _, err := strconv.Atoi(s); err == nil {
		return s
	}

	switch s {
	case "a":
		if rand.Intn(2) == 0 {
			return "4"
		} else {
			return "a"
		}
	case "e":
		if rand.Intn(2) == 0 {
			return "3"
		} else {
			return "e"
		}
	case "i":
		if rand.Intn(2) == 0 {
			return "1"
		} else {
			return "i"
		}
	default:
		return changeCase(s)
	}
}

func changeCase(s string) string {
	if rand.Intn(2) == 0 {
		return strings.ToUpper(s)
	}
	return s
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
