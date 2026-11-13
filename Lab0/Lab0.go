package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"

	"crypto/sha3"
)

func main() {
	var value1, value2 string

	fmt.Print("Please input value 1: ")
	fmt.Scan(&value1)

	fmt.Print("Please input value 2: ")
	fmt.Scan(&value2)

	displayHash("MD5", hashMD5(value1), hashMD5(value2))
	displayHash("SHA1", hashSHA1(value1), hashSHA1(value2))
	displayHash("SHA256", hashSHA256(value1), hashSHA256(value2))
	displayHash("SHA512", hashSHA512(value1), hashSHA512(value2))
	displayHash("SHA3-256", hashSHA3(value1), hashSHA3(value2))
}

func displayHash(method, value1, value2 string) {
	fmt.Printf("Hash (%s):\n", method)
	fmt.Printf("  Output A = %s\n", value1)
	fmt.Printf("  Output B = %s\n", value2)
	if value1 == value2 {
		fmt.Println("  => Match!")
	} else {
		fmt.Println("  => No Match!")
	}
}

func hashMD5(input string) string {
	h := md5.Sum([]byte(input))
	return fmt.Sprintf("%x", h)
}

func hashSHA1(input string) string {
	h := sha1.Sum([]byte(input))
	return fmt.Sprintf("%x", h)
}

func hashSHA256(input string) string {
	h := sha256.Sum256([]byte(input))
	return fmt.Sprintf("%x", h)
}

func hashSHA512(input string) string {
	h := sha512.Sum512([]byte(input))
	return fmt.Sprintf("%x", h)
}

func hashSHA3(input string) string {
	h := sha3.Sum256([]byte(input))
	return fmt.Sprintf("%x", h)
}
