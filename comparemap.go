package main

func comparemap() map[string]string {

	mapone := createmapfilehash()
	maptwo := createmapfilehash()

	var samefiles map[string]string
	samefiles = make(map[string]string)

	for k1, v1 := range mapone {
		for k2, v2 := range maptwo {
			if v1 == v2 {
				samefiles[k1] = k2
			}
		}
	}
	return samefiles
}
