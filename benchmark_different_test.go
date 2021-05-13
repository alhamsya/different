package different

import "testing"

func BenchmarkGenerateDiff(b *testing.B) {
	type User struct {
		Name string
		Age  int `diff:"-"`
	}

	origin := &User{
		Name: "Alhamsya",
		Age:  10,
	}

	new := &User{
		Name: "Bintang",
		Age:  10,
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GenerateDiff(origin, new)
	}
}
