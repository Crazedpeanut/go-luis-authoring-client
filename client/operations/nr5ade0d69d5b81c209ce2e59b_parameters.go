// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewNr5ade0d69d5b81c209ce2e59bParams creates a new Nr5ade0d69d5b81c209ce2e59bParams object
// with the default values initialized.
func NewNr5ade0d69d5b81c209ce2e59bParams() *Nr5ade0d69d5b81c209ce2e59bParams {
	var ()
	return &Nr5ade0d69d5b81c209ce2e59bParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNr5ade0d69d5b81c209ce2e59bParamsWithTimeout creates a new Nr5ade0d69d5b81c209ce2e59bParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNr5ade0d69d5b81c209ce2e59bParamsWithTimeout(timeout time.Duration) *Nr5ade0d69d5b81c209ce2e59bParams {
	var ()
	return &Nr5ade0d69d5b81c209ce2e59bParams{

		timeout: timeout,
	}
}

// NewNr5ade0d69d5b81c209ce2e59bParamsWithContext creates a new Nr5ade0d69d5b81c209ce2e59bParams object
// with the default values initialized, and the ability to set a context for a request
func NewNr5ade0d69d5b81c209ce2e59bParamsWithContext(ctx context.Context) *Nr5ade0d69d5b81c209ce2e59bParams {
	var ()
	return &Nr5ade0d69d5b81c209ce2e59bParams{

		Context: ctx,
	}
}

// NewNr5ade0d69d5b81c209ce2e59bParamsWithHTTPClient creates a new Nr5ade0d69d5b81c209ce2e59bParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNr5ade0d69d5b81c209ce2e59bParamsWithHTTPClient(client *http.Client) *Nr5ade0d69d5b81c209ce2e59bParams {
	var ()
	return &Nr5ade0d69d5b81c209ce2e59bParams{
		HTTPClient: client,
	}
}

/*Nr5ade0d69d5b81c209ce2e59bParams contains all the parameters to send to the API endpoint
for the 5ade0d69d5b81c209ce2e59b operation typically these are written to a http.Request
*/
type Nr5ade0d69d5b81c209ce2e59bParams struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*Body
	  A model object containing the name for the new role.

	The name of the role must be unique across the application version.

	*/
	Body interface{}
	/*EntityID
	  Format - guid. The closed list entity ID.

	*/
	EntityID string
	/*VersionID
	  The task version ID.

	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the 5ade0d69d5b81c209ce2e59b params
func (o *Nr5ade0d69d5b81c209ce2e59bParams) WithTimeout(timeout time.Duration) *Nr5ade0d69d5b81c209ce2e59bParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the 5ade0d69d5b81c209ce2e59b params
func (o *Nr5ade0d69d5b81c209ce2e59bParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the 5ade0d69d5b81c209ce2e59b params
func (o *Nr5ade0d69d5b81c209ce2e59bParams) WithContext(ctx context.Context) *Nr5ade0d69d5b81c209ce2e59bParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the 5ade0d69d5b81c209ce2e59b params
func (o *Nr5ade0d69d5b81c209ce2e59bParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the 5ade0d69d5b81c209ce2e59b params
func (o *Nr5ade0d69d5b81c209ce2e59bParams) WithHTTPClient(client *http.Client) *Nr5ade0d69d5b81c209ce2e59bParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the 5ade0d69d5b81c209ce2e59b params
func (o *Nr5ade0d69d5b81c209ce2e59bParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the 5ade0d69d5b81c209ce2e59b params
func (o *Nr5ade0d69d5b81c209ce2e59bParams) WithAppID(appID string) *Nr5ade0d69d5b81c209ce2e59bParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the 5ade0d69d5b81c209ce2e59b params
func (o *Nr5ade0d69d5b81c209ce2e59bParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithBody adds the body to the 5ade0d69d5b81c209ce2e59b params
func (o *Nr5ade0d69d5b81c209ce2e59bParams) WithBody(body interface{}) *Nr5ade0d69d5b81c209ce2e59bParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the 5ade0d69d5b81c209ce2e59b params
func (o *Nr5ade0d69d5b81c209ce2e59bParams) SetBody(body interface{}) {
	o.Body = body
}

// WithEntityID adds the entityID to the 5ade0d69d5b81c209ce2e59b params
func (o *Nr5ade0d69d5b81c209ce2e59bParams) WithEntityID(entityID string) *Nr5ade0d69d5b81c209ce2e59bParams {
	o.SetEntityID(entityID)
	return o
}

// SetEntityID adds the entityId to the 5ade0d69d5b81c209ce2e59b params
func (o *Nr5ade0d69d5b81c209ce2e59bParams) SetEntityID(entityID string) {
	o.EntityID = entityID
}

// WithVersionID adds the versionID to the 5ade0d69d5b81c209ce2e59b params
func (o *Nr5ade0d69d5b81c209ce2e59bParams) WithVersionID(versionID string) *Nr5ade0d69d5b81c209ce2e59bParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the 5ade0d69d5b81c209ce2e59b params
func (o *Nr5ade0d69d5b81c209ce2e59bParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *Nr5ade0d69d5b81c209ce2e59bParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param entityId
	if err := r.SetPathParam("entityId", o.EntityID); err != nil {
		return err
	}

	// path param versionId
	if err := r.SetPathParam("versionId", o.VersionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
