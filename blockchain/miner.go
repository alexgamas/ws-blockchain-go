package blockchain

import (
	"fmt"
	"regexp"
)

const (
	difficulty = 6
	pattern    = "^0{%d}\\w+$"
)

var matcher *regexp.Regexp

//ValidateProof Validate a hash regard difficulty
func ValidateProof(hash string) bool {
	return matcher.MatchString(hash)
}

func init() {
	matcher, _ = regexp.Compile(fmt.Sprintf(pattern, difficulty))
}
