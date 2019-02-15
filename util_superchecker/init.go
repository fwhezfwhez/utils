package util_superchecker

import "github.com/fwhezfwhez/superChecker"

var Spc *superChecker.Checker

func init() {
	Spc = superChecker.GetChecker()
}
