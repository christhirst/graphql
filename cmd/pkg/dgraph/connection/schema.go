package connection

import (
	"context"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
)

func Setup(c *dgo.Dgraph) error {
	// Install a schema into dgraph. Accounts have a `name` and a `balance`.
	err := c.Alter(context.Background(), &api.Operation{
		Schema: `
			name: string @index(term) .
			balance: int .
		`,
	})
	return err
}
