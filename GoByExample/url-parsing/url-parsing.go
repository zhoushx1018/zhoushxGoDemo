// URLs provide a [uniform way to locate resources](http://adam.heroku.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/).
// Here's how to parse URLs in Go.

package main

import "fmt"
import "net"
import "net/url"
import . "zhoushxGoDemo/GoByExample/fileline"

func main() {

	// We'll parse this example URL, which includes a
	// scheme, authentication info, host, port, path,
	// query params, and query fragment.
	s := "postgres://user:pass@host.com:5432/path?key=value1&key=value2#f"

	// Parse the URL and ensure there are no errors.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// Accessing the scheme is straightforward.
	fmt.Println(FileLine(), u.Scheme)

	// `User` contains all authentication info; call
	// `Username` and `Password` on this for individual
	// values.
	fmt.Println(FileLine(), "user=", u.User)
	fmt.Println(FileLine(), "UserName=", u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(FileLine(), "passwd=", p)

	// The `Host` contains both the hostname and the port,
	// if present. Use `SplitHostPort` to extract them.
	fmt.Println(FileLine(), "urlHost=", u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(FileLine(), "host=", host)
	fmt.Println(FileLine(), "port=", port)

	// Here we extract the `path` and the fragment after
	// the `#`.
	fmt.Println(FileLine(), "Path=", u.Path)
	fmt.Println(FileLine(), "Fragment=", u.Fragment)

	// To get query params in a string of `k=v` format,
	// use `RawQuery`. You can also parse query params
	// into a map. The parsed query param maps are from
	// strings to slices of strings, so index into `[0]`
	// if you only want the first value.
	fmt.Println(FileLine(), "RawQuery=", u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(FileLine(), "m=", m)
	fmt.Println(FileLine(), "m[key][0]=", m["key"][0])
	fmt.Println(FileLine(), "m[key][1]=", m["key"][1])
}
