package seki

import (
	"bufio"
	"log"
	"os"
	"path"
	"strings"
)

func LoadEnvFile() error {
	workingDir, _ := os.Getwd()
	envFilePath := path.Join(workingDir, "/.env")
	envFile, err := os.Open(envFilePath)

	if err != nil {
		return err
	}
	defer envFile.Close()

	// create a new scanner to read each row
	scanner := bufio.NewScanner(envFile)
	scanner.Split(bufio.ScanLines)

	//loop through each row of the .env file
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, "=") {
			continue
		}

		envVar := strings.Split(line, "=")
		err := os.Setenv(envVar[0], envVar[1])

		if err != nil {
			return err
		}
	}

	//	check for errors with scanner.Scan
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}
