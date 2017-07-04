package main

import "fmt"

const (
	cStartYear  = 1901
	cEndYear    = 2000
	cJanuary    = 1
	cFebruary   = 2
	cDecember   = 12
	cSaturday   = 6
	cSunday     = 7
	cDaysInWeek = 7
)

var smallMonthsMap map[int]bool

func main() {
	var firstSundaysCount int
	smallMonthsMap = setupSmallMonths()
	firstSunday := find1901FirstSunday()

	//Account for days of the first month first week in the previous year
	totalDays := cDaysInWeek - firstSunday

	for year := cStartYear; year <= cEndYear; year++ {
		for month := cJanuary; month <= cDecember; month++ {

			prevMonthEndingDay := totalDays % cDaysInWeek

			//If prev month ends on Saturday, this starts on Sunday
			if prevMonthEndingDay == cSaturday {
				firstSundaysCount++
			}

			totalDays += getMonthDays(year, month)
		}
	}

	fmt.Println(firstSundaysCount)

}

func setupSmallMonths() map[int]bool {
	smallMonths := make(map[int]bool)
	smallMonths[4] = true
	smallMonths[6] = true
	smallMonths[9] = true
	smallMonths[11] = true
	return smallMonths
}

func getMonthDays(year int, month int) int {
	days := 31
	if smallMonthsMap[month] {
		days--
	} else if month == cFebruary {
		days = 28
		if isLeapYear(year) {
			days++
		}
	}

	return days
}

func isLeapYear(year int) bool {
	return isFourthYear(year) && !(isCentuary(year) && !isCentuaryDivisibleBy400(year))
}

func isFourthYear(year int) bool {
	return year%4 == 0
}

func isCentuary(year int) bool {
	return year%100 == 0
}

func isCentuaryDivisibleBy400(year int) bool {
	return year%400 == 0
}

func find1901FirstSunday() int {
	const year1900 = 1900
	days := 365
	if isLeapYear(year1900) {
		days++
	}

	return cSunday - (days % cDaysInWeek)
}

/*
You are given the following information, but you may prefer to do some research for yourself.

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
