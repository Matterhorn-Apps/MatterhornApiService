package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/Matterhorn-Apps/MatterhornApiService/graph/generated"
	"github.com/Matterhorn-Apps/MatterhornApiService/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) IncrementCounter(ctx context.Context, id *int) (*model.Counter, error) {
	// Query the database for the current counter value
	log.Print("Querying...")
	readRows, readErr := r.DB.Query("SELECT Value from Counters WHERE ID='1';")
	if readErr != nil {
		log.Printf("Failed to query database: %v", readErr)
		return nil, readErr
	}
	defer readRows.Close()

	// Read value from row response
	var value int64
	readRows.Next()
	readRows.Scan(&value)

	// Query the database to update the counter value
	updateRows, updateErr := r.DB.Query(fmt.Sprintf("UPDATE Counters SET Value='%d' WHERE ID='%d';", value+1, 1))
	if updateErr != nil {
		log.Printf("Failed to update counter value: %v", updateErr)
		return nil, updateErr
	}
	defer updateRows.Close()

	// Construct response data object
	log.Print("Returning response...")
	return &model.Counter{
		ID:    strconv.Itoa(1),
		Value: int(value + 1),
	}, nil
}

func (r *queryResolver) Counter(ctx context.Context, id *int) (*model.Counter, error) {
	// Query the database for the current counter value
	readRows, readErr := r.DB.Query("SELECT Value from Counters WHERE ID='1';")
	if readErr != nil {
		log.Printf("Failed to query database: %v", readErr)
		return nil, readErr
	}
	defer readRows.Close()

	// Read value from row response
	var value int64
	readRows.Next()
	readRows.Scan(&value)

	// Construct response data object
	return &model.Counter{
		ID:    strconv.Itoa(1),
		Value: int(value),
	}, nil
}

func (r *queryResolver) User(ctx context.Context, id *int) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
