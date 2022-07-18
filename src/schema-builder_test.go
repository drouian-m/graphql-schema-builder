package schemabuilder

import (
	"github.com/graphql-go/graphql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ShouldConvertSimpleStruct(t *testing.T) {
	type Simple struct {
		FirstName string `gql:"type=string"`
		LastName  string `gql:"type=string"`
	}

	gqlResult, _ := ConvertStructToGraphqlSchema(Simple{})
	expectedResult := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Simple",
			Fields: graphql.Fields{
				"FirstName": &graphql.Field{
					Name: "FirstName",
					Type: graphql.String,
				},
				"LastName": &graphql.Field{
					Name: "LastName",
					Type: graphql.String,
				},
			},
		},
	)
	assert.Equal(t, expectedResult, gqlResult)
}

func Test_ShouldHaveTagError(t *testing.T) {
	type Simple struct {
		FirstName string
		LastName  string `gql:"type=string,name=LastName"`
	}

	_, err := ConvertStructToGraphqlSchema(Simple{})
	assert.Error(t, err, "GenerationError - Cause : FirstName does not have a valid tag")
}

func Test_ShouldOverrideAttributeName(t *testing.T) {
	type Simple struct {
		FirstName string `gql:"type=string"`
		LastName  string `gql:"type=string,name=Name"`
	}

	gqlResult, _ := ConvertStructToGraphqlSchema(Simple{})
	expectedResult := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Simple",
			Fields: graphql.Fields{
				"FirstName": &graphql.Field{
					Name: "FirstName",
					Type: graphql.String,
				},
				"Name": &graphql.Field{
					Name: "Name",
					Type: graphql.String,
				},
			},
		},
	)
	assert.Equal(t, expectedResult, gqlResult)
}

func Test_ShouldConvertSimpleArray(t *testing.T) {
	type Simple struct {
		FirstName string   `gql:"type=string"`
		LastName  string   `gql:"type=string"`
		Tags      []string `gql:"type=string"`
	}

	gqlResult, _ := ConvertStructToGraphqlSchema(Simple{})
	expectedResult := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Simple",
			Fields: graphql.Fields{
				"FirstName": &graphql.Field{
					Name: "FirstName",
					Type: graphql.String,
				},
				"LastName": &graphql.Field{
					Name: "LastName",
					Type: graphql.String,
				},
				"Tags": &graphql.Field{
					Name: "Tags",
					Type: graphql.NewList(graphql.String),
				},
			},
		},
	)
	assert.Equal(t, expectedResult, gqlResult)
}

func Test_ShouldConvertStructArray(t *testing.T) {
	type Todo struct {
		Task string `gql:"type=string"`
		Done bool   `gql:"type=bool"`
	}
	type Simple struct {
		FirstName string   `gql:"type=string"`
		LastName  string   `gql:"type=string"`
		Tags      []string `gql:"type=string"`
		Todos     []Todo
	}

	gqlResult, _ := ConvertStructToGraphqlSchema(Simple{})
	expectedResult := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Simple",
			Fields: graphql.Fields{
				"FirstName": &graphql.Field{
					Name: "FirstName",
					Type: graphql.String,
				},
				"LastName": &graphql.Field{
					Name: "LastName",
					Type: graphql.String,
				},
				"Tags": &graphql.Field{
					Name: "Tags",
					Type: graphql.NewList(graphql.String),
				},
				"Todos": &graphql.Field{
					Name: "Todos",
					Type: graphql.NewList(
						graphql.NewObject(
							graphql.ObjectConfig{
								Name: "Todo",
								Fields: graphql.Fields{
									"Task": &graphql.Field{
										Name: "Task",
										Type: graphql.String,
									},
									"Done": &graphql.Field{
										Name: "Done",
										Type: graphql.Boolean,
									},
								},
							}),
					),
				},
			},
		},
	)
	assert.Equal(t, expectedResult, gqlResult)
}
