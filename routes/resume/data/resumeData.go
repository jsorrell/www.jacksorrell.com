package data

import (
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/jsorrell/www.jacksorrell.com/data"
)

type month struct {
	Month string `yaml:"month"`
	Year  string `yaml:"year"`
}

// ResumeData contains all data in resume.yaml used to build /resume/.
type ResumeData struct {
	Profile struct {
		Name struct {
			First string `yaml:"first"`
			Last  string `yaml:"last"`
		} `yaml:"name"`
		Title     string    `yaml:"title"`
		Birthdate time.Time `yaml:"birthdate"`
		Residence struct {
			City    string `yaml:"city"`
			State   string `yaml:"state"`
			Country string `yaml:"country"`
		} `yaml:"residence"`
		Bio string `yaml:"bio"`
	} `yaml:"profile"`

	Degrees []struct {
		Name      string `yaml:"name"`
		StartDate month  `yaml:"startDate"`
		EndDate   month  `yaml:"endDate"`
		Location  struct {
			City  string `yaml:"city"`
			State string `ymal:"state"`
		} `yaml:"location"`
		Type        string   `yaml:"type"`
		Degree      string   `yaml:"degree"`
		DegreeShort string   `yaml:"degreeShort"`
		Major       string   `yaml:"major"`
		Minors      []string `yaml:"minors"`
	} `yaml:"degrees"`

	CourseCategories []struct {
		Category string `yaml:"category"`
		Courses  []struct {
			Name        string `yaml:"name"`
			ShortName   string `yaml:"shortName"`
			ID          string `yaml:"id"`
			Description string `yaml:"description"`
		} `yaml:"courses"`
	} `yaml:"courseCategories"`

	Experiences []struct {
		Title    string `yaml:"title"`
		Company  string `yaml:"company"`
		Location struct {
			City  string `yaml:"city"`
			State string `yaml:"state"`
		} `yaml:"location"`
		StartDate   month  `yaml:"startDate"`
		EndDate     month  `yaml:"endDate"`
		Description string `yaml:"description"`
	} `yaml:"experiences"`

	Projects []struct {
		Name        string `yaml:"name"`
		StartDate   month  `yaml:"startDate"`
		EndDate     month  `yaml:"endDate"`
		Description string `yaml:"description"`
	} `yaml:"projects"`

	SkillCategories []struct {
		Category string `yaml:"category"`
		Skills   []struct {
			Name  string `yaml:"name"`
			Level int    `yaml:"level"`
		} `yaml:"skills"`
	} `yaml:"skillCategories"`

	Links []struct {
		Name string `yaml:"name"`
		Icon string `yaml:"icon"`
		Href string `yaml:"href"`
	} `yaml:"links"`
}

// ParseResumeData parses resume.yaml and returns the data.
func ParseResumeData() (*ResumeData, error) {
	resumeDataYaml, err := data.Assets.Open("resume_data.yaml")
	if err != nil {
		return nil, err
	}
	defer resumeDataYaml.Close()

	resumeDataBytes, err := ioutil.ReadAll(resumeDataYaml)

	if err != nil {
		return nil, err
	}

	var data ResumeData
	err = yaml.Unmarshal(resumeDataBytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
