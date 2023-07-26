package test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	for time := range ticker.C {
		fmt.Println(time)
	}

	time.Tick(1 * time.Second)

	runtime.CPUProfile()
}
