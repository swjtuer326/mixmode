package sophon

import (
	"fmt"
	"strconv"
	"testing"
)

func TestGetTemperature(t *testing.T) {
	// GetTemperature()
	a := []byte{51, 57, 48, 48, 48, 10}
	fmt.Println(string(a[0 : len(a)-1]))
	if val, err := strconv.Atoi(string(a[0 : len(a)-1])); err == nil {
		fmt.Println(val / 1000)
	}
}
