package domain

import "testing"

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		name  string
		email string
		want  bool
	}{
		{"email gmail valid", "budi@gmail.com", true},
		{"email icloud valid", "sari@icloud.com", true},
		{"tanpa domain", "bukanemail", false},
		{"hanya teks", "abc", false},
		{"format salah", "foo@bar", false},
		{"email kosong", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{Email: tt.email}
			got := u.IsValidEmail()
			if got != tt.want {
				t.Errorf("IsValidEmail(%q) = %v, want %v", tt.email, got, tt.want)
			}
		})
	}
}

func TestValidTelephone(t *testing.T) {
	tests := []struct {
		name  string
		nomor string
		want  bool
	}{
		{"nomor tidak valid", "nfdja83884343", false},
		{"nomor valid", "09384383843", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{Phone: tt.nomor}
			got, _ := u.IsValidPhone()
			if got != tt.want {
				t.Errorf("IsValidPhone(%q) = %v, want %v", tt.nomor, got, tt.want)
			}
		})
	}
}

func TestIsValidPhone(t *testing.T) {
	tests := []struct {
		name    string
		phone   string
		wantOK  bool
		wantMsg string
	}{
		{"nomor valid 12 digit", "081234567890", true, ""},
		{"nomor valid panjang minimal", "0812345678", true, ""},
		{"nomor valid panjang maksimal", "081234567890123", true, ""},
		{"nomor kosong", "", false, "mohon untuk mengisi nomor telephone terlebih dahulu"},
		{"huruf semua", "abcdefghij", false, "Nomor telephone tidak boleh mengandung selain angka"},
		{"campur huruf dan angka", "0812 abc def", false, "Nomor telephone tidak boleh mengandung selain angka"},
		{"terlalu pendek", "0812", false, "Nomor telephone tidak valid"},
		{"terlalu panjang", "0812345678901234", false, "Nomor telephone tidak valid"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{Phone: tt.phone}
			gotOK, gotMsg := u.IsValidPhone()
			if gotOK != tt.wantOK {
				t.Errorf("IsValidPhone(%q) ok = %v, want %v", tt.phone, gotOK, tt.wantOK)
			}
			if gotMsg != tt.wantMsg {
				t.Errorf("IsValidPhone(%q) msg = %q, want %q", tt.phone, gotMsg, tt.wantMsg)
			}
		})
	}
}

func TestIsValidLocation(t *testing.T) {
	office := OfficeLocation{Latitude: -6.2, Longitude: 106.8} // Jakarta

	tests := []struct {
		name      string
		lat       float64
		lon       float64
		maxRadius float64
		wantMsg   string
	}{
		{"dalam radius", -6.21, 106.81, 5, "Lokasi Valid"},
		{"di luar radius", -6.9, 107.6, 5, "Lokasi Tidak Valid"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{Latitude: tt.lat, Longitude: tt.lon}
			km, msg := u.IsValidLocation(office, tt.maxRadius)
			if msg != tt.wantMsg {
				t.Errorf("IsValidLocation() msg = %q, want %q (jarak %.2f km)", msg, tt.wantMsg, km)
			}
		})
	}
}
