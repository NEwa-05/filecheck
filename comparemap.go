package main

import "log"

var sameFilesList []string

func comparemap(fileshamap map[string]string) []string {

	tempmap := make(map[string]struct{})

	for k, v := range fileshamap {
		_, has := tempmap[v]
		if has {
			log.Println(k)
			sameFilesList = append(sameFilesList, k)
		}
		tempmap[v] = struct{}{}
	}

	// 2 maps (folders)
	//for k1, v1 := range fileshamap {
	//	for k2, v2 := range fileshamap {
	//		if v1 == v2 {
	//			samefiles[k1] = k2
	//		}
	//	}
	//}
	return sameFilesList
}
