package config

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	RootPath = "../" // for test
	err := get()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Configuration:")
	fmt.Println(Data)
}
