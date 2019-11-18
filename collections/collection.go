package collections

import "errors"

var language = map[string]string {
	"es":"SPANISH",
	"en": "ENGLISH",
}

// All values should be same type
// Fixed size, never grows

func Names() []string {
	return []string{"Alice", "Marcus", "Gab"}
}

func Sum(a, b int) int {
	return 10
}

func Ages(age uint8) (stage string, err error) {
	if age == 0 {
		err = errors.New("age should be greater than zero")
		return
	}
	if age <= 12 {
		stage = "Child"
	} else if age <= 18 {
		stage = "Teenager"
	} else if age <= 60 {
		stage = "Adult"
	} else {
		stage = "Old man"
	}
	return
}

func WeekDays(day uint8) (dayName string, err error) {
	if day > 6 {
		err = errors.New("day number shouldn't be greater than 6")
	}
	switch day {
	case 0:
		dayName = "SUNDAY"
	case 1:
		dayName = "MONDAY"
	case 2:
		dayName = "TUESDAY"
	case 3:
		dayName = "WEDNESDAY"
	case 4:
		dayName = "THURSDAY"
	case 5:
		dayName = "FRIDAY"
	case 6:
		dayName = "SATURDAY"
	}
	return
}

func Lang(l string) (v string, err error) {
	v, ok := language[l]
	if !ok {
		err = errors.New("language not supported")
	}
	return
}
