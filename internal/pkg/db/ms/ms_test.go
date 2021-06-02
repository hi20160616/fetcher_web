package ms

import (
	"fmt"
	"testing"

	"github.com/hi20160616/fetchnews/config"
)

func TestOpen(t *testing.T) {
	config.Reset("../../../../")
	if err := Open(); err != nil {
		t.Error(err)
	}
	fmt.Println(Conns["bbc"])
}

func TestList(t *testing.T) {
	config.Reset("../../../../")
	if err := Open(); err != nil {
		t.Error(err)
	}
	if err := List("bbc"); err != nil {
		t.Error(err)
	}
}
