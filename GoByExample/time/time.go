// Go offers extensive support for times and durations;
// here are some examples.

package main

import "fmt"
import "time"
import . "zhoushxGoDemo/GoByExample/fileline"

func main() {
	p := fmt.Println

	// We'll start by getting the current time.
	now := time.Now()
	p(FileLine(), now)

	// You can build a `time` struct by providing the
	// year, month, day, etc. Times are always associated
	// with a `Location`, i.e. time zone.
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(FileLine(), then)

	// You can extract the various components of the time
	// value as expected.
	p(FileLine(), then.Year())
	p(FileLine(), then.Month())
	p(FileLine(), then.Day())
	p(FileLine(), then.Hour())
	p(FileLine(), then.Minute())
	p(FileLine(), then.Second())
	p(FileLine(), then.Nanosecond())
	p(FileLine(), then.Location())

	// The Monday-Sunday `Weekday` is also available.
	p(FileLine(), then.Weekday())

	// These methods compare two times, testing if the
	// first occurs before, after, or at the same time
	// as the second, respectively.
	p(FileLine(), then.Before(now))
	p(FileLine(), then.After(now))
	p(FileLine(), then.Equal(now))

	// The `Sub` methods returns a `Duration` representing
	// the interval between two times.
	diff := now.Sub(then)
	p(FileLine(), diff)

	// We can compute the length of the duration in
	// various units.
	p(FileLine(), "diff.Hour=", diff.Hours())
	p(FileLine(), "diff.Minutes=", diff.Minutes())
	p(FileLine(), "diff.Seconds=", diff.Seconds())
	p(FileLine(), "diff.Nanoseconds=", diff.Nanoseconds())

	// You can use `Add` to advance a time by a given
	// duration, or with a `-` to move backwards by a
	// duration.
	p(FileLine(), then.Add(diff))
	p(FileLine(), then.Add(-diff))
}
