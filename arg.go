package main

import (
	"flag"
	"strings"

	"github.com/rickb777/date/v2"
)

type dateArgType struct {
	date.Date
}

var _ flag.Value = &dateArgType{}

func (d *dateArgType) Set(s string) error {
	if strings.EqualFold(s, "today") {
		d.Date = date.Today()
		return nil
	}

	parsed, err := date.AutoParseUS(s)
	if err != nil {
		return err
	}
	d.Date = parsed
	return nil
}

func (d *dateArgType) Copy() flag.Value {
	return &dateArgType{Date: d.Date}
}
