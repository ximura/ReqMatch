package main

import (
	"github.com/pkg/profile"
	"github.com/ximura/ReqMatch/internal/sorted"
)

type statInfo struct {
	id    int
	delay int
}

func main() {
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1), profile.ProfilePath(".")).Stop()

	stats := sorted.NewArray(10, func(a, b statInfo) int {
		if a.delay == b.delay {
			return 0
		}
		if a.delay > b.delay {
			return 1
		}

		return -1
	}, nil)

	for i := 0; i < 1_000_000; i++ {
		stats.Insert(statInfo{
			id:    i,
			delay: i,
		})
	}
}
