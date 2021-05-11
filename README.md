## Golang Lichess Package

[![Go Report Card](https://goreportcard.com/badge/github.com/chrisbarnes2000/lichess)](https://goreportcard.com/report/github.com/chrisbarnes2000/lichess)
[![GitHub](https://img.shields.io/github/forks/ChrisBarnes2000/lichess.svg?style=flat-square)](https://github.com/ChrisBarnes2000/lichess/network)
[![GitHub](https://img.shields.io/github/issues/ChrisBarnes2000/lichess.svg?style=flat-square)](https://github.com/ChrisBarnes2000/lichess/issues)

<!-- [![NPM version](https://img.shields.io/npm/v/chris-barnes-lichess.svg?style=flat)](https://www.npmjs.com/package/chris-barnes-lichess) -->
<!-- [![NPM downloads](https://img.shields.io/npm/dm/chris-barnes-lichess.svg?style=flat)](https://npmjs.org/package/chris-barnes-lichess) -->
<!-- [![Build Status](https://img.shields.io/travis/ChrisBarnes2000/chris-barnes-lichess.svg?style=flat)](https://travis-ci.org/ChrisBarnes2000/chris-barnes-lichess) -->

A Go implementation of [Lichess](https://lichess.org)'s [API](https://lichess.org/api)

### Usage

```go
package main

import (
 "fmt"
 "github.com/chrisbarnes2000/lichess"
 "log"
 "net/http"
 "net/url"
)

func main() {
 baseURL, err := url.Parse("https://lichess.org")
 if err != nil {
  log.Fatal(err)
 }
 client := new(lichess.Client)

 client.BaseURL = baseURL
 client.APIKey = "<API Key>"
 client.UserAgent = "Golang Client"
 client.HttpClient = new(http.Client)

 user, err := client.GetProfile()
 if err != nil {
  log.Fatal(err)
 }
 fmt.Println(user.ID)
}
```

### Available Features

- [x] Accounts
- [X] Users
- [X] Unit Testing
- [X] Games
- [X] Board
- [X] User IDs
- [ ] Teams
- [ ] Tournaments \
    Ideas For Tournament Main Code

    Afetr setting things up for the Chess club at MakeSchool to compete and learn from eachother I have noticed these features are missing from the tournament sections of lichess.org and will be looking into how one mighth add these

    1. Ability to Kick/Pause Members
    2. Ability to Invite people you follow
    3. Ability to Pick who you compete against
    4. Ability to Change How Matches are paird up
    5. Ability to Link to Twitter/Twitch/Discord/Slack Group Channel for voice/disscussion
    6. Dispay Tournaments you've joined via quick button on mobile and desktop versions

- [ ] Relations
- [ ] Challenges
- [ ] Chessbot

## 💻 Local Development

## Getting Started

- First, Create an API Token from [here](https://lichess.org/account/oauth/token/create)
- Second, [fork this repo](https://github.com/chrisbarnes2000/lichess/fork),
- Third, run these commands to clone it locally and get started

```zsh
# Clone and CD into/Open this project
$ git clone git@github.com:YOUR_GITHUB_USERNAME/lichess.git && cd lichess
# Create a .env file to hold our secrets
$ touch .env > LICHESS_TOKEN=YOUR_API_TOKEN
# Download & Install the dependancies. Then Compile the program
$ go build
# Run the program locally
$ go run lichess.go
```

## 📚 Resources

- [Go Template CheeatSheet](https://curtisvermeeren.github.io/2017/09/14/Golang-Templates-Cheatsheet)
- [Intro to go testing](https://tutorialedge.net/golang/intro-testing-in-go/)
- [Intro to go benchmark testing](https://tutorialedge.net/golang/benchmarking-your-go-programs/)
- [Go: tests with HTML coverage report](https://kenanbek.medium.com/go-tests-with-html-coverage-report-f977da09552d)
- [Makefiles for Go Developers](https://tutorialedge.net/golang/makefiles-for-go-developers/)
- [Runes in golang](https://www.geeksforgeeks.org/rune-in-golang/)
- [Default Vaules for Struct Fields in golang](https://www.geeksforgeeks.org/how-to-assign-default-value-for-struct-field-in-golang/)
