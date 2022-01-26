// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteObjectOKCode is the HTTP code returned for type DeleteObjectOK
const DeleteObjectOKCode int = 200

/*DeleteObjectOK OK

swagger:response deleteObjectOK
*/
type DeleteObjectOK struct {
}

// NewDeleteObjectOK creates DeleteObjectOK with default headers values
func NewDeleteObjectOK() *DeleteObjectOK {

	return &DeleteObjectOK{}
}

// WriteResponse to the client
func (o *DeleteObjectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteObjectNotFoundCode is the HTTP code returned for type DeleteObjectNotFound
const DeleteObjectNotFoundCode int = 404

/*DeleteObjectNotFound Object not found

swagger:response deleteObjectNotFound
*/
type DeleteObjectNotFound struct {
}

// NewDeleteObjectNotFound creates DeleteObjectNotFound with default headers values
func NewDeleteObjectNotFound() *DeleteObjectNotFound {

	return &DeleteObjectNotFound{}
}

// WriteResponse to the client
func (o *DeleteObjectNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// DeleteObjectInternalServerErrorCode is the HTTP code returned for type DeleteObjectInternalServerError
const DeleteObjectInternalServerErrorCode int = 500

/*DeleteObjectInternalServerError Failed to delete object.

swagger:response deleteObjectInternalServerError
*/
type DeleteObjectInternalServerError struct {
}

// NewDeleteObjectInternalServerError creates DeleteObjectInternalServerError with default headers values
func NewDeleteObjectInternalServerError() *DeleteObjectInternalServerError {

	return &DeleteObjectInternalServerError{}
}

// WriteResponse to the client
func (o *DeleteObjectInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
