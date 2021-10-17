package connection

import (
	"log"
	"testing"

	"google.golang.org/grpc"
)

func TestNewClient(t *testing.T) {
	_, err := grpc.Dial("cold-meadow.grpc.eu-central-1.aws.cloud.dgraph.io:443", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

}
