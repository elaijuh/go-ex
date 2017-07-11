package main

type TimeRangeFilter struct {
	Timezone string
	Ranges   []struct{ start, end string }
}
