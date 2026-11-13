package main

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"regexp"
)

func main() {
	query := "%62%36%63%63%62%34%65%63%65%35%34%35%34%64%63%61%65%35%31%37%37%38%62%33%65%32%33%39%65%62%63%32"
	leakedHash, _ := url.QueryUnescape(query)

	fmt.Println("Leaked Hash:", leakedHash)

	password := "meowmeow"

	h := md5.Sum([]byte(password))
	myHash := fmt.Sprintf("%x", h)

	fmt.Println("Password:", password)
	fmt.Println("MD5 Hash:", myHash)

	if myHash == leakedHash {
		fmt.Println("Hash match!")
	} else {
		fmt.Println("Hash not match!")
	}

	flag := "cryptoCTF{" + password + "}"
	fmt.Println("Flag:", flag)

	r, _ := regexp.Compile(`^cryptoCTF\{(?:\x6d\x65\x6f\x77){2}\}$`)
	fmt.Println("Verify regex:", r.MatchString(flag))
}
