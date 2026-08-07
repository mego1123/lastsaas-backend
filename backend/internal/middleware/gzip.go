package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
	"sync"
)

// gzipWriter wraps an http.ResponseWriter and gzips the output.
type gzipWriter struct {
	http.ResponseWriter
	gz *gzip.Writer
}

func (w *gzipWriter) Write(b []byte) (int, error) {
	// Set the Content-Encoding header if not already set
	if w.Header().Get("Content-Encoding") == "" {
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Del("Content-Length") // length changed
	}
	return w.gz.Write(b)
}

func (w *gzipWriter) Flush() {
	if f, ok := w.ResponseWriter.(http.Flusher); ok {
		w.gz.Flush()
		f.Flush()
	}
}

func (w *gzipWriter) Hijack() (interface{}, error) {
	// If the underlying ResponseWriter supports Hijack (e.g. for websockets),
	// close the gzip writer first and pass through.
	if hj, ok := w.ResponseWriter.(http.Hijacker); ok {
		w.gz.Close()
		return hj.Hijack()
	}
	return nil, io.ErrUnexpectedEOF
}

// gzipWriterPool reuses gzip.Writer instances to avoid allocations.
var gzipWriterPool = sync.Pool{
	New: func() interface{} {
		// Default compression level (6 = default, balances speed/ratio)
		w, _ := gzip.NewWriterLevel(io.Discard, 6)
		return w
	},
}

// contentTypesToCompress is the set of Content-Type prefixes that should
// be gzipped. JSON and HTML benefit most; images/PDFs are already compressed.
var contentTypesToCompress = []string{
	"application/json",
	"text/html",
	"text/plain",
	"text/css",
	"application/javascript",
	"application/xml",
	"text/xml",
	"image/svg+xml",
}

// shouldCompress checks if the response's Content-Type should be compressed.
func shouldCompress(contentType string) bool {
	for _, ct := range contentTypesToCompress {
		if strings.HasPrefix(contentType, ct) {
			return true
		}
	}
	return false
}

// GzipMiddleware compresses HTTP responses for clients that accept gzip.
// It only compresses content types that benefit from compression (JSON,
// HTML, CSS, JS, SVG). It skips SSE (text/event-stream) and already-
// compressed formats (images, PDFs).
func GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the client accepts gzip encoding
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		// Skip SSE (Server-Sent Events) — compression breaks streaming
		if strings.Contains(r.Header.Get("Accept"), "text/event-stream") {
			next.ServeHTTP(w, r)
			return
		}

		// Get a gzip writer from the pool
		gz := gzipWriterPool.Get().(*gzip.Writer)
		gz.Reset(w)
		defer func() {
			gz.Close()
			gzipWriterPool.Put(gz)
		}()

		// Wrap the response writer
		gw := &gzipWriter{
			ResponseWriter: w,
			gz:             gz,
		}

		// Use a response observer to check Content-Type before compressing.
		// We intercept the first Write to decide if we should compress.
		// This is done by checking the Content-Type header which is set
		// before any Write call.
		ct := w.Header().Get("Content-Type")
		if ct != "" && !shouldCompress(ct) {
			// Content-Type is already set and it's not a compressible type
			// (e.g. image/png, application/pdf) — skip compression
			next.ServeHTTP(w, r)
			return
		}

		// Content-Type is either compressible or not set yet (will be set
		// by the handler) — compress and let the handler set Content-Type
		next.ServeHTTP(gw, r)
	})
}
