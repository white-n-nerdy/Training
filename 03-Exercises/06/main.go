package main

import "fmt"

func main() {
	months := []string{"jan", "feb", "march", "april", "may", "june", "july", "aug", "sept", "oct", "nov", "dec"}
	weekdays := []string{"sun", "mon", "tues", "wed", "thurs", "fri", "sat"}
	totalDays := 0 //Total number of days
	init := 1      //Initializer for 'for' loop
	answer := 0    // Initialized variable for counting number of applicable Sundays

	//Iterate through years
	for year := 1900; year <= 2000; year++ {
		//Determine if this is a leap year
		leap := leapYear(year)
		fmt.Println("Year: ", year, "Leap year?:", leap)

		for _, month := range months {
			days := daysInMonth(month, leap) // Return # of days in the month
			totalDays += days
			dayOfMonth := 1 //Instantiate 1st day of month
			// To keep week days and month days lined up, keep tabs of total number of days
			for i := init; i <= totalDays; i++ {
				wkDayNum := i % 7 //Iterate through weekdays

				//Determine if this is a Sunday that falls on the 1st of the month after 1901
				if year >= 1901 && dayOfMonth == 1 && weekdays[wkDayNum] == "sun" {
					answer += 1
					fmt.Println("Answer:", answer)
					fmt.Println(i, " ", wkDayNum, " ", weekdays[wkDayNum], "-", month, "-", dayOfMonth, "-", year)
				}

				dayOfMonth += 1
			}
			init = totalDays + 1 // Start next loop at first of next month
		}
	}

	fmt.Println("Final Answer:", answer)
}

func daysInMonth(month string, leap bool) int {
	var days int
	switch month {
	case "sept", "april", "june", "nov":
		days = 30
	case "feb":
		days = 28
	case "jan", "march", "may", "july", "aug", "oct", "dec":
		days = 31
	}
	if month == "feb" && leap == true {
		days = 29 //Return correct days for a leap year
	}
	return days
}

func leapYear(year int) bool {
	var leap bool

	if year%100 == 0 {
		//Year is a century, needs to be divisible by 400 to be a leap year
		if year%400 == 0 {
			leap = true //Year is a leap year
		}
	} else if year%4 == 0 {
		leap = true //Year is a leap year
	} else {
		leap = false //Not a leap year
	}
	return leap
}

/*
You are given the following information, but you may prefer to do some research for yourself.
https://projecteuler.net/problem=19
    1 Jan 1900 was a Monday.
    Thirty days has September,
    April, June and November.
    All the rest have thirty-one,
    Saving February alone,
    Which has twenty-eight, rain or shine.
    And on leap years, twenty-nine.
    A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.

How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
*/
