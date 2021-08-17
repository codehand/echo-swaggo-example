# Example echo with swaggo/swag

## Author

- codehand

## Support

- Golang 1.15 or later 
- Echo/Gin
- Swagger gen by swaggo
- validator gopkg.in/go-playground/validator.v9

## How to run
The first
```
    go get github.com/swaggo/swag
```

Then run
```
    swag init
```

## Description

```
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

```

Swag understand struct tags and will read the example and format struct tags. It also supports the validate:"required" struct tag which is used by the go-playground/validator library.

On the last line you see @name ExampleValidatorRequest. This is where you can give the name of response/model. By default, Swag gives the name of package.Struct, which in this case was main.ExampleValidatorRequest.