# Symetric crytptography

This shows how to perform the twofish and the AES encryption and decryption
The language of choice is Golang so you should have it installed


## Build

To build the executable you should either do a 

```go get golang.org/x/crypto/twofish```

to make sure you have this library on your GOPATH or just use the go mod

```go mod tidy```

if you have the modules enabled in GO 1.11 or higher

## Usage

```
./symetric.exe
  -d    decrypt message
  -k string
        the key to use for twofish (default "averylongandsecurekey12345678key")
  -m string
        the message to encrypt with twofish
  -p string
        crypto protocol to use: aes or twofish (default "aes")
You can use https://codebeautify.org/encrypt-decrypt for testing
```
