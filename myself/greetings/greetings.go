package greetings

import (
	"time"
)

func GoodDay() string {
	return "Good Day"
}

func GoodNight() string {
	return "Good Night"
}

func IsAM() bool {
	localTime := time.Now()
	return localTime.Hour() <= 12
}

func IsAfternoon() bool {
	localTime := time.Now()
	return localTime.Hour() <= 18
}

func IsEvening() bool {
	localTime := time.Now()
	return localTime.Hour() <= 22
}
