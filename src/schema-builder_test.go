package schemabuilder

import (
	"github.com/graphql-go/graphql"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_ShouldConvertSimpleStruct(t *testing.T) {
	type Simple struct {
		FirstName string
		LastName  string
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

func Test_ShouldOverrideAttributeName(t *testing.T) {
	type Simple struct {
		FirstName string
		LastName  string `gql:"name=Name"`
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
		FirstName string
		LastName  string
		Tags      []string
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
		Task string
		Done bool
	}
	type Simple struct {
		FirstName string
		LastName  string
		Tags      []string
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

func Test_ShouldOverrideSubArrayAttributeName(t *testing.T) {
	type Todo struct {
		Task string `gql:"name=Action"`
		Done bool
	}
	type Simple struct {
		FirstName string
		LastName  string
		Tags      []string
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
									"Action": &graphql.Field{
										Name: "Action",
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

func Test_ShouldConvertTimeAttribute(t *testing.T) {
	type Simple struct {
		FirstName string
		LastName  string
		CreatedAt time.Time
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
				"CreatedAt": &graphql.Field{
					Name: "CreatedAt",
					Type: graphql.DateTime,
				},
			},
		},
	)
	assert.Equal(t, expectedResult, gqlResult)
}

func Test_ShouldConvertTimeInStructArray(t *testing.T) {
	type Todo struct {
		Task string
		Done bool
		End  time.Time
	}
	type Simple struct {
		FirstName string
		LastName  string
		Tags      []string
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
									"End": &graphql.Field{
										Name: "End",
										Type: graphql.DateTime,
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

func Test_ShouldConvertFloatAttribute(t *testing.T) {
	type Simple struct {
		Num  float32
		Num2 float64
	}

	gqlResult, _ := ConvertStructToGraphqlSchema(Simple{})
	expectedResult := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Simple",
			Fields: graphql.Fields{
				"Num": &graphql.Field{
					Name: "Num",
					Type: graphql.Float,
				},
				"Num2": &graphql.Field{
					Name: "Num2",
					Type: graphql.Float,
				},
			},
		},
	)
	assert.Equal(t, expectedResult, gqlResult)
}

func Test_ShouldConvertIntAttribute(t *testing.T) {
	type Simple struct {
		Num  int
		Num2 int8
		Num3 int16
		Num4 int32
		Num5 int64
	}

	gqlResult, _ := ConvertStructToGraphqlSchema(Simple{})
	expectedResult := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Simple",
			Fields: graphql.Fields{
				"Num": &graphql.Field{
					Name: "Num",
					Type: graphql.Int,
				},
				"Num2": &graphql.Field{
					Name: "Num2",
					Type: graphql.Int,
				},
				"Num3": &graphql.Field{
					Name: "Num3",
					Type: graphql.Int,
				},
				"Num4": &graphql.Field{
					Name: "Num4",
					Type: graphql.Int,
				},
				"Num5": &graphql.Field{
					Name: "Num5",
					Type: graphql.Int,
				},
			},
		},
	)
	assert.Equal(t, expectedResult, gqlResult)
}

func Test_ShouldConvertUIntAttribute(t *testing.T) {
	type Simple struct {
		Num  uint
		Num2 uint8
		Num3 uint16
		Num4 uint32
		Num5 uint64
	}

	gqlResult, _ := ConvertStructToGraphqlSchema(Simple{})
	expectedResult := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Simple",
			Fields: graphql.Fields{
				"Num": &graphql.Field{
					Name: "Num",
					Type: graphql.Int,
				},
				"Num2": &graphql.Field{
					Name: "Num2",
					Type: graphql.Int,
				},
				"Num3": &graphql.Field{
					Name: "Num3",
					Type: graphql.Int,
				},
				"Num4": &graphql.Field{
					Name: "Num4",
					Type: graphql.Int,
				},
				"Num5": &graphql.Field{
					Name: "Num5",
					Type: graphql.Int,
				},
			},
		},
	)
	assert.Equal(t, expectedResult, gqlResult)
}
