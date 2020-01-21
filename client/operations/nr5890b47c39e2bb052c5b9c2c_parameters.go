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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewNr5890b47c39e2bb052c5b9c2cParams creates a new Nr5890b47c39e2bb052c5b9c2cParams object
// with the default values initialized.
func NewNr5890b47c39e2bb052c5b9c2cParams() *Nr5890b47c39e2bb052c5b9c2cParams {
	var ()
	return &Nr5890b47c39e2bb052c5b9c2cParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNr5890b47c39e2bb052c5b9c2cParamsWithTimeout creates a new Nr5890b47c39e2bb052c5b9c2cParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNr5890b47c39e2bb052c5b9c2cParamsWithTimeout(timeout time.Duration) *Nr5890b47c39e2bb052c5b9c2cParams {
	var ()
	return &Nr5890b47c39e2bb052c5b9c2cParams{

		timeout: timeout,
	}
}

// NewNr5890b47c39e2bb052c5b9c2cParamsWithContext creates a new Nr5890b47c39e2bb052c5b9c2cParams object
// with the default values initialized, and the ability to set a context for a request
func NewNr5890b47c39e2bb052c5b9c2cParamsWithContext(ctx context.Context) *Nr5890b47c39e2bb052c5b9c2cParams {
	var ()
	return &Nr5890b47c39e2bb052c5b9c2cParams{

		Context: ctx,
	}
}

// NewNr5890b47c39e2bb052c5b9c2cParamsWithHTTPClient creates a new Nr5890b47c39e2bb052c5b9c2cParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNr5890b47c39e2bb052c5b9c2cParamsWithHTTPClient(client *http.Client) *Nr5890b47c39e2bb052c5b9c2cParams {
	var ()
	return &Nr5890b47c39e2bb052c5b9c2cParams{
		HTTPClient: client,
	}
}

/*Nr5890b47c39e2bb052c5b9c2cParams contains all the parameters to send to the API endpoint
for the 5890b47c39e2bb052c5b9c2c operation typically these are written to a http.Request
*/
type Nr5890b47c39e2bb052c5b9c2cParams struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*ClEntityID
	  Format - guid. The closed list entity extractor ID.

	*/
	ClEntityID string
	/*SubListID
	  The sublist ID.

	*/
	SubListID int64
	/*VersionID
	  The task version ID.

	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the 5890b47c39e2bb052c5b9c2c params
func (o *Nr5890b47c39e2bb052c5b9c2cParams) WithTimeout(timeout time.Duration) *Nr5890b47c39e2bb052c5b9c2cParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the 5890b47c39e2bb052c5b9c2c params
func (o *Nr5890b47c39e2bb052c5b9c2cParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the 5890b47c39e2bb052c5b9c2c params
func (o *Nr5890b47c39e2bb052c5b9c2cParams) WithContext(ctx context.Context) *Nr5890b47c39e2bb052c5b9c2cParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the 5890b47c39e2bb052c5b9c2c params
func (o *Nr5890b47c39e2bb052c5b9c2cParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the 5890b47c39e2bb052c5b9c2c params
func (o *Nr5890b47c39e2bb052c5b9c2cParams) WithHTTPClient(client *http.Client) *Nr5890b47c39e2bb052c5b9c2cParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the 5890b47c39e2bb052c5b9c2c params
func (o *Nr5890b47c39e2bb052c5b9c2cParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the 5890b47c39e2bb052c5b9c2c params
func (o *Nr5890b47c39e2bb052c5b9c2cParams) WithAppID(appID string) *Nr5890b47c39e2bb052c5b9c2cParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the 5890b47c39e2bb052c5b9c2c params
func (o *Nr5890b47c39e2bb052c5b9c2cParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithClEntityID adds the clEntityID to the 5890b47c39e2bb052c5b9c2c params
func (o *Nr5890b47c39e2bb052c5b9c2cParams) WithClEntityID(clEntityID string) *Nr5890b47c39e2bb052c5b9c2cParams {
	o.SetClEntityID(clEntityID)
	return o
}

// SetClEntityID adds the clEntityId to the 5890b47c39e2bb052c5b9c2c params
func (o *Nr5890b47c39e2bb052c5b9c2cParams) SetClEntityID(clEntityID string) {
	o.ClEntityID = clEntityID
}

// WithSubListID adds the subListID to the 5890b47c39e2bb052c5b9c2c params
func (o *Nr5890b47c39e2bb052c5b9c2cParams) WithSubListID(subListID int64) *Nr5890b47c39e2bb052c5b9c2cParams {
	o.SetSubListID(subListID)
	return o
}

// SetSubListID adds the subListId to the 5890b47c39e2bb052c5b9c2c params
func (o *Nr5890b47c39e2bb052c5b9c2cParams) SetSubListID(subListID int64) {
	o.SubListID = subListID
}

// WithVersionID adds the versionID to the 5890b47c39e2bb052c5b9c2c params
func (o *Nr5890b47c39e2bb052c5b9c2cParams) WithVersionID(versionID string) *Nr5890b47c39e2bb052c5b9c2cParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the 5890b47c39e2bb052c5b9c2c params
func (o *Nr5890b47c39e2bb052c5b9c2cParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *Nr5890b47c39e2bb052c5b9c2cParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param subListId
	if err := r.SetPathParam("subListId", swag.FormatInt64(o.SubListID)); err != nil {
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
