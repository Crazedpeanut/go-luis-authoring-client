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

	models "github.com/crazedpeanut/luis/models"
)

// NewCreatePhraseListFeatureParams creates a new CreatePhraseListFeatureParams object
// with the default values initialized.
func NewCreatePhraseListFeatureParams() *CreatePhraseListFeatureParams {
	var ()
	return &CreatePhraseListFeatureParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePhraseListFeatureParamsWithTimeout creates a new CreatePhraseListFeatureParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreatePhraseListFeatureParamsWithTimeout(timeout time.Duration) *CreatePhraseListFeatureParams {
	var ()
	return &CreatePhraseListFeatureParams{

		timeout: timeout,
	}
}

// NewCreatePhraseListFeatureParamsWithContext creates a new CreatePhraseListFeatureParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreatePhraseListFeatureParamsWithContext(ctx context.Context) *CreatePhraseListFeatureParams {
	var ()
	return &CreatePhraseListFeatureParams{

		Context: ctx,
	}
}

// NewCreatePhraseListFeatureParamsWithHTTPClient creates a new CreatePhraseListFeatureParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreatePhraseListFeatureParamsWithHTTPClient(client *http.Client) *CreatePhraseListFeatureParams {
	var ()
	return &CreatePhraseListFeatureParams{
		HTTPClient: client,
	}
}

/*CreatePhraseListFeatureParams contains all the parameters to send to the API endpoint
for the create phrase list feature operation typically these are written to a http.Request
*/
type CreatePhraseListFeatureParams struct {

	/*AppID
	  Format - guid. The application ID.

	*/
	AppID string
	/*PhraselistCreateObject
	  A JSON object containing Name, comma-separated Phrases and the isExchangeable boolean.

	Default value for isExchangeable is true.

	*/
	PhraselistCreateObject *models.PhraselistCreateObject
	/*VersionID
	  The version ID of the task.

	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create phrase list feature params
func (o *CreatePhraseListFeatureParams) WithTimeout(timeout time.Duration) *CreatePhraseListFeatureParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create phrase list feature params
func (o *CreatePhraseListFeatureParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create phrase list feature params
func (o *CreatePhraseListFeatureParams) WithContext(ctx context.Context) *CreatePhraseListFeatureParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create phrase list feature params
func (o *CreatePhraseListFeatureParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create phrase list feature params
func (o *CreatePhraseListFeatureParams) WithHTTPClient(client *http.Client) *CreatePhraseListFeatureParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create phrase list feature params
func (o *CreatePhraseListFeatureParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the create phrase list feature params
func (o *CreatePhraseListFeatureParams) WithAppID(appID string) *CreatePhraseListFeatureParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the create phrase list feature params
func (o *CreatePhraseListFeatureParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithPhraselistCreateObject adds the phraselistCreateObject to the create phrase list feature params
func (o *CreatePhraseListFeatureParams) WithPhraselistCreateObject(phraselistCreateObject *models.PhraselistCreateObject) *CreatePhraseListFeatureParams {
	o.SetPhraselistCreateObject(phraselistCreateObject)
	return o
}

// SetPhraselistCreateObject adds the phraselistCreateObject to the create phrase list feature params
func (o *CreatePhraseListFeatureParams) SetPhraselistCreateObject(phraselistCreateObject *models.PhraselistCreateObject) {
	o.PhraselistCreateObject = phraselistCreateObject
}

// WithVersionID adds the versionID to the create phrase list feature params
func (o *CreatePhraseListFeatureParams) WithVersionID(versionID string) *CreatePhraseListFeatureParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the create phrase list feature params
func (o *CreatePhraseListFeatureParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePhraseListFeatureParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appId
	if err := r.SetPathParam("appId", o.AppID); err != nil {
		return err
	}

	if o.PhraselistCreateObject != nil {
		if err := r.SetBodyParam(o.PhraselistCreateObject); err != nil {
			return err
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