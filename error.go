package different

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

//GenerateError generate builder error
func GenerateError(err error, flag string, args ...string) error {
	var msg string
	var strPkg string
	if strPkg = strings.TrimSpace(flag); strPkg != "" {
		strPkg = fmt.Sprintf("[%s] ", strings.ToUpper(flag))
	}

	msg = strings.Join(args, " | ")
	msg = strPkg + msg

	if err == nil {
		err = errors.New(fmt.Sprintf("[FOR DEVELOPER] forget to set error in %s", flag))
	}

	return errors.Wrap(err, msg)
}
