package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"goweb-quick-start/api/model"
	"goweb-quick-start/api/model/base"
	module "goweb-quick-start/module/student"
	"net/http"
	"runtime/debug"
	"strconv"
)

func FrontendGetStudent(c *gin.Context) {
	//id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var req model.StudentListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, base.ResponseErr(errors.New(http.StatusText(http.StatusBadRequest))))
		return
	}
	students, _ := module.FindStudent(req)
	c.JSON(http.StatusOK, base.ResponseData(students))
	return
}

func FrontendPostStudent(c *gin.Context) {
	args := model.Student{}
	if err := c.ShouldBindBodyWith(&args, binding.JSON); err != nil {
		c.JSON(http.StatusOK, base.ResponseErr(err))
		return
	}
	if err := module.SaveStudent(args); err != nil {
		logrus.Errorf("SaveStudent error:%+v,堆栈信息:%+v", err, string(debug.Stack()))
		c.JSON(http.StatusInternalServerError, base.ResponseErr(errors.New(http.StatusText(http.StatusInternalServerError))))
		return
	}
	c.JSON(http.StatusOK, base.ResponseSuccess())
	return
}

func FrontendDeleteStudent(c *gin.Context) {
	//args := c.Param("id") // .../:id
	//args := c.Query("id") // ...?id=
	//id, _ := strconv.Atoi(c.Param("id"))
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	module.DeleteStudent(id)
	c.JSON(http.StatusOK, base.ResponseSuccess())
	return
}

func FrontendPutStudent(c *gin.Context) {
	args := model.Student{}
	if err := c.ShouldBindBodyWith(&args, binding.JSON); err != nil {
		c.JSON(http.StatusOK, base.ResponseErr(err))
		return
	}
	if err := module.UpdateStudent(args); err != nil {
		logrus.Errorf("SaveStudent error:%+v,堆栈信息:%+v", err, string(debug.Stack()))
		c.JSON(http.StatusInternalServerError, base.ResponseErr(errors.New(http.StatusText(http.StatusInternalServerError))))
		return
	}
	c.JSON(http.StatusOK, base.ResponseSuccess())
	return
}
