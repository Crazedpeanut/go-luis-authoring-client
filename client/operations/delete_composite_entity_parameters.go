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

// NewDeleteCompositeEntityParams creates a new DeleteCompositeEntityParams object
// with the default values initialized.
func NewDeleteCompositeEntityParams() *DeleteCompositeEntityParams {
	var ()
	return &DeleteCompositeEntityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCompositeEntityParamsWithTimeout creates a new DeleteCompositeEntityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCompositeEntityParamsWithTimeout(timeout time.Duration) *DeleteCompositeEntityParams {
	var ()
	return &DeleteCompositeEntityParams{

		timeout: timeout,
	}
}

// NewDeleteCompositeEntityParamsWithContext creates a new DeleteCompositeEntityParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCompositeEntityParamsWithContext(ctx context.Context) *DeleteCompositeEntityParams {
	var ()
	return &DeleteCompositeEntityParams{

		Context: ctx,
	}
}

// NewDeleteCompositeEntityParamsWithHTTPClient creates a new DeleteCompositeEntityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCompositeEntityParamsWithHTTPClient(client *http.Client) *DeleteCompositeEntityParams {
	var ()
	return &DeleteCompositeEntityParams{
		HTTPClient: client,
	}
}

/*DeleteCompositeEntityParams contains all the parameters to send to the API endpoint
for the delete composite entity operation typically these are written to a http.Request
*/
type DeleteCompositeEntityParams struct {

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

// WithTimeout adds the timeout to the delete composite entity params
func (o *DeleteCompositeEntityParams) WithTimeout(timeout time.Duration) *DeleteCompositeEntityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete composite entity params
func (o *DeleteCompositeEntityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete composite entity params
func (o *DeleteCompositeEntityParams) WithContext(ctx context.Context) *DeleteCompositeEntityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete composite entity params
func (o *DeleteCompositeEntityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete composite entity params
func (o *DeleteCompositeEntityParams) WithHTTPClient(client *http.Client) *DeleteCompositeEntityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete composite entity params
func (o *DeleteCompositeEntityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the delete composite entity params
func (o *DeleteCompositeEntityParams) WithAppID(appID string) *DeleteCompositeEntityParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the delete composite entity params
func (o *DeleteCompositeEntityParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithCEntityID adds the cEntityID to the delete composite entity params
func (o *DeleteCompositeEntityParams) WithCEntityID(cEntityID string) *DeleteCompositeEntityParams {
	o.SetCEntityID(cEntityID)
	return o
}

// SetCEntityID adds the cEntityId to the delete composite entity params
func (o *DeleteCompositeEntityParams) SetCEntityID(cEntityID string) {
	o.CEntityID = cEntityID
}

// WithVersionID adds the versionID to the delete composite entity params
func (o *DeleteCompositeEntityParams) WithVersionID(versionID string) *DeleteCompositeEntityParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the delete composite entity params
func (o *DeleteCompositeEntityParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCompositeEntityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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