package configs

import (
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	fmt.Println("Configuration:")
	fmt.Println(Data.WebServer)
	fmt.Println(Data.API)
	// for k, v := range Data.MS {
	//         fmt.Printf("Key: %s, Value: %s", k, v)
	// }
}
