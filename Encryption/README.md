# Don8: Blockchain


## General info
IMPORTANT: This encryption method is not included in the website, but is capable of becoming a part of the blockchain/webserver

This simple encryption method uses RSA keys to encrypt and decrypt messages. The idea of this protocol was to allow the website to display
a public key in which people can use to encrypt messages before adding to the blockchain, this keeps all blocks private until it reaches the website.

However, it is implemented that the website would change this public key after every reboot or shutdown. This would lead the past blocks to be tampered with (i.e. unvalid)


## Technologies
Project is created with:
* Golang 1.11.7