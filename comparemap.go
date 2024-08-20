package main

func comparemap(fileshamap map[string]string) map[string]string {

	var samefiles map[string]string
	samefiles = make(map[string]string)

	for k1, v1 := range fileshamap {
		for k2, v2 := range fileshamap {
			if v1 == v2 {
				samefiles[k1] = k2
			}
		}
	}
	return samefiles
}
