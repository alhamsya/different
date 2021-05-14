package different

import (
	"errors"
	"reflect"

	"github.com/r3labs/diff/v2"

	jsoniter "github.com/json-iterator/go"
)

//GenerateDiff generate different data
func GenerateDiff(originData interface{}, newData interface{}) ([]byte, error) {
	//compare deep equal struct
	if reflect.DeepEqual(originData, newData) {
		return nil, nil
	}

	//compare data
	changelog, err := diff.Diff(originData, newData)
	if err != nil {
		return nil, GenerateError(err, "library")
	}

	//builder
	var temp []interface{}
	for _, data := range changelog {
		if data.Type != diff.UPDATE {
			return nil, GenerateError(errors.New("type different struct not update"), "originData")
		}

		var res interface{}
		for i := len(data.Path) - 1; i >= 0; i-- {
			if i == len(data.Path)-1 {
				res = map[string]interface{}{
					data.Path[i]: map[string]interface{}{
						"before": data.From,
						"after":  data.To,
					},
				}
			} else {
				res = map[string]interface{}{
					data.Path[i]: res,
				}
			}
		}
		temp = append(temp, res)
	}

	//generate json marshal
	result, err := jsoniter.Marshal(&temp)
	if err != nil {
		return nil, GenerateError(err, "jsoniter")
	}

	//result different data
	return result, nil
}
