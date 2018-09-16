package configloader

// Configuration the type of Config
type configurationSchema struct {
	Dev      bool   `yaml:"dev"`
	LogLevel string `yaml:"logLevel"`
	Server   struct {
		Port uint16 `yaml:"port"`
	} `yaml:"server"`
	Contact struct {
		Mailgun struct {
			PrivateAPIKey       string `yaml:"privateAPIKey"`
			PublicValidationKey string `yaml:"publicValidationKey"`
		} `yaml:"mailgun"`
		Email struct {
			Domain    string `yaml:"domain"`
			ToAddress string `yaml:"toAddress"`
			Subject   string `yaml:"subject"`
		} `yaml:"email"`
		MaxLengths struct {
			Name    uint `yaml:"name"`
			Email   uint `yaml:"email"`
			Message uint `yaml:"message"`
		} `yaml:"maxLengths"`
	} `yaml:"contact"`
}
