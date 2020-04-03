package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/Matterhorn-Apps/MatterhornApiService/database"
	"github.com/Matterhorn-Apps/MatterhornApiService/graph/generated"
	"github.com/Matterhorn-Apps/MatterhornApiService/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	// Query the database for matching exercise records
	query := fmt.Sprintf("INSERT INTO Users (UserID) VALUES (%s);", input.ID)
	_, readErr := r.DB.Exec(query)
	if readErr != nil {
		if errCode, ok := database.TryExtractMySQLErrorCode(readErr); ok {
			switch *errCode {
			case 1062:
				// User ID exists
				return nil, readErr
			}
		}

		log.Printf("Failed to query database: %v", readErr)
		return nil, readErr
	}

	return &model.User{
		ID: input.ID,
	}, nil
}

func (r *mutationResolver) CreateExerciseRecord(ctx context.Context, input *model.NewExerciseRecord) (*model.ExerciseRecord, error) {
	if *input.Calories < 0 {
		return nil, errors.New("Invalid calorie value provided")
	}

	// Query the database for matching exercise records
	query := fmt.Sprintf("INSERT INTO ExerciseRecords (UserID, Label, Calories) VALUES ('%s', '%s', %d);",
		input.UserID, *input.Label, *input.Calories)
	_, readErr := r.DB.Exec(query)
	if readErr != nil {
		if errCode, ok := database.TryExtractMySQLErrorCode(readErr); ok {
			switch *errCode {
			case 1452:
				// User not found
				return nil, readErr
			}
		}

		log.Printf("Failed to query database: %v", readErr)
		return nil, readErr
	}

	return &model.ExerciseRecord{
		Calories: *input.Calories,
		Label:    *input.Label,
	}, nil
}

func (r *mutationResolver) CreateFoodRecord(ctx context.Context, input *model.NewFoodRecord) (*model.FoodRecord, error) {
	if *input.Calories < 0 {
		return nil, errors.New("Invalid calorie value provided")
	}

	// Query the database for matching Food records
	query := fmt.Sprintf("INSERT INTO FoodRecords (UserID, Label, Calories) VALUES ('%s', '%s', %d);",
		input.UserID, *input.Label, *input.Calories)
	_, readErr := r.DB.Exec(query)
	if readErr != nil {
		if errCode, ok := database.TryExtractMySQLErrorCode(readErr); ok {
			switch *errCode {
			case 1452:
				// User not found
				return nil, readErr
			}
		}

		log.Printf("Failed to query database: %v", readErr)
		return nil, readErr
	}

	return &model.FoodRecord{
		Calories: *input.Calories,
		Label:    *input.Label,
	}, nil
}

func (r *mutationResolver) SetCalorieGoal(ctx context.Context, input model.CalorieGoalInput) (*model.CalorieGoal, error) {
	if input.Calories < 0 {
		return nil, errors.New("Invalid calorie goal provided")
	}

	// Query the database for matching exercise records
	query := fmt.Sprintf(
		"INSERT INTO CalorieGoals (UserID, Calories) VALUES (%s, %d) ON DUPLICATE KEY UPDATE Calories = %[2]d;",
		input.UserID, input.Calories)
	_, readErr := r.DB.Exec(query)
	if readErr != nil {
		if errCode, ok := database.TryExtractMySQLErrorCode(readErr); ok {
			switch *errCode {
			case 1452:
				// User not found
				return nil, readErr
			}
		}

		log.Printf("Failed to query database: %v", readErr)
		return nil, readErr
	}

	return &model.CalorieGoal{}, nil
}

