// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// NewV2UpdateHostParams creates a new V2UpdateHostParams object
// with the default values initialized.
func NewV2UpdateHostParams() *V2UpdateHostParams {
	var ()
	return &V2UpdateHostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV2UpdateHostParamsWithTimeout creates a new V2UpdateHostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV2UpdateHostParamsWithTimeout(timeout time.Duration) *V2UpdateHostParams {
	var ()
	return &V2UpdateHostParams{

		timeout: timeout,
	}
}

// NewV2UpdateHostParamsWithContext creates a new V2UpdateHostParams object
// with the default values initialized, and the ability to set a context for a request
func NewV2UpdateHostParamsWithContext(ctx context.Context) *V2UpdateHostParams {
	var ()
	return &V2UpdateHostParams{

		Context: ctx,
	}
}

// NewV2UpdateHostParamsWithHTTPClient creates a new V2UpdateHostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV2UpdateHostParamsWithHTTPClient(client *http.Client) *V2UpdateHostParams {
	var ()
	return &V2UpdateHostParams{
		HTTPClient: client,
	}
}

/*V2UpdateHostParams contains all the parameters to send to the API endpoint
for the v2 update host operation typically these are written to a http.Request
*/
type V2UpdateHostParams struct {

	/*HostUpdateParams
	  The properties to update.

	*/
	HostUpdateParams *models.HostUpdateParams
	/*HostID
	  The host that should be updated.

	*/
	HostID strfmt.UUID
	/*InfraEnvID
	  The infra_env_id of the host to be updated.

	*/
	InfraEnvID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v2 update host params
func (o *V2UpdateHostParams) WithTimeout(timeout time.Duration) *V2UpdateHostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 update host params
func (o *V2UpdateHostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 update host params
func (o *V2UpdateHostParams) WithContext(ctx context.Context) *V2UpdateHostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 update host params
func (o *V2UpdateHostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 update host params
func (o *V2UpdateHostParams) WithHTTPClient(client *http.Client) *V2UpdateHostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 update host params
func (o *V2UpdateHostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHostUpdateParams adds the hostUpdateParams to the v2 update host params
func (o *V2UpdateHostParams) WithHostUpdateParams(hostUpdateParams *models.HostUpdateParams) *V2UpdateHostParams {
	o.SetHostUpdateParams(hostUpdateParams)
	return o
}

// SetHostUpdateParams adds the hostUpdateParams to the v2 update host params
func (o *V2UpdateHostParams) SetHostUpdateParams(hostUpdateParams *models.HostUpdateParams) {
	o.HostUpdateParams = hostUpdateParams
}

// WithHostID adds the hostID to the v2 update host params
func (o *V2UpdateHostParams) WithHostID(hostID strfmt.UUID) *V2UpdateHostParams {
	o.SetHostID(hostID)
	return o
}

// SetHostID adds the hostId to the v2 update host params
func (o *V2UpdateHostParams) SetHostID(hostID strfmt.UUID) {
	o.HostID = hostID
}

// WithInfraEnvID adds the infraEnvID to the v2 update host params
func (o *V2UpdateHostParams) WithInfraEnvID(infraEnvID strfmt.UUID) *V2UpdateHostParams {
	o.SetInfraEnvID(infraEnvID)
	return o
}

// SetInfraEnvID adds the infraEnvId to the v2 update host params
func (o *V2UpdateHostParams) SetInfraEnvID(infraEnvID strfmt.UUID) {
	o.InfraEnvID = infraEnvID
}

// WriteToRequest writes these params to a swagger request
func (o *V2UpdateHostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.HostUpdateParams != nil {
		if err := r.SetBodyParam(o.HostUpdateParams); err != nil {
			return err
		}
	}

	// path param host_id
	if err := r.SetPathParam("host_id", o.HostID.String()); err != nil {
		return err
	}

	// path param infra_env_id
	if err := r.SetPathParam("infra_env_id", o.InfraEnvID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
