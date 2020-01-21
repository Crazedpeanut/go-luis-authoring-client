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

// NewNr5890b47c39e2bb052c5b9c1fParams creates a new Nr5890b47c39e2bb052c5b9c1fParams object
// with the default values initialized.
func NewNr5890b47c39e2bb052c5b9c1fParams() *Nr5890b47c39e2bb052c5b9c1fParams {
	var ()
	return &Nr5890b47c39e2bb052c5b9c1fParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNr5890b47c39e2bb052c5b9c1fParamsWithTimeout creates a new Nr5890b47c39e2bb052c5b9c1fParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNr5890b47c39e2bb052c5b9c1fParamsWithTimeout(timeout time.Duration) *Nr5890b47c39e2bb052c5b9c1fParams {
	var ()
	return &Nr5890b47c39e2bb052c5b9c1fParams{

		timeout: timeout,
	}
}

// NewNr5890b47c39e2bb052c5b9c1fParamsWithContext creates a new Nr5890b47c39e2bb052c5b9c1fParams object
// with the default values initialized, and the ability to set a context for a request
func NewNr5890b47c39e2bb052c5b9c1fParamsWithContext(ctx context.Context) *Nr5890b47c39e2bb052c5b9c1fParams {
	var ()
	return &Nr5890b47c39e2bb052c5b9c1fParams{

		Context: ctx,
	}
}

// NewNr5890b47c39e2bb052c5b9c1fParamsWithHTTPClient creates a new Nr5890b47c39e2bb052c5b9c1fParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNr5890b47c39e2bb052c5b9c1fParamsWithHTTPClient(client *http.Client) *Nr5890b47c39e2bb052c5b9c1fParams {
	var ()
	return &Nr5890b47c39e2bb052c5b9c1fParams{
		HTTPClient: client,
	}
}

/*Nr5890b47c39e2bb052c5b9c1fParams contains all the parameters to send to the API endpoint
for the 5890b47c39e2bb052c5b9c1f operation typically these are written to a http.Request
*/
type Nr5890b47c39e2bb052c5b9c1fParams struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*EntityID
	  Format - guid. The entity extractor ID.

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

// WithTimeout adds the timeout to the 5890b47c39e2bb052c5b9c1f params
func (o *Nr5890b47c39e2bb052c5b9c1fParams) WithTimeout(timeout time.Duration) *Nr5890b47c39e2bb052c5b9c1fParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the 5890b47c39e2bb052c5b9c1f params
func (o *Nr5890b47c39e2bb052c5b9c1fParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the 5890b47c39e2bb052c5b9c1f params
func (o *Nr5890b47c39e2bb052c5b9c1fParams) WithContext(ctx context.Context) *Nr5890b47c39e2bb052c5b9c1fParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the 5890b47c39e2bb052c5b9c1f params
func (o *Nr5890b47c39e2bb052c5b9c1fParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the 5890b47c39e2bb052c5b9c1f params
func (o *Nr5890b47c39e2bb052c5b9c1fParams) WithHTTPClient(client *http.Client) *Nr5890b47c39e2bb052c5b9c1fParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the 5890b47c39e2bb052c5b9c1f params
func (o *Nr5890b47c39e2bb052c5b9c1fParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the 5890b47c39e2bb052c5b9c1f params
func (o *Nr5890b47c39e2bb052c5b9c1fParams) WithAppID(appID string) *Nr5890b47c39e2bb052c5b9c1fParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the 5890b47c39e2bb052c5b9c1f params
func (o *Nr5890b47c39e2bb052c5b9c1fParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithEntityID adds the entityID to the 5890b47c39e2bb052c5b9c1f params
func (o *Nr5890b47c39e2bb052c5b9c1fParams) WithEntityID(entityID string) *Nr5890b47c39e2bb052c5b9c1fParams {
	o.SetEntityID(entityID)
	return o
}

// SetEntityID adds the entityId to the 5890b47c39e2bb052c5b9c1f params
func (o *Nr5890b47c39e2bb052c5b9c1fParams) SetEntityID(entityID string) {
	o.EntityID = entityID
}

// WithVersionID adds the versionID to the 5890b47c39e2bb052c5b9c1f params
func (o *Nr5890b47c39e2bb052c5b9c1fParams) WithVersionID(versionID string) *Nr5890b47c39e2bb052c5b9c1fParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the 5890b47c39e2bb052c5b9c1f params
func (o *Nr5890b47c39e2bb052c5b9c1fParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *Nr5890b47c39e2bb052c5b9c1fParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
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
