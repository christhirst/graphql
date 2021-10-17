package connection

import (
	"fmt"
	"log"
	"testing"
)

func TestNewClient(t *testing.T) {
	var dg DGraph
	err := dg.NewClient()

	if err != nil {
		log.Fatal(err)
	}

	if dg.client == nil && dg.conn == nil {
		t.Errorf("Getting Account failed")
		fmt.Println("##")
		fmt.Println(dg.conn)
	}

}
