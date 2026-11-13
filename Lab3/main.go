package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"lab3/utils/crack"
)

func main() {
	target := "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"

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

	fmt.Println("Cracking SHA512:", target)

	scanner := bufio.NewScanner(file)
	attempt := 0
	found := false

	for scanner.Scan() {
		attempt++
		password := strings.TrimSpace(scanner.Text())
		hash := crack.HashSHA512(password)

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
