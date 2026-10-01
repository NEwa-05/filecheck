package check

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"io"
	"log"
	"os"
)

func Createmapfilehash(dir []string) map[string]string {
	fileshamap := make(map[string]string)
	for _, element := range dir {
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
		fileshamap[base64.StdEncoding.EncodeToString([]byte(element))] = hn
	}
	for k, v := range fileshamap {
		log.Printf("Maphash: %s, %s", k, v)
	}

	return fileshamap
}
