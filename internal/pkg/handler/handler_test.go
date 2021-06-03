package handler

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/hi20160616/fetchnews/config"
)

func TestValidReq(t *testing.T) {
	// path prepare
	if err := config.Reset("../../../"); err != nil {
		t.Error(err)
	}
	// test section
	u, err := url.Parse("http://localhost:8080/list/bbc")
	if err != nil {
		t.Error(err)
	}

	r := &http.Request{URL: u}
	a := validReq(r)
	fmt.Println(a)
}
