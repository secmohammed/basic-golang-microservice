package middlewares

import (
    "compress/gzip"
    "net/http"
    "strings"
)

// GzipHandler struct is used to define the handler of gzip's content.
type GzipHandler struct {
}

//WrappedResponseWriter struct.
type WrappedResponseWriter struct {
    rw http.ResponseWriter
    gw *gzip.Writer
}

func NewWrappedResponseWriter(rw http.ResponseWriter) *WrappedResponseWriter {
    gw := gzip.NewWriter(rw)
    return &WrappedResponseWriter{
        rw: rw,
        gw: gw,
    }
}

//GzipMiddleware is used to check if the request header's contain gzip header.
func (gzip *GzipHandler) GzipMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
        if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
            wrw := NewWrappedResponseWriter(rw)
            wrw.Header().Set("Content-Encoding", "gzip")
            next.ServeHTTP(wrw, r)
            defer wrw.Flush()
            return
        }
        next.ServeHTTP(rw, r)
    })
}
func (wr *WrappedResponseWriter) Header() http.Header {
    return wr.rw.Header()
}
func (wr *WrappedResponseWriter) WriteHeader(statuscode int) {
    wr.rw.WriteHeader(statuscode)
}
func (wr *WrappedResponseWriter) Write(d []byte) (int, error) {
    return wr.gw.Write(d)
}
func (wr *WrappedResponseWriter) Flush() {
    wr.gw.Flush()
    wr.gw.Close()
}
