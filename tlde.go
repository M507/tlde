package main

import (
	"fmt"
	"flag"
	"os"
	"io"
	"bufio"
	"strings"
	"github.com/M507/tlde/src"
)

func pipe_print_(r io.Reader, w io.Writer, subPtr1 *bool,subPtr2 *bool,rootPtr1 *bool,rootPtr2 *bool, tldPtr1 *bool,tldPtr2 *bool) error {
    scanner := bufio.NewScanner(bufio.NewReader(r))
	read_line := ""
    for scanner.Scan() {
		read_line = strings.TrimSuffix(scanner.Text(), "\r\n")
		cache := "/tmp/tld.cache"
		extract, _ := tldextract.New(cache,false)
		result := extract.Extract(read_line)
		print_(result, *subPtr1 ,*subPtr2 ,*rootPtr1 ,*rootPtr2 , *tldPtr1 ,*tldPtr2 )
    }
    return nil
}

func print_(result *tldextract.Result, subPtr1 bool,subPtr2 bool,rootPtr1 bool,rootPtr2 bool, tldPtr1 bool,tldPtr2 bool) {
	
	var result_str string
	result_str = ""

	if (len(result.Sub) > 0) {
		if (subPtr1 == true){
			result_str = result_str + result.Sub
		} else if (subPtr2 == true) {
			result_str = result_str + result.Sub
		}
	}

	if (rootPtr1 == true){
		if (subPtr1 == true){
			if (len(result.Sub) > 0) {
				result_str = result_str + "." + result.Sub
			} else {
				result_str = result_str + result.Sub
			}
		} else if (subPtr2 == true) {
			if (len(result.Sub) > 0) {
				result_str = result_str + "." + result.Sub
			} else {
				result_str = result_str + result.Sub
			}
		} else {
			fmt.Printf("%s",result.Root)
		}
	} else if (rootPtr2 == true) {
		if (subPtr1 == true){
			if (len(result.Sub) > 0) {
				result_str = result_str + "." + result.Root
			} else {
				result_str = result_str + result.Root
			}
		} else if (subPtr2 == true) {
			if (len(result.Sub) > 0) {
				result_str = result_str + "." + result.Root
			} else {
				result_str = result_str + result.Root
			}
		} else {
			result_str = result_str + result.Root
		}
	}

	if (tldPtr1 == true){
		if (rootPtr1 == true){
			result_str = result_str + "." + result.Tld
		} else if (rootPtr2 == true) {
			result_str = result_str + "." + result.Tld
		} else {
			result_str = result_str + result.Tld
		}
	} else if (tldPtr2 == true) {
		if (rootPtr1 == true){
			result_str = result_str + "." + result.Tld
		} else if (rootPtr2 == true) {
			result_str = result_str + "." + result.Tld
		} else {
			result_str = result_str + result.Tld
		}
	}
	if (len(result_str) > 0) {
		fmt.Printf("%s\n",result_str)
	}
}

func main() {
	subPtr1 := flag.Bool("sub", false, "Print subdomain")
	subPtr2 := flag.Bool("s", false, "Print subdomain")
	rootPtr1 := flag.Bool("root", false, "Print hostname")
	rootPtr2 := flag.Bool("r", false, "Print hostname")
	tldPtr1 := flag.Bool("tld", false, "Print Top-Level Domain")
	tldPtr2 := flag.Bool("t", false, "Print Top-Level Domain")

	fi, err := os.Stdin.Stat()
	if err != nil {
	  panic(err)
	}
	if fi.Mode() & os.ModeNamedPipe == 0 {
		wordPtr := flag.String("u", "http://www1.web.google.com.sh/", "URL")
		flag.Parse()
		cache := "/tmp/tld.cache"
		extract, _ := tldextract.New(cache,false)
		result := extract.Extract(*wordPtr)	
		print_(result, *subPtr1 ,*subPtr2 ,*rootPtr1 ,*rootPtr2 , *tldPtr1 ,*tldPtr2 )
	} else {
		flag.Parse()
		pipe_print_(os.Stdin, os.Stdout, subPtr1, subPtr2, rootPtr1, rootPtr2, tldPtr1, tldPtr2)
	}
}