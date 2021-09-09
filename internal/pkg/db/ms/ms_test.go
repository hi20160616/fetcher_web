package ms

import (
	"fmt"
	"testing"
)

func TestOpen(t *testing.T) {
	if err := Open(); err != nil {
		t.Error(err)
	}
	fmt.Println(Conns["bbc"])
}

func TestList(t *testing.T) {
	if err := Open(); err != nil {
		t.Error(err)
	}
	if err := List("bbc"); err != nil {
		t.Error(err)
	}
}
