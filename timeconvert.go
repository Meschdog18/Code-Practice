package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {

reader := bufio.NewReader(os.Stdin)
fmt.Println("Enter")
text, _:= reader.ReadString('\n')
text = strings.Replace(text, "\n", "", -1)
textint, _:= strconv.Atoi(text)
	TimeConvert(textint)
}
func TimeConvert(num int){
	minutes := num;
	hours := 0;
	for minutes >= 60{
	if(minutes >= 60){
	minutes = minutes - 60;
	hours++;
	}	
	}
	fmt.Println(hours,":",minutes)
}
