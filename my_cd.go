package main

import (
	"fmt"
	"os"
	"strings"
)

func available(a []string, c string) bool {
	var result bool = false
	for _, k := range a {
		if c == k {
			fmt.Println(k, c)
			result = true
			break
		}
	}
	return result
}

func main() {
	var current string
	var new string
	var joined string
	//var prev string
	if len(os.Args) <= 1 {
		fmt.Println("[ARGUMENT ERROR] mycd takes only two input paths.")
		os.Exit(1)
	}
	if len(os.Args) > 3 {
		fmt.Println("[ARGUMENT ERROR] mycd takes only two input paths.")
		os.Exit(1)
	}
	current = os.Args[1]
	new = os.Args[2]

	curs := strings.Split(current, "/")
	news := strings.Split(new, "/")
	temp := curs
	//err := false
	//curl := len(news)

	for _, m := range news {
		if !available(curs, m) && m != ".." && m != "" {
			temp = append(temp, m) // only for appending the ones with the / tag
		} else if !available(curs, m) && m == ".." && m != "" {
			temp = temp[:len(temp)-1] // for going back a directory behind
		}
	}

	for _, k := range temp {
		joined += k + "/" //joining the directories
	}

	if len(joined) == 0 {
		joined = "/"
	}

	fmt.Println(joined) //finally printing the directory
}
