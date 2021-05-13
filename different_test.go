package different

import (
	"reflect"
	"testing"
)

func TestGenerateDiff(t *testing.T) {
	type User struct {
		Name string
		Age  int `diff:"-"`
	}

	result := `[{"Name":{"before":"Alhamsya","after":"Bintang"}}]`

	type args struct {
		origin interface{}
		new    interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "When_test1",
			args: args{
				origin: &User{
					Name: "Alhamsya",
					Age:  10,
				},
				new: &User{
					Name: "Bintang",
					Age:  10,
				},
			},
			want:    []byte(result),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateDiff(tt.args.origin, tt.args.new)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateDiff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
