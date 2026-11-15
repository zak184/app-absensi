package domain

import (
	"strings"

	"github.com/google/uuid"
	"github.com/umahmood/haversine"
)

type Role string

const (
	RoleUser Role = "user"
	RoleHr   Role = "HR"
)

func (r Role) IsValid() bool {
	switch r {
	case RoleUser, RoleHr:
		return true
	}
	return false
}

func (r Role) String() string {
	return string(r)
}

type User struct {
	ID            uuid.UUID
	Nama          string
	Email         string
	Password_hash string
	Latitude      float64
	Longitude     float64
	Phone         string
	Role          Role
}

type OfficeLocation struct {
	Latitude  float64
	Longitude float64
}

func (r *User) IsValidEmail() bool {
	email := strings.ToLower(r.Email)

	if email != "" {
		return true
	} else if strings.HasSuffix(email, "@gmail.com") || strings.HasSuffix(email, "@icloud.com") {
		return true
	} else {
		return false
	}
}

func (r *User) IsValidLocation(office OfficeLocation, maxRadius float64) (float64, string) {
	officeLocation := haversine.Coord{Lat: office.Latitude, Lon: office.Longitude}
	userLocation := haversine.Coord{Lat: r.Latitude, Lon: r.Longitude}
	_, Km := haversine.Distance(officeLocation, userLocation)

	if Km <= maxRadius {
		return Km, "Lokasi Valid"
	}
	return Km, "Lokasi Tidak Valid"
}
