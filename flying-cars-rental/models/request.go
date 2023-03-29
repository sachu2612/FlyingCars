package models

import (
	"errors"
	"time"
)

type RequestStatus string

const (
	Pending  RequestStatus = "PENDING"
	Approved RequestStatus = "APPROVED"
	Rejected RequestStatus = "REJECTED"
)

type Request struct {
	ID          int64          `json:"id" db:"id"`
	UserID      int64          `json:"user_id" db:"user_id"`
	CarID       int64          `json:"car_id" db:"car_id"`
	PickupDate  time.Time      `json:"pickup_date" db:"pickup_date"`
	ReturnDate  time.Time      `json:"return_date" db:"return_date"`
	Status      RequestStatus  `json:"status" db:"status"`
	CreatedDate time.Time      `json:"created_date" db:"created_date"`
	UpdatedDate time.Time      `json:"updated_date" db:"updated_date"`
	Car         *Car           `json:"car,omitempty" db:"-"`
	User        *User          `json:"user,omitempty" db:"-"`
}

func (r *Request) Validate() error {
	if r.UserID <= 0 {
		return errors.New("user id is required")
	}

	if r.CarID <= 0 {
		return errors.New("car id is required")
	}

	if r.PickupDate.IsZero() {
		return errors.New("pickup date is required")
	}

	if r.ReturnDate.IsZero() {
		return errors.New("return date is required")
	}

	if r.PickupDate.After(r.ReturnDate) {
		return errors.New("return date must be after pickup date")
	}

	if r.Status != Pending && r.Status != Approved && r.Status != Rejected {
		return errors.New("invalid request status")
	}

	return nil
}
