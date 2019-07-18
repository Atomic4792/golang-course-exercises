package main

import (
	"fmt"
	"time"
)

func main() {
	today:=time.Now().Weekday()
	if today==time.Monday {
		fmt.Printf("today is %v\n",time.Monday)

	} else if today==time.Tuesday{
		fmt.Printf("today is %v\n",time.Tuesday)

	} else if today==time.Wednesday{
		fmt.Printf("today is %v\n",time.Wednesday)

	} else if today==time.Thursday{
		fmt.Printf("today is %v\n",time.Thursday)

	} else if today==time.Friday{
		fmt.Printf("today is %v\n",time.Friday)

	} else if today==time.Saturday{
		fmt.Printf("today is %v\n",time.Saturday)

	} else {
		fmt.Printf("today is %v\n",time.Sunday)

	}


	}


/*
create a program that uses “else if” and “else”.
 */
