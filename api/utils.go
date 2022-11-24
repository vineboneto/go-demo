package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

func TimeExec(start time.Time) {
	fmt.Println("total: ", time.Since(start))
}

func Validator(payload any) []map[string]string {
	var validate = validator.New()

	err := validate.Struct(payload)

	if err != nil {
		var errors []map[string]string
		for _, e := range err.(validator.ValidationErrors) {
			mapErr := make(map[string]string)
			mapErr["field"] = strings.ToLower(string(e.StructField()[0])) + e.StructField()[1:]
			mapErr["tag"] = e.Tag()
			mapErr["value"] = e.Param()
			errors = append(errors, mapErr)
		}
		return errors
	}
	return nil
}
