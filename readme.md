## createDay script

This script can be used to create the template for how 
I've been solving the puzzles to start a new day.

This script uses the following environment variables:
- `ADVENT_COOKIE` the session cookie used to fetch the correct data
- `ADVENT_YEAR` the default year to use if none is provided

The script takes a day and an optional year in order to create
the file and get the correct data. If no year is provided it uses the value
set by `ADVENT_YEAR` env variable.

This script creates the year/day directories and copies across the `main.go` and
`main_test.go` templates as defined in `./templates`.

