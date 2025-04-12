package seki

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func DecodeValid[T Validator](r *http.Request) (T, map[string]string) {
	var v T

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&v); err != nil {
		if unmarshalTypeError, ok := err.(*json.UnmarshalTypeError); ok {
			unmarshalErrorMessage := fmt.Sprintf("JSON Decode error: received %v where expected %v", unmarshalTypeError.Value, unmarshalTypeError.Type)
			return v, map[string]string{unmarshalTypeError.Field: unmarshalErrorMessage}
		} else {
			return v, map[string]string{"JSON error": err.Error()}
		}
	}

	if problems := v.Valid(); len(problems) > 0 {
		return v, problems
	}
	return v, nil
}
