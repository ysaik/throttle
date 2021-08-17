# throttle
Throttle go functions
- This package can be used as a rate-limiter on the client side which connects server frequently
- This package can also be used to reduce load on the channels by limitting the rate of requests sending

```
package main

import (
	"fmt"
	"myexp/throttle"
	"time"
)

func greetings(msg string) {
	fmt.Printf("%v \t ", msg)
}
func main() {
	//quit := make(chan int)
	t := throttle.Init(5, time.Millisecond*500)
	t.Start()

	time.Sleep(2 * time.Second)
	for {
		ti := time.Now()
		t.Check()
		greetings("Hello !!!")
		time.Sleep(time.Millisecond * 50)
		val := time.Since(ti)
		fmt.Println(val)
	}

}

```
