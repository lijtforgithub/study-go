package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zht "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"reflect"
	"strings"
)

type validStruct struct {
	ID   string `validate:"required"`
	Name string `validate:"required,len=2"`
	Age  int    `validate:"required,gte=1"`
}

type validForm struct {
	ID    string `form:"id" binding:"required,uuid"`
	Name  string `form:"name" binding:"required,min=2,max=5"`
	Email string `form:"email" binding:"required,email"`
	Age   int    `form:"age" binding:"required,gte=1"`
}

var (
	validate  *validator.Validate
	trans     ut.Translator
	bindTrans ut.Translator
)

func main() {
	err := initValidate()
	fmt.Println("校验结构体", err)
	validStructMethod()

	server := gin.Default()

	// 校验表单
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhCn := zh.New()
		uni := ut.New(zhCn, zhCn)
		bindTrans, _ = uni.GetTranslator("zh")

		err := zht.RegisterDefaultTranslations(v, bindTrans)
		fmt.Println("校验表单", err)

		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			return field.Tag.Get("form")
		})
	}
	server.POST("/valid/form", func(c *gin.Context) {
		var vf validForm
		err := c.ShouldBind(&vf)
		if err != nil {
			var verr validator.ValidationErrors
			if errors.As(err, &verr) {
				//c.JSON(http.StatusBadRequest, verr.Translate(bindTrans))
				c.JSON(http.StatusBadRequest, formatErrors(verr))
			} else {
				c.String(http.StatusInternalServerError, err.Error())
			}
			return
		}

		c.String(http.StatusOK, "表单校验通过")
	})

	server.Run()
}

// 格式化错误信息，去除字段名
func formatErrors(err error) string {
	var errs []string
	for _, e := range err.(validator.ValidationErrors) {
		errs = append(errs, e.Translate(bindTrans))
	}
	return strings.Join(errs, ", ")
}

func initValidate() error {
	zhCn := zh.New()
	uni := ut.New(zhCn, zhCn)
	trans, _ = uni.GetTranslator("zh")

	// 校验结构体
	validate = validator.New()
	return zht.RegisterDefaultTranslations(validate, trans)
}

func validStructMethod() {
	vs := validStruct{
		ID:   "xxoo",
		Name: "你吃饭了吗",
	}

	err := validate.Struct(vs)
	if err != nil {
		var errs validator.ValidationErrors
		if errors.As(err, &errs) {
			for k, v := range errs.Translate(trans) {
				fmt.Println(k, v)
			}
		} else {
			fmt.Println(err)
		}
	}
}
