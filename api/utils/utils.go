package utils

import (
	"fmt"
	"time"
)

func TimeExec(start time.Time, name string) {
	fmt.Printf("total %s: %s\n", name, time.Since(start))
}
