package main

var sameFilesList []string

func oneFolderDup(fileshamap map[string]string) []string {

	tempmap := make(map[string]struct{})

	for k, v := range fileshamap {
		_, has := tempmap[v]
		if has {
			sameFilesList = append(sameFilesList, k)
		}
		tempmap[v] = struct{}{}
	}
	return sameFilesList
}

func twoFolderDup(srcFileshamap map[string]string, dupsrcFileshamap map[string]string) []string {

	for _, v1 := range srcFileshamap {
		for k2, v2 := range dupsrcFileshamap {
			if v1 == v2 {
				sameFilesList = append(sameFilesList, k2)
			}
		}
	}
	return sameFilesList
}
