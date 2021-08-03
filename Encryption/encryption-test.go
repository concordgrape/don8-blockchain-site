/*
*
*	PROJECT			:	Don8 - Encrypted BlockChain Application
*	FILE			:	encryption-test.go
*	PROGRAMER		:	Sky Roth
*	CREATED			:	April 10, 2021
*
*	DESCRIPTION		:	This file won't be used for production, it's used to test our encryption method
*							Encrypts with RSA and stores in a temporary file using PEM
 */

package main

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func main() {

	//	Step 1.	Check command line arguments
	if len(os.Args) < 4 {
		fmt.Println("Error: Please supply more arguments - <TITLE> <DESCRIPTION> <WALLET ADDRESS>")
		os.Exit(1)
	} else if len(os.Args) == 1 {
		fmt.Println("Error: You are missing command line arguments! Please supply them as follows: <TITLE> <DESCRIPTION> <WALLET ADDRESS>")
		fmt.Println("Remember you can get the public key from the website!")
		os.Exit(1)
	}

	//	Step 2.	Open the text file that stores the public key
	file, err := os.Open("key.txt")
	if err != nil {
		os.Exit(-1)
	}

	//	Step 3.	Read the text file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	//	close the file after reading it
	file.Close()

	var pubPEM string
	for _, word := range text {
		pubPEM += word + "\n"
	}

	//	Step 4.	Decode the key with pem
	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		panic("failed to parse PEM block containing the public key")
	}

	//	Step 5.	Parse the public key so we can use it for encryption
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic("failed to parse DER encoded public key: " + err.Error())
	}

	//	Step 6.	Create the message and encrypt it with the public key
	fullPub := pub.(*rsa.PublicKey)

	// define the message
	message := []byte("{'title': '" + os.Args[1] + "', 'description': '" + os.Args[2] + "', 'wallet': '" + os.Args[3] + "'}")

	// store the encrypted message
	encrypted, _ := rsa.EncryptPKCS1v15(rand.Reader, fullPub, message)

	//	Step 7.	Create the JSON string to post it to the block chain
	reqBody, err := json.Marshal(map[string]string{
		"data": string(encrypted),
	})
	if err != nil {
		print(err)
		os.Exit(3)
	}

	//	Step 8.	POST the encrypted string
	s := "http://localhost:3001/mineBlock"
	u, err := url.Parse(s)
	if err != nil {
		print(err)
		os.Exit(2)
	}

	//	Step 9.	Set the POST request as JSON
	resp, err := http.Post(u.String(),
		"application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		print(err)
		os.Exit(-1)
	}
	defer resp.Body.Close()

	//	Step 10. Finish the program
	fmt.Printf("Success! Fund was posted to the blockchain")
	os.Exit(0)
}
