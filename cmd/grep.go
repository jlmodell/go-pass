package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/atotto/clipboard"
	"github.com/jlmodell/go-pass/utils"
	"github.com/spf13/cobra"
)

func init() {		
	rootCmd.AddCommand(grepCmd)
}

var grepCmd = &cobra.Command{
	Use: "grep",
	Short: "retrieve a password",
	Long: "retrieve a password using key provided via args",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Println("No password key was presented")
			os.Exit(1)
		}		
		key := args[0]
		passwordKeyFile, err := filepath.Abs(fmt.Sprintf("%v\\.%v", file, key))		
		if err != nil {
			log.Printf("'%v' is not present", key)
			os.Exit(1)
		}		

		pwd,err := ioutil.ReadFile(passwordKeyFile)
		if err != nil {
			panic(err)
		}

		decryptedPwd := utils.Decrypt(string(pwd),SecretKey)
		
		clipboard.WriteAll(decryptedPwd)
		fmt.Printf("'%v' password was saved to your clipboard [`ctrl+v`]", key)
	},
}

