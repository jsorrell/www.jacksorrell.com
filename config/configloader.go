package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/jsorrell/www.jacksorrell.com/utils/copy"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
)

func init() {
	app := kingpin.New("www.jacksorrell.com", "The private web server of Jack Sorrell.")
	genConfig := app.Command("genconfig", "Generates an example config and exits immediately.")
	genConfigFileName := genConfig.Arg("config", "The YAML configuration file to create.").Default("config.yaml").String()

	server := app.Command("server", "Runs the server.").Default()

	optionalConfigFileName := server.Arg("config", "The YAML configuration file. Defaults to config.yaml.").String()

	logLevel := server.Flag("log-level", "The logging level.").Short('l').HintOptions("panic", "fatal", "error", "warning", "warn", "info", "debug").PlaceHolder("warn").String()

	server.Flag("port", "The port to run the server on.").Short('p').HintOptions("80", "443", "3000", "8080").PlaceHolder("3000").Uint16Var(&Server.Port)
	server.Flag("mailgun-validation-key", "Mailgun public validation key. Required.").PlaceHolder("pubkey-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx").StringVar(&Contact.Mailgun.PublicValidationKey)
	server.Flag("mailgun-api-key", "Mailgun private API key. Required.").PlaceHolder("key-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx").StringVar(&Contact.Mailgun.PrivateAPIKey)
	server.Flag("email-domain", "Domain to send contact email to. Required.").PlaceHolder("mg.example.com").StringVar(&Contact.Email.Domain)
	server.Flag("email-to-address", "Address to send contact email to. Required.").PlaceHolder("contact@example.com").StringVar(&Contact.Email.ToAddress)
	server.Flag("email-subject", "Subject to use for contact email.").PlaceHolder("Contact Form Message").StringVar(&Contact.Email.Subject)
	server.Flag("contact-name-maxlength", "The maximum length allowed for the contact name.").PlaceHolder("70").UintVar(&Contact.MaxLengths.Name)
	server.Flag("contact-email-maxlength", "The maximum length allowed for the 'from' email address.").PlaceHolder("254").UintVar(&Contact.MaxLengths.Email)
	server.Flag("contact-message-maxlength", "The maximum length allowed for the email message").PlaceHolder("10000").UintVar(&Contact.MaxLengths.Message)

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case genConfig.FullCommand():
		if _, err := os.Stat(*genConfigFileName); os.IsNotExist(err) {
			err = copy.WriteAssetToDisk("config.example.yaml", *genConfigFileName)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s generated.\n", *genConfigFileName)
			os.Exit(0)
		}

		kingpin.Fatalf("%s exists. Exiting with no file changes.\n", *genConfigFileName)
	case server.FullCommand():
		configFileName := pickFirstDefault("config.yaml", *optionalConfigFileName).(string)
		configFile, err := ioutil.ReadFile(configFileName)
		if err != nil && !(os.IsNotExist(err) && *optionalConfigFileName == "") {
			kingpin.Fatalf("Could not read config file %s.\n%v\n", configFileName, err)
		}

		var yamlFile configurationSchema
		kingpin.FatalIfError(yaml.Unmarshal(configFile, &yamlFile), "Invald YAML in config file %s.\n%v\n", configFileName, err)
		genConfiguration(yamlFile)

		// Log Level
		*logLevel = pickFirstDefault("warn", *logLevel, yamlFile.LogLevel).(string)
		logLevelVal, err := log.ParseLevel(*logLevel)
		kingpin.FatalIfError(err, "Invald log level %s.\n", *logLevel)
		log.SetLevel(logLevelVal)
	}
}

func genConfiguration(yamlFile configurationSchema) {
	Server.Port = pickFirstDefault(uint16(3000), Server.Port, yamlFile.Server.Port).(uint16)
	Contact.Mailgun.PublicValidationKey = pickFirstRequired("mailgun validation key", Contact.Mailgun.PublicValidationKey, yamlFile.Contact.Mailgun.PublicValidationKey).(string)
	Contact.Mailgun.PrivateAPIKey = pickFirstRequired("mailgun api key", Contact.Mailgun.PrivateAPIKey, yamlFile.Contact.Mailgun.PrivateAPIKey).(string)
	Contact.Email.Domain = pickFirstRequired("email domain", Contact.Email.Domain, yamlFile.Contact.Email.Domain).(string)
	Contact.Email.ToAddress = pickFirstRequired("email 'to' address", Contact.Email.ToAddress, yamlFile.Contact.Email.ToAddress).(string)
	Contact.Email.Subject = pickFirstDefault("Contact Form Message", Contact.Email.Subject, yamlFile.Contact.Email.Subject).(string)
	Contact.MaxLengths.Name = pickFirstDefault(uint(70), Contact.MaxLengths.Name, yamlFile.Contact.MaxLengths.Name).(uint)
	Contact.MaxLengths.Email = pickFirstDefault(uint(254), Contact.MaxLengths.Email, yamlFile.Contact.MaxLengths.Email).(uint)
	Contact.MaxLengths.Message = pickFirstDefault(uint(10000), Contact.MaxLengths.Message, yamlFile.Contact.MaxLengths.Message).(uint)
}

func pickFirstDefault(def interface{}, args ...interface{}) interface{} {
	for _, val := range args {
		if val != reflect.Zero(reflect.TypeOf(val)).Interface() {
			return val
		}
	}
	return def
}

func pickFirstRequired(name string, args ...interface{}) interface{} {
	for _, val := range args {
		if val != reflect.Zero(reflect.TypeOf(val)).Interface() {
			return val
		}
	}
	kingpin.Fatalf("%s is required.", name)
	return nil // Will never get here
}
