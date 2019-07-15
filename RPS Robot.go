package main

import(
	"fmt"
	"bufio"
	"os"
)

func main(){
	filereader()
}
func filereader(){
	file, _=os.Open("RPS.txt")
	scanner:=bufio.NewScanner(file)


	for scanner.Scan(){
		line:=scanner.Text()
		fmt.Println(line)
	}
}