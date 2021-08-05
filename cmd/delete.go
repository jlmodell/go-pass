package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() {		
	rootCmd.AddCommand(delCmd)
}

var delCmd = &cobra.Command{
	Use: "del",
	Short: "del a password",
	Long: "del a password using key provided via args",
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

		if err := os.Remove(passwordKeyFile); err != nil {
			log.Printf("'%v' was not able to be found",key)
			os.Exit(1)
		}		
	},
}

