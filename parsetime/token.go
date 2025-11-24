package parsetime

type tokenType int

const (
	_ tokenType = iota
	illegal
	eof
	literal
	directive
)

type token struct {
	tokenType tokenType
	start     int
	end       int
	value     string
}

//

// use similar to Ruby's strftime
// https://ruby-doc.org/core-2.6.4/Time.html
// 	"N", //fractional seconds in digits 9 or fewer digits, representing nanosec, can be %3N (milli), %6 (micro), %9N (nano or larger)
// 	"L", // millisecond
// 	"S", //  01, 02, 25, 99

// 	// year
// 	"Y", // 2025
// 	"y", // 25

// 	// month
// 	"m", // 01, 02
// 	"b", // jan, feb, etc...
// 	"B", // january, february, etc...

// 	// day
// 	"d", // 01, 02
// }
