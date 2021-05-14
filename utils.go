package different

import "github.com/r3labs/diff/v2"

//BuildBeforeAfter
func BuildBeforeAfter(data diff.Change) (result interface{}) {
	for i := len(data.Path) - 1; i >= 0; i-- {
		if i == len(data.Path)-1 {
			result = map[string]interface{}{
				data.Path[i]: map[string]interface{}{
					"before": data.From,
					"after":  data.To,
				},
			}
		} else {
			result = map[string]interface{}{
				data.Path[i]: result,
			}
		}
	}

	return result
}
