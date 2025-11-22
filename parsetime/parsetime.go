package parsetime

import (
	"fmt"
	"log"
	"strings"
	"time"
)

//python-like str(p/f)time
// ref: https://docs.python.org/3/library/datetime.html#strftime-and-strptime-format-codes
// ref https://pkg.go.dev/time#pkg-constants

// Mon Jan 2 15:04:05 MST 2006
var ParseMap = map[string]string{
	"%Y":     "2006",      // year
	"%y":     "06",        // years (2digit)
	"%d":     "02",        // day
	"%m":     "01",        // month num
	"%b":     "Jan",       // month abbrev
	"%B":     "January",   //month full
	"%H":     "15",        // hour (24 hr clock)
	"%I":     "03",        // hour (12hr clock)
	"%M":     "04",        //minute
	"%S":     "05",        //second
	"%milli": "000",       // milliseconds
	"%micro": "000000",    // microseconds
	"%nano":  "000000000", //nanosec

	"%Z": "MST",   // timezone name
	"%z": "-0700", // UTC offset
}

func Strptime(format string, input string) (time.Time, error) {
	return time.Parse(toGoLayout(format), input)
}

func Strftime(format string, t time.Time) string {
	ti, err := time.Parse(time.StampMicro, "Feb 12 01:02:03.123456")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SDFSDF: ", ti.Nanosecond())
	return t.Format(toGoLayout(format))
}

func toGoLayout(formatString string) string {
	var goLayout string = formatString
	for k, v := range ParseMap {
		goLayout = strings.ReplaceAll(goLayout, k, v)
	}
	return goLayout
}
