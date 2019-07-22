package seeder

var formats = []byte(`{
	"nameFormat": [
		"{{firstNameMale}} {{lastName}}",	
		"{{titleMale}} {{firstNameMale}} {{lastName}}",
        "{{suffix}} {{firstNameMale}} {{lastName}}",
        "{{suffix}} {{titleMale}} {{firstNameMale}} {{lastName}}",

		"{{firstNameFemale}} {{lastName}}",
		"{{titleFemale}} {{firstNameFemale}} {{lastName}}",
        "{{suffix}} {{firstNameFemale}} {{lastName}}",
        "{{suffix}} {{titleFemale}} {{firstNameFemale}} {{lastName}}"
	],
	"emailFormat": [
		"{{firstNameMale}}@{{freeEmailDomain}}",
        "{{firstNameFemale}}@{{freeEmailDomain}}",
		"{{lastName}}@{{freeEmailDomain}}"
	]
}`)
