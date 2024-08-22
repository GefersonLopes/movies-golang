package middlewares

import (
	"golang-movie/models"
	"net/http"
	"reflect"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type ValidationError struct {
	Message string
	StatusCode int
}

func (e *ValidationError) Error() string {
	return e.Message
}

func ValidateCreateMovie(movie models.Movie) error {
	errosList := []interface{}{"", 0, nil}
	
	value := reflect.ValueOf(movie)

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i).Interface()

		for _, error := range errosList {
			if field == error {
				return &ValidationError{
					Message: value.Type().Field(i).Name + " is required",
					StatusCode: http.StatusBadRequest,
				}
			}
		}
	}
	return nil
}

func ValidadeParamSearchMovie(id primitive.ObjectID) error {
	if id == primitive.NilObjectID {
		return &ValidationError{
			Message: "id is required",
			StatusCode: http.StatusBadRequest,
		}
	}
	return nil
}


func ReturnNotFoundMovie() error {
	return &ValidationError{
		Message: "movie not found",
		StatusCode: http.StatusNotFound,
	}
}
	