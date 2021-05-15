package different

import (
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
		return nil, generateError(err, "library")
	}

	//builder different value
	var temp []interface{}
	for _, data := range changelog {
		var res interface{}
		if len(data.Path) == 0 {
			res = map[string]interface{}{
				"before": data.From,
				"after":  data.To,
			}
		} else {
			res = buildBeforeAfter(data)
		}
		temp = append(temp, res)
	}

	//generate json marshal
	result, err := jsoniter.Marshal(&temp)
	if err != nil {
		return nil, generateError(err, "jsoniter")
	}

	//result different data
	return result, nil
}
