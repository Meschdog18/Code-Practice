package main
import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"math"
)

func main(){
	filereader()
	
}
func primefinder(N int){
//var primecheck int
//first odd prime should be 1 but it starts off at 3
//var oddprimes []int




//this isnt my code i tried for like 2 hours to get my method to work and even this code doesn't work eaither
var primenumber bool
fmt.Println(N)
	for l:=2;l<103;l++{
			for i := 2; i<= int(math.Floor(float64(N)/2)); i++{
				if N%i == 0{
					primenumber = false
				}else{
					primenumber = true
				}
				
			}
			if(primenumber == true){
				fmt.Println(l)
			}
		}
		
		
		}
	



func filereader(){
	fileloop := 0
	var N[]string
	var L[]string
	var NumbersString[]string
	var NUMB[]int
	file, _ := os.Open("deinput.txt")
	scanner := bufio.NewScanner(file)
	sendN := 0

	





	for scanner.Scan(){
		line := scanner.Text()

		

		

		if(fileloop == 0){
				
				

			T, _ := strconv.Atoi(line)
			fmt.Println(T)
		} 
		
		if(fileloop == 1){
			//spilts line two to var N and L(length)
			N =strings.Split(line, " ")
			L = append(L,N[1])
			//convert to int
			N, _:=strconv.Atoi(N[0])
			L, _:=strconv.Atoi(L[0])
			sendN = N

			fmt.Println(L)
			fmt.Println(N)
			}
		if(fileloop == 2){
			//sperate and convert to ints
			NumbersString = strings.Split(line, " ")
			for i:=0;i<len(NumbersString);i++{
				Numbers, _:=strconv.Atoi(NumbersString[i])
				NUMB = append(NUMB, Numbers)
				
				//fmt.Println(NUMB[i])
			}
			//N, _:= strconv.Atoi(line)
		}
		fileloop++;
		
		
	}
	primefinder(sendN)

}