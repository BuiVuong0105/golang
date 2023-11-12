package pattern

import (
	"fmt"
)

type SaftyPlacer interface {
	placeSafeties()
}

type ICESafetyPlacer struct {
}

func (sp *ICESafetyPlacer) placeSafeties() {
	fmt.Println("ICE Placing my safeties....")
}

type NOPafetyPlacer struct {
}

func (sp *NOPafetyPlacer) placeSafeties() {
	fmt.Println("NOP Placing my safeties....")
}

type RockClimber struct {
	rocksClimber int
	saftyPlacer  SaftyPlacer
}

func CreateRockClimber(sp SaftyPlacer) *RockClimber {
	return &RockClimber{
		saftyPlacer: sp,
	}
}

func (rc *RockClimber) ClimbRock() {
	rc.rocksClimber++
	if rc.rocksClimber == 10 {
		rc.saftyPlacer.placeSafeties()
	}
}
