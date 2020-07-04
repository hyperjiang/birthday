# birthday

[![GoDoc](https://godoc.org/github.com/hyperjiang/birthday?status.svg)](https://pkg.go.dev/github.com/hyperjiang/birthday?tab=doc)
[![Build Status](https://travis-ci.org/hyperjiang/birthday.svg?branch=master)](https://travis-ci.org/hyperjiang/birthday)
[![](https://goreportcard.com/badge/github.com/hyperjiang/birthday)](https://goreportcard.com/report/github.com/hyperjiang/birthday)
[![codecov](https://codecov.io/gh/hyperjiang/birthday/branch/master/graph/badge.svg)](https://codecov.io/gh/hyperjiang/birthday)
[![Release](https://img.shields.io/github/release/hyperjiang/birthday.svg)](https://github.com/hyperjiang/birthday/releases)

Parse a birthday to get the age and constellation.

## Requirements
- Go 1.13+
- Go Modules

## Usage

```
import "github.com/hyperjiang/birthday"

b, _ := birthday.Parse("2020-01-01")

fmt.Println(b.Age)
fmt.Println(b.Constellation)
fmt.Println(b.GetConstellation("zh")) // print constellation in Chinese
fmt.Println(b.Format("Jan 2nd, 2006"))
fmt.Println(b.String())
```
