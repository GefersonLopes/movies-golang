package middlewares

import (
	"golang-movie/models"
	"net/http"
	"reflect"
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
