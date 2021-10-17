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

	qq := `query Alice($a: string){
		me(func: eq(name, $a)) {
			name
		}
	}`
	variables := make(map[string]string)
	variables["$a"] = "Alice"
	fmt.Println(c)

	resp, err := c.NewReadOnlyTxn().QueryWithVars(context.Background(), qq, variables)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(resp.Json))
	fmt.Println(resp.Size())
	json.Unmarshal(resp.GetJson(), &result)

	jsonString, _ := json.Marshal(result)
	fmt.Println("xx")

	fmt.Println(string(jsonString))

	if err := json.Unmarshal(resp.GetJson(), &decode); err != nil {
		log.Fatal(err)
	}

	fmt.Println("sss")
	fmt.Println(resp)
	fmt.Println(decode)
	fmt.Println("sss")

	return result

}
