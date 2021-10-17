package connection

import (
	"context"
	"encoding/json"
	"log"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
)

func (dg DGraph) Mutate(c *dgo.Dgraph) error {

	p := ""
	pb, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	_, err = c.NewTxn().Mutate(context.Background(), &api.Mutation{SetJson: pb})

	return err
}
