package main

import (
	"github.com/vhvb1989/survey/v2"
)

var value = ""

var table = []TestUtil.TestTableEntry{
	{
		"standard", &survey.Password{Message: "Please type your password:"}, &value, nil,
	},
	{
		"please make sure paste works", &survey.Password{Message: "Please paste your password:"}, &value, nil,
	},
	{
		"no help, send '?'", &survey.Password{Message: "Please type your password:"}, &value, nil,
	},
}

func main() {
	TestUtil.RunTable(table)
}
