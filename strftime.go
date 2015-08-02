package timeutil

import (
	"fmt"
	"time"
)

var longDayNames = []string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

var shortDayNames = []string{
	"Sun",
	"Mon",
	"Tue",
	"Wed",
	"Thu",
	"Fri",
	"Sat",
}

var shortMonthNames = []string{
	"---",
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec",
}

var longMonthNames = []string{
	"---",
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

func weekNumber(t *time.Time, char int) int {
	weekday := int(t.Weekday())

	if char == 'W' {
		// Monday as the first day of the week
		if weekday == 0 {
			weekday = 6
		} else {
			weekday -= 1
		}
	}

	return (t.YearDay() + 6 - weekday) / 7
}

// Strftime formats time.Date according to the directives in the given format string. The directives begins with a percent (%) character.
func Strftime(t *time.Time, format string) string {
	var result string

	for i := 0; i < len(format); i++ {
		switch format[i] {
		case '%':
			if i < len(format)-1 {
				switch format[i+1] {
				case 'a':
					result += shortDayNames[t.Weekday()]
				case 'A':
					result += longDayNames[t.Weekday()]
				case 'w':
					result += fmt.Sprintf("%d", t.Weekday())
				case 'd':
					result += fmt.Sprintf("%02d", t.Day())
				case 'b':
					result += shortMonthNames[t.Month()]
				case 'B':
					result += longMonthNames[t.Month()]
				case 'm':
					result += fmt.Sprintf("%02d", t.Month())
				case 'y':
					result += fmt.Sprintf("%02d", t.Year()%100)
				case 'Y':
					result += fmt.Sprintf("%02d", t.Year())
				case 'H':
					result += fmt.Sprintf("%02d", t.Hour())
				case 'I':
					if t.Hour() == 0 {
						result += fmt.Sprintf("%02d", 12)
					} else if t.Hour() > 12 {
						result += fmt.Sprintf("%02d", t.Hour()-12)
					} else {
						result += fmt.Sprintf("%02d", t.Hour())
					}
				case 'p':
					if t.Hour() < 12 {
						result += "AM"
					} else {
						result += "PM"
					}
				case 'M':
					result += fmt.Sprintf("%02d", t.Minute())
				case 'S':
					result += fmt.Sprintf("%02d", t.Second())
				case 'f':
					result += fmt.Sprintf("%06d", t.Nanosecond()/1000)
				case 'z':
					result += t.Format("-0700")
				case 'Z':
					result += t.Location().String()
				case 'j':
					result += fmt.Sprintf("%03d", t.YearDay())
				case 'U':
					result += fmt.Sprintf("%02d", weekNumber(t, 'U'))
				case 'W':
					result += fmt.Sprintf("%02d", weekNumber(t, 'W'))
				case 'c':
					result += t.Format("Mon Jan 2 15:04:05 2006")
				case 'x':
					result += fmt.Sprintf("%02d/%02d/%02d", t.Month(), t.Day(), t.Year()%100)
				case 'X':
					result += fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
				case '%':
					result += "%"
				}
				i += 1
			}
		default:
			result += fmt.Sprintf("%c", format[i])
		}
	}

	return result
}
