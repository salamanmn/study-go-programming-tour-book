package errcode

var (
	//新增业务错误码
	ErrorGetTagListFail = NewError(20000001, "获取标签列表失败")
	ErrorCreateTagFail  = NewError(20000002, "创建标签失败")
	ErrorUpdateTagFail  = NewError(20000003, "更新标签失败")
	ErrorDeleteTagFail  = NewError(20000004, "删除标签失败")
	ErrorCountTagFail   = NewError(20000005, "统计标签失败")

	ErrorUploadFileFail = NewError(20030001, "上传文件失败")
)
