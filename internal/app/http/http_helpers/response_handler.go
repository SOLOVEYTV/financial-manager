package http_helpers

import (
	"encoding/json"
	"net/http"
)

// FieldViolation описывает одну ошибку валидации конкретного поля запроса.
// Используется для передачи клиенту, какое поле не прошло проверку и почему.
type FieldViolation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidationError накапливает список нарушений валидации входных данных.
// Решает задачу пошагового сбора ошибок по полям перед формированием HTTP-ответа.
type ValidationError struct {
	violations []FieldViolation
}

// NewValidationError создаёт пустой контейнер для ошибок валидации.
func NewValidationError() *ValidationError {
	return &ValidationError{}
}

// Add добавляет нарушение для указанного поля и возвращает тот же объект
// для цепочки вызовов (fluent API).
func (e *ValidationError) Add(field, message string) *ValidationError {
	e.violations = append(e.violations, FieldViolation{
		Field:   field,
		Message: message,
	})

	return e
}

// Violations возвращает собранный список нарушений для включения в JSON-ответ клиенту.
func (e *ValidationError) Violations() []FieldViolation {
	return e.violations
}

// Error реализует интерфейс error и возвращает машиночитаемый код ошибки валидации.
func (e *ValidationError) Error() string {
	return CodeValidation
}

// ResponseHandler инкапсулирует запись HTTP-ответов в едином JSON-формате.
// Централизует установку заголовков, статус-кода и сериализацию тела ответа.
type ResponseHandler struct {
	rw http.ResponseWriter
}

// NewResponseHandler создаёт обработчик ответов для конкретного HTTP-запроса.
func NewResponseHandler(rw http.ResponseWriter) *ResponseHandler {
	return &ResponseHandler{rw: rw}
}

// errorResponseBody — внутренняя структура JSON-тела ошибочного HTTP-ответа.
// Содержит человекочитаемое сообщение, код ошибки и опциональный список нарушений валидации.
type errorResponseBody struct {
	Message    string           `json:"message"`
	Error      string           `json:"error"`
	Violations []FieldViolation `json:"violations,omitempty"`
}

// ErrorResponse отправляет клиенту JSON-ответ об ошибке с заданным HTTP-статусом.
// При наличии violations добавляет детализацию по полям, не прошедшим валидацию.
func (h *ResponseHandler) ErrorResponse(statusCode int, message, errorCode string, violations ...FieldViolation) {
	resp := errorResponseBody{
		Message: message,
		Error:   errorCode,
	}
	if len(violations) > 0 {
		resp.Violations = violations
	}

	h.rw.Header().Set("Content-Type", "application/json")
	h.rw.WriteHeader(statusCode)

	err := json.NewEncoder(h.rw).Encode(resp)
	if err != nil {
		return
	}
}

// SuccessResponse отправляет успешный HTTP-ответ.
// Без body записывает только статус-код; с body сериализует первый аргумент в JSON.
func (h *ResponseHandler) SuccessResponse(statusCode int, body ...any) {
	if len(body) == 0 {
		h.rw.WriteHeader(statusCode)

		return
	}

	h.rw.Header().Set("Content-Type", "application/json")
	h.rw.WriteHeader(statusCode)

	if err := json.NewEncoder(h.rw).Encode(body[0]); err != nil {
		return
	}
}
