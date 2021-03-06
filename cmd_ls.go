package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

const historyFileName = ".ghcc_history"

func ls(c *cli.Context) error {
	file, err := os.Open(getHistoryPath())
	if os.IsNotExist(err) {
		fmt.Println("No history")
		return nil
	}

	if err != nil {
		return err
	}

	defer file.Close()

	history, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	fmt.Print(string(history))
	return nil
}

func getHistoryPath() string {
	homePath, _ := os.UserHomeDir()
	return filepath.Join(homePath, historyFileName)
}

// Add new result to history
// If a history file doesn't exist, create new one
func writeResultToHistory(result string) {
	file, _ := os.OpenFile(getHistoryPath(), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	defer file.Close()

	fmt.Fprintln(file, result)
}
