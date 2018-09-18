package contact

import (
	"fmt"
	"net/http"
	"strings"

	"unicode/utf8"

	"github.com/gorilla/mux"
	"github.com/goware/emailx"
	"github.com/jsorrell/www.jacksorrell.com/configloader"
	"github.com/jsorrell/www.jacksorrell.com/templates"
	weberror "github.com/jsorrell/www.jacksorrell.com/web/error"
	"gopkg.in/mailgun/mailgun-go.v1"
)

// RegisterRoutesTo registers routes to router
func RegisterRoutesTo(router *mux.Router) {
	sub := router.Path("/contact/").Subrouter()
	sub.Methods("GET", "HEAD").Handler(templates.Contact)
	sub.Methods("POST").HandlerFunc(handleContactFormSubmission)
}

func handleContactFormSubmission(res http.ResponseWriter, req *http.Request) {
	ajax := req.URL.Query().Get("ajax") == "1"
	var errorHandler weberror.ErrorHandler
	if ajax {
		errorHandler = weberror.Plain
	} else {
		errorHandler = weberror.HTML
	}
	maxLengths := configloader.Config.Contact.MaxLengths
	req.Body = http.MaxBytesReader(res, req.Body, int64(maxLengths.Name+maxLengths.Email+maxLengths.Message+1000))

	if errorMessage, statusCode := validateFormSubmission(req); errorMessage != "" {
		errorHandler.Error(res, req, statusCode, errorMessage, fmt.Sprintf("Invalid Contact Submission: "+errorMessage))
		return
	}

	err := sendMailgunEmail(req.PostFormValue("name"), req.PostFormValue("email"), req.PostFormValue("message"))
	if err != nil {
		errorHandler.Error(res, req, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}
	//TODO pretty message sent page
	res.Write([]byte("Message sent"))
}

func validateFormSubmission(req *http.Request) (string, int) {
	err := req.ParseForm()
	if err != nil {
		return "Invalid Message", http.StatusBadRequest
	}

	values := map[string]string{
		"name":    getFieldValue(req, "name"),
		"email":   getFieldValue(req, "email"),
		"message": getFieldValue(req, "message"),
	}

	for fieldName, fieldValue := range values {
		if fieldValue == "" {
			return "Missing " + strings.ToTitle(fieldName), http.StatusBadRequest
		}

		maxLength, erro := configloader.ContactMaxLength(fieldName)
		if erro != nil {
			return "Internal Server Error", http.StatusInternalServerError
		}
		if uint(utf8.RuneCountInString(fieldValue)) > maxLength {
			return strings.ToTitle(fieldValue) + " is too long", http.StatusBadRequest
		}
	}

	email := req.PostFormValue("email")
	err = emailx.Validate(email)
	if err != nil {
		return "Email <" + email + "> is invalid", http.StatusBadRequest
	}

	return "", 200
}

func getFieldValue(req *http.Request, field string) string {
	val := req.PostFormValue(field)
	return strings.Replace(val, "\r\n", "\n", -1)
}

func sendMailgunEmail(name, email, message string) error {
	mg := mailgun.NewMailgun(
		configloader.Config.Contact.Email.Domain,
		configloader.Config.Contact.Mailgun.PrivateAPIKey,
		configloader.Config.Contact.Mailgun.PublicValidationKey,
	)

	m := mg.NewMessage(
		fmt.Sprintf("%s<%s>", name, email),
		configloader.Config.Contact.Email.Subject,
		message,
		configloader.Config.Contact.Email.ToAddress,
	)
	_, _, err := mg.Send(m)

	return err
}
