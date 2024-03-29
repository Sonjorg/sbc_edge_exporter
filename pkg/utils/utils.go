/* Copyright (C) 2023 Sondre Jørgensen - All Rights Reserved
 * You may use, distribute and modify this code under the
 * terms of the CC BY 4.0 license
*/
package utils

import (
"log"
"time"
)
func Expired(hours float64, previoustime time.Time) bool{
	var timeSchedule time.Duration = time.Duration(hours)
	duration := timeSchedule*time.Hour
	now := time.Now().Format(time.RFC3339)
	timeNowParsed, err := time.Parse(time.RFC3339, now)
	if err != nil {
		log.Print(err)
		return false
	}
	if err != nil {
		log.Print(err)
		return false
	}
	//If previous time + 24 hours is before now: database for routingentries has expired
	return previoustime.Add(duration).Before(timeNowParsed)
}