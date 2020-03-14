// Code generated by go-swagger; DO NOT EDIT.

package products

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

// NewStoreProductParams creates a new StoreProductParams object
// with the default values initialized.
func NewStoreProductParams() *StoreProductParams {

	return &StoreProductParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStoreProductParamsWithTimeout creates a new StoreProductParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStoreProductParamsWithTimeout(timeout time.Duration) *StoreProductParams {

	return &StoreProductParams{

		timeout: timeout,
	}
}

// NewStoreProductParamsWithContext creates a new StoreProductParams object
// with the default values initialized, and the ability to set a context for a request
func NewStoreProductParamsWithContext(ctx context.Context) *StoreProductParams {

	return &StoreProductParams{

		Context: ctx,
	}
}

// NewStoreProductParamsWithHTTPClient creates a new StoreProductParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStoreProductParamsWithHTTPClient(client *http.Client) *StoreProductParams {

	return &StoreProductParams{
		HTTPClient: client,
	}
}

/*StoreProductParams contains all the parameters to send to the API endpoint
for the store product operation typically these are written to a http.Request
*/
type StoreProductParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the store product params
func (o *StoreProductParams) WithTimeout(timeout time.Duration) *StoreProductParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the store product params
func (o *StoreProductParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the store product params
func (o *StoreProductParams) WithContext(ctx context.Context) *StoreProductParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the store product params
func (o *StoreProductParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the store product params
func (o *StoreProductParams) WithHTTPClient(client *http.Client) *StoreProductParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the store product params
func (o *StoreProductParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StoreProductParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}