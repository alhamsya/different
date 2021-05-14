package different

import (
	"testing"
)

func TestGenerateDiff(t *testing.T) {
	type User struct {
		Name string
		Age  int `diff:"-"`
	}

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

	type Dummy struct {
		Error string
	}

	type args struct {
		originData interface{}
		newData    interface{}
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
				originData: &User{
					Name: "Alhamsya",
					Age:  10,
				},
				newData: &User{
					Name: "Bintang",
					Age:  12,
				},
			},
			want:    []byte(`[{"Name":{"before":"Alhamsya","after":"Bintang"}}]`),
			wantErr: false,
		},
		{
			name: "When_nestedStructureDataIsDifferent_ReturnSuccess",
			args: args{
				originData: &ContactInfo{
					Name: "Alhamsya",
					Phone: Phone{
						PhoneNumber: "085xxxxxxxx",
					},
					Address: Address{
						City: "Jakarta",
					},
				},
				newData: &ContactInfo{
					Name: "Alhamsya Bintang Dyasta",
					Phone: Phone{
						PhoneNumber: "082xxxxxxxx",
					},
					Address: Address{
						City: "Kediri",
					},
				},
			},
			want:    []byte(`[{"Name":{"before":"Alhamsya","after":"Alhamsya Bintang Dyasta"}},{"Phone":{"PhoneNumber":{"before":"085xxxxxxxx","after":"082xxxxxxxx"}}},{"Address":{"City":{"before":"Jakarta","after":"Kediri"}}}]`),
			wantErr: false,
		},
		{
			name: "When_deepEqual_ReturnNil",
			args: args{
				originData: nil,
				newData:    nil,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "When_TypeInteger_ReturnNil",
			args: args{
				originData: 1,
				newData:    2,
			},
			want:    []byte(`[{"before":1,"after":2}]`),
			wantErr: false,
		},
		{
			name: "When_diff_ReturnError",
			args: args{
				originData: &User{
					Name: "Alhamsya",
					Age:  10,
				},
				newData: 1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "When_typeDiffNotUpdate_ReturnError",
			args: args{
				originData: &User{
					Name: "Alhamsya",
					Age:  10,
				},
				newData: nil,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "When_originParamIsNil_ReturnError",
			args: args{
				originData: &User{
					Name: "Alhamsya",
					Age:  10,
				},
				newData: &Dummy{},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GenerateDiff(tt.args.originData, tt.args.newData)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateDiff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
