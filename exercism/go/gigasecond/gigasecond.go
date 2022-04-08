package gigasecond

import "time"

// AddGigasecond determine the moment that
// would be after a gigasecond has passed
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1_000_000_000 * time.Second)
}
