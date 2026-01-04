package seedgo

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w CustomResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func sliceContain(sli []string, k string) bool {
	for _, i := range sli {
		if i == k {
			return true
		}
	}
	return false
}

func AccessLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//headerList := []string{"Authorization", "X-Tid"}
		headerList := []string{"X-Tid"}

		t := time.Now()
		body, _ := ctx.GetRawData()
		var header string
		for k, v := range ctx.Request.Header {
			if sliceContain(headerList, k) {
				header = header + k + ":" + strings.Join(v, ",") + ";"
			}
		}

		bodystr := ""
		// ingore file upload body, beacuse is too big
		contentType := ctx.Request.Header.Get("Content-Type")
		if !strings.Contains(contentType, "multipart/form-data") {
			bodystr = string(body)
		}
		Infof(ctx, "access log request, uri: %s, method: %s, header: %s, params: %s",
			ctx.Request.RequestURI,
			ctx.Request.Method,
			header,
			bodystr,
		)
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		blw := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = blw
		ctx.Next()
		// after request

		costtime := time.Since(t).Microseconds()
		Infof(ctx, "access log response, costtime: %dms, result: %s", costtime, blw.body.String())
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				Fail(c, fmt.Sprint(err), 400)
				return
			}
		}()
		c.Next()
	}
}

func TraceMiddware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tid := ctx.GetHeader("X-Tid")
		if len(tid) == 0 {
			tid = NextUid()
		}

		ctx.Set("tid", tid)

		ctx.Next()
	}

}
