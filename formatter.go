package seeder

import (
	"encoding/json"
	"log"
	"math/rand"
	"regexp"
	"strings"
	"sync"
	"time"
)

type Formats struct {
	Name  []string `json:"nameFormat"`
	Email []string `json:"emailFormat"`
}

type Data struct {
	FirstNameMale   []string `json:"firstNameMale"`
	FirstNameFemale []string `json:"firstNameFemale"`
	LastName        []string `json:"lastName"`
	TitleMale       []string `json:"titleMale"`
	TitleFemale     []string `json:"titleFemale"`
	Suffix          []string `json:"suffix"`
	FreeEmailDomain []string `json:"freeEmailDomain"`
	Tld             []string `json:"tld"`
}

type serviceProvider struct {
	randomize *rand.Rand
	mutex     *sync.Mutex
}

func (provider *serviceProvider) randomIndex(count int) int {
	provider.mutex.Lock()
	defer provider.mutex.Unlock()
	return provider.randomize.Intn(count)
}

var AllFormat = Formats{}

var AllData = Data{}

var randomGenerator *serviceProvider

var functionMap = map[string]func() string{
	"firstNameMale":   firstNameMale,
	"firstNameFemale": firstNameFemale,
	"lastName":        lastName,
	"titleMale":       titleMale,
	"titleFemale":     titleFemale,
	"suffix":          suffix,
	"freeEmailDomain": freeEmailDomain,
}

func init() {
	randomGenerator = &serviceProvider{rand.New(rand.NewSource(time.Now().UnixNano())), &sync.Mutex{}}
	AllFormat = Formats{}
	formatErr := json.Unmarshal(formats, &AllFormat)
	if formatErr != nil {
		log.Fatalln(formatErr.Error())
	}

	AllData = Data{}
	dataErr := json.Unmarshal(data, &AllData)
	if dataErr != nil {
		log.Fatalln(dataErr.Error())
	}

}

func getRandomData(source []string) string {
	formats := getRandomFormat(source)
	parseFormat := parseFormat(formats)
	return parseFormat
}

func getRandomFormat(source []string) string {
	return source[randomIndex(len(source))]
}

func getRandomElement(source []string) string {
	return source[randomIndex(len(source))]
}

func randomIndex(count int) int {
	return randomGenerator.randomIndex(count)
}

func parseFormat(format string) string {

	var response *string = &format
	matches := regex(format)

	for _, param := range matches {
		functionName := separator(param)
		oldValue := param
		newValue := functionMap[functionName]()
		*response = replaceTag(format, oldValue, newValue)
	}

	return *response
}

func regex(format string) []string {
	pattern := "{{([a-zA-Z]+)}}"
	r, _ := regexp.Compile(pattern)
	matches := r.FindAllString(format, -1)
	return matches
}

func replaceTag(content string, oldValue string, newValue string) string {
	return strings.ReplaceAll(content, oldValue, newValue)
}

func separator(content string) string {
	return strings.Replace(strings.Replace(content, "{{", "", -1), "}}", "", -1)
}

func firstNameMale() string {

	return getRandomElement(AllData.FirstNameMale)
}

func firstNameFemale() string {
	return getRandomElement(AllData.FirstNameFemale)
}

func lastName() string {
	return getRandomElement(AllData.LastName)
}

func titleMale() string {
	return getRandomElement(AllData.TitleMale)
}

func titleFemale() string {
	return getRandomElement(AllData.TitleFemale)
}

func suffix() string {
	return getRandomElement(AllData.Suffix)
}

func tld() string {
	return getRandomElement(AllData.Tld)
}

func freeEmailDomain() string {
	return getRandomElement(AllData.FreeEmailDomain)
}

func domainName() string {
	name := strings.ToLower(lastName())
	return name + "." + tld()
}
