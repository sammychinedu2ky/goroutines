package proj

import (
	"fmt"
	// "proj/helpers"
	"sync"

)
var wg sync.WaitGroup

func sendToChannel(c chan int, value int){
   c <- value
   c <- value*2
}
func GetFromChannel()  {
	mychan := make(chan int)
	go sendToChannel(mychan, 1)
	go sendToChannel(mychan, 2)
	
     <- mychan
	fmt.Println(<- mychan)
}

func UseRange(){
	mychan := make(chan int,10)
	for i:=0; i<10; i++ {
		wg.Add(1)
		go sendToChannel2(mychan, i)
	}
	wg.Wait()
	close(mychan)
	for v := range mychan {
		fmt.Println(v)
	}
}

func sendToChannel2(mychan chan int, i int) {
	defer wg.Done()
	mychan <- i
}
