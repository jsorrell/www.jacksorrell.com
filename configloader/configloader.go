package configloader

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jsorrell/www.jacksorrell.com/utils/copy"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
)

// Configuration configuration object
type Configuration struct {
	Server struct {
		Port uint16
	}
	Contact struct {
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
}

// Config the configuration object storing the settings for the app
var Config Configuration

// TODO better error messages?
func init() {
	configFileName := kingpin.Arg("config", "The YAML configuration file. Defaults to config.yaml.").Default("config.yaml").String()
	generateConfig := kingpin.Flag("genConfig", "Generates an example config and exits immediately.").Short('g').Bool()
	kingpin.Parse()

	if *generateConfig {
		if _, err := os.Stat(*configFileName); os.IsNotExist(err) {
			err = copy.WriteAssetToDisk("config.example.yaml", *configFileName)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s generated.\n", *configFileName)
			os.Exit(0)
		}

		fmt.Fprintf(os.Stderr, "%s exists. Exiting with no file changes.\n", *configFileName)
		os.Exit(1)
	}

	configFile, err := ioutil.ReadFile(*configFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read config file %s.\n%v\n", *configFileName, err)
		os.Exit(1)
	}

	var yamlFile configurationSchema
	err = yaml.Unmarshal(configFile, &yamlFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invald YAML in config file %s.\n%v\n", *configFileName, err)
		os.Exit(1)
	}

	if err = Config.genConfiguration(yamlFile); err != nil {
		fmt.Fprintf(os.Stderr, "Invald fields in config file %s.\n%v\n", *configFileName, err)
		os.Exit(1)
	}
}

func (config *Configuration) genConfiguration(yamlFile configurationSchema) error {
	/* Defaults */
	if yamlFile.Server.Port == 0 {
		yamlFile.Server.Port = 3000
	}
	config.Server.Port = yamlFile.Server.Port

	if yamlFile.Contact.Email.Subject == "" {
		yamlFile.Contact.Email.Subject = "Contact Form Message"
	}
	config.Contact.Email.Subject = yamlFile.Contact.Email.Subject

	if yamlFile.Contact.MaxLengths.Name == 0 {
		yamlFile.Contact.MaxLengths.Name = 70
	}
	config.Contact.MaxLengths.Name = yamlFile.Contact.MaxLengths.Name

	if yamlFile.Contact.MaxLengths.Email == 0 {
		yamlFile.Contact.MaxLengths.Email = 254
	}
	config.Contact.MaxLengths.Email = yamlFile.Contact.MaxLengths.Email

	if yamlFile.Contact.MaxLengths.Message == 0 {
		yamlFile.Contact.MaxLengths.Message = 10000
	}
	config.Contact.MaxLengths.Message = yamlFile.Contact.MaxLengths.Message

	/* Check Validity */
	if yamlFile.LogLevel == "" {
		yamlFile.LogLevel = "warn"
	}
	level, err := log.ParseLevel(yamlFile.LogLevel)
	if err != nil {
		return err
	}
	log.SetLevel(level)

	if yamlFile.Contact.Mailgun.PublicValidationKey == "" {
		return errors.New("mailgun publicValidationKey is required")
	}
	config.Contact.Mailgun.PublicValidationKey = yamlFile.Contact.Mailgun.PublicValidationKey

	if yamlFile.Contact.Mailgun.PrivateAPIKey == "" {
		return errors.New("mailgun privateAPIKey is required")
	}
	config.Contact.Mailgun.PrivateAPIKey = yamlFile.Contact.Mailgun.PrivateAPIKey

	if yamlFile.Contact.Email.Domain == "" {
		return errors.New("email domain is required")
	}
	config.Contact.Email.Domain = yamlFile.Contact.Email.Domain

	if yamlFile.Contact.Email.ToAddress == "" {
		return errors.New("email toAddress is required")
	}
	config.Contact.Email.ToAddress = yamlFile.Contact.Email.ToAddress

	return nil
}

// ContactMaxLength get the max length of a contact field from the field name
func ContactMaxLength(field string) (uint, error) {
	switch field {
	case "name":
		return Config.Contact.MaxLengths.Name, nil
	case "email":
		return Config.Contact.MaxLengths.Email, nil
	case "message":
		return Config.Contact.MaxLengths.Message, nil
	default:
		return 0, errors.New("\"" + field + "\" is not a valid field with a max length")
	}
}
