package util_jsoncrack

import "github.com/fwhezfwhez/jsoncrack"

var Jc jsoncrack.JsonCracker
func init(){
	Jc =jsoncrack.NewCracker(nil)
}
