package timeutil

import (
	"reflect"
	"testing"
	"time"
)

func AssertEqual(t *testing.T, x, y interface{}) {
	if !reflect.DeepEqual(x, y) {
		t.Error("Expected ", y, ", got ", x)
	}
}

func TestTimedelta(t *testing.T) {
	base := time.Date(1980, 1, 6, 0, 0, 0, 0, time.UTC)
	result := base.Add((&Timedelta{days: 1, seconds: 66355, weeks: 1722}).Duration())
	AssertEqual(t, result.String(), "2013-01-07 18:25:55 +0000 UTC")

	result = result.Add((&Timedelta{microseconds: 3, milliseconds: 10, minutes: 1}).Duration())
	AssertEqual(t, result.String(), "2013-01-07 18:26:55.010003 +0000 UTC")

	td := Timedelta{days: 10, minutes: 17, seconds: 56}
	td2 := Timedelta{days: 15, minutes: 13, seconds: 42}
	td.Add(&td2)

	base = time.Date(2015, 2, 3, 0, 0, 0, 0, time.UTC)
	result = base.Add(td.Duration())
	AssertEqual(t, result.String(), "2015-02-28 00:31:38 +0000 UTC")

	td.Subtract(&td2)

	result = base.Add(td.Duration())
	AssertEqual(t, result.String(), "2015-02-13 00:17:56 +0000 UTC")

	AssertEqual(t, td.String(), "240h17m56s")

	td = Timedelta{days: -1, seconds: -1, microseconds: -1, milliseconds: -1, minutes: -1, hours: -1, weeks: -1}
	td2 = td
	td.Abs()
	td.Add(&td2)
	AssertEqual(t, td.String(), "0")
}
