package domain

import "errors"

// Sentinel error untuk domain user.
// Layer luar cukup memakai errors.Is(err, domain.ErrXxx)
// tanpa perlu tahu error ini berasal dari database atau bukan.
var (
	ErrUserNotFound       = errors.New("user tidak ditemukan")
	ErrEmailAlreadyExists = errors.New("email sudah terdaftar")
	ErrInvalidCredentials = errors.New("email atau password salah")
	ErrForbidden          = errors.New("akses ditolak: role tidak mencukupi")
	ErrLocationNotSet     = errors.New("lokasi belum diisi")
)
