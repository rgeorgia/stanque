package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

type FilesStats struct {
	name     string
	checksum int64
}

var filesToCheck = map[string]string{
	"file_checker":  "f62f0949ffa5a041e31b2fc706d93708ce93b62b7099f516380f8d39d15380c6",
	"text_file.txt": "539181b57717e6f3430d599ca3b11fe18c9d405c8cfe3ff163037694e4864b1a",
	"pid_checker":   "b97ab6fabafba27199d50a190a2ad6513ccf8ee722558e86d2a45fd2ac535c67"}

func main() {

	for fileName, fileHash := range filesToCheck {
		if fileExists(fileName) {
			if fileHash == getCheckSum(fileName) {
				fmt.Printf("SHA-256 for %s: %s\n", fileName, getCheckSum(fileName))
			} else {
				fmt.Printf("Hash miss match %s\n expected: %s\n recivied %s\n",
					fileName, fileHash, getCheckSum(fileName))
			}

		} else {
			fmt.Printf("%s does not exist\n", fileName)
		}
	}

}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

/* Need to include error handling */
func getCheckSum(filename string) string {
	hash := sha256.New()
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	defer f.Close()
	_, err = io.Copy(hash, f)
	if err != nil {
		fmt.Println("Error:", err)
		log.Fatal("Error: ", err)
	}

	hashSum := hash.Sum(nil)
	return fmt.Sprintf("%x", hashSum)
}
