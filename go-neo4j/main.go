package main

import (
	"context"
	"fmt"
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func main() {
	// replace the following with your credentials
	dbUri := "neo4j+s://your-uri"
	driver, err := neo4j.NewDriverWithContext(dbUri, neo4j.BasicAuth("neo4j", "your-password", ""))
	if err != nil {
		panic(err)
	}
	// Starting with 5.0, you can control the execution of most driver APIs
	// To keep things simple, we create here a never-cancelling context
	// Read https://pkg.go.dev/context to learn more about contexts
	ctx := context.Background()
	// Handle driver lifetime based on your application lifetime requirements.
	// driver's lifetime is usually bound by the application lifetime, which usually implies one driver instance per
	// application

	defer driver.Close(ctx) // Make sure to handle errors during deferred calls

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		panic(err)
	}

	//readTransaction(ctx, driver)
	read(ctx, driver)
}

func readTransaction(ctx context.Context, driver neo4j.DriverWithContext) {
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	start := time.Now()
	creatorResult, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, _ := tx.Run(ctx, "profile MATCH (n:Creator {id: $idParam}) return n", map[string]any{
			"idParam": "323ce0f7-7",
		})
		records, _ := result.Collect(ctx)
		return records, nil
	})
	elapsed := time.Since(start)

	fmt.Printf("Time taken with read transaction: %v\n", elapsed)

	for _, creator := range creatorResult.([]*neo4j.Record) {
		fmt.Println(creator.AsMap())
	}

	fmt.Println(err)
}

func read(ctx context.Context, driver neo4j.DriverWithContext) {
	start := time.Now()

	result, _ := neo4j.ExecuteQuery(ctx, driver, "profile MATCH (n:Creator {id: $idParam}) return n",
		map[string]any{
			"idParam": "323ce0f7-7",
		}, neo4j.EagerResultTransformer)
	elapsed := time.Since(start)

	fmt.Printf("Time taken: %v\n", elapsed)

	// Loop through results and do something with them
	for _, record := range result.Records {
		fmt.Println(record.AsMap())
	}

	// Summary information
	fmt.Printf("The query `%v` returned %v records in %+v.\n",
		result.Summary.Query().Text(), len(result.Records),
		result.Summary.ResultAvailableAfter())
}
