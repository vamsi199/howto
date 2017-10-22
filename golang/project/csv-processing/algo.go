package main

func algo(input []float32) []chain { // TODO: need a better name

	dedup := dedupFloat32s(input)

	sortFloat32(dedup)

	return prepareChains(dedup)

}
