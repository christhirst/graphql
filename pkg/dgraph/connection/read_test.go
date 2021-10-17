package connection

import (
	"log"
	"testing"
)

func TestGetJson(t *testing.T) {
	var c DGraph
	err := c.NewClient()

	if err != nil {
		log.Fatal(err)
	}

	c.GetJson(c.client)
	//defer c.conn.Close()

}

/* tests := map[string]struct {
   input string
   sep   string
   want  []string
}{
   "simple":       {input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
   "wrong sep":    {input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
   "no sep":       {input: "abc", sep: "/", want: []string{"abc"}},
   "trailing sep": {input: "a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
}

for name, tc := range tests {
   t.Run(name, func(t *testing.T) {
       got := Split(tc.input, tc.sep)
       diff := cmp.Diff(tc.want, got)
       if diff != "" {
           t.Fatalf(diff)
       }
   })
} */
