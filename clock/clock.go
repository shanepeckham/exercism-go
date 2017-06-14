package clock

import (
	"fmt"
	"strings"
)

const testVersion = 4

type Clock struct{ hour, minute int }

var c Clock

func New(hour, minute int) Clock {
	c := Clock{hour, minute}
	cv := Validate(c)
	return cv

}

func (c Clock) String() string {
	tr := fmt.Sprintf("%02d", c)
	tr = strings.TrimLeft(tr, "{")
	tr = strings.Replace(tr, " ", ":", 2)
	tr = strings.TrimRight(tr, "}")
	return tr
}

func (c Clock) Add(minutes int) Clock {
	c.minute += minutes
	cv := Validate(c)
	return cv
}

func Validate(c Clock) Clock {
	switch {
	case c.minute >= 60:
		m := c.minute
		for m >= 60 {
			m -= 60
			c.hour++
		}
		c.minute = m
	case c.minute < 0:
		m := c.minute
		for m < 0 {
			m += 60
			c.hour--
		}
		c.minute = m
	}

	switch {
	case c.hour >= 24:
		c.hour = c.hour % 24
	case c.hour < 0:
		h := c.hour
		for h < 0 {
			h += 24
		}
		c.hour = h
	}

	return c
}
