# Dates

This is dates,
a command-line tool for doing date arithmetic.

## Installation

```sh
go install github.com/bobg/dates@latest
```

## Usage

In all of the following cases,
a date argument can be the string “today”
to mean today’s date.

```sh
dates add DATE N
```

Show the date resulting from adding N dates to the given DATE.
N can be negative.

```sh
dates delta DATE1 DATE2
```

Show the number of days from DATE1 to DATE2.
The result is positive if DATE1 is before DATE2
and negative if DATE1 is after DATE2.

```sh
dates since DATE
```

Show the number of days since DATE.
This is the same as `dates delta DATE today`.
