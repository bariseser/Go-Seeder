package seeder

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"
)

type Formats struct {
	Name []string `json:"nameFormat"`
}

type Data struct {
	FirstNameMale   []string `json:"firstNameMale"`
	FirstNameFemale []string `json:"firstNameFemale"`
	LastName        []string `json:"lastName"`
	TitleMale       []string `json:"titleMale"`
	TitleFemale     []string `json:"titleFemale"`
	Suffix          []string `json:"suffix"`
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
	fmt.Println(formats)
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

	formatPattern := strings.Fields(format)

	for _, param := range formatPattern {
		oldValue := param
		newValue := functionMap[param]()
		*response = replaceTag(format, oldValue, newValue)
	}

	return *response
}

func replaceTag(content string, oldValue string, newValue string) string {
	return strings.ReplaceAll(content, oldValue, newValue)
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
