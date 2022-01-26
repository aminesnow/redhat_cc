// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteObjectHandlerFunc turns a function with the right signature into a delete object handler
type DeleteObjectHandlerFunc func(DeleteObjectParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteObjectHandlerFunc) Handle(params DeleteObjectParams) middleware.Responder {
	return fn(params)
}

// DeleteObjectHandler interface for that can handle valid delete object params
type DeleteObjectHandler interface {
	Handle(DeleteObjectParams) middleware.Responder
}

// NewDeleteObject creates a new http.Handler for the delete object operation
func NewDeleteObject(ctx *middleware.Context, handler DeleteObjectHandler) *DeleteObject {
	return &DeleteObject{Context: ctx, Handler: handler}
}

/* DeleteObject swagger:route DELETE /objects/{bucket}/{objectID} deleteObject

Deletes an object

*/
type DeleteObject struct {
	Context *middleware.Context
	Handler DeleteObjectHandler
}

func (o *DeleteObject) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteObjectParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
