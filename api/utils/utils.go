package utils

import (
	"log"
	"time"
)

func TimeExec(start time.Time, name string) {
	log.Fatalf("total %s: %s\n", name, time.Since(start))
}
