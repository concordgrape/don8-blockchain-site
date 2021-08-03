/*
*
*	PROJECT			:	Doon8 - Encrypted BlockChain Application
*	FILE			:	webserver.go
*	PROGRAMER		:	Sky Roth
*	DESCRIPTION		:	Create the web server for the Doon8 site, this will also generate the public/private key
*
*	CREATED			:	April 9, 2021
*	LAST UPDATED	:	April 11, 2021
*
 */

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

// What each block in the blockchain will hold
type Block struct {
	Index        int
	PreviousHash string
	Timestamp    int
	Data         string
	Hash         string
	Difficult    int
	Nonce        int
}

// Data within the data block
type Data struct {
	Title        string
	Description  string
	CryptoWallet string
}

func main() {

	//	Step 1. Which file to send when the user enters a specific URL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// get the blocks from the block chain when the user goes to the site
		if r.Header.Get("Referer") == "http://localhost:8080/" {
			getBlocks()
		}

		// send the file the user was requesting from the URL
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	//	Step 2.	If the user "Adds a Fund", handle the POST request
	http.HandleFunc("/postDat", PostHandler)
	http.HandleFunc("/postDat/", PostHandler)

	//	Step 3.	Listen on port 8080
	http.ListenAndServe(":8080", nil)
}

/*
*	FUNCTION	:	PostHandler()
*	DESCRIPTION	:	How to handle the user's post request after "adding a fund"
*	PARAMETERS	:
*		w http.ResponseWriter	:	What to respond to the client
*		r *http.Request			:	The contents of what was requested (parameters)
*
*	RETURN		:	NIL
 */
func PostHandler(w http.ResponseWriter, r *http.Request) {

	//	Step 1.	Parse the form that was posted
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	//	Step 2.	Get specific form values
	title := r.PostFormValue("title")
	description := r.PostFormValue("description")
	wallet := r.PostFormValue("wallet")

	//	Step 3.	Verify the values
	if len(title) == 0 || len(description) == 0 || len(wallet) == 0 {
		fmt.Fprintf(w, `Failed to post, no data was supplied`)
		return
	}

	//	Step 4.	Add a new block to the blockchain
	reqBody, err := json.Marshal(map[string]string{
		"data": "{\"title\":\"" + title + "\", \"description\": \"" + description + "\", \"wallet\": \"" + wallet + "\"}",
	})
	if err != nil {
		os.Exit(3)
	}

	//	Step 5.	Define the URL to post the new block
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

	//	Step 6.	Send a response to the clients browser
	fmt.Fprintf(w, `Succesfully posted to the blockchain!`)
}

/*
*	FUNCTION	:	getBlocks()
*	DESCRIPTION	:	Retrieve the blocks in the blockchain
*	PARAMETERS	:	NIL
*	RETURN		:	NIL
 */
func getBlocks() {

	//	Step 1.	Get the URL where the blocks are located
	s := "http://localhost:3001/blocks"
	u, err := url.Parse(s)
	if err != nil {
		os.Exit(-1)
	}

	//	Step 2.	Get the response from the GET request
	resp, err := http.Get(u.String())
	if err != nil {
		os.Exit(-1)
	}

	//	Step 3.	Read all contents of the GET request
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		os.Exit(-1)
	}

	//	Step 4.	Convert body to a string
	sb := string(body)

	//	Step 5. Convert the string into a Block struct
	bodyByte := []byte(sb)
	bodies := make([]Block, 0)
	json.Unmarshal(bodyByte, &bodies)

	//	Step 6.	Create a new text file that will hold the blocks from the blockchain
	f, err := os.Create("funds.txt")
	if err != nil {
		os.Exit(-1)
	}
	f.WriteString(sb)
}
