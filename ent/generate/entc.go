//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/oSethoum/entgqlplus"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithWhereFilters(true),
		entgql.WithConfigPath("../../gqlgen.yml"),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("../../graph/schemas/schema.graphqls"),
	)

	exs := entgqlplus.NewExtension(
		entgqlplus.WithEchoServer(true),
		entgqlplus.WithDatabase(entgqlplus.SQLite),
		entgqlplus.WithConfigPath("../../gqlgen.yml"),
		entgqlplus.WithMutation(true),
		entgqlplus.WithJWTAuth(true),
	)

	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	options := []entc.Option{
		entc.Extensions(ex, exs),
	}
	if err := entc.Generate("../schema", &gen.Config{}, options...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
