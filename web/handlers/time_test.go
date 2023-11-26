package handlers

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	add := now.Add(time.Hour * 24 * 7)

	fmt.Println(add.Format("2006-01-02 15:04:05"))
}
