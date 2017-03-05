// A common requirement in programs is getting the number
// of seconds, milliseconds, or nanoseconds since the
// [Unix epoch](http://en.wikipedia.org/wiki/Unix_time).
// Here's how to do it in Go.

package main

import "fmt"
import "time"
import . "zhoushxGoDemo/GoByExample/fileline"

func main() {

	// Use `time.Now` with `Unix` or `UnixNano` to get
	// elapsed time since the Unix epoch in seconds or
	// nanoseconds, respectively.
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(FileLine(), now)

	// Note that there is no `UnixMillis`, so to get the
	// milliseconds since epoch you'll need to manually
	// divide from nanoseconds.
	millis := nanos / 1000000
	fmt.Println(FileLine(), secs)
	fmt.Println(FileLine(), millis)
	fmt.Println(FileLine(), nanos)

	// You can also convert integer seconds or nanoseconds
	// since the epoch into the corresponding `time`.
	fmt.Println(FileLine(), time.Unix(secs, 0))
	fmt.Println(FileLine(), time.Unix(0, nanos))
}