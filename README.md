# GoSeeder

GoSeeder is a Go library that generates fake data for you. Whether you need to bootstrap your database, create good-looking XML documents, fill-in your persistence to stress test it, or anonymize data taken from a production service.

GoSeeder is heavily inspired by fzaninotto's [fzaninotto/Faker](https://github.com/fzaninotto/Faker)

[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/bariseser/Go-Seeder/issues)
[![GoDoc](https://godoc.org/github.com/Pallinder/go-randomdata?status.svg)](https://godoc.org/github.com/bariseser/Go-Seeder)
[![Go Report Card](https://goreportcard.com/badge/github.com/bariseser/Go-Seeder)](https://goreportcard.com/report/github.com/bariseser/Go-Seeder)
[![Build Status](https://travis-ci.org/bariseser/Go-Seeder.svg?branch=master)](https://travis-ci.org/bariseser/Go-Seeder)

# Table of Contents

- [Installation](#installation)
- [Basic Usage](#basic-usage)
- [Formatters](#formatters)
	- [Name](#name)
	- [Internet](#internet)
- [Contributors](#contributors)
- [License](#license)


## Installation

```go get github.com/bariseser/Go-Seeder```

## Basic Usage

### Name
````go

package main

import (
    "fmt"
    "github.com/bariseser/Go-Seeder"
)

func main() {
	//Earnest Abshire
	//PhD Mrs Patrick Emard
	seeder.Name()
	
	// Randall
	seeder.FirstNameMale()
	
     //marcellus@yahoo.com
    seeder.Email()
    
    //wilderman.info
    seeder.Domain()
	
	//Daisha
	seeder.FirstNameFemale()
	
	//Abernathy
	seeder.LastName()
	
	//Mr
	seeder.TitleMale()
	
	//Miss
	seeder.TitleFemale()
	
	//PhD
	seeder.Suffix()
}
````

## Formatters

### Name

    - Name
    - FirstNameMale
    - FirstNameFemale
    - LastName
    - titleMale
    - TitleFemale
    - Suffix
    
### Internet
    
    - Email
    - Domain
     
### Contributors
* [Bariseser](https://github.com/bariseser)

##License

GoSeeder is released under the MIT Licence. See the bundled LICENSE file for details.