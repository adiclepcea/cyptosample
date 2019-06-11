package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/crypto/twofish"
)

func main() {

	key := flag.String("k", "averylongandsecurekey12345678key", "the key to use for your crypto")
	msg := flag.String("m", "", "the message to encrypt with your algorithm")
	decryptMsg := flag.Bool("d", false, "decrypt message")
	protocol := flag.String("p", "aes", "crypto protocol to use: aes or twofish")

	flag.Parse()

	if *msg == "" {
		flag.PrintDefaults()
		fmt.Println("You can use https://codebeautify.org/encrypt-decrypt for testing")
		os.Exit(0)
	}

	//Decrypt
	if *decryptMsg {
		decrypt(protocol, msg, key)
	}

	//Encrypt
	encrypt(protocol, msg, key)
}

func decrypt(protocol, msg, key *string) {
	if *protocol == "twofish" {
		decrypted, err := decryptTwofish(*msg, *key)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(decrypted)
		return
	} else if *protocol == "aes" {
		decrypted, err := decryptAes(*msg, *key)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(decrypted)
		return
	} else {
		log.Fatal("Only aes and twofish are supported")
	}
}

func encrypt(protocol, msg, key *string) {
	if *protocol == "twofish" {
		iv := make([]byte, twofish.BlockSize)
		if _, err := io.ReadFull(rand.Reader, iv); err != nil {
			log.Fatalf("Could not create the iv")
		}

		crypted, err := encryptTwofish(*msg, *key, iv)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(crypted)
	} else if *protocol == "aes" {
		iv := make([]byte, aes.BlockSize)
		if _, err := io.ReadFull(rand.Reader, iv); err != nil {
			log.Fatalf("Could not create the iv")
		}

		crypted, err := encryptAes(*msg, *key, iv)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(crypted)
	} else {
		log.Fatal("Only aes and twofish are supported")
	}
}

func encryptTwofish(msg, key string, iv []byte) (string, error) {
	b := PKCS5Padding([]byte(msg), aes.BlockSize, len(msg))
	msg = string(b)
	block, err := twofish.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, twofish.BlockSize+len(msg))
	copy(ciphertext[:twofish.BlockSize], iv)
	cbc := cipher.NewCBCEncrypter(block, iv)
	cbc.CryptBlocks(ciphertext[twofish.BlockSize:], []byte(msg))
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func decryptTwofish(crypted, key string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(crypted)
	if err != nil {
		return "", err
	}

	block, err := twofish.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	iv := ciphertext[:twofish.BlockSize]
	decrypted := ciphertext[twofish.BlockSize:]
	cbc := cipher.NewCBCDecrypter(block, iv)
	cbc.CryptBlocks(decrypted, decrypted)
	s := string(PKCS5UnPadding(decrypted[:]))
	return s, nil
}

func encryptAes(msg, key string, iv []byte) (string, error) {
	b := PKCS5Padding([]byte(msg), aes.BlockSize, len(msg))
	msg = string(b)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(msg))
	copy(ciphertext[:aes.BlockSize], iv)
	cbc := cipher.NewCBCEncrypter(block, iv)
	cbc.CryptBlocks(ciphertext[aes.BlockSize:], []byte(msg))
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func decryptAes(crypted, key string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(crypted)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	iv := ciphertext[:aes.BlockSize]
	decrypted := ciphertext[aes.BlockSize:]
	cbc := cipher.NewCBCDecrypter(block, iv)
	cbc.CryptBlocks(decrypted, decrypted)
	s := string(PKCS5UnPadding(decrypted[:]))
	return s, nil
}

// PKCS5Padding fills the text until it reaches the required multiple
func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS5UnPadding removes the padding
func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
