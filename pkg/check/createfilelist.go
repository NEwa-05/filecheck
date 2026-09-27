package check

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func Createfilelist(folder string) []string {
	fileList := make([]string, 0)
	e := filepath.WalkDir(folder, func(path string, f fs.DirEntry, errors error) error {
		unusedir, err := os.Lstat(path)
		if unusedir.Mode().IsRegular() {
			fileList = append(fileList, path)
		}
		log.Print(err)
		return err
	})
	if e != nil {
		log.Print(e)
	}
	return fileList
}
