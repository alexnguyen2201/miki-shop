package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/nguyenvanson2201/miki-shop/util"
)

// Custom validator
var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
