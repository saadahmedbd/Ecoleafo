package orderservice

import "strings"

// getShippingZone determines shipping zone from address
func (s *orderService) getShippingZone(address string) string {
	addressLower := strings.ToLower(address)

	// Local zones (Dhaka city)
	localAreas := []string{"dhaka", "uttara", "mirpur", "dhanmondi", "mohammadpur",
		"tejgaon", "shantinagar", "shaheenbag", "ashulia", "banani", "wari", "kawran bazar",
		"motijheel", "gulshan", "bashundhara", "khailgaon", "farmgate", "dilkusha",
		"nakhalpara"}
	for _, area := range localAreas {
		if strings.Contains(addressLower, area) {
			return "local"
		}
	}

	// Nationwide (other Bangladesh cities)
	nationwideAreas := []string{"bagerhat", "bandarban", "barguna", "barisal", "bhola",
		"jhalokati", "patuakhali", "pirojpur", "chattogram", "brahmanbaria", "chandpur",
		"cumilla", "cox's bazar", "feni", "khagrachhari", "lakshmipur", "noakhali",
		"rangamati", "faridpur", "gazipur", "gopalganj", "kishoreganj", "madaripur",
		"manikganj", "munshiganj", "narayanganj", "narsingdi", "rajbari", "shariatpur",
		"tangail", "khulna", "chuadanga", "jashore", "jhenaidah", "kushtia", "magura",
		"meherpur", "narail", "satkhira", "jamalpur", "mymensingh", "netrokona",
		"sherpur", "bogra", "joypurhat", "naogaon", "natore", "nawabganj", "pabna",
		"rajshahi", "sirajganj", "dinajpur", "gaibandha", "kurigram", "lalmonirhat",
		"nilphamari", "panchagarh", "rangpur", "thakurgaon", "habiganj", "moulvibazar",
		"sunamganj", "sylhet", "bangladesh"}
	for _, area := range nationwideAreas {
		if strings.Contains(addressLower, area) {
			return "nationwide"
		}
	}

	return "international"
}

// getZoneMultiplier returns cost multiplier for shipping zone
func (s *orderService) getZoneMultiplier(zone string) float64 {
	multipliers := map[string]float64{
		"local":         1.0,
		"nationwide":    1.5,
		"international": 3.0,
	}
	if mult, exists := multipliers[zone]; exists {
		return mult
	}
	return 1.0
}
