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

	type Dummy struct {
		Error string
	}

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
			name: "When_structureDataIsDifferent_ReturnSuccess",
			args: args{
				origin: &User{
					Name: "Alhamsya",
					Age:  10,
				},
				new: &User{
					Name: "Bintang",
					Age:  12,
				},
			},
			want:    []byte(`[{"Name":{"before":"Alhamsya","after":"Bintang"}}]`),
			wantErr: false,
		},
		{
			name: "When_deepEqual_ReturnNil",
			args: args{
				origin: nil,
				new:    nil,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "When_deepEqual_ReturnNil",
			args: args{
				origin: nil,
				new:    nil,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "When_originParamIsNil_ReturnError",
			args: args{
				origin: &User{
					Name: "Alhamsya",
					Age:  10,
				},
				new: &Dummy{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "When_originParamIsNil_ReturnError",
			args: args{
				origin: &User{},
				new: &User{
					Name: "Alhamsya",
					Age:  10,
				},
			},
			want:    []byte(`[{"Name":{"before":"","after":"Alhamsya"}}]`),
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
