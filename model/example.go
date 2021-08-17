package model

// An ExampleRequest model.
//
// This is used for operations that want an Example as body of the request
type ExampleRequest struct {
	// Fiels string comment
	FieldString string `json:"field_string"`
	// Field int comment
	FieldInt int `json:"field_int"`
	// Field bool comment
	FieldBool bool `json:"field_bool"`
	// Field array comment
	FieldArray []string `json:"field_array"`
} // @name ExampleRequest

// An ExampleRequest model.
//
// This is used for operations that want an Example as body of the request
type ExampleValidatorRequest struct {
	// Fiels string comment
	FieldString string `json:"field_string" validate:"required"`
	// Field int comment
	FieldInt int `json:"field_int" validate:"required"`
	// Field bool comment
	FieldBool bool `json:"field_bool" validate:"required"`
	// Field array comment
	FieldArray []string `json:"field_array"`
} // @name ExampleValidatorRequest

// Example is model
type Example struct {
	// Fiels string comment
	FieldString string `json:"field_string" example:"something"`
	// Field int comment
	FieldInt int `json:"field_int" example:"1"`
	// Field bool comment
	FieldBool bool `json:"field_bool" example:"true"`
	// Field array comment
	FieldArray []string `json:"field_array"`
}
