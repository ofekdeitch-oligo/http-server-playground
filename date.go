package main

type Instant struct {
	milliseconds int
}

type Duration struct {
	value int
}

type NewDurationOptions struct {
	seconds int
	minutes int
}

// func NewDuration(options NewDurationOptions) *Duration {
// 	milliseconds := 0

// 	if options.seconds != nil {
// 		milliseconds += options.seconds * 60
// 	}

// 	duration := new
// }
