// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// ListClusterHostsOKCode is the HTTP code returned for type ListClusterHostsOK
const ListClusterHostsOKCode int = 200

/*ListClusterHostsOK Success.

swagger:response listClusterHostsOK
*/
type ListClusterHostsOK struct {

	/*
	  In: Body
	*/
	Payload models.HostList `json:"body,omitempty"`
}

// NewListClusterHostsOK creates ListClusterHostsOK with default headers values
func NewListClusterHostsOK() *ListClusterHostsOK {

	return &ListClusterHostsOK{}
}

// WithPayload adds the payload to the list cluster hosts o k response
func (o *ListClusterHostsOK) WithPayload(payload models.HostList) *ListClusterHostsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list cluster hosts o k response
func (o *ListClusterHostsOK) SetPayload(payload models.HostList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListClusterHostsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.HostList{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListClusterHostsUnauthorizedCode is the HTTP code returned for type ListClusterHostsUnauthorized
const ListClusterHostsUnauthorizedCode int = 401

/*ListClusterHostsUnauthorized Unauthorized.

swagger:response listClusterHostsUnauthorized
*/
type ListClusterHostsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewListClusterHostsUnauthorized creates ListClusterHostsUnauthorized with default headers values
func NewListClusterHostsUnauthorized() *ListClusterHostsUnauthorized {

	return &ListClusterHostsUnauthorized{}
}

// WithPayload adds the payload to the list cluster hosts unauthorized response
func (o *ListClusterHostsUnauthorized) WithPayload(payload *models.InfraError) *ListClusterHostsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list cluster hosts unauthorized response
func (o *ListClusterHostsUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListClusterHostsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListClusterHostsForbiddenCode is the HTTP code returned for type ListClusterHostsForbidden
const ListClusterHostsForbiddenCode int = 403

/*ListClusterHostsForbidden Forbidden.

swagger:response listClusterHostsForbidden
*/
type ListClusterHostsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewListClusterHostsForbidden creates ListClusterHostsForbidden with default headers values
func NewListClusterHostsForbidden() *ListClusterHostsForbidden {

	return &ListClusterHostsForbidden{}
}

// WithPayload adds the payload to the list cluster hosts forbidden response
func (o *ListClusterHostsForbidden) WithPayload(payload *models.InfraError) *ListClusterHostsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list cluster hosts forbidden response
func (o *ListClusterHostsForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListClusterHostsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListClusterHostsNotFoundCode is the HTTP code returned for type ListClusterHostsNotFound
const ListClusterHostsNotFoundCode int = 404

/*ListClusterHostsNotFound Error.

swagger:response listClusterHostsNotFound
*/
type ListClusterHostsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListClusterHostsNotFound creates ListClusterHostsNotFound with default headers values
func NewListClusterHostsNotFound() *ListClusterHostsNotFound {

	return &ListClusterHostsNotFound{}
}

// WithPayload adds the payload to the list cluster hosts not found response
func (o *ListClusterHostsNotFound) WithPayload(payload *models.Error) *ListClusterHostsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list cluster hosts not found response
func (o *ListClusterHostsNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListClusterHostsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListClusterHostsInternalServerErrorCode is the HTTP code returned for type ListClusterHostsInternalServerError
const ListClusterHostsInternalServerErrorCode int = 500

/*ListClusterHostsInternalServerError Error.

swagger:response listClusterHostsInternalServerError
*/
type ListClusterHostsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListClusterHostsInternalServerError creates ListClusterHostsInternalServerError with default headers values
func NewListClusterHostsInternalServerError() *ListClusterHostsInternalServerError {

	return &ListClusterHostsInternalServerError{}
}

// WithPayload adds the payload to the list cluster hosts internal server error response
func (o *ListClusterHostsInternalServerError) WithPayload(payload *models.Error) *ListClusterHostsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list cluster hosts internal server error response
func (o *ListClusterHostsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListClusterHostsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
