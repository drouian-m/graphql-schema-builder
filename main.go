package main

import (
	schemabuilder "github.com/drouian-m/graphql-schema-builder/src"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
	"net/http"
	"time"
)

type Toto struct {
	Name string
	ID   int64
	//CreateAt time.Time
}

type Complex struct {
	Title       string `gql:"type=string,name=Title"`
	ID          int64  `gql:"type=int,name=ID"`
	Description string `gql:"type=string,name=Description"`
	//CreatedAt   time.Time `gql:"type=time,name=CreatedAt"`
}

type Simple struct {
	Valid bool      `gql:"type=boolean,name=Valid"`
	Time  time.Time `gql:"type=time,name=Time"`
}

func main() {
	//fmt.Println("Hello world !")
	//schemabuilder.ConvertStructToGraphqlSchema(Simple{Valid: true})

	//toto1 := Toto{Name: "Hi", ID: 123}
	//toto2 := Toto{Name: "qlwkej", ID: 132}
	//schemabuilder.ConvertStructToGraphqlSchema(toto1)
	//res := schemabuilder.ConvertStructToGraphqlSchema(Complex{
	//	Title:       "awkledjqwe",
	//	ID:          132890,
	//	Description: "asdoiuqwkdhjqwoidu",
	//	Items:       []Toto{toto1, toto2},
	//})

	res, _ := schemabuilder.ConvertStructToGraphqlSchema(Complex{})

	//fmt.Println()

	schemaConfig := graphql.SchemaConfig{Query: res}
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
	//http.ListenAndServe(":8080", nil)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
