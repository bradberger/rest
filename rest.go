// Package rest implements a simple REST microframework designed for Google App Engine
package rest

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/context"

	"github.com/gorilla/mux"

	"google.golang.org/appengine"
)

// AppHandler is the wrapper for all HTTP requests. It provides a valid context, authorization information, and route parameters.
// The returned interface is written to the response, unless an error is returned.
type AppHandler func(ctx context.Context) error

// Decode decodes the http request body in JSON format into the dst variable.
func Decode(ctx context.Context, dst interface{}) error {
	b := ctx.Value(ContextKeyRequestBody)
	if b == nil {
		return errors.New("no request body")
	}
	return json.Unmarshal(b.([]byte), dst)
}

// FormValue returns the form value (or mux.Vars value) from the request
func FormValue(ctx context.Context, key string) string {
	v := ctx.Value(ContextKeyRequestVars)
	if v == nil {
		return ""
	}
	return v.(map[string]string)[key]
}

// ResponseWriter returns the response writer for the given context
func ResponseWriter(ctx context.Context) http.ResponseWriter {
	return ctx.Value(ContextKeyResponseWriter).(http.ResponseWriter)
}

// Request returns the http.Request asso
func Request(ctx context.Context) *http.Request {
	return ctx.Value(ContextKeyRequest).(*http.Request)
}

// Header returns the ResponseWriter headers
func Header(ctx context.Context) http.Header {
	return ResponseWriter(ctx).Header()
}

// Headers returns the http.Request headers
func Headers(ctx context.Context) http.Header {
	return Request(ctx).Header
}

// Hostname returns the hostname of the current instance
func Hostname(ctx context.Context) (string, error) {
	return appengine.ModuleHostname(ctx, "", "", "")
}

func setRequest(ctx context.Context, r *http.Request) context.Context {
	return context.WithValue(ctx, ContextKeyRequest, r)
}

func setVars(ctx context.Context) (context.Context, error) {

	v := map[string]string{}
	r := Request(ctx)

	err := r.ParseForm()
	if err != nil {
		return ctx, err
	}

	for i := range r.Form {
		v[i] = r.FormValue(i)
	}
	for i := range mux.Vars(r) {
		v[i] = mux.Vars(r)[i]
	}

	ctx = context.WithValue(ctx, ContextKeyRequestVars, v)
	return ctx, nil
}

func setBody(ctx context.Context) (context.Context, error) {

	r := Request(ctx)
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, ContextKeyRequestBody, bodyBytes), nil
}

func setWriter(ctx context.Context, w http.ResponseWriter) context.Context {
	return context.WithValue(ctx, ContextKeyResponseWriter, w)
}

// initRequest returns a context with the user and other context variables set
func initRequest(w http.ResponseWriter, r *http.Request) (ctx context.Context, err error) {

	ctx = appengine.NewContext(r)
	ctx = setRequest(ctx, r)
	ctx = setWriter(ctx, w)

	// If custom namespace handler set then execute it here
	if Namespace != nil {
		ns, err := Namespace(ctx)
		if err != nil {
			return ctx, err
		}
		ctx, err = appengine.Namespace(ctx, ns)
		if err != nil {
			return ctx, err
		}
	}

	if ctx, err = setVars(ctx); err != nil {
		return
	}
	if ctx, err = setBody(ctx); err != nil {
		return
	}
	return
}

// processRequest calls the controller function. If there's an error, then return a 500 template or a basic message.
func processRequest(ctx context.Context, fn func(ctx context.Context) error) {

	err := fn(ctx)
	if err == nil {
		return
	}

	OnError(ctx, GetErrorCode(ctx, err), err)
}

// ServeHTTP provides authorization and write wrappers for http controllers and responses
func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var ctx context.Context
	var err error

	// Catch panics with the custom handler, if it's not nil
	defer func() {
		if OnPanic != nil {
			if ctx == nil {
				ctx = appengine.NewContext(r)
			}
			defer OnPanic(ctx)
		}
	}()

	ctx, err = initRequest(w, r)
	if err != nil {
		Criticalf(ctx, "Error bootstrapping request: %v", err)
	}

	processRequest(ctx, fn)
}

func New() *mux.Router {
	r := mux.NewRouter()
	r.KeepContext = true
	return r
}
