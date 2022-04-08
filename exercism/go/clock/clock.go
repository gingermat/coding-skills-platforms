package clock

import "fmt"

// Clock implement a clock that handles times
type Clock struct {
	minute int
}

// New is return ref of Clock struct
func New(hour int, minute int) Clock {
	minute = (hour*60 + minute) % 1440

	if minute < 0 {
		minute += 1440
	}

	return Clock{minute: minute}
}

// Add is add minutes
func (c Clock) Add(minute int) Clock {
	return New(0, c.minute+minute)
}

// Subtract is subtract minutes
func (c Clock) Subtract(minute int) Clock {
	return New(0, c.minute-minute)
}

// String is represent text format Clock struct
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minute/60, c.minute%60)
}
