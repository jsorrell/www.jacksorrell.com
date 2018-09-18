package config

import (
	"errors"
)

var Server struct {
	Port uint16
}

var Contact struct {
	Mailgun struct {
		PublicValidationKey string
		PrivateAPIKey       string
	}
	Email struct {
		Domain    string
		ToAddress string
		Subject   string
	}
	MaxLengths struct {
		Name    uint
		Email   uint
		Message uint
	}
}

// ContactMaxLength get the max length of a contact field from the field name
func ContactMaxLength(field string) (uint, error) {
	switch field {
	case "name":
		return Contact.MaxLengths.Name, nil
	case "email":
		return Contact.MaxLengths.Email, nil
	case "message":
		return Contact.MaxLengths.Message, nil
	default:
		return 0, errors.New("\"" + field + "\" is not a valid field with a max length")
	}
}
