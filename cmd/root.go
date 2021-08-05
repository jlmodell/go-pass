package cmd

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const file = "c:\\temp\\.go-pass"

var (
	err error	
	SecretKey string = ""
	Len int	
	rootCmd = &cobra.Command{
		Use: "cobra",
		Short: "init with key or random will be given",
		Long: "initialize with a key or random will be generated for you",		
	}
)

// Execute executes the root cmd
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	var secretKey []byte
	var skf string

	if err := ensureDir(file); err != nil {
		fmt.Println("Directory creation failed with error: " + err.Error())
		os.Exit(1)
	}

	skf, err := filepath.Abs(file + "\\.skf")

	_, err = os.Stat(skf)
	if err != nil {
		os.Create(skf)
	}
	
	secretKey, err = ioutil.ReadFile(skf)
	if err != nil {
		log.Println("Master key is not initialized.\n\trun again without arguements")
		os.Exit(1)
	}	

	if len(secretKey) == 0 {
		fmt.Println("initializing your secret key...")
		bytes := make([]byte, 32)
		if _, err := rand.Read(bytes); err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}	
		SecretKey = hex.EncodeToString(bytes)				
		
		if err = ioutil.WriteFile(skf, []byte(SecretKey), 0); err != nil {
			panic(err)
		}
	} else {
		SecretKey = string(secretKey)		
	}		
}

func ensureDir(dirName string) error {
    err := os.Mkdir(dirName, os.ModeDir)
    if err == nil {
        return nil
    }
    if os.IsExist(err) {
        // check that the existing path is a directory
        info, err := os.Stat(dirName)
        if err != nil {
            return err
        }
        if !info.IsDir() {
            return errors.New("path exists but is not a directory")
        }
        return nil
    }
    return err  
}