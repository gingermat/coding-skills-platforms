package license_key

import "testing"

var tests = []struct {
	s    string
	k    int
	want string
}{
	{
		"5F3Z-2e-9-w", 4, "5F3Z-2E9W",
	},
	{
		"2-5g-3-J", 2, "2-5G-3J",
	},
}

func TestLicenseKeyFormatting(t *testing.T) {
	for _, tt := range tests {
		got := licenseKeyFormatting(tt.s, tt.k)
		if got != tt.want {
			t.Fatalf("licenseKeyFormatting(%v, %v) = %v, want %v", tt.s, tt.k, got, tt.want)
		}
	}
}
