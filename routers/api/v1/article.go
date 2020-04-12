package v1

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/jeristiano/wendo/model"
	setting "github.com/jeristiano/wendo/pkg"
	error2 "github.com/jeristiano/wendo/pkg/error"
	"github.com/jeristiano/wendo/pkg/util"
	"github.com/unknwon/com"
	"log"
	"net/http"
)

/**
获得单个文章
*/
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id必须大于0")
	code := error2.INVALID_PARAMS

	var data interface{}

	if !valid.HasErrors() {
		if model.ExistArticleByID(id) {
			data = model.GetArticle(id)
			code = error2.SUCCESS
		} else {
			code = error2.ERROR_NOT_EXIST_ARTICLE
		}

	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key:%s,err.message:%s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error2.GetMsg(code),
		"data": data,
	})
}

func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
		valid.Range(state, 0, 1, "state").Message("状态只允许0，1")
	}
	var tagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId

		valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	}
	code := error2.INVALID_PARAMS

	if !valid.HasErrors() {
		code = error2.SUCCESS
		data["lists"] = model.GetArticles(util.GetPage(c), setting.PageSize, maps)
		data["total"] = model.GetArticleTotal(maps)
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error2.GetMsg(code),
		"data": data,
	})
}

func AddArticle(c *gin.Context) {
	tagId := com.StrTo(c.PostForm("tag_id")).MustInt()
	title := c.PostForm("title")
	desc := c.PostForm("desc")
	content := c.PostForm("content")
	createdBy := c.PostForm("created_by")
	state := com.StrTo(c.DefaultPostForm("state", "0")).MustInt()
	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	code := error2.INVALID_PARAMS
	if !valid.HasErrors() {
		code = error2.SUCCESS
		data := make(map[string]interface{})
		data["tag_id"] = tagId
		data["title"] = title
		data["desc"] = desc
		data["content"] = content
		data["created_by"] = createdBy
		data["state"] = state
		model.AddArticle(data)

	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	fmt.Println(code)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error2.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

func EditArticle(c *gin.Context) {
	valid := validation.Validation{}

	id := com.StrTo(c.PostForm("id")).MustInt()
	tagId := com.StrTo(c.PostForm("tag_id")).MustInt()
	title := c.PostForm("title")
	desc := c.PostForm("desc")
	content := c.PostForm("content")
	modifiedBy := c.PostForm("modified_by")

	var state int = -1
	if arg := c.PostForm("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	code := error2.INVALID_PARAMS

	if !valid.HasErrors() {
		if model.ExistArticleByID(id) {
			if model.ExistTagByID(tagId) {
				data := make(map[string]interface{})
				if tagId > 0 {
					data["tag_id"] = tagId
				}
				if title != "" {
					data["title"] = title
				}
				if desc != "" {
					data["desc"] = desc
				}
				if content != "" {
					data["content"] = content
				}

				data["modified_by"] = modifiedBy

				model.EditArticle(id, data)
				code = error2.SUCCESS
			} else {
				code = error2.ERROR_NOT_EXIST_TAG
			}
		} else {
			code = error2.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error2.GetMsg(code),
		"data": make(map[string]string),
	})

}

func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := error2.INVALID_PARAMS
	if !valid.HasErrors() {
		if model.ExistArticleByID(id) {
			model.DeleteArticle(id)
			code = error2.SUCCESS
		} else {
			code = error2.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error2.GetMsg(code),
		"data": make(map[string]string),
	})
}
