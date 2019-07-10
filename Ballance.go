package main

import(
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
//io/ioutil"
	
)
func main(){
numi := 0
// n := 0
holdi := 0
var numbers []int

file, _:=os.Open("input.txt")
scanner := bufio.NewScanner(file)
for scanner.Scan(){
	line := scanner.Text()
    lineint, _:= strconv.Atoi(line) 

	numbers = append(numbers, lineint)
}






//numbers, _= strconv.Atoi(data) 

//numbers = append(numbers,data)


fmt.Println(len(numbers))

for i := range(numbers) {
 	

for tru:=0;tru == 0;{
	holdi = numbers[i] - numi
 	hold:= strconv.Itoa(holdi)
 	num:= strconv.Itoa(numi)
 	if(strings.Contains(hold, "4") ==  false && strings.Contains(num, "4") == false){

 		tru = 1;
 		fmt.Println(hold," ",num)
 	}else{
 		numi++;

 	}

}
 	
}

}	




