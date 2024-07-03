package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/bobg/subcmd/v2"
	"github.com/rickb777/date/v2"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	return subcmd.Run(context.Background(), maincmd{}, os.Args[1:])
}

type maincmd struct{}

func (maincmd) Subcmds() subcmd.Map {
	dflt := &dateArgType{Date: date.Today()}

	return subcmd.Commands(
		"add", doAdd, "add days to a date", subcmd.Params(
			"date", subcmd.Value, dflt, "the date",
			"days", subcmd.Int, 0, "number of days to add",
		),
		"delta", doDelta, "compute the difference between two dates", subcmd.Params(
			"date1", subcmd.Value, dflt, "first date",
			"date2", subcmd.Value, dflt, "second date",
		),
		"since", doSince, "compute the number of days since a date", subcmd.Params(
			"date", subcmd.Value, dflt, "the date",
		),
	)
}

func doAdd(_ context.Context, dv flag.Value, days int, _ []string) error {
	d, ok := dv.(*dateArgType)
	if !ok {
		return fmt.Errorf("unexpected type %T for date", dv)
	}
	dd := d.Date.AddDate(0, 0, days)
	fmt.Println(dd)
	return nil
}

func doDelta(_ context.Context, date1v, date2v flag.Value, _ []string) error {
	date1, ok := date1v.(*dateArgType)
	if !ok {
		return fmt.Errorf("unexpected type %T for date1", date1v)
	}
	date2, ok := date2v.(*dateArgType)
	if !ok {
		return fmt.Errorf("unexpected type %T for date2", date2v)
	}

	delta := int64(date2.Date) - int64(date1.Date)
	fmt.Println(delta)
	return nil
}

func doSince(_ context.Context, dv flag.Value, _ []string) error {
	d, ok := dv.(*dateArgType)
	if !ok {
		return fmt.Errorf("unexpected type %T for date", dv)
	}
	delta := int64(date.Today()) - int64(d.Date)
	fmt.Println(delta)
	return nil
}
