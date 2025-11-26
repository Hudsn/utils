package parsetime

import "time"

type directiveCategory string

const (
	dcat_date    = "date"
	dcat_time    = "time"
	dcat_year    = "timezone"
	dcat_weekday = "weekday"
)

var directiveCategoryBaseList []directiveCategory = []directiveCategory{
	dcat_date, dcat_time, dcat_year, dcat_weekday,
}

type directiveRegistry struct {
	directiveMap  map[rune]*directiveInfo
	directiveList []*directiveInfo
}

type directiveInfo struct {
	useNumericModifiers bool
	useTextModifiers    bool
	char                rune
	overrideModifiers   []rune
	category            directiveCategory
	parseFn             func(input string) (time.Time, error)
	formatFn            func(input time.Time) (string, error)
}

var directivePrefixesNumeric []rune = []rune{
	'-', // dont pad
	'_', // use spaces for paddin
	'0', // use zeros for padding
}
var directivePrefixesText []rune = []rune{
	'^', // uppercase
	'#', // swap case
}

var directiveCharList = []rune{'Y', 'y', 'm', 'b', 'B', 'd', 'e', 'k', 'I', 'l', 'N'}

func newDirectiveRegistry() *directiveRegistry {
	ret := &directiveRegistry{
		directiveMap:  make(map[rune]*directiveInfo),
		directiveList: []*directiveInfo{},
	}
	ret.generateDefaultRegistry()
	return ret
}

func (dr *directiveRegistry) addInfo(di *directiveInfo) {
	dr.directiveMap[di.char] = di
	dr.directiveList = append(dr.directiveList, di)
}

func (dr *directiveRegistry) generateDefaultRegistry() {
	add := &directiveInfo{
		useNumericModifiers: true,
		useTextModifiers:    false,
		char:                'Y',
		overrideModifiers:   nil,
		category:            dcat_year,
	}
	dr.addInfo(add)
}

//

// use similar to Ruby's strftime
// https://ruby-doc.org/core-2.6.4/Time.html

// 	// year
// 	"Y", // 2025
// 	"y", // 25

// 	// month
// 	"m", // 01, 02
// 	"b", // jan, feb, etc...
// 	"B", // january, february, etc...

// 	// day
// 	"d", // 01, 02
// "e", // 1, 2, 3
// }

// hour
// %H - Hour of the day, 24-hour clock, zero-padded (00..23)
//   %k - Hour of the day, 24-hour clock, blank-padded ( 0..23)
//   %I - Hour of the day, 12-hour clock, zero-padded (01..12)
//   %l - Hour of the day, 12-hour clock, blank-padded ( 1..12)

// min and sec
// %M - Minute of the hour (00..59)

//   %S - Second of the minute (00..60)

// subsecond
// 	"N", //fractional seconds in digits 9 or fewer digits, representing nanosec, can be %3N (milli), %6 (micro), %9N (nano or larger)

//   %P - Meridian indicator, lowercase (``am'' or ``pm'')
//   %p - Meridian indicator, uppercase (``AM'' or ``PM'')

// timezone
// %z - Time zone as hour and minute offset from UTC (e.g. +0900)
//   %:z - hour and minute offset from UTC with a colon (e.g. +09:00)
//   %::z - hour, minute and second offset from UTC (e.g. +09:00:00)

//  %A - The full weekday name (``Sunday'')
//  %^A  uppercased (``SUNDAY'')
//  %a - The abbreviated name (``Sun'')
//  %^a  uppercased (``SUN'')
