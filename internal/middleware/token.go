package middleware

import (
	"context"
	"net/http"
	"projectName/internal/model"
	"projectName/internal/repository"
	"projectName/pkg/helper/resp"
	"strings"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
)

const (
	defaultMemory = 32 << 20
)

func TokenMiddleware(repository *repository.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {

		user := model.User{}
		params := GetQueryMap(c.Request)
		token, ok := params["token"]
		if !ok {
			resp.HandleError(c, http.StatusBadRequest, model.CodeParamErr, "", nil)
			c.Abort()
			return
		}
		exist, err := repository.GetData(token, &user)
		if err != nil {
			resp.HandleError(c, http.StatusBadRequest, model.TokenExpErr, "", nil)
			c.Abort()
			return
		}
		if !exist || user.ID <= 0 {
			resp.HandleError(c, http.StatusBadRequest, model.TokenExpErr, "", nil)
			c.Abort()
			return
		}

		context.WithValue(c, "user_id", user.ID)
		c.Next()
	}
}

func GetQueryMap(r *http.Request) map[string]string {

	method := strings.ToUpper(r.Method)
	var params map[string][]string

	if method == http.MethodGet {
		// get方法获取所有url params
		params = r.URL.Query()

	} else if method == http.MethodPost {
		//根据 post Content-Type 类型选择获取参数的方法
		contentType := ContentType(r)
		switch contentType {
		case binding.MIMEMultipartPOSTForm:
			//post方法获取所有的form-data的参数
			err := r.ParseMultipartForm(defaultMemory)
			if err != nil {
				return nil
			}
			MultiParam := r.MultipartForm
			if len(MultiParam.Value) == 0 {
				break
			}
			params = MultiParam.Value
		default:
			//post方法获取所有的x-www-form-urlencoded的参数
			err := r.ParseForm()
			if err != nil {
				return nil
			}
			params = r.PostForm
		}

		//如果params为空，尝试通过url params方式获取
		if len(params) == 0 {
			params = r.URL.Query()
		}
	}
	if len(params) == 0 {
		return nil
	}
	res := make(map[string]string)
	for key, value := range params {
		res[key] = value[0]
	}
	return res
}
func ContentType(r *http.Request) string {
	return filterFlags(requestHeader(r, "Content-Type"))
}
func requestHeader(r *http.Request, key string) string {
	return r.Header.Get(key)
}

func filterFlags(content string) string {
	for i, char := range content {
		if char == ' ' || char == ';' {
			return content[:i]
		}
	}
	return content
}
