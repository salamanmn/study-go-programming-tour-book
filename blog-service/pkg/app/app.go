/*响应处理*/
package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	Page int `json:"page"` 				//当前页号
	PageSize int `json:"page_size"`		//页大小
	TotalRows int `json:"total_rows"`	//符合查询条件的所有记录数
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{ctx}
}

//将数据data，以json格式存到响应体中，发送响应报文
func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

//这里的list是什么意思？
func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list": list, 
		"pager": Pager{
			Page:	GetPage(r.Ctx),
			PageSize: GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

//将错误的信息，以json格式存到响应体中，发送响应报文
func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}

	r.Ctx.JSON(err.StatusCode(), response)
}
