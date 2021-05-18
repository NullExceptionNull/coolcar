package id

type TripID string

func (t TripID) String() string {
	return string(t)
}

type IdentityId string

func (i IdentityId) String() string {
	return string(i)
}

type CarID string

func (c CarID) String() string {
	return string(c)
}
