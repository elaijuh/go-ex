package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"
)

var stringFlag *string
var intFlag int

func init() {
	// Example 1: Defind a string flag named "stringFlag"
	stringFlag = flag.String("stringFlag", "defaultStringFlag", "String Flag")

	// Example 2: Bind the flag to a variable
	flag.IntVar(&intFlag, "intFlag", 123, "Int Flag")

}

type interval []time.Duration

func (i *interval) String() string {
	return fmt.Sprint(*i)
}

func (i *interval) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("interval flag already set")
	}
	for _, dt := range strings.Split(value, ",") {
		duration, err := time.ParseDuration(dt)
		if err != nil {
			return err
		}
		*i = append(*i, duration)
	}
	return nil
}

var intervalFlag interval

func init() {

	// Example 3: Bind the flag to a cutomer type which implements flag.Value
	// flag.Value has set() method
	flag.Var(&intervalFlag, "delta", "Time Duration")

}

func main() {
	flag.Parse()

	// Print flags after they are processed
	fmt.Printf("%s, %d, %s\n", *stringFlag, intFlag, intervalFlag)

	// Print arguments which are not flagged
	fmt.Println(flag.Args())

}
