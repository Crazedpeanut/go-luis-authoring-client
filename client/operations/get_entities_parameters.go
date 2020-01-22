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

// NewGetEntitiesParams creates a new GetEntitiesParams object
// with the default values initialized.
func NewGetEntitiesParams() *GetEntitiesParams {
	var (
		skipDefault = int64(0)
		takeDefault = int64(100)
	)
	return &GetEntitiesParams{
		Skip: &skipDefault,
		Take: &takeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEntitiesParamsWithTimeout creates a new GetEntitiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEntitiesParamsWithTimeout(timeout time.Duration) *GetEntitiesParams {
	var (
		skipDefault = int64(0)
		takeDefault = int64(100)
	)
	return &GetEntitiesParams{
		Skip: &skipDefault,
		Take: &takeDefault,

		timeout: timeout,
	}
}

// NewGetEntitiesParamsWithContext creates a new GetEntitiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEntitiesParamsWithContext(ctx context.Context) *GetEntitiesParams {
	var (
		skipDefault = int64(0)
		takeDefault = int64(100)
	)
	return &GetEntitiesParams{
		Skip: &skipDefault,
		Take: &takeDefault,

		Context: ctx,
	}
}

// NewGetEntitiesParamsWithHTTPClient creates a new GetEntitiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEntitiesParamsWithHTTPClient(client *http.Client) *GetEntitiesParams {
	var (
		skipDefault = int64(0)
		takeDefault = int64(100)
	)
	return &GetEntitiesParams{
		Skip:       &skipDefault,
		Take:       &takeDefault,
		HTTPClient: client,
	}
}

/*GetEntitiesParams contains all the parameters to send to the API endpoint
for the get entities operation typically these are written to a http.Request
*/
type GetEntitiesParams struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*Skip
	  The number of entries to skip. Default value is 0.

	*/
	Skip *int64
	/*Take
	  The number of entries to return. Maximum page size is 500. Default is 100.

	*/
	Take *int64
	/*VersionID
	  The task version ID.

	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get entities params
func (o *GetEntitiesParams) WithTimeout(timeout time.Duration) *GetEntitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get entities params
func (o *GetEntitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get entities params
func (o *GetEntitiesParams) WithContext(ctx context.Context) *GetEntitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get entities params
func (o *GetEntitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get entities params
func (o *GetEntitiesParams) WithHTTPClient(client *http.Client) *GetEntitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get entities params
func (o *GetEntitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the get entities params
func (o *GetEntitiesParams) WithAppID(appID string) *GetEntitiesParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the get entities params
func (o *GetEntitiesParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithSkip adds the skip to the get entities params
func (o *GetEntitiesParams) WithSkip(skip *int64) *GetEntitiesParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the get entities params
func (o *GetEntitiesParams) SetSkip(skip *int64) {
	o.Skip = skip
}

// WithTake adds the take to the get entities params
func (o *GetEntitiesParams) WithTake(take *int64) *GetEntitiesParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the get entities params
func (o *GetEntitiesParams) SetTake(take *int64) {
	o.Take = take
}

// WithVersionID adds the versionID to the get entities params
func (o *GetEntitiesParams) WithVersionID(versionID string) *GetEntitiesParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the get entities params
func (o *GetEntitiesParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetEntitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	if o.Skip != nil {

		// query param skip
		var qrSkip int64
		if o.Skip != nil {
			qrSkip = *o.Skip
		}
		qSkip := swag.FormatInt64(qrSkip)
		if qSkip != "" {
			if err := r.SetQueryParam("skip", qSkip); err != nil {
				return err
			}
		}

	}

	if o.Take != nil {

		// query param take
		var qrTake int64
		if o.Take != nil {
			qrTake = *o.Take
		}
		qTake := swag.FormatInt64(qrTake)
		if qTake != "" {
			if err := r.SetQueryParam("take", qTake); err != nil {
				return err
			}
		}

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