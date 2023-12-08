package validator

import (
	"regexp"
	"unicode/utf8"
)

/*
when user submit an input for logging or sign up
it is good idea to validate the data.

The following check would be performed

1. Check that the provided name, email address and password are not blank.
2. Sanity check the format of the email address.
3. Ensure that the password is at least 8 characters long.
4. Make sure that the email address isn’t already in use.

*/

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// return true if a value contain at least n characters
func MinChars(value string, n int) bool {
	return utf8.RuneCountInString(value) >= n

}

// returns true if a value matches a provided compiled
// regular expression pattern
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}
