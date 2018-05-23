package main

import (
	"flag"
	"fmt"
	"goNotes/utils"
	"os"
)

func main() {
	leninline := 0
	linenum := 0
	filepath := ""

	flag.IntVar(&leninline, "leninline", 64, "len in a line")
	flag.IntVar(&linenum, "linenum", 16, "line num")
	flag.StringVar(&filepath, "filepath", "file.txt", "filepath")
	flag.Parse()

	file, err := os.Create(filepath)
	defer file.Close()

	if nil != err {
		fmt.Println(err)
		return
	}

	for l := 0; l < linenum; l++ {
		header := fmt.Sprintf("%04d * %04d : ", l, leninline)
		_, err := file.WriteString(header)
		if nil != err {
			fmt.Println(err)
			return
		}

		alphaString, err := utils.AlphaString(leninline - len(header) - 1)

		if nil != err {
			fmt.Println(err)
			return
		}
		fmt.Println(alphaString)
		_, err = file.WriteString(alphaString + "\n")
		if nil != err {
			fmt.Println(err)
			return
		}

	}

}
