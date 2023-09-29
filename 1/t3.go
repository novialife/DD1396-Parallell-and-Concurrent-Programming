// Mervan Kaya palinda21 Task 3
package main

import (
	"fmt"
	"time"
)

func Remind(text string, delay time.Duration) {
	for {
		hours, minutes, _ := time.Now().Clock()                  // Omit seconds and split time into Hours and Minutes
		time.Sleep(delay)                                        // Sleep for the specified amount of time, ex. 3hrs for food
		fmt.Printf("Klockan är %v.%v: %s", hours, minutes, text) // After the time, print time and text.
	}
}

func main() {
	go Remind("Dags att äta", time.Duration(3*time.Hour))    // Create 3 concurrent threads. One for each "alarm"
	go Remind("Dags att arbeta", time.Duration(8*time.Hour)) // These will all run in the background simultaneously
	go Remind("Dags att sova", time.Duration(24*time.Hour))  // And perform their print after the given amount of time
	select {}                                                // Block closing of application
}
