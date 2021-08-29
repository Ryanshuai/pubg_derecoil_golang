package src

import (
	"strconv"
	"strings"
)

type Gun struct {
	name         string
	fireInterval float32
	pushSequence []float32
	fireTime     float32
}

func (gun Gun) fire() {
	fireLength := gun.fireTime / 0.015
	print(fireLength)

	press := src.Press{}
	press.distance = gun.pushSequence
	go press.Act()
}

func getTime(name string) float32 {
	state := strings.Split(name, "_")
	time, _ := strconv.ParseFloat(state[1], 32)
	return float32(time)
}

func getFactor(scope string, muzzle string, grip string) float32 {
	scopeFactor := getTime(scope)
	muzzleFactor := getTime(muzzle)
	gripFactor := getTime(grip)
	return scopeFactor * muzzleFactor * gripFactor
}

func (gun Gun) calculateDistance(fireSequence []float32, factor float32, fireLength int) []int32 {
	gun.pushSequence = fireSequence
	for i := 0; i < len(gun.pushSequence); i++ {
		gun.pushSequence[i] *= gun.pushSequence[i] * factor
	}
	gun.pushSequence = gun.pushSequence[:fireLength]

	deltas := []int32{}
	for i := 0; i < len(gun.pushSequence); i++ {
		deltaNum := gun.fireInterval / 0.15
		deltaDist := int32(gun.pushSequence[i] / deltaNum)

		for i := 0; i < int(deltaNum); i++ {
			deltas = append(deltas, deltaDist)
		}

		deltaNumInt := int32(deltaNum)
		remainDist := int32(gun.pushSequence[i]) - deltaDist*deltaNumInt
		deltas = append(deltas, remainDist)
	}
	return deltas
}
