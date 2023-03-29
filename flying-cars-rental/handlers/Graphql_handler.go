package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
)

type GraphQLRequest struct {
	Query string `json:"query"`
}

type GraphQLResponse struct {
	Result interface{} `json:"result"`
}

func graphqlHandler(schema graphql.Schema) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "GraphQL endpoint only accepts POST requests", http.StatusMethodNotAllowed)
			return
		}

		var req GraphQLRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid GraphQL request", http.StatusBadRequest)
			return
		}
        getUser, _ := getUserFromContext(r.Context())
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: req.Query,
			Context:       context.WithValue(r.Context(), "user", ),
		})

		res := GraphQLResponse{
			Result: result,
		}

		json.NewEncoder(w).Encode(res)
	}
}