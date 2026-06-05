package middleware

import (
	"net/http"
	"time"
	"log"
)

type wrappedWriter struct {
    http.ResponseWriter
    statusCode int
}

func (w *wrappedWriter) WriterHeader(statusCode int){
    w.ResponseWriter.WriteHeader(statusCode)
    w.statusCode = statusCode
}


func Logging (next http.Handler) http.Handler {
    //闭包内部的代码构成了中间件
    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
        start := time.Now()     //捕获请求开始的时间
        wrapped := &wrappedWriter{
            ResponseWriter: w,
            statusCode: http.StatusOK,
        }
        next.ServeHTTP(wrapped,r) //ServeHTTP由http.Handler实现
        log.Println(wrapped.statusCode,r.Method,r.URL.Path,time.Since(start))
    })
}
