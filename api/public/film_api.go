package public

import (
	"fmt"
	"gin_api_02/global"
	"gin_api_02/model/common/response"
	res "gin_api_02/model/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetFilmInfo
// @Tags      Film
// @Summary   分页获取Film列表
// @accept    application/json
// @Produce   application/json
// @Param     data  query     res.FilmQuery                       true  "页码, 每页大小, 搜索条件"
// @Success   200   {object}  response.Response{data=response.QueryResult,msg=string}  "FilmInfo列表,返回包括列表,总数,页码,每页数量"
// @Router    /public/film [get]
func (s *PublicApi) GetFilmInfo(c *gin.Context) {

	var queryModel res.FilmQuery
	err := c.ShouldBindQuery(&queryModel)
	fmt.Println(c.Request.URL.Query())
	fmt.Println(queryModel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := filmService.GetMfilmInfo(queryModel)
	if err != nil {
		global.SYS_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.QueryResult{
		List:       list,
		Total:      total,
		PageNumber: queryModel.PageNumber,
		PageSize:   queryModel.PageSize,
	}, "获取成功", c)
}
