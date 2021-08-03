/*
*
*	PROJECT			:	Don8 - Encrypted BlockChain Application
*	FILE			:	client.go
*	PROGRAMER		:	Sky Roth
*	CREATED			:	April 10, 2021
*
*	DESCRIPTION		:	This file allows users to enter new blocks for the blockchain without access to our website
*							Everything can be connected through the blockchain without a middle man
 */

package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func main() {

	//	Step 1.	Define the contents of each block
	var title string
	var description string
	var walletAddress string

	//	Step 2. Get contents from user
	fmt.Println("\n\nWELCOME TO DON8 DONATIONS!\nTo get started please enter a title for the fund: ")

	// get the title
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Println("Error: Scanning failed")
		os.Exit(-1)
	}
	title = scanner.Text()

	// get the description
	fmt.Println("Next, please enter a description: ")
	scanner = bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Println("Error: Scanning failed")
		os.Exit(-1)
	}
	description = scanner.Text()

	// get the wallet address
	fmt.Println("And finally, please enter your wallet address so funds can be deposited")
	scanner = bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Println("Error: Scanning failed")
		os.Exit(-1)
	}
	walletAddress = scanner.Text()

	//	Step 3.	Post the new block to the blockchain
	reqBody, err := json.Marshal(map[string]string{
		"data": "{\"title\":\"" + title + "\", \"description\": \"" + description + "\", \"wallet\": \"" + walletAddress + "\"}",
	})
	if err != nil {
		os.Exit(3)
	}

	// get the url that we should post to
	s := "http://localhost:3001/mineBlock"
	u, err := url.Parse(s)
	if err != nil {
		os.Exit(2)
	}
	resp, err := http.Post(u.String(),
		"application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		os.Exit(-1)
	}
	defer resp.Body.Close()

	//	Step 4.	Prompt user if the POST was succesful or not
	fmt.Printf("Success! Fund was posted to the blockchain")
	os.Exit(0)
}
