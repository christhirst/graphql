package connection

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo/v210"
)

const q = `
{
	michael(func: anyofterms(name, "Michael")) {
	  name
	}
  }
 
`

var result map[string]interface{}

var decode struct {
	Michael []struct {
		Name  string `json:"name"`
		Name2 string `json:"name2"`
	} `json:"michael"`
}

func (dg DGraph) GetJson(c *dgo.Dgraph) map[string]interface{} {

/* 	qq := `query all($a: string) {
		all(func: eq(name, $a)) {
		  name
		}
	  }` */

	g, _ := c.NewReadOnlyTxn().Query(context.Background(), q)
	json.Unmarshal(g.GetJson(), &result)

	jsonString, _ := json.Marshal(result)
	fmt.Println("xx")

	fmt.Println(string(jsonString))

	if err := json.Unmarshal(g.GetJson(), &decode); err != nil {
		log.Fatal(err)
	}

	fmt.Println("sss")
	fmt.Println(g)
	fmt.Println(decode)
	fmt.Println("sss")

	return result

}
