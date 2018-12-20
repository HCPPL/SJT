package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "fmt"
    "io"
    "log"
    "encoding/json"
)

//var secretKey []byte = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpX")

/*
func main() {
    originalText := "Hello World"
    fmt.Println(originalText)

    //secretKey := []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpX")

    // encrypt value to base64

    cryptoText := encrypt(originalText)
    fmt.Println(cryptoText)

    // encrypt base64 crypto to original value
    text := decrypt( cryptoText)
    fmt.Printf(text)
}
*/
// encrypt string to base64 crypto using AES
func encrypt( text string) string {
	CallingLog("encrypt()","encrypt-code.go", "NULL", "NULL",text)
    
    // secretKey := []byte(keyText)

    plaintext := []byte(text)
    block, err := aes.NewCipher(secretKey)
    if err != nil {
        panic(err)
    }

    // The IV needs to be unique, but not secure. Therefore it's common to
    // include it at the beginning of the ciphertext.

    ciphertext := make([]byte, aes.BlockSize+len(plaintext))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        panic(err)
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
    b, err := json.Marshal(ciphertext)
	
		if err != nil {
            throwException("encrypt()", "encrypt-code.go", "NULL", "NULL","Fatal Error")
			log.Println("Error: %s", err)
	
			
			
	
		}
	
		var s5 string= string(b)
    // convert to base64
	ResultLog ("encrypt()","encrypt-code.go", s5,text)
    
    return base64.URLEncoding.EncodeToString(ciphertext)
}

// decrypt from base64 to decrypted string
func decrypt( cryptoText string) string {
    CallingLog("decrypt()","encrypt-code.go", "NULL", "NULL",cryptoText)
    
    ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)
    block, err := aes.NewCipher(secretKey)
    if err != nil {
        panic(err)
    }

// The IV needs to be unique, but not secure. Therefore it's common to
// include it at the beginning of the ciphertext.
    if len(ciphertext) < aes.BlockSize {
        panic("ciphertext too short")
    }

    iv := ciphertext[:aes.BlockSize]
    ciphertext = ciphertext[aes.BlockSize:]
    stream := cipher.NewCFBDecrypter(block, iv)

    // XORKeyStream can work in-place if the two arguments are the same.
    stream.XORKeyStream(ciphertext, ciphertext)
    b, err := json.Marshal(ciphertext)
	
		if err != nil {
            throwException("decrypt()", "encrypt-code.go", "NULL", "NULL","Fatal Error")
			log.Println("Error: %s", err)
	
		
			
	
		}
	
		var s5 string= string(b)
	ResultLog ("decrypt()","encrypt-code.go", s5,cryptoText)
    
    return fmt.Sprintf("%s", ciphertext)
}