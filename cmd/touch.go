package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/jlmodell/go-pass/utils"
	"github.com/spf13/cobra"
)

func init() {
	// rootCmd.PersistentFlags().StringP("touch", "t", "", "generate a new password")
	// rootCmd.PersistentFlags().StringP("grep", "g", "", "find a password")
	touchCmd.PersistentFlags().IntVarP(&Len,"len","l",22,"length of password <22>")

	rootCmd.AddCommand(touchCmd)
}

var touchCmd = &cobra.Command{
	Use: "touch",
	Short: "generate a password",
	Long: "generate a password, store as hash, return plainstring to stdout",
	Run: func(cmd *cobra.Command, args []string) {
		var key string = ""
		var pwd string = ""		

		if len(args) < 1 {
			log.Println("No password key was presented")
			os.Exit(1)
		}		

		key = args[0]
		
		if len(args) >= 2 && args[1] != "" {
			pwd = args[1]
		}

		passwordKeyFile, err := filepath.Abs(fmt.Sprintf("%v\\.%v", file, key))
		_, err = os.Stat(passwordKeyFile)
		if err != nil {
			os.Create(passwordKeyFile)
		} else {
			log.Printf("'%v' already exists, [grep]",key)
			os.Exit(1)
		}

		if pwd == "" {
			pwd = string(generatePassword(Len))
		}

		encryptedPwd := utils.Encrypt(pwd,SecretKey)
		
		if err = ioutil.WriteFile(passwordKeyFile, []byte(encryptedPwd), 0); err != nil {
			panic(err)
		}
	},
}

func generatePassword(length int) []byte {
	buf := make([]byte, length)

	rand.Seed(time.Now().UnixNano())		
	
	digits := "0123456789"
	specials := "~=+%^*/()[]{}/!@#$?|"
	
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz" + digits + specials

	buf[0] = digits[rand.Intn(len(digits))]
	buf[1] = specials[rand.Intn(len(specials))]

	for i := 2; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}

	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})

	return buf
}