func (r *mutationResolver) IncrementCounter(ctx context.Context, id *string) (*model.Counter, error) {
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
	// Query the database for the User row
	query := fmt.Sprintf("SELECT Displayname, Age, Height, Sex, Weight from Users WHERE UserID='%d';", *id)
	readRows, readErr := r.DB.Query(query)
	if readErr != nil {
		if errCode, ok := database.TryExtractMySQLErrorCode(readErr); ok {
			switch *errCode {
			case 1452:
				// User not found
				return nil, readErr
			}
		}

		log.Printf("Failed to query database: %v", readErr)
		return nil, readErr
	}
	defer readRows.Close()

	// Read value from row response

	var displayName string
	var age, height, weight int64
	var sex model.Sex
	readRows.Next()
	readRows.Scan(&displayName, &age, &height, &sex, &weight)

	// Construct response data object
	return &model.User{
		ID:          strconv.Itoa(*id),
		DisplayName: displayName,
		Age:         int(age),
		Height:      int(height),
		Sex:         sex,
		Weight:      int(weight),
	}, nil
}

func (r *userResolver) CalorieGoal(ctx context.Context, obj *model.User) (*model.CalorieGoal, error) {
	// Query the database for the CalorieGoal row for the given user
	query := fmt.Sprintf("SELECT ID, Calories from CalorieGoals WHERE UserID='%s';", obj.ID)
	readRows, readErr := r.DB.Query(query)
	if readErr != nil {
		if errCode, ok := database.TryExtractMySQLErrorCode(readErr); ok {
			switch *errCode {
			case 1452:
				// User not found
				return nil, readErr
			}
		}

		log.Printf("Failed to query database: %v", readErr)
		return nil, readErr
	}
	defer readRows.Close()

	var id, calories int
	readRows.Next()
	readRows.Scan(&id, &calories)

	// Construct response data object
	return &model.CalorieGoal{
		User:     obj,
		Calories: calories,
	}, nil
}

func (r *userResolver) ExerciseRecords(ctx context.Context, obj *model.User) ([]*model.ExerciseRecord, error) {
	// Query the database for matching exercise records
	query := fmt.Sprintf(
		"SELECT Calories, Label, Timestamp from ExerciseRecords WHERE UserID=%s;", obj.ID)
	readRows, readErr := r.DB.Query(query)
	if readErr != nil {
		if errCode, ok := database.TryExtractMySQLErrorCode(readErr); ok {
			switch *errCode {
			case 1292:
				// Time range invalid
				return nil, readErr
			case 1452:
				// User not found
				return nil, readErr
			}
		}

		log.Printf("Failed to query database: %v", readErr)
		return nil, readErr
	}
	defer readRows.Close()

	records := []*model.ExerciseRecord{}
	for readRows.Next() {
		var calories int32
		var label string
		var timestamp string
		readErr = readRows.Scan(&calories, &label, &timestamp)
		if readErr != nil {
			log.Printf("Failed to read row returned from query: %v", readErr)
			return nil, readErr
		}

		records = append(records, &model.ExerciseRecord{
			Calories:  int(calories),
			Label:     label,
			Timestamp: timestamp,
		})
	}

	return records, nil
}

func (r *userResolver) FoodRecords(ctx context.Context, obj *model.User) ([]*model.FoodRecord, error) {
	// Query the database for matching exercise records
	query := fmt.Sprintf(
		"SELECT Calories, Label, Timestamp from FoodRecords WHERE UserID=%s;", obj.ID)
	readRows, readErr := r.DB.Query(query)
	if readErr != nil {
		if errCode, ok := database.TryExtractMySQLErrorCode(readErr); ok {
			switch *errCode {
			case 1292:
				// Time range invalid
				return nil, readErr
			case 1452:
				// User not found
				return nil, readErr
			}
		}

		log.Printf("Failed to query database: %v", readErr)
		return nil, readErr
	}
	defer readRows.Close()

	records := []*model.FoodRecord{}
	for readRows.Next() {
		var calories int32
		var label string
		var timestamp string
		readErr = readRows.Scan(&calories, &label, &timestamp)
		if readErr != nil {
			log.Printf("Failed to read row returned from query: %v", readErr)
			return nil, readErr
		}

		records = append(records, &model.FoodRecord{
			Calories:  int(calories),
			Label:     label,
			Timestamp: timestamp,
		})
	}

	return records, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
