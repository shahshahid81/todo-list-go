package utils

import (
	"encoding/json"
	"fmt"
	"time"
)

type IsoDate time.Time

func (isoDate *IsoDate) Marshal() ([]byte, error) {
	t := time.Time(*isoDate)
	date := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	fmt.Printf("date is %v", date)
	data, err := json.Marshal(date)
	if err != nil {
		fmt.Print(
			"Error in marshal",
		)
		fmt.Println(err)
		return []byte{}, err
	}
	return data, nil
}

func (isoDate *IsoDate) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation("2006-01-02", s, time.UTC)
	if err != nil {
		return err
	}
	*isoDate = IsoDate(t)
	return nil
}
