package consts

import (
	"regexp"
)

const httpMarReg = "(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|].mar"
const localMarReg = ".*.mar"

var HttpMarRegexp *regexp.Regexp
var LocalMarRegexp *regexp.Regexp

func Init() {
	HttpMarRegexp = regexp.MustCompile(httpMarReg)
	LocalMarRegexp = regexp.MustCompile(localMarReg)
}
