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

// NewGetPatternFeatureParams creates a new GetPatternFeatureParams object
// with the default values initialized.
func NewGetPatternFeatureParams() *GetPatternFeatureParams {
	var ()
	return &GetPatternFeatureParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPatternFeatureParamsWithTimeout creates a new GetPatternFeatureParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPatternFeatureParamsWithTimeout(timeout time.Duration) *GetPatternFeatureParams {
	var ()
	return &GetPatternFeatureParams{

		timeout: timeout,
	}
}

// NewGetPatternFeatureParamsWithContext creates a new GetPatternFeatureParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPatternFeatureParamsWithContext(ctx context.Context) *GetPatternFeatureParams {
	var ()
	return &GetPatternFeatureParams{

		Context: ctx,
	}
}

// NewGetPatternFeatureParamsWithHTTPClient creates a new GetPatternFeatureParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPatternFeatureParamsWithHTTPClient(client *http.Client) *GetPatternFeatureParams {
	var ()
	return &GetPatternFeatureParams{
		HTTPClient: client,
	}
}

/*GetPatternFeatureParams contains all the parameters to send to the API endpoint
for the get pattern feature operation typically these are written to a http.Request
*/
type GetPatternFeatureParams struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*PatternID
	  The pattern feature ID.

	*/
	PatternID int64
	/*VersionID
	  The version ID of the task.

	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get pattern feature params
func (o *GetPatternFeatureParams) WithTimeout(timeout time.Duration) *GetPatternFeatureParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pattern feature params
func (o *GetPatternFeatureParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pattern feature params
func (o *GetPatternFeatureParams) WithContext(ctx context.Context) *GetPatternFeatureParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pattern feature params
func (o *GetPatternFeatureParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pattern feature params
func (o *GetPatternFeatureParams) WithHTTPClient(client *http.Client) *GetPatternFeatureParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pattern feature params
func (o *GetPatternFeatureParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the get pattern feature params
func (o *GetPatternFeatureParams) WithAppID(appID string) *GetPatternFeatureParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the get pattern feature params
func (o *GetPatternFeatureParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithPatternID adds the patternID to the get pattern feature params
func (o *GetPatternFeatureParams) WithPatternID(patternID int64) *GetPatternFeatureParams {
	o.SetPatternID(patternID)
	return o
}

// SetPatternID adds the patternId to the get pattern feature params
func (o *GetPatternFeatureParams) SetPatternID(patternID int64) {
	o.PatternID = patternID
}

// WithVersionID adds the versionID to the get pattern feature params
func (o *GetPatternFeatureParams) WithVersionID(versionID string) *GetPatternFeatureParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the get pattern feature params
func (o *GetPatternFeatureParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPatternFeatureParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	// path param patternId
	if err := r.SetPathParam("patternId", swag.FormatInt64(o.PatternID)); err != nil {
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