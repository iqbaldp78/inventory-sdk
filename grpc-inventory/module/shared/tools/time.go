package tools

import "time"

//DateFormat used as constant for date format
const DateFormat = "2006-01-02"

//TimestampFormat used as constant for timestamp format
const TimestampFormat = "2006-01-02 15:04:05"

//SetUTCPlus7 used for set base time to UTC+7
func SetUTCPlus7(input time.Time) time.Time {
	temp := input.Format(TimestampFormat)
	zone, _ := time.LoadLocation("Asia/Jakarta")
	result, _ := time.ParseInLocation(TimestampFormat, temp, zone)
	return result
}

//LocalTime used for change current date time to local date time
func LocalTime(input time.Time, location string) time.Time {
	zone, _ := time.LoadLocation(location)
	_, oo := input.In(zone).Zone()
	return input.Add(time.Duration(oo) * time.Second)
}

//TimeToString used for convert time to string
func TimeToString(input time.Time, to string) string {
	var result string
	switch to {
	case "date":
		if result = input.Format("2006-01-02"); result == "0001-01-01" {
			result = ""
		}
	case "timestamp":
		if result = input.Format("2006-01-02 15:04:05"); result == "0001-01-01 00:00:00" {
			result = ""
		}
	}
	return result
}
