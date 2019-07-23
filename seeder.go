package seeder

import "strings"

func Name() string {
	return getRandomData(AllFormat.Name)
}

func FirstNameMale() string {
	return getRandomElement(AllData.FirstNameMale)
}

func FirstNameFemale() string {
	return getRandomElement(AllData.FirstNameFemale)
}

func LastName() string {
	return getRandomElement(AllData.LastName)
}

func TitleMale() string {
	return getRandomElement(AllData.TitleMale)
}

func TitleFemale() string {
	return getRandomElement(AllData.TitleFemale)
}

func Suffix() string {
	return getRandomElement(AllData.Suffix)
}

func Email() string {
	return strings.ToLower(getRandomData(AllFormat.Email))
}

func Domain() string {
	return domainName()
}
