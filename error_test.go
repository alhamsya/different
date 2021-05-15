package different

import (
	"errors"
	"testing"
)

func TestGenerateError(t *testing.T) {
	type args struct {
		err  error
		flag string
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "When_errorWrapUsingFlagReturnSuccess_Expect_Success",
			args: args{
				err:  errors.New("some error"),
				flag: "tes",
				args: []string{"unmarshal"},
			},
			wantErr: true,
		},
		{
			name: "When_errParamNilReturnSuccess_Expect_Success",
			args: args{
				err:  nil,
				flag: "tes",
				args: []string{"unmarshal"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GenerateError(tt.args.err, tt.args.flag, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("GenerateError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
