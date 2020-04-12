package model

type User struct {
	ID                string   `json:"id"`
	DisplayName       string   `json:"displayName"`
	Age               int      `json:"age"`
	Height            int      `json:"height"`
	Sex               Sex      `json:"sex"`
	Weight            int      `json:"weight"`
	CalorieGoal       int      `json:"calorieGoal"`
	FoodRecordIDs     []string `json:"exerciseRecords"`
	ExerciseRecordIDs []string `json:"exerciseRecords"`
}
