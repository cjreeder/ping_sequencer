package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Reading File and Getting ping sequence")
	var filepath string
	var order []int
	var num int

	flag.StringVar(&filepath, "opt", "", "Usage")
	flag.Parse()
	fmt.Println(filepath)

	if filepath == "" {
		fmt.Sprintf("Please put in your files!")
	} else {
		content, err := ioutil.ReadFile(filepath)
		if err != nil {
			fmt.Sprint("Filed to read the file you provided")
		}
		s := string([]byte(content))
		//fmt.Printf(s)
		re := regexp.MustCompile(`(?m)icmp_seq=\d*`)
		//re := regexp.MustCompile(`(?m)/icmp_seq=`)
		//fmt.Printf("%q\n", re.FindString(s))
		//fmt.Printf("%q\n", re.FindString(str))

		for _, match := range re.FindAllString(s, -1) {
			//fmt.Println(match, "is at index", i)
			newlist := strings.Split(match, "=")
			fmt.Println(newlist[1])
			num, err = strconv.Atoi(newlist[1])
			if err != nil {
				fmt.Sprintf("Failed!")
			}
			order = append(order, num)

		}
		//fmt.Printf("Here is the final array: %v", order)
		for i := 0; i < len(order)-1; i++ {
			//fmt.Printf("Testings %v\n", order[i])
			if order[i]+1 != order[i+1] {
				fmt.Printf("Ordering is missing at this point: %v\n", order[i])
				fmt.Printf("The current ping is %v and the next ping is %v\n", order[i], order[i+1])
				diff := order[i+1] - order[i]
				fmt.Printf("The Difference between them is %v\n", diff)
			}
		}
		//for i, next := range order {
		//	if i++ == next++ {
		//		fmt.Println("You're all good")
		//	}
		//}
	}
}
