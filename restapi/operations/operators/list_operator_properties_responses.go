// Code generated by go-swagger; DO NOT EDIT.

package operators

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// ListOperatorPropertiesOKCode is the HTTP code returned for type ListOperatorPropertiesOK
const ListOperatorPropertiesOKCode int = 200

/*ListOperatorPropertiesOK Success.

swagger:response listOperatorPropertiesOK
*/
type ListOperatorPropertiesOK struct {

	/*
	  In: Body
	*/
	Payload models.OperatorProperties `json:"body,omitempty"`
}

// NewListOperatorPropertiesOK creates ListOperatorPropertiesOK with default headers values
func NewListOperatorPropertiesOK() *ListOperatorPropertiesOK {

	return &ListOperatorPropertiesOK{}
}

// WithPayload adds the payload to the list operator properties o k response
func (o *ListOperatorPropertiesOK) WithPayload(payload models.OperatorProperties) *ListOperatorPropertiesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list operator properties o k response
func (o *ListOperatorPropertiesOK) SetPayload(payload models.OperatorProperties) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListOperatorPropertiesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.OperatorProperties{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListOperatorPropertiesUnauthorizedCode is the HTTP code returned for type ListOperatorPropertiesUnauthorized
const ListOperatorPropertiesUnauthorizedCode int = 401

/*ListOperatorPropertiesUnauthorized Unauthorized.

swagger:response listOperatorPropertiesUnauthorized
*/
type ListOperatorPropertiesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewListOperatorPropertiesUnauthorized creates ListOperatorPropertiesUnauthorized with default headers values
func NewListOperatorPropertiesUnauthorized() *ListOperatorPropertiesUnauthorized {

	return &ListOperatorPropertiesUnauthorized{}
}

// WithPayload adds the payload to the list operator properties unauthorized response
func (o *ListOperatorPropertiesUnauthorized) WithPayload(payload *models.InfraError) *ListOperatorPropertiesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list operator properties unauthorized response
func (o *ListOperatorPropertiesUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListOperatorPropertiesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListOperatorPropertiesForbiddenCode is the HTTP code returned for type ListOperatorPropertiesForbidden
const ListOperatorPropertiesForbiddenCode int = 403

/*ListOperatorPropertiesForbidden Forbidden.

swagger:response listOperatorPropertiesForbidden
*/
type ListOperatorPropertiesForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewListOperatorPropertiesForbidden creates ListOperatorPropertiesForbidden with default headers values
func NewListOperatorPropertiesForbidden() *ListOperatorPropertiesForbidden {

	return &ListOperatorPropertiesForbidden{}
}

// WithPayload adds the payload to the list operator properties forbidden response
func (o *ListOperatorPropertiesForbidden) WithPayload(payload *models.InfraError) *ListOperatorPropertiesForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list operator properties forbidden response
func (o *ListOperatorPropertiesForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListOperatorPropertiesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListOperatorPropertiesNotFoundCode is the HTTP code returned for type ListOperatorPropertiesNotFound
const ListOperatorPropertiesNotFoundCode int = 404

/*ListOperatorPropertiesNotFound Error.

swagger:response listOperatorPropertiesNotFound
*/
type ListOperatorPropertiesNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListOperatorPropertiesNotFound creates ListOperatorPropertiesNotFound with default headers values
func NewListOperatorPropertiesNotFound() *ListOperatorPropertiesNotFound {

	return &ListOperatorPropertiesNotFound{}
}

// WithPayload adds the payload to the list operator properties not found response
func (o *ListOperatorPropertiesNotFound) WithPayload(payload *models.Error) *ListOperatorPropertiesNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list operator properties not found response
func (o *ListOperatorPropertiesNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListOperatorPropertiesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListOperatorPropertiesInternalServerErrorCode is the HTTP code returned for type ListOperatorPropertiesInternalServerError
const ListOperatorPropertiesInternalServerErrorCode int = 500

/*ListOperatorPropertiesInternalServerError Error.

swagger:response listOperatorPropertiesInternalServerError
*/
type ListOperatorPropertiesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListOperatorPropertiesInternalServerError creates ListOperatorPropertiesInternalServerError with default headers values
func NewListOperatorPropertiesInternalServerError() *ListOperatorPropertiesInternalServerError {

	return &ListOperatorPropertiesInternalServerError{}
}

// WithPayload adds the payload to the list operator properties internal server error response
func (o *ListOperatorPropertiesInternalServerError) WithPayload(payload *models.Error) *ListOperatorPropertiesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list operator properties internal server error response
func (o *ListOperatorPropertiesInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListOperatorPropertiesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
