package isdayoff

import (
	"net/http"
	"testing"
	"time"
)

func TestIsDayOff(t *testing.T) {
	currentTime := time.Now()
	tests := []struct {
		date     time.Time
		isDayOff bool
		wantErr  bool
	}{
		{
			date:     time.Date(2019, 11, 8, 15, 0, 0, 0, currentTime.Location()),
			isDayOff: false,
			wantErr:  false,
		},
		{
			date:     time.Date(2019, 11, 9, 15, 0, 0, 0, currentTime.Location()),
			isDayOff: true,
			wantErr:  false,
		},
		{
			date:     time.Date(9999, 11, 9, 15, 0, 0, 0, currentTime.Location()),
			isDayOff: false,
			wantErr:  true,
		},
	}
	for _, test := range tests {
		isDayOff, err := IsDayOff(http.DefaultClient, test.date)
		if (err != nil) != test.wantErr {
			t.Errorf("IsDayOff() error=%v wantErr=%v", err, test.wantErr)
		}
		if isDayOff != test.isDayOff {
			t.Errorf("IsDayOff().isDayOff != %v", test.isDayOff)
		}
	}
}

func TestIsDayOffToday(t *testing.T) {
	_, err := IsDayOffToday(http.DefaultClient)
	if err != nil {
		t.Error(err)
	}
}
