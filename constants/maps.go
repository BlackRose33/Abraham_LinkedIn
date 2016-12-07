package constants

// OfficeNames maps office codes to names
var OfficeNames map[int]string

// BaseInitialization maps office positions to list
func BaseInitialization() {
	OfficeNames = map[int]string{
		1: "Mayor",
		2: "Public Advocate",
		3: "Comptroller",
		4: "Borough President",
		5: "City Council",
		6: "Other",
	}
}
