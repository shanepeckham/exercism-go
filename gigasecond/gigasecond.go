package gigasecond

import (
	"time"
)

const testVersion = 4

func AddGigasecond(tm time.Time) time.Time {

	s := tm.Unix()
	s += 1000000000
	tm = time.Unix(s, 0)
	return tm

}
