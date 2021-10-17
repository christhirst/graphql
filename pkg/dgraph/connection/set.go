package connection

import (
	"github.com/dgraph-io/dgo/v210"
)

func (dg DGraph) SetData(c *dgo.Dgraph) {

	//_, errr := app.dg.NewTxn().QueryWithVars(ctx, q, variables)
	/*
		q := `query all($a: string) {
			all(func: eq(name, $a)) {
			  name
			}
		  }`
		variables := make(map[string]string)
		variables["$a"] = "Alice"
		ctx := context.Background() */
}
