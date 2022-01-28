package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
)

const(
	key = "abcdefghijklmnopqrstuvwxyz012345"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> Enter your password: ")
		in, _ := reader.ReadString('\n')
		out, err := encrypt(in, []byte(key))
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("> Your encrypted password is %s\n", out)
	}
}

// encrypt will encrypt plaintext using AES-GCM
// - Requires AES-128 or AES-256 keys
// - Uses nil padding nonce for convenience
func encrypt(data string, key []byte) (string, error) {
	plaintext := []byte(data)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, 12)
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)
	return string(ciphertext), nil
}
