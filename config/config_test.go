package config

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	err := get()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Configuration:")
	fmt.Println(Data)
}
