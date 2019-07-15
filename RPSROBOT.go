package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"unicode"
	"strings"
)

func main(){
	testcase,T := filereader()
	logic(testcase,T)
}
func logic(testcase1 []string, T int){
var move string
choices := []string{"R","P","S"}
//number of oppenents
var opponents [] int
//where we are in array
var currentqueue int
for i:=0;i<len(testcase1);i++{
		r := []rune(testcase1[i])
		for l:=0;l<len(r);l++{
			//pulls all the number of oppenents from the array
			if(unicode.IsLetter(r[l]) == false){
			opp, _:=strconv.Atoi(testcase1[i])
			opponents = append(opponents, opp)
			//fmt.Println(opponents)
		}
		}
}



for i:=0;i<T;i++{
	var resultstring string
	var impossible bool
		//reset case move var
		move = ""

		currentqueue = opponents[i] 

	for l:=0;l<opponents[i];l++{
		for q:=0;q<opponents[i];q++{
			//if other programs arnt all the same its impossible, because programs memory resets each round so it doesnt know which bot its facing
			if(strings.Contains(testcase1[currentqueue+l],testcase1[currentqueue+q])!=true){
				impossible = true
			}
		}
		//splits current set of moves to first move,since we only need to know first move to win
		moves:=strings.Split(testcase1[currentqueue + l],"")
		

		if(strings.Contains(moves[0], "R") == true){
			move = move + "P"
		}
		if(strings.Contains(moves[0],"P") == true){
			move = move + "S"
		}
		if(strings.Contains(moves[0],"S") == true){
			move = move + "R"
		}
	}
//compresses all duplcates to one
var result []string
for i:=0;i<len(choices);i++{
	
	
	if(strings.Count(move, choices[i])>1){
		result = append(result,choices[i])
	}
	if(strings.Count(move, choices[i])==1){
		result = append(result,choices[i])
	}

}
//prints results
for i:=0;i<len(result);i++{
resultstring =resultstring + result[i]
}
	if(impossible == true){
		fmt.Println("Case #",i,":","IMPOSSIBLE")
	}
	if(impossible != true){
		fmt.Println("Case #",i,":",resultstring)
	}
	
	

}
}




func filereader()(testcase1 []string,T int){
	file, _:=os.Open("RPS.txt")
	scanner:=bufio.NewScanner(file)
	fileloop := 0
	//fileloopmax:=0
	//numb of test cases
	//numb of opponents
	N:= 0
	//moves
	

	for scanner.Scan(){
		line:=scanner.Text()
		//fmt.Println(line)
		if(fileloop == 0){
			T,_=strconv.Atoi(line)
			fmt.Println(T)
		}
		

			

			if(fileloop == 1){
				N,_=strconv.Atoi(line)
			}
			for i:=0;i<N;i++{
				testcase1 = append(testcase1, line)
			}											

			
		
		
			
			fileloop++

	}
	return testcase1,T
}
