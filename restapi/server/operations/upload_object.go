// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UploadObjectHandlerFunc turns a function with the right signature into a upload object handler
type UploadObjectHandlerFunc func(UploadObjectParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UploadObjectHandlerFunc) Handle(params UploadObjectParams) middleware.Responder {
	return fn(params)
}

// UploadObjectHandler interface for that can handle valid upload object params
type UploadObjectHandler interface {
	Handle(UploadObjectParams) middleware.Responder
}

// NewUploadObject creates a new http.Handler for the upload object operation
func NewUploadObject(ctx *middleware.Context, handler UploadObjectHandler) *UploadObject {
	return &UploadObject{Context: ctx, Handler: handler}
}

/* UploadObject swagger:route PUT /objects/{bucket}/{objectID} uploadObject

Upload object to the service.

*/
type UploadObject struct {
	Context *middleware.Context
	Handler UploadObjectHandler
}

func (o *UploadObject) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUploadObjectParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}