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

func (t *Timedelta) Add(t2 *Timedelta) Timedelta {
	return Timedelta{
		days:         t.days + t2.days,
		seconds:      t.seconds + t2.seconds,
		microseconds: t.microseconds + t2.microseconds,
		milliseconds: t.milliseconds + t2.milliseconds,
		minutes:      t.minutes + t2.minutes,
		hours:        t.hours + t2.hours,
		weeks:        t.weeks + t2.weeks,
	}
}

func (t *Timedelta) Subtract(t2 *Timedelta) Timedelta {
	return Timedelta{
		days:         t.days - t2.days,
		seconds:      t.seconds - t2.seconds,
		microseconds: t.microseconds - t2.microseconds,
		milliseconds: t.milliseconds - t2.milliseconds,
		minutes:      t.minutes - t2.minutes,
		hours:        t.hours - t2.hours,
		weeks:        t.weeks - t2.weeks,
	}
}

func (t *Timedelta) Abs() Timedelta {
	return Timedelta{
		days:         abs(t.days),
		seconds:      abs(t.seconds),
		microseconds: abs(t.microseconds),
		milliseconds: abs(t.milliseconds),
		minutes:      abs(t.minutes),
		hours:        abs(t.hours),
		weeks:        abs(t.weeks),
	}
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
