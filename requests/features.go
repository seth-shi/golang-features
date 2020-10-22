package requests

import (
	"github.com/go-playground/validator/v10"
)

type FeatureRequest struct {
	Title 		string 		`form:"title"  binding:"required,lt=100"`
	Description  string	    `form:"description"   binding:"required,lt=100"`
	Code  string	    	`form:"code"   binding:"required,lt=2000"`
}

func (r *FeatureRequest) GetError (err validator.ValidationErrors) string {


	for _, val := range err {

		switch val.Field() + "." + val.Tag() {

		case "Title.required":
			return "标题不能为空"
		case "Title.lt":
			return "标题太多了"
		case "Description.required":
			return "描述不能为空"
		case "Description.lt":
			return "描述太多了"
		case "Code.required":
			return "代码不能为空"
		case "Code.lt":
			return "代码太多了"

		}
	}

	return "未知错误"
}