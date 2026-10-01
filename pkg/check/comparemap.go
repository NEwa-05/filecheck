package check

import (
	"encoding/base64"
	"log"
)

var sameFilesList []string

func OneFolderDup(srcFileshamap map[string]string) []string {

	tempmap := make(map[string]struct{})

	for k, v := range srcFileshamap {
		_, has := tempmap[v]
		if has {
			b64k, err := base64.StdEncoding.DecodeString(k)
			if err != nil {
				log.Print("can't decode b64")
			}
			sameFilesList = append(sameFilesList, string(b64k))
		}
		tempmap[v] = struct{}{}

	}

	return sameFilesList
}

func TwoFolderDup(srcFileshamap map[string]string, srcdupFileshamap map[string]string) []string {

	tempmap := make(map[string]struct{})
	var tempSameFilesList []string

	for _, v1 := range srcFileshamap {
		for k2, v2 := range srcdupFileshamap {
			if v1 == v2 {
				b64k, err := base64.StdEncoding.DecodeString(k2)
				if err != nil {
					log.Print("can't decode b64")
				}
				tempSameFilesList = append(tempSameFilesList, string(b64k))
			}
		}
	}

	for _, v := range tempSameFilesList {
		_, has := tempmap[v]
		if !has {
			sameFilesList = append(sameFilesList, string(v))
		}
		tempmap[v] = struct{}{}
	}

	return sameFilesList
}
