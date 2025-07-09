package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// ---- Sample Data (to act like a database) ----
var sampleSpaces = []map[string]interface{}{
	{
		"id":      "balancer.eth",
		"name":    "Balancer",
		"about":   "A decentralized exchange protocol.",
		"network": "1",
		"symbol":  "BAL",
		"admins":  []string{"0xAdmin1"},
		"members": []string{"0xMember1"},
	},
	{
		"id":      "yam.eth",
		"name":    "Yam Finance",
		"about":   "A decentralized farming protocol.",
		"network": "1",
		"symbol":  "YAM",
		"admins":  []string{"0xAdmin2"},
		"members": []string{"0xMember2"},
	},
}

var sampleProposals = []map[string]interface{}{
	{
		"id":     "proposal1",
		"title":  "New Feature for Balancer",
		"body":   "Details about the new feature.",
		"state":  "closed",
		"author": "0xAuthor1",
		"space":  sampleSpaces[0], // Link to Balancer space
	},
}

// ---- Define GraphQL Object Types ----

var strategyType *graphql.Object
var filtersType *graphql.Object
var spaceType *graphql.Object
var proposalType *graphql.Object
var voteType *graphql.Object
var followType *graphql.Object

func init() {
	// Nested Type: Strategy (inside Space)
	strategyType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Strategy",
		Fields: graphql.Fields{
			"name":    &graphql.Field{Type: graphql.String},
			"network": &graphql.Field{Type: graphql.String},
			"params":  &graphql.Field{Type: graphql.String}, // Simplified as a string
		},
	})

	// Nested Type: Filters (inside Space)
	filtersType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Filters",
		Fields: graphql.Fields{
			"minScore":    &graphql.Field{Type: graphql.Float},
			"onlyMembers": &graphql.Field{Type: graphql.Boolean},
		},
	})

	// Main Type: Space
	spaceType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Space",
		Fields: graphql.Fields{
			"id":         &graphql.Field{Type: graphql.String},
			"name":       &graphql.Field{Type: graphql.String},
			"about":      &graphql.Field{Type: graphql.String},
			"network":    &graphql.Field{Type: graphql.String},
			"symbol":     &graphql.Field{Type: graphql.String},
			"strategies": &graphql.Field{Type: graphql.NewList(strategyType)},
			"admins":     &graphql.Field{Type: graphql.NewList(graphql.String)},
			"moderators": &graphql.Field{Type: graphql.NewList(graphql.String)},
			"members":    &graphql.Field{Type: graphql.NewList(graphql.String)},
			"filters":    &graphql.Field{Type: filtersType},
			"plugins":    &graphql.Field{Type: graphql.String}, // Simplified
		},
	})

	// Main Type: Proposal
	proposalType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Proposal",
		Fields: graphql.Fields{
			"id":       &graphql.Field{Type: graphql.String},
			"title":    &graphql.Field{Type: graphql.String},
			"body":     &graphql.Field{Type: graphql.String},
			"choices":  &graphql.Field{Type: graphql.NewList(graphql.String)},
			"start":    &graphql.Field{Type: graphql.Int},
			"end":      &graphql.Field{Type: graphql.Int},
			"snapshot": &graphql.Field{Type: graphql.String},
			"state":    &graphql.Field{Type: graphql.String},
			"author":   &graphql.Field{Type: graphql.String},
			"space":    &graphql.Field{Type: spaceType},
		},
	})

	// You would define voteType and followType similarly...
}

func main() {
	// Define the arguments for the 'proposals' query
	proposalsWhereInput := graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "ProposalsWhereInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"space_in": &graphql.InputObjectFieldConfig{Type: graphql.NewList(graphql.String)},
			"state":    &graphql.InputObjectFieldConfig{Type: graphql.String},
		},
	})

	// This is the Root Query - the entry point for all READ operations
	fields := graphql.Fields{
		// 1. "spaces" query field
		"spaces": &graphql.Field{
			Type:        graphql.NewList(spaceType),
			Description: "Get a list of spaces.",
			Args: graphql.FieldConfigArgument{
				"first":          &graphql.ArgumentConfig{Type: graphql.Int},
				"skip":           &graphql.ArgumentConfig{Type: graphql.Int},
				"orderBy":        &graphql.ArgumentConfig{Type: graphql.String},
				"orderDirection": &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// In a real app, you would use the args to query a database.
				// Here, we just return all our sample spaces.
				return sampleSpaces, nil
			},
		},
		// 2. "proposals" query field
		"proposals": &graphql.Field{
			Type:        graphql.NewList(proposalType),
			Description: "Get a list of proposals.",
			Args: graphql.FieldConfigArgument{
				"first":          &graphql.ArgumentConfig{Type: graphql.Int},
				"skip":           &graphql.ArgumentConfig{Type: graphql.Int},
				"orderBy":        &graphql.ArgumentConfig{Type: graphql.String},
				"orderDirection": &graphql.ArgumentConfig{Type: graphql.String},
				"where":          &graphql.ArgumentConfig{Type: proposalsWhereInput},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// We can inspect the "where" argument here
				if where, ok := p.Args["where"].(map[string]interface{}); ok {
					if state, ok := where["state"].(string); ok {
						fmt.Printf("Filtering proposals by state: %s\n", state)
					}
				}
				// Return sample data
				return sampleProposals, nil
			},
		},
		// 3. You would define "votes" and "follows" queries here similarly...
	}

	// ---- The rest is the same setup code ----
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	fmt.Println("Server is running at http://localhost:8080/graphql")
	http.ListenAndServe(":8080", nil)
}
