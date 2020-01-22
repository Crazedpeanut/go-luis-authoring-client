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

// NewDeleteClosedListParams creates a new DeleteClosedListParams object
// with the default values initialized.
func NewDeleteClosedListParams() *DeleteClosedListParams {
	var ()
	return &DeleteClosedListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteClosedListParamsWithTimeout creates a new DeleteClosedListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteClosedListParamsWithTimeout(timeout time.Duration) *DeleteClosedListParams {
	var ()
	return &DeleteClosedListParams{

		timeout: timeout,
	}
}

// NewDeleteClosedListParamsWithContext creates a new DeleteClosedListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteClosedListParamsWithContext(ctx context.Context) *DeleteClosedListParams {
	var ()
	return &DeleteClosedListParams{

		Context: ctx,
	}
}

// NewDeleteClosedListParamsWithHTTPClient creates a new DeleteClosedListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteClosedListParamsWithHTTPClient(client *http.Client) *DeleteClosedListParams {
	var ()
	return &DeleteClosedListParams{
		HTTPClient: client,
	}
}

/*DeleteClosedListParams contains all the parameters to send to the API endpoint
for the delete closed list operation typically these are written to a http.Request
*/
type DeleteClosedListParams struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*ClEntityID
	  Format - guid. The closed list entity extractor ID.

	*/
	ClEntityID string
	/*VersionID
	  The task version ID.

	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete closed list params
func (o *DeleteClosedListParams) WithTimeout(timeout time.Duration) *DeleteClosedListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete closed list params
func (o *DeleteClosedListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete closed list params
func (o *DeleteClosedListParams) WithContext(ctx context.Context) *DeleteClosedListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete closed list params
func (o *DeleteClosedListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete closed list params
func (o *DeleteClosedListParams) WithHTTPClient(client *http.Client) *DeleteClosedListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete closed list params
func (o *DeleteClosedListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the delete closed list params
func (o *DeleteClosedListParams) WithAppID(appID string) *DeleteClosedListParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the delete closed list params
func (o *DeleteClosedListParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithClEntityID adds the clEntityID to the delete closed list params
func (o *DeleteClosedListParams) WithClEntityID(clEntityID string) *DeleteClosedListParams {
	o.SetClEntityID(clEntityID)
	return o
}

// SetClEntityID adds the clEntityId to the delete closed list params
func (o *DeleteClosedListParams) SetClEntityID(clEntityID string) {
	o.ClEntityID = clEntityID
}

// WithVersionID adds the versionID to the delete closed list params
func (o *DeleteClosedListParams) WithVersionID(versionID string) *DeleteClosedListParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the delete closed list params
func (o *DeleteClosedListParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteClosedListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	// path param clEntityId
	if err := r.SetPathParam("clEntityId", o.ClEntityID); err != nil {
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