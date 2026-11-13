package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"lab1/utils/crack"
)

func main() {
	target := "6a85dfd77d9cb35770c9dc6728d73d3f"

	file, err := os.Open("nord_vpn.txt")
	if err != nil {
		fmt.Println("Error opening wordlist:", err)
		return
	}
	defer file.Close()

	verboseFile, err := os.Create("verbose.txt")
	if err != nil {
		fmt.Println("Error creating verbose.txt:", err)
		return
	}
	defer verboseFile.Close()

	fmt.Println("Cracking MD5:", target)

	scanner := bufio.NewScanner(file)
	attempt := 0
	found := false

	for scanner.Scan() {
		attempt++
		password := strings.TrimSpace(scanner.Text())
		hash := crack.HashMD5(password)

		line := fmt.Sprintf("Attempt %d: %s -> %s\n", attempt, password, hash)
		fmt.Print(line)
		verboseFile.WriteString(line)

		if hash == target {
			result := fmt.Sprintf("\nFound password: %s at attempt %d\n", password, attempt)
			fmt.Print(result)
			verboseFile.WriteString(result)
			found = true
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	if !found {
		fmt.Println("Password not found.")
	}
}
