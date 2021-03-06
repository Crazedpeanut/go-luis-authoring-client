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

// NewGetCompositeEntityParams creates a new GetCompositeEntityParams object
// with the default values initialized.
func NewGetCompositeEntityParams() *GetCompositeEntityParams {
	var ()
	return &GetCompositeEntityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCompositeEntityParamsWithTimeout creates a new GetCompositeEntityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCompositeEntityParamsWithTimeout(timeout time.Duration) *GetCompositeEntityParams {
	var ()
	return &GetCompositeEntityParams{

		timeout: timeout,
	}
}

// NewGetCompositeEntityParamsWithContext creates a new GetCompositeEntityParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCompositeEntityParamsWithContext(ctx context.Context) *GetCompositeEntityParams {
	var ()
	return &GetCompositeEntityParams{

		Context: ctx,
	}
}

// NewGetCompositeEntityParamsWithHTTPClient creates a new GetCompositeEntityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCompositeEntityParamsWithHTTPClient(client *http.Client) *GetCompositeEntityParams {
	var ()
	return &GetCompositeEntityParams{
		HTTPClient: client,
	}
}

/*GetCompositeEntityParams contains all the parameters to send to the API endpoint
for the get composite entity operation typically these are written to a http.Request
*/
type GetCompositeEntityParams struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*CEntityID
	  Format - guid. The composite entity extractor ID.

	*/
	CEntityID string
	/*VersionID
	  The task version ID.

	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get composite entity params
func (o *GetCompositeEntityParams) WithTimeout(timeout time.Duration) *GetCompositeEntityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get composite entity params
func (o *GetCompositeEntityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get composite entity params
func (o *GetCompositeEntityParams) WithContext(ctx context.Context) *GetCompositeEntityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get composite entity params
func (o *GetCompositeEntityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get composite entity params
func (o *GetCompositeEntityParams) WithHTTPClient(client *http.Client) *GetCompositeEntityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get composite entity params
func (o *GetCompositeEntityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the get composite entity params
func (o *GetCompositeEntityParams) WithAppID(appID string) *GetCompositeEntityParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the get composite entity params
func (o *GetCompositeEntityParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithCEntityID adds the cEntityID to the get composite entity params
func (o *GetCompositeEntityParams) WithCEntityID(cEntityID string) *GetCompositeEntityParams {
	o.SetCEntityID(cEntityID)
	return o
}

// SetCEntityID adds the cEntityId to the get composite entity params
func (o *GetCompositeEntityParams) SetCEntityID(cEntityID string) {
	o.CEntityID = cEntityID
}

// WithVersionID adds the versionID to the get composite entity params
func (o *GetCompositeEntityParams) WithVersionID(versionID string) *GetCompositeEntityParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the get composite entity params
func (o *GetCompositeEntityParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCompositeEntityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	// path param cEntityId
	if err := r.SetPathParam("cEntityId", o.CEntityID); err != nil {
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
