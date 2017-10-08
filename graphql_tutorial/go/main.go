package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
)

// Schema
var fields graphql.Fields = graphql.Fields{
	"hello": &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "world", nil
		},
	},
}

var rootQuery graphql.ObjectConfig = graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
var schemaConfig graphql.SchemaConfig = graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
var schema, _ = graphql.NewSchema(schemaConfig)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	r := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(r.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", r.Errors)
	}

	j, _ := json.Marshal(r)
	fmt.Printf("%s \n", j)

	return r
}

func handler(w http.ResponseWriter, r *http.Request) {
	bufBody := new(bytes.Buffer)
	bufBody.ReadFrom(r.Body)
	query := bufBody.String()

	result := executeQuery(query, schema)
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
