package connection

import (
	"log"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"google.golang.org/grpc"
)

type DGraph struct {
	con *dgo.Dgraph
}

func (dg DGraph) NewClient() error {

	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	d, err := grpc.Dial("cold-meadow.grpc.eu-central-1.aws.cloud.dgraph.io:443", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	dg.con = dgo.NewDgraphClient(api.NewDgraphClient(d))
	return err
}
