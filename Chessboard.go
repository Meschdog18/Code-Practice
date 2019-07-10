package main

import(
	"fmt"
	"strconv"
	"strings"
)

func main(){
 
	ChessboardTraveling("1,1.2,2")
}
func ChessboardTraveling(str string){
	 //spilt[0] is current pos coords
	 //split[1] is goal pos coords
	 //CorrectPath := 0
	//splits the two sets of coords
reset := false
	foundx := false
	foundy := false
	var historyx []int
	var historyy []int
	split := strings.SplitAfterN(str, ".",2)
	split[0] = strings.Replace(split[0], ".","", -1) 
	//splits the set of coords to x and y
	currentxy := strings.SplitAfterN(split[0],",",-1)
	previousxy := strings.SplitAfterN(split[1],",",-1)
	//assign each x and y
	currentx, _:= strconv.Atoi(currentxy[0])
		currenty, _:= strconv.Atoi(currentxy[1])
		previousx, _:= strconv.Atoi(previousxy[0])
		previousy, _:= strconv.Atoi(previousxy[1])
	
	
	

	for {
   	if(reset == true){
   		currentx, _ = strconv.Atoi(currentxy[0])
		currenty, _ = strconv.Atoi(currentxy[1])
		previousx, _ = strconv.Atoi(previousxy[0])
		previousy, _ = strconv.Atoi(previousxy[1])
		reset = false;
   	}
   	//try going straight up



//go right
    if(previousx > currentx && foundx == false){
		currentx++;

	}
//go left
	if(previousy > currenty && foundy == false){
		currenty++;
	}
	
	if(previousy == currenty){
		foundy = true
		
	}
	if(previousx == currentx){
		foundx = true
	}
	if(foundx == true && foundy == true){
		for i:=0;i<len(historyy);i++{
			//if this current path is differant from previous, save it to history
			if(historyy[i] != currenty && historyx[i] != currentx){
				fmt.Println("made it")
				// historyy = append(historyy, currenty)
			}
		}
	}
	}
	
	
	
	

}