// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type CalorieGoal struct {
	Calories int `json:"calories"`
}

type CalorieGoalInput struct {
	UserID   string `json:"userId"`
	Calories int    `json:"calories"`
}

type ExerciseRecord struct {
	User      *User  `json:"user"`
	Calories  int    `json:"calories"`
	Label     string `json:"label"`
	Timestamp string `json:"timestamp"`
}

type FoodRecord struct {
	User      *User  `json:"user"`
	Calories  int    `json:"calories"`
	Label     string `json:"label"`
	Timestamp string `json:"timestamp"`
}

type NewExerciseRecord struct {
	UserID   string  `json:"userId"`
	Label    *string `json:"label"`
	Calories *int    `json:"calories"`
}

type NewFoodRecord struct {
	UserID   string  `json:"userId"`
	Label    *string `json:"label"`
	Calories *int    `json:"calories"`
}

type NewUser struct {
	ID string `json:"id"`
}

type Sex string

const (
	SexFemale Sex = "FEMALE"
	SexMale   Sex = "MALE"
)

var AllSex = []Sex{
	SexFemale,
	SexMale,
}

func (e Sex) IsValid() bool {
	switch e {
	case SexFemale, SexMale:
		return true
	}
	return false
}

func (e Sex) String() string {
	return string(e)
}

func (e *Sex) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Sex(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Sex", str)
	}
	return nil
}

func (e Sex) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
