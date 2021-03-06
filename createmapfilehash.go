package main

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"log"
	"os"
)

func createmapfilehash() map[string]string {
	var fileshamap map[string]string
	fileshamap = make(map[string]string)
	for _, element := range createfilelist() {
		f, err := os.Open(element)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		h := sha1.New()
		if _, err := io.Copy(h, f); err != nil {
			log.Fatal(err)
		}
		hn := hex.EncodeToString(h.Sum(nil))
		fileshamap[element] = hn
	}
	return fileshamap
}
