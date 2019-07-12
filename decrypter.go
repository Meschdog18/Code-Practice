package main
import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	//"math"
)

func main(){
	sendN, NUMB:=filereader()
	primefinder(sendN)
	oddprimes := primefinder(sendN)
	decrypter(oddprimes,NUMB)
}
func primefinder(N int)(oddprimes []int){
//var primecheck int
//first odd prime should be 1 but it starts off at 3
//var oddprimes []int
var divisible float64
var l float64
//var primenumbers[]int
//appending all the primes that get multiplied out
oddprimes = append(oddprimes,3)
oddprimes = append(oddprimes,5)
oddprimes = append(oddprimes,7)
for i:=0;i<=N;i++{
		
		//follows the Sieve of Eratosthenes
		//each layer has a divisble value, and if the number's decimal doesnt contain 0 (which would point to a whole number), than it gets passed through another filter until it makes it to the end and is added as a prime
		//the deceimal value(yeah i know i spelled it wrong), is used to just compare the decimal's value not the whole number left over after division


		divisible = 2
		
		contain := fmt.Sprintf("%.1f",l/divisible)
		
		deceimal := strings.SplitAfter(contain,".")
		if(strings.Contains(deceimal[1],"0")!=true){
			divisible = 3
			contain2 := fmt.Sprintf("%.1f",l/divisible)
			deceimal2 := strings.SplitAfter(contain2,".")
			if(strings.Contains(deceimal2[1],"0")!=true){
				divisible = 5
				contain3 := fmt.Sprintf("%.1f", l/divisible)
				deceimal3 := strings.SplitAfter(contain3,".")
				if(strings.Contains(deceimal3[1],"0")!=true){
					divisible = 7
					contain4 := fmt.Sprintf("%.1f",l/divisible)
					deceimal4 := strings.SplitAfter(contain4, ".")
					if(strings.Contains(deceimal4[1],"0")!=true){
							//check to see if 1, because 1 is not a prime

							if(i != 1){
								

								oddprimes = append(oddprimes,i)
							}

					}

				}
			}
			
		
		}
		l++

}
	
	return oddprimes
}

func decrypter(oddprimes []int, NUMB []int){

var potential []string
 var potential2 []string

//var last bool


//letters:=[]string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
//var panagram []string
for l:=0;l<len(NUMB);l++{
		
	f1 := float64(NUMB[l])

		for i:=0;i<len(oddprimes);i++{
				f2 := float64(oddprimes[i])

			f := fmt.Sprintf("%.4f",f1/f2)
			decimal:=strings.Split(f,".")
			if(strings.Contains(decimal[1],"0000")==true){
				//ive tried multiple things not sure what todo
				//it prints out the two prime numbers it equals but not sure where to go from their 
			}

		}

		//fmt.Println(oddprimes[i],":",letters[i])
	}
	
}

func filereader()(int, []int){
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
	return sendN,NUMB

}