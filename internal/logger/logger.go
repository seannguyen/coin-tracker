package logger

import (
	"fmt"
	"time"
)

func Info(message string) {
	fmt.Printf("[%s] INFO %s\n", time.Now().Format(time.RFC3339), message)
}

func Err(message string) {
	fmt.Printf("[%s] ERRO %s\n", time.Now().Format(time.RFC3339), message)
}