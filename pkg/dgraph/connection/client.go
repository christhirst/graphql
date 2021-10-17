package connection

import (
	"fmt"
	"log"
	"os"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"google.golang.org/grpc"
)

type DGraph struct {
	client *dgo.Dgraph
	conn   *grpc.ClientConn
}

func (dg *DGraph) NewClient() error {

	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.

	conn, err := dgo.DialCloud("https://cold-meadow.eu-central-1.aws.cloud.dgraph.io/graphql", os.Getenv("dgraph_admin"))

	if err != nil {
		log.Fatal(err)
	}

	dg.conn = conn
	fmt.Println(conn)
	dg.client = dgo.NewDgraphClient(api.NewDgraphClient(conn))
	return err
}
