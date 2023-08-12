package main

import (
	"fmt"
	"regexp"
)

var p1 = fmt.Println

func main() {
	// simple match string
	reStr := "The ape was at the apex"
	match, _ := regexp.MatchString("(ape[^ ]?", reStr)
	p1(match)

	// compiled regexp
	reStr2 := "Cat rat mat fat pat"
	r, _ := regexp.Compile("([crmfp]at)")

	p1("Match string:", r.MatchString(reStr2))
	p1("Find string:", r.FindString(reStr2))
	p1("Find string index:", r.FindStringIndex(reStr2))
	p1("Find all string:", r.FindAllString(reStr2, -1))
	p1("First two string:", r.FindAllString(reStr2, 2))
	p1("All submatch string:", r.FindAllStringSubmatchIndex(reStr2, -1))
	p1(r.ReplaceAllString(reStr2, "Dog"))
}
