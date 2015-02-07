// Package tisikkirlir provides ...
package tisikkirlir

import (
	"regexp"
)

func Tsk(s string) string {
	regLowerCase, _ := regexp.Compile("[aeoöuüıi]")
	regUpperCase, _ := regexp.Compile("[AEOÖUÜI]")
	return regUpperCase.ReplaceAllString(regLowerCase.ReplaceAllString(s, "i"), "İ")
}
