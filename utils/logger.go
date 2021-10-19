package utils

import (
	"fmt"
	"time"
)

func Log(message string) {
	fmt.Printf("[%s] : %s\n", time.Now().Format("01.02.2006 - 15:04:05"), message)
}
