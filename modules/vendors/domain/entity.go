package domain

import (
	"time"
)

type Vendor struct {
	ID            int
	ContactPerson *string
	Email         *string
	ContactNumber *string
	Location      *string
	SIUPNumber    *string
	NIBNumber     *string
	NPWPNumber    *string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
