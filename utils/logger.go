package utils

import (
	"fmt"
	"time"
)

func Log(message string) {
	fmt.Printf("[%s] : %s\n", time.Now().Format("02.01.2006 - 15:04:05"), message)
}
