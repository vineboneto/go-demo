package utils

import (
	"fmt"
	"time"
)

func TimeExec(start time.Time) {
	fmt.Println("total: ", time.Since(start))
}
