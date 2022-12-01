package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
)

var servedCust int
var unservedCust int
var wrCapacity = 3
var waiting = make(chan int, wrCapacity) // waiting room
var wrClosed = false
var wg sync.WaitGroup
var numWaiting int
var mu sync.Mutex
var shopOpenDuration = 25
var newCustInterval = 1
var custServiceDuration = 2
var barberAsleep = true

func main() {
	wg.Add(1)
	go barber()
	time.Sleep(3 * time.Second)
	go customers()
	wg.Wait()

	fmt.Println()
	color.Cyan(fmt.Sprintf("Customers served: %v\n", servedCust))
	color.Cyan(fmt.Sprintf("Customers passed: %v\n", unservedCust))
}

func barber() {
	color.Cyan("\n= = = Shop OPEN! = = =\n\n")

	go func() {
		time.Sleep(time.Duration(shopOpenDuration) * time.Second)
		color.Cyan(fmt.Sprintln("\n= = = Waiting Room closed! = = ="))
		wrClosed = true
	}()

	for nextCust := range waiting {
		if barberAsleep {
			barberAsleep = false
			color.Blue("\t\t🔔 Barber wakes up...\n")
			time.Sleep(50 * time.Millisecond) // wake up time
		}
		time.Sleep(100 * time.Millisecond) // invite customer into the barber chair
		color.Cyan(fmt.Sprintf("\t🪒 Serving 👨 #%v...\n", nextCust))
		mu.Lock()
		numWaiting--
		printCustomers()
		mu.Unlock()
		time.Sleep(time.Duration(custServiceDuration) * time.Second)
		color.Cyan(fmt.Sprintf("\t✅ Finished 👨 #%v...\n", nextCust))
		if numWaiting == 0 {
			barberAsleep = true
			color.Blue("\t\t💤 Barber snoozing...\n")
		}
	}

	color.Cyan(fmt.Sprintln("\n= = = Shop CLOSED! = = ="))
	wg.Done()
}

func customers() {
	custNo := 0

	for {
		custNo++
		if wrClosed {
			close(waiting)
			break
		}
		select {
		case waiting <- custNo:
			servedCust++
			fmt.Printf(" 👨 #%v ENTER\n", custNo)
			mu.Lock()
			numWaiting++
			printCustomers()
			mu.Unlock()
		default:
			unservedCust++
			color.Red(fmt.Sprintf("\t\t\t\t👨 #%v PASS\n", custNo))
		}
		time.Sleep(time.Duration(newCustInterval) * time.Second)
		if custNo == wrCapacity+2 {
			newCustInterval = newCustInterval * 5
		}
	}
}

func printCustomers() {
	color.Green(fmt.Sprintf(" WAITING ROOM  *  *  *  *  *  -%v- \n", strings.Repeat(" |👨| ", numWaiting)))
}
