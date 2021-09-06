package main

import (
	"fmt"
	"flag"
	"github.com/M507/tlde/src"
)


func main() {
	rootPtr1 := flag.Bool("root", false, "a bool")
	rootPtr2 := flag.Bool("r", false, "a bool")
	subPtr1 := flag.Bool("sub", false, "a bool")
	subPtr2 := flag.Bool("s", false, "a bool")
	tldPtr1 := flag.Bool("tld", false, "a bool")
	tldPtr2 := flag.Bool("t", false, "a bool")
	wordPtr := flag.String("u", "foo", "a string")
	flag.Parse()

	cache := "/tmp/tld.cache"
	extract, _ := tldextract.New(cache,false)

	result:=extract.Extract(*wordPtr)
	if (*subPtr1 == true){
		fmt.Printf("%s",result.Sub)
	} else if (*subPtr2 == true) {
		fmt.Printf("%s",result.Sub)
	}

	if (*rootPtr1 == true){
		fmt.Printf("%s",result.Root)
	} else if (*rootPtr2 == true) {
		fmt.Printf("%s",result.Root)
	}

	if (*tldPtr1 == true){
		fmt.Printf("%s",result.Tld)
	} else if (*tldPtr2 == true) {
		fmt.Printf("%s",result.Tld)
	}
	fmt.Printf("\n")
}