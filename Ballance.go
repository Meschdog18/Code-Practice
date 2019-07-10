package main

import(
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
	
)
func main(){
numi := 0
// n := 0
holdi := 0
var numbers []int
//Opens Input File
file, _:=os.Open("input.txt")
scanner := bufio.NewScanner(file)
for scanner.Scan(){
	line := scanner.Text()
    //converts string from txt file to ints 
    lineint, _:= strconv.Atoi(line) 
    //adds each line of numbers to the array
	numbers = append(numbers, lineint)
}








fmt.Println(len(numbers))

for i := range(numbers) {
 	

for tru:=0;tru == 0;{
	//trys to find number without 4 by subtracting its value by numi(which increase by one each loop)
	holdi = numbers[i] - numi
 	hold:= strconv.Itoa(holdi)
 	num:= strconv.Itoa(numi)
 	//if both number a and b, dont have a 4 in them than it resets the loop and goes to tthe next number in the input file
 	if(strings.Contains(hold, "4") ==  false && strings.Contains(num, "4") == false){

 		tru = 1;
 		fmt.Println(hold," ",num)
 	}else{
 		numi++;

 	}

}
 	
}

}	




