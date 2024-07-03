package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	Hour   int
	Minute int
}

func ChangeToStandartTime(time interface{}) string {
	switch v := time.(type) {
	case string:
		hour, minute, err := parseTimeString(v)
		if err != nil {
			return "Invalid input"
		}
		return formatTime(hour, minute)
	case []int:
		if len(v) != 2 {
			return "Invalid input"
		}
		return formatTime(v[0], v[1])
	case map[string]int:
		hour, okHour := v["hour"]
		minute, okMinute := v["minute"]
		if !okHour || !okMinute {
			return "Invalid input"
		}
		return formatTime(hour, minute)
	case Time:
		return formatTime(v.Hour, v.Minute)
	default:
		return "Invalid input"
	}
}

func parseTimeString(timeStr string) (int, int, error) {
	parts := strings.Split(timeStr, ":")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid time format")
	}
	hour, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}
	minute, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, err
	}
	return hour, minute, nil
}

func formatTime(hour, minute int) string {
	if hour < 0 || hour > 23 || minute < 0 || minute > 59 {
		return "Invalid input"
	}
	period := "AM"
	if hour >= 12 {
		period = "PM"
		if hour > 12 {
			hour -= 12
		}
	}
	return fmt.Sprintf("%02d:%02d %s", hour, minute, period)
}

func main() {
	fmt.Println(ChangeToStandartTime("16:00"))
	fmt.Println(ChangeToStandartTime([]int{16, 0}))
	fmt.Println(ChangeToStandartTime(map[string]int{"hour": 16, "minute": 0}))
	fmt.Println(ChangeToStandartTime(Time{16, 0}))
}
