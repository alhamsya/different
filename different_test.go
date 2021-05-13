package different

import (
	"errors"
	"reflect"
	"testing"

	"github.com/r3labs/diff/v2"
	"github.com/undefinedlabs/go-mpatch"
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
		patch   func()
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
			patch:   func() {},
			want:    []byte(`[{"Name":{"before":"Alhamsya","after":"Bintang"}}]`),
			wantErr: false,
		},
		{
			name: "When_deepEqual_ReturnNil",
			args: args{
				origin: nil,
				new:    nil,
			},
			patch:   func() {},
			want:    nil,
			wantErr: false,
		},
		{
			name: "When_diff_ReturnError",
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
			patch: func() {
				var guard *mpatch.Patch
				guard, _ = mpatch.PatchMethod(diff.Diff, func(a, b interface{}, opts ...func(d *diff.Differ) error) (diff.Changelog, error) {
					defer guard.Unpatch()
					return nil, errors.New("error diff.Diff")
				})
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "When_typeDiffNotUpdate_ReturnError",
			args: args{
				origin: &User{
					Name: "Alhamsya",
					Age:  10,
				},
				new: nil,
			},
			patch:   func() {},
			want:    nil,
			wantErr: true,
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
			patch:   func() {},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.patch()
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
