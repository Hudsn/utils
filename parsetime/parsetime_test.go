package parsetime

import (
	"testing"
)

func TestTimeParsing(t *testing.T) {
	in := "August 12, 2036 @ 01:23:45.987654321 UTC"
	gotTime, err := Strptime("%B %d, %Y @ %H:%M:%S.%nano %Z", in)
	if err != nil {
		t.Fatal(err)
	}
	fmtString := "year:%Y, month:%m, day:%d, hour:%H, min:%M, sec:%S, ns:%nano, tz:%Z"
	gotStr := Strftime(fmtString, gotTime)
	wantStr := "year:2036, month:08, day:12, hour:01, min:23, sec:45, ns:987654321, tz:UTC"
	if wantStr != gotStr {
		t.Errorf("wrong value.\nExpected:\n\t%s\nGot:\n\t%s", wantStr, gotStr)
	}
}
