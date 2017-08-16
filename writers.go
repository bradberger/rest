package rest

import (
	"encoding/json"
	"encoding/xml"
	"image"
	"image/jpeg"
	"image/png"

	"net/http"

	"golang.org/x/net/context"
)

// Text writes the string to the HTTP connection as text/plain content type
func Text(ctx context.Context, code int, str string) error {
	w := ResponseWriter(ctx)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(code)
	_, err := w.Write([]byte(str))
	return err
}

// Error writes the error string as the HTTP response
func Error(ctx context.Context, code int, err error) error {
	return Text(ctx, code, err.Error())
}

// Status writes the HTTP response code with the default status text for that code
func Status(ctx context.Context, code int) error {
	return Text(ctx, code, http.StatusText(code))
}

// Redirect redirects the request to the given location
func Redirect(ctx context.Context, urlStr string, code int) error {
	http.Redirect(ResponseWriter(ctx), Request(ctx), urlStr, code)
	return nil
}

// HTML writes the raw HTML to the HTTP connection
func HTML(ctx context.Context, code int, html string) error {
	w := ResponseWriter(ctx)
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(code)
	if _, err := w.Write([]byte(html)); err != nil {
		return err
	}
	return nil
}

// CSS writes the raw CSS to the HTTP connection
func CSS(ctx context.Context, code int, css string) error {
	w := ResponseWriter(ctx)
	w.Header().Set("Content-Type", "text/css")
	w.WriteHeader(code)
	if _, err := w.Write([]byte(css)); err != nil {
		return err
	}
	return nil
}

// JSON writes the encoded data to the HTTP connection
func JSON(ctx context.Context, code int, data interface{}) error {
	w := ResponseWriter(ctx)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	switch data.(type) {
	case []byte:
		_, err := w.Write(data.([]byte))
		return err
	default:
		return json.NewEncoder(w).Encode(data)
	}
}

// XML writes the XML encoded data to the HTTP connection
func XML(ctx context.Context, code int, data interface{}) error {
	w := ResponseWriter(ctx)
	w.Header().Set("Content-Type", "text/xml; charset=utf-8")
	w.WriteHeader(code)
	if err := xml.NewEncoder(w).Encode(data); err != nil {
		return err
	}
	return nil
}

// PNG writes the image to the HTTP connection
func PNG(ctx context.Context, code int, img image.Image) error {
	w := ResponseWriter(ctx)
	w.Header().Set("Content-Type", "image/png")
	enc := &png.Encoder{CompressionLevel: png.BestCompression}
	return enc.Encode(w, img)
}

// JPEG writes the image to the HTTP connection
func JPEG(ctx context.Context, code int, img image.Image) error {
	w := ResponseWriter(ctx)
	w.Header().Set("Content-Type", "image/jpeg")
	return jpeg.Encode(w, img, &jpeg.Options{Quality: 100})
}

// NoContent handles responses without any content
func NoContent(ctx context.Context) error {
	w := ResponseWriter(ctx)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusNoContent)
	return nil
}

// Bytes writes the bytes to the HTTP response with the given code and content type
func Bytes(ctx context.Context, code int, contentType string, b []byte) error {
	w := ResponseWriter(ctx)
	w.WriteHeader(code)
	w.Header().Set("Content-Type", contentType)
	_, err := w.Write(b)
	return err
}

// Font serves the byte slice as a truetype font
func Font(ctx context.Context, b []byte) error {
	return Bytes(ctx, 200, "application/x-font-truetype", b)
}

// Unauthorized returns a 401 error with 401 status text
func Unauthorized(ctx context.Context) error {
	return Status(ctx, http.StatusUnauthorized)
}
