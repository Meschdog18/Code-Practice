package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func main(){
	T,N,Ci,Ji:=filereader()
	solver(T,N,Ci,Ji)

}
func solver(T int,N[]int,Ci[]int,Ji[]int){
	fmt.Println(T,Ji[0])
	for i:=0;i<3;i++{
			//Ci
			min := 8
			for l:=0;l<N[i];l++{
					if(Ci[l]<min){
						min = Ci[l] 
						
					}
					
			}
				fmt.Println(min)
		}
			
	}
	

func filereader()(T int,N[]int,Ci[]int, Ji[]int){
	repeat:=0
	file, _:=os.Open("Elements.txt")
	scanner:=bufio.NewScanner(file)
	fileloop:=0
	
	for scanner.Scan(){
		
		line:=scanner.Text()
		if(fileloop == 0){
			T,_=strconv.Atoi(line)
			fileloop = 1
		}
		
		if(fileloop == 1){
			//if not a N value, split it 
			if(strings.Contains(line, " ")){
				Hold := strings.Split(line," ")
				Holdint,_ := strconv.Atoi(Hold[0])
				Holdint2,_ := strconv.Atoi(Hold[1])
				//Gets converted to ints and added to each array
				Ci = append(Ci,Holdint)
				Ji = append(Ji,Holdint2)
				
				
			}
			if(strings.Contains(line," ") == false){
				//stops from repeating T
				if(repeat == 1){
					Hold,_:= strconv.Atoi(line)
				N = append(N,Hold)
				}
				repeat = 1

			}
			//conv, _:=strconv.Atoi(line)
			//Rawinput = append(Rawinput,conv)
		}
		
	}
	return T,N,Ci,Ji
}

