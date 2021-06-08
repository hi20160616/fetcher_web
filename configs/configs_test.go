package configs

import (
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	// RootPath = "../" // for test
	// err := get()
	// if err != nil {
	//         t.Error(err)
	// }
	if err := Reset("../"); err != nil {
		t.Error(err)
	}
	if err := load(); err != nil {
		t.Error(err)
	}
	fmt.Println("Configuration:")
	for k, v := range Data.MS {
		fmt.Printf("Key: %s, Value: %s", k, v)
	}
}
