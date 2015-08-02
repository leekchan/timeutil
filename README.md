# timeutil - useful extensions to the golang's time package
[![Build Status](https://travis-ci.org/leekchan/timeutil.svg?branch=master)](https://travis-ci.org/leekchan/timeutil)
[![Coverage Status](https://coveralls.io/repos/leekchan/timeutil/badge.svg?branch=master&service=github)](https://coveralls.io/github/leekchan/timeutil?branch=master)
[![GoDoc](https://godoc.org/github.com/leekchan/timeutil?status.svg)](https://godoc.org/github.com/leekchan/timeutil)

timeutil provides useful extensions (Timedelta, Strftime, ...) to the golang's time package.

## Timedelta

Timedelta represents a duration between two dates. (inspired by python's timedelta)

### Timedelta struct

```Go
type Timedelta struct {
    days, seconds, microseconds, milliseconds, minutes, hours, weeks time.Duration
}
```

### Initialization

All fields are optional and default to 0. You can initialize any type of timedelta by specifying field values which you want to use.

**Examples:**

```Go
td := timeutil.Timedelta{days: 10}
td = timeutil.Timedelta{minutes: 17}
td = timeutil.Timedelta{seconds: 56}
td = timeutil.Timedelta{days: 10, minutes: 17, seconds: 56}
td = timeutil.Timedelta{days: 1, seconds: 1, microseconds: 1, milliseconds: 1, minutes: 1, hours: 1, weeks: 1}
```

### func (t *Timedelta) Duration() time.Duration

Duration() returns time.Duration. time.Duration can be added to time.Date.

**Examples:**

```Go
base := time.Date(2015, 2, 3, 0, 0, 0, 0, time.UTC)
td := timeutil.Timedelta{days: 10, minutes: 17, seconds: 56}

result := base.Add(td.Duration())
fmt.Println(result) // "2015-02-28 00:31:38 +0000 UTC"
```

### Operations

#### func (t *Timedelta) Add(t2 *Timedelta)

Add returns the Timedelta t+t2.

**Examples:**

```Go
base = time.Date(2015, 2, 3, 0, 0, 0, 0, time.UTC)

td := timeutil.Timedelta{days: 10, minutes: 17, seconds: 56}
td2 := timeutil.Timedelta{days: 15, minutes: 13, seconds: 42}
td = td.Add(&td2)

result = base.Add(td.Duration())
```

#### func (t *Timedelta) Subtract(t2 *Timedelta) Timedelta

Subtract returns the Timedelta t-t2.

**Examples:**

```Go
base = time.Date(2015, 2, 3, 0, 0, 0, 0, time.UTC)

td := timeutil.Timedelta{days: 10, minutes: 17, seconds: 56}
td2 := timeutil.Timedelta{days: 15, minutes: 13, seconds: 42}
td = td.Subtract(&td2)

result = base.Add(td.Duration())
```

#### func (t *Timedelta) Abs() Timedelta

Abs returns the absolute value of t

**Examples:**

```Go
td := timeutil.Timedelta{days: -10, minutes: -17, seconds: -56}
td = td.Abs(&td)
```


## Strftime

Strftime formats time.Date according to the directives in the given format string. The directives begins with a percent (%) character.


Directive | Meaning | Example
-------------| ------------- | -------------
%a | Weekday as locale’s abbreviated name. | Sun, Mon, ..., Sat
%A | Weekday as locale’s full name.     | Sunday, Monday, ..., Saturday 
%w | Weekday as a decimal number, where 0 is Sunday and 6 is Saturday | 0, 1, ..., 6     
%d | Day of the month as a zero-padded decimal number. | 01, 02, ..., 31 
%b | Month as locale’s abbreviated name. | Jan, Feb, ..., Dec
%B | Month as locale’s full name. | January, February, ..., December
%m | Month as a zero-padded decimal number. | 01, 02, ..., 12
%y | Year without century as a zero-padded decimal number. | 00, 01, ..., 99
%Y | Year with century as a decimal number. |   1970, 1988, 2001, 2013
%H | Hour (24-hour clock) as a zero-padded decimal number. | 00, 01, ..., 23
%I | Hour (12-hour clock) as a zero-padded decimal number. | 01, 02, ..., 12 
%p | Meridian indicator. (AM or PM.) | AM, PM
%M | Minute as a zero-padded decimal number. | 00, 01, ..., 59
%S | Second as a zero-padded decimal number. | 00, 01, ..., 59
%f | Microsecond as a decimal number, zero-padded on the left. | 000000, 000001, ..., 999999
%z | UTC offset in the form +HHMM or -HHMM | +0000
%Z | Time zone name | UTC
%j | Day of the year as a zero-padded decimal number | 001, 002, ..., 366
%U | Week number of the year (Sunday as the first day of the week) as a zero padded decimal number. All days in a new year preceding the first Sunday are considered to be in week 0. | 00, 01, ..., 53 
%W | Week number of the year (Monday as the first day of the week) as a decimal number. All days in a new year preceding the first Monday are considered to be in week 0.   | 00, 01, ..., 53
%c | Date and time representation. | Tue Aug 16 21:30:00 1988
%x | Date representation. | 08/16/88
%X | Time representation. | 21:30:00
%% | A literal '%' character. | %


## TODO

* Locale support
* Strptime - a function which returns a Date parsed according to a format string
* Auto date parser - a generic string parser which is able to parse most known formats to represent a date
* And other useful features...