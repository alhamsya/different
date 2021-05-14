package different

import "testing"

func BenchmarkGenerateDiff(b *testing.B) {
	type Phone struct {
		Type        string
		PhoneNumber string
	}

	type Address struct {
		City   string
		Street string
	}

	type ContactInfo struct {
		Name  string
		Phone Phone
		Address
	}

	origin := &ContactInfo{
		Name: "Alhamsya",
		Phone: Phone{
			PhoneNumber: "085xxxxxxxx",
		},
		Address: Address{
			City: "Jakarta",
		},
	}

	newData := &ContactInfo{
		Name: "Alhamsya Bintang Dyasta",
		Phone: Phone{
			PhoneNumber: "082xxxxxxxx",
		},
		Address: Address{
			City: "Kediri",
		},
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GenerateDiff(origin, newData)
	}
}
