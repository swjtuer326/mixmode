package logic

import (
	"github.com/gin-gonic/gin"
	"sophgo.com/mixmode/model/response"
	"sophgo.com/mixmode/utils/sophon"
)

type SophonTestLogic struct{}

func (l SophonTestLogic) TestCDMA(c *gin.Context) (interface{}, interface{}) {
	res, err := sophon.TestCDMA()
	if err != nil {
		return nil, err
	}
	return response.SophonTestRsp{Res: res}, nil
}
func (l SophonTestLogic) TestGDMA(c *gin.Context) (interface{}, interface{}) {
	res, err := sophon.TestGDMA()
	if err != nil {
		return nil, err
	}
	return response.SophonTestRsp{Res: res}, nil
}
