package response

import "sophgo.com/mixmode/model"

type LogListRsp struct {
	Total int64                `json:"total"`
	Logs  []model.OperationLog `json:"logs"`
}
