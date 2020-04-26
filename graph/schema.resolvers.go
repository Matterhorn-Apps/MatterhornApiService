package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/Matterhorn-Apps/MatterhornApiService/auth"
	"github.com/Matterhorn-Apps/MatterhornApiService/database"
	"github.com/Matterhorn-Apps/MatterhornApiService/graph/generated"
	"github.com/Matterhorn-Apps/MatterhornApiService/graph/model"
)

func (r *mutationResolver) CreateMe(ctx context.Context) (*model.User, error) {
	// Get the user ID from context
	// If not found, request is not authenticated and should fail
	tokenSub, err := auth.GetUserIDFromContext(ctx)
	if tokenSub == nil || err != nil {
		return nil, errors.New("request failed: caller not authenticated")
	}

	// Try to get user with given token sub as ID
	// This will fail if the user already exists
	return r.CreateUser(ctx, model.NewUser{
		ID: *tokenSub,
	})
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	// TODO: Require authorization to access protected APIs #13
	// https://github.com/Matterhorn-Apps/MatterhornApiService/issues/13

	// Execute query to insert new user row
	query := fmt.Sprintf("INSERT INTO users (user_id) VALUES ( '%s' );", input.ID)
	_, readErr := r.DB.Query(query)
	if readErr != nil {
		if errCode, ok := database.TryExtractMySQLErrorCode(readErr); ok {
			switch *errCode {
			case 1062:
				// User ID already exists
				return nil, fmt.Errorf("user with ID '%s' already exists", input.ID)
			}
		}

		log.Printf("Failed to query database: %v", readErr)
		return nil, readErr
	}

	return &model.User{
		ID: input.ID,
	}, nil
}

func (r *mutationResolver) CreateExerciseRecord(ctx context.Context, input model.NewExerciseRecord) (*model.ExerciseRecord, error) {
	// Get the user ID from context
	// If not found, request is not authenticated and should fail
	userId, err := auth.GetUserIDFromContext(ctx)
	if userId == nil || err != nil {
		return nil, errors.New("request failed: caller not authenticated")
	}

	if *input.Calories < 0 {
		return nil, errors.New("Invalid calorie value provided")
	}

	// Query the database for matching exercise records
	query := fmt.Sprintf("INSERT INTO exercise_records (user_id, label, calories) VALUES ('%s', '%s', %d);",
		*userId, *input.Label, *input.Calories)
	_, readErr := r.DB.Query(query)
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

func (r *mutationResolver) CreateFoodRecord(ctx context.Context, input model.NewFoodRecord) (*model.FoodRecord, error) {
	// Get the user ID from context
	// If not found, request is not authenticated and should fail
	userId, err := auth.GetUserIDFromContext(ctx)
	if userId == nil || err != nil {
		return nil, errors.New("request failed: caller not authenticated")
	}

	if *input.Calories < 0 {
		return nil, errors.New("Invalid calorie value provided")
	}

	// Query the database for matching Food records
	query := fmt.Sprintf("INSERT INTO food_records (user_id, label, calories) VALUES ('%s', '%s', %d);",
		*userId, *input.Label, *input.Calories)
	_, readErr := r.DB.Query(query)
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
	// Get the user ID from context
	// If not found, request is not authenticated and should fail
	userId, err := auth.GetUserIDFromContext(ctx)
	if userId == nil || err != nil {
		return nil, errors.New("request failed: caller not authenticated")
	}

	if input.Calories < 0 {
		return nil, errors.New("Invalid calorie goal provided")
	}

	// Update calorie goal in the user row
	query := fmt.Sprintf("UPDATE users SET calorie_goal=%d WHERE user_id='%s'",
		input.Calories, *userId)
	_, readErr := r.DB.Query(query)
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

	return &model.CalorieGoal{
		Calories: input.Calories,
	}, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*bool, error) {
	// Delete the row in the users table with an ID matching the one passed in
	query := fmt.Sprintf(`DELETE FROM users WHERE user_id='%s';`, id)
	_, readErr := r.DB.Exec(query)
	if readErr != nil {
		if errCode, ok := database.TryExtractMySQLErrorCode(readErr); ok {
			switch *errCode {
			case 1452:
				// User not found
				temp := false
				return &temp, readErr
			}
		}

		log.Printf("Failed to query database: %v", readErr)
		temp := false
		return &temp, readErr
	}

	// Allocating an address so we can return a pointer
	temp := true
	return &temp, nil
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	// Get the user ID from context
	// If not found, request is not authenticated and should fail
	userId, err := auth.GetUserIDFromContext(ctx)
	if userId == nil || err != nil {
		return nil, errors.New("request failed: caller not authenticated")
	}

	return r.User(ctx, *userId)
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	// Query the database for the User row
	query := fmt.Sprintf(
		`SELECT display_name, age, height_inches, sex, weight_pounds, calorie_goal from users 
			WHERE user_id='%s';`, id)
	readRows, readErr := r.DB.Query(query)
	if readErr != nil {
		if errCode, ok := database.TryExtractMySQLErrorCode(readErr); ok {
			switch *errCode {
			case 1452:
				// User does not exist - return nil value without error
				return nil, nil
			}
		}

		log.Printf("Failed to query database: %v", readErr)
		return nil, readErr
	}
	defer readRows.Close()

	// Read value from row response

	var displayName, sex sql.NullString
	var age, height, weight, calorieGoal sql.NullInt64
	readRows.Next()
	readRows.Scan(&displayName, &age, &height, &sex, &weight, &calorieGoal)

	// Construct response data object
	return &model.User{
		ID:          id,
		DisplayName: displayName.String,
		Age:         int(age.Int64),
		CalorieGoal: int(calorieGoal.Int64),
		Height:      int(height.Int64),
		Sex:         model.Sex(sex.String),
		Weight:      int(weight.Int64),
	}, nil
}

func (r *userResolver) ExerciseRecords(ctx context.Context, obj *model.User, startTime *string, endTime *string) ([]*model.ExerciseRecord, error) {
	MinTimestamp := "1000-00-00 00:00:00"
	MaxTimestamp := "9999-12-31 23:59:59"

	if startTime == nil {
		startTime = &MinTimestamp
	}
	if endTime == nil {
		endTime = &MaxTimestamp
	}

	// Query the database for matching exercise records
	query := fmt.Sprintf(
		"SELECT Calories, Label, timestamp from exercise_records WHERE user_id='%s' AND timestamp BETWEEN '%s' AND '%s';",
		obj.ID, *startTime, *endTime)
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

func (r *userResolver) FoodRecords(ctx context.Context, obj *model.User, startTime *string, endTime *string) ([]*model.FoodRecord, error) {
	MinTimestamp := "1000-00-00 00:00:00"
	MaxTimestamp := "9999-12-31 23:59:59"

	if startTime == nil {
		startTime = &MinTimestamp
	}
	if endTime == nil {
		endTime = &MaxTimestamp
	}

	// Query the database for matching food records
	query := fmt.Sprintf(
		"SELECT calories, label, timestamp from food_records WHERE user_id='%s' AND timestamp BETWEEN '%s' AND '%s';",
		obj.ID, *startTime, *endTime)
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
