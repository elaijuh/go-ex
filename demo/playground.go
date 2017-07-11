package main

import "fmt"

type condition struct {
	Operator string
	Fact     string
	Value    interface{}
	Negate   bool
}

type timeRange struct {
	start string
	end   string
}

type timeRangeFilter struct {
	Timezone string
	Ranges   []timeRange
}

type operatorFunc func(string, interface{}) (bool, error)

func operatorIn(fact string, value interface{}) (bool, error) {
	fmt.Printf("use operatorIn on %s\n", fact)
	for _, v := range value.([]string) {
		if fact == v {
			return true, nil
		}
	}
	return false, nil
}

func operatorIs(fact string, value interface{}) (bool, error) {
	fmt.Printf("use operatorIs on %s\n", fact)
	return fact == value.(string), nil
}

func operatorDefault(fact string, value interface{}) (bool, error) {
	return false, nil
}

// even if we have different operators from frontend, we can use same operator func (here operatorIn ) in backend
// in this way it's more loose coupling in terms of code structure

var opMap = map[string]operatorFunc{
	"in":             operatorIn,
	"contain":        operatorIn,
	"is":             operatorIs,
	"time_in":        operatorDefault,
	"day_of_week_in": operatorDefault,
	"default":        operatorDefault,
}

// use Fact and Operator to determin the special operator (time_in) function
// if not found then use normal operator (in) function
// if still not found use default operator (default) function

func (c *condition) getOpFunc(fact, operator string) (operatorFunc, error) {
	// combine face and operator to a new operator in operator map
	newop := fmt.Sprintf("%s_%s", fact, operator)
	fmt.Printf("Checking Operator %s in opMap\n", newop)
	var opFunc operatorFunc
	if f, ok := opMap[newop]; ok {
		opFunc = f
	} else if f, ok := opMap[operator]; ok {
		opFunc = f
	} else {
		opFunc = opMap["default"]
	}
	return opFunc, nil
}

func (c *condition) check2(fact string) (bool, error) {
	if opFunc, e := c.getOpFunc(c.Fact, c.Operator); e != nil {
		return false, e
	} else {
		return opFunc(fact, c.Value)
	}

	return false, nil
}

func main() {

	rules := []condition{
		{"in", "country", []string{"us", "uk"}, false},
		{"contain", "user_agent", "Chrome", false},
		{"in", "time", timeRangeFilter{"utc", []timeRange{{"0000", "0100"}}}, false},
		{"in", "day_in_week", []int{1, 2, 3}, false},
		{"is", "os", "ios", false},
	}

	var res bool
	r1 := rules[0]
	res, _ = r1.check2("us")
	fmt.Println(res)

	r2 := rules[4]
	res, _ = r2.check2("ios")
	fmt.Println(res)

}
