/*
*
*	PROJECT			:	Don8 - Encrypted BlockChain Application
*	FILE			:	new-peer.go
*	PROGRAMER		:	Sky Roth
*	CREATED			:	April 10, 2021
*
*	DESCRIPTION		:
 */

/*
*
*	PROJECT			:	Don8 - Encrypted BlockChain Application
*	FILE			:	new-peer.go
*	PROGRAMER		:	Sky Roth
*	CREATED			:	April 10, 2021
*
*	DESCRIPTION		:	This simple file allows users to create a peer for the blockchain
 */

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type Data struct {
	data string
}

func main() {

	reqBody, err := json.Marshal(map[string]string{
		"peer": "ws://localhost:6001",
	})
	if err != nil {
		os.Exit(3)
	}

	s := "http://localhost:3001/addPeer"
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

	fmt.Println("Success! Peer added")
	os.Exit(0)
}
