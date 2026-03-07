package annalyn

func CanFastAttack(knightIsAwake bool) bool {
	if knightIsAwake {
		return false
	}
	return true
}

func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	if knightIsAwake || archerIsAwake || prisonerIsAwake {
		return true
	}
	return false
}

func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	if prisonerIsAwake && !archerIsAwake {
		return true
	}
	return false
}

func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	if petDogIsPresent && !archerIsAwake {
		return true
	}
	if prisonerIsAwake && !petDogIsPresent && !knightIsAwake && !archerIsAwake {
		return true
	}
	return false
}
