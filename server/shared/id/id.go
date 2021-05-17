package id

type TripID string

func (t TripID) String() string {
	return string(t)
}
