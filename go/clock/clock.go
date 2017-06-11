package clock

import "fmt"

const testVersion = 4

type Clock struct {
	minutes int
}

func New(hour, minutes int) Clock {
	amount := Clock{(hour*60 + minutes) % 1440}
	if amount.minutes < 0 {
		amount.minutes += 1440
	}
	return amount
}

func (amount Clock) String() string {
	hours := amount.minutes / 60
	minutes := amount.minutes % 60
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}

func (amount Clock) Add(minutes int) Clock {
	amount.minutes = (amount.minutes + minutes) % 1440
	if amount.minutes < 0 {
		amount.minutes += 1440
	}
	return amount
}
