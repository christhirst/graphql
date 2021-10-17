package connection

import (
	"log"
	"testing"
)

func TestNewClient(t *testing.T) {
	var dg DGraph
	err := dg.NewClient()
	if err != nil {
		log.Fatal(err)
	}

}
