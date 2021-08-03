# Don8: Blockchain Donation Website
(get it? "donate")

## General info
Don8 is a simple blockchain based off of Naivechain [(Github)](https://github.com/lhartikk/naivechain)
This website allows users to add fundraisers through a blockchain, or donate to a specific cryptocurrency address.
These fundraisers will include a cryptocurrency wallet address so that everything stays anonymous.

## Technologies
Project is created with:
* Typescript 2.4.1
* Golang 1.11.7
* JavaScript ES5
* jQuery 3.5.1
* HTML/CSS
	
## Setup
This project includes a blockchain server as well as a web server that hosts the website.
Both need to be running in order for the project to work correctly.

```
$ cd /Blockchain          -> Change your directory to the blockchain server
$ npm start               -> Start the blockchain server
$ cd ../UI/webserver      -> Change directory to the web server
$ go run webserver.go     -> Run the web server for the website
$ index.html              -> Open the main webpage
```

## Client Setup
These are simple instructions to run the client programs

**client.go**
This program allows users to add to the blockchain without access to the website
```
$ go run client.go     -> Run the program, follow the prompts
```

**new-peer.go**
This program allows users to add a peer to the blockchain
```
$ go run new-peer.go     -> Run the program, follow the prompts
```
