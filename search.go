package search

import (
	"bufio"
	"path/filepath"
	"os"
	"flag"
	"fmt"
	"strings"
	"strconv"
	"log"
)

func visit(path string, f os.FileInfo, err error) error {
	if(!f.IsDir()) {
		fil, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(fil)
		line := 1
		for scanner.Scan() {
			if strings.Contains(scanner.Text(), "Sumeria") {
				fmt.Println(path + " at line " + strconv.Itoa(line) + ": " + scanner.Text())
			}
			line++
		}
		if err := scanner.Err(); err != nil {
			// Handle the error
		}
	}
	return nil
}


func main() {
	flag.Parse()
	root := flag.Arg(0)
	filepath.Walk(root, visit)
	//err := filepath.Walk(root, visit)
	//fmt.Printf("filepath.Walk() returned %v\n", err)
}
