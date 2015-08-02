package timeutil

import (
	"time"
)

func abs(v time.Duration) time.Duration {
	if v < 0 {
		v *= -1
	}
	return v
}

type Timedelta struct {
	days, seconds, microseconds, milliseconds, minutes, hours, weeks time.Duration
}

func (t *Timedelta) Add(t2 *Timedelta) {
	t.days += t2.days
	t.seconds += t2.seconds
	t.microseconds += t2.microseconds
	t.milliseconds += t2.milliseconds
	t.minutes += t2.minutes
	t.hours += t2.hours
	t.weeks += t2.weeks
}

func (t *Timedelta) Subtract(t2 *Timedelta) {
	t.days -= t2.days
	t.seconds -= t2.seconds
	t.microseconds -= t2.microseconds
	t.milliseconds -= t2.milliseconds
	t.minutes -= t2.minutes
	t.hours -= t2.hours
	t.weeks -= t2.weeks
}

func (t *Timedelta) Abs() {
	t.days = abs(t.days)
	t.seconds = abs(t.seconds)
	t.microseconds = abs(t.microseconds)
	t.milliseconds = abs(t.milliseconds)
	t.minutes = abs(t.minutes)
	t.hours = abs(t.hours)
	t.weeks = abs(t.weeks)
}

func (t *Timedelta) Duration() time.Duration {
	return t.days*24*time.Hour +
		t.seconds*time.Second +
		t.microseconds*time.Microsecond +
		t.milliseconds*time.Millisecond +
		t.minutes*time.Minute +
		t.hours*time.Hour +
		t.weeks*7*24*time.Hour
}

func (t *Timedelta) String() string {
	return t.Duration().String()
}
