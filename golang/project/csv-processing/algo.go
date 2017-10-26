package main


func algo(input []float32) []chain { // TODO: need a better name
	log.Debugln("input:",input)

	dedup := dedupFloat32s(input)
	log.Debugln("after dedup:",dedup)

	sortFloat32(dedup)
	log.Debugln("after sort:",dedup)
	return prepareChains(dedup)

}
