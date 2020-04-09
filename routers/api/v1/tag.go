package v1

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	_ "github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/jeristiano/wendo/model"
	setting "github.com/jeristiano/wendo/pkg"
	"github.com/jeristiano/wendo/pkg/error"
	"github.com/jeristiano/wendo/pkg/util"
	"github.com/unknwon/com"
	"net/http"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	code := error.SUCCESS
	data["lists"] = model.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["totals"] = model.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})

}

//添加文章标签
func AddTag(c *gin.Context) {

	name, _ := c.GetPostForm("name")
	state := com.StrTo(c.DefaultPostForm("state", "0")).MustInt()
	createdBY, _ := c.GetPostForm("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名字不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBY, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBY, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	code := error.INVALID_PARAMS

	if !valid.HasErrors() {
		if !model.ExistTagByName(name) {
			code = error.SUCCESS
			model.AddTag(name, state, createdBY)
		} else {
			code = error.ERROR_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": make(map[string]string),
	})
}

//改
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	name := c.PostForm("name")
	modifiedBy := c.PostForm("modified_by")
	valid := validation.Validation{}

	var state int = -1
	if arg := c.PostForm("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Required(name, "name").Message("名字不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	code := error.INVALID_PARAMS

	fmt.Println(id, name, modifiedBy, valid.ErrorMap())

	if !valid.HasErrors() {
		code = error.SUCCESS
		if model.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}

			model.EditTag(id, data)
		} else {
			code = error.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": make(map[string]string),
	})
}

//删除
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	code := error.INVALID_PARAMS
	if !valid.HasErrors() {
		code = error.SUCCESS
		if model.ExistTagByID(id) {
			model.DeleteTag(id)
		} else {
			code = error.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": make(map[string]string),
	})
}
