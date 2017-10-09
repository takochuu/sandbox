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
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: resolve,
	},
}

// MutationSchema
var mutaiotnFields graphql.Fields = graphql.Fields{
	"create": &graphql.Field{
		Type: graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
			Name: "mutation",
			Fields: graphql.Fields{
				"id":   &graphql.Field{Type: graphql.ID},
				"name": &graphql.Field{Type: graphql.String},
			},
		})),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: resolveMutation,
	},
}

var rootQuery graphql.ObjectConfig = graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
var rootMutation graphql.ObjectConfig = graphql.ObjectConfig{Name: "RootMutation", Fields: mutaiotnFields}
var schemaConfig graphql.SchemaConfig = graphql.SchemaConfig{
	Query:    graphql.NewObject(rootQuery),
	Mutation: graphql.NewObject(rootMutation),
}
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

func resolve(p graphql.ResolveParams) (interface{}, error) {
	return p.Args["id"], nil
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func resolveMutation(p graphql.ResolveParams) (interface{}, error) {
	return []User{
		{ID: p.Args["id"].(int), Name: "Fugafuga"},
		{ID: p.Args["id"].(int), Name: "Hogehoge"},
	}, nil
}
