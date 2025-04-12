package seki

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Seki) SuccessResponse(w http.ResponseWriter, v any) error {
	s.Log.Debug(fmt.Sprintf("Success reponse"))
	return respond(s, w, 200, v)
}
func (s *Seki) ErrorResponse(w http.ResponseWriter, message string) error {
	s.Log.Error(fmt.Sprintf("Error reponse: %v", message))
	return respond(s, w, 500, errorResponse{Error: message})
}
func (s *Seki) BadRequestResponse(w http.ResponseWriter, message string) error {
	s.Log.Warn(fmt.Sprintf("Bad request reponse: %v", message))
	return respond(s, w, 400, errorResponse{Error: message})
}
func (s *Seki) NotFoundResponse(w http.ResponseWriter) error {
	s.Log.Warn(fmt.Sprintf("Not found reponse"))
	return respond(s, w, 404, errorResponse{Error: "Not found"})
}
func (s *Seki) UnauthorizedResponse(w http.ResponseWriter) error {
	s.Log.Warn(fmt.Sprintf("Unauthorized reponse"))
	return respond(s, w, 401, errorResponse{Error: "Unauthorized"})
}
func (s *Seki) ValidationErrorResponse(w http.ResponseWriter, validationErrors map[string]string) error {
	s.Log.Warn(fmt.Sprintf("Validation error response: %v", validationErrors))
	return respond(s, w, 400, validationErrorResponse{Error: "Input Validation error", Details: validationErrors})
}

func respond(s *Seki, w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		s.Log.Error("JSON Encode error " + err.Error())
		return fmt.Errorf("encode json: %w", err)
	}

	return nil
}

type validationErrorResponse struct {
	Error   string            `json:"error"`
	Details map[string]string `json:"details"`
}

type errorResponse struct {
	Error string `json:"error"`
}

// type paginatedResponse[T any] struct {
// 	Page      int `json:"page"`
// 	PageCount int `json:"pageCount"`
// 	Data      []T `json:"data"`
// }
