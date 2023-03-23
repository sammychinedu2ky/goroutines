package helpers

import (
	"fmt"
	
	"sync"
	"time"
)

var wg sync.WaitGroup
func GoRoutines1() {
 go say("world")
	say("hello")

}

func GoRoutines2(){
	wg.Add(2)
	go say2("hello")
   // wg.Add(1)
	go say2("world")
	wg.Wait()
	wg.Add(1)
	go say2("sing")
	wg.Wait()
}

func GoRoutines3(){
	wg.Add(1)
	go say3("hello")
    wg.Add(1)
	go say3("world")
	
	wg.Wait()
	wg.Add(1)
	go say3("sing")
	wg.Wait()
}
func cleanUp(){
	if r:= recover(); r!= nil{
		fmt.Println("recovered in cleanUp", r)
	}
	defer wg.Done()
}
func say3(s string) {
	defer cleanUp()
	for i := 0; i < 2; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
	panic("ooh dear I have failed " + s )
}
func say(s string) {
	for i := 0; i < 2; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
}
func say2(s string) {
	for i := 0; i < 2; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
	
}
