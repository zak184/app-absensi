package domain

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/umahmood/haversine"
)

type Role string

const (
	RoleUser Role = "USER"
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
	ID           uuid.UUID
	Nama         string
	Email        string
	PasswordHash string
	Latitude     float64
	Longitude    float64
	Phone        string
	Role         Role
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type OfficeLocation struct {
	Latitude  float64
	Longitude float64
}

// validasi email
func (r *User) IsValidEmail() bool {
	email := strings.ToLower(r.Email)

	if email == "" {
		return false
	}

	if strings.HasSuffix(email, "@gmail.com") || strings.HasSuffix(email, "@icloud.com") {
		return true
	} else {
		return false
	}
}

// validasi nomor telephone
func (r *User) IsValidPhone() (bool, string) {
	phoneLen := len(r.Phone)

	for _, char := range r.Phone {
		if char < '0' || char > '9' {
			return false, "Nomor telephone tidak boleh mengandung selain angka"
		}
	}

	if r.Phone == "" {
		return false, "mohon untuk mengisi nomor telephone terlebih dahulu"
	}

	if phoneLen > 15 || phoneLen < 10 {
		return false, "Nomor telephone tidak valid"
	}

	return true, ""
}

// validasi jarak lokasi user dengan office
func (r *User) IsValidLocation(office OfficeLocation, maxRadius float64) (float64, string) {

	if r.Latitude == 0.0 || r.Longitude == 0.0 {
		return 0.0, "Mohon untuk mengisi lokasi"
	}

	officeLocation := haversine.Coord{Lat: office.Latitude, Lon: office.Longitude}
	userLocation := haversine.Coord{Lat: r.Latitude, Lon: r.Longitude}
	_, Km := haversine.Distance(officeLocation, userLocation)

	if Km <= maxRadius {
		return Km, "Lokasi Valid"
	}
	return Km, "Lokasi Tidak Valid"
}
