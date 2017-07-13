// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "appengine": Account Resource Client
//
// Command:
// $ goagen
// --design=github.com/tikasan/goa-datastore/design
// --out=$(GOPATH)/src/github.com/tikasan/goa-datastore
// --version=v1.2.0-dirty

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// CreateAccountPayload is the Account create action payload.
type CreateAccountPayload struct {
	// name
	Name string `form:"name" json:"name" xml:"name"`
}

// CreateAccountPath computes a request path to the create action of Account.
func CreateAccountPath() string {

	return fmt.Sprintf("/account")
}

// create
func (c *Client) CreateAccount(ctx context.Context, path string, payload *CreateAccountPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateAccountRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateAccountRequest create the request corresponding to the create action endpoint of the Account resource.
func (c *Client) NewCreateAccountRequest(ctx context.Context, path string, payload *CreateAccountPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// DeleteAccountPayload is the Account delete action payload.
type DeleteAccountPayload struct {
	// name
	Name string `form:"name" json:"name" xml:"name"`
}

// DeleteAccountPath computes a request path to the delete action of Account.
func DeleteAccountPath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/account/%s", param0)
}

// delete
func (c *Client) DeleteAccount(ctx context.Context, path string, payload *DeleteAccountPayload, contentType string) (*http.Response, error) {
	req, err := c.NewDeleteAccountRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteAccountRequest create the request corresponding to the delete action endpoint of the Account resource.
func (c *Client) NewDeleteAccountRequest(ctx context.Context, path string, payload *DeleteAccountPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// ListAccountPath computes a request path to the list action of Account.
func ListAccountPath() string {

	return fmt.Sprintf("/account")
}

// list
func (c *Client) ListAccount(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListAccountRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListAccountRequest create the request corresponding to the list action endpoint of the Account resource.
func (c *Client) NewListAccountRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowAccountPath computes a request path to the show action of Account.
func ShowAccountPath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/account/%s", param0)
}

// show
func (c *Client) ShowAccount(ctx context.Context, path string, name string) (*http.Response, error) {
	req, err := c.NewShowAccountRequest(ctx, path, name)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowAccountRequest create the request corresponding to the show action endpoint of the Account resource.
func (c *Client) NewShowAccountRequest(ctx context.Context, path string, name string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("name", name)
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// UpdateAccountPayload is the Account update action payload.
type UpdateAccountPayload struct {
	// name
	Name string `form:"name" json:"name" xml:"name"`
}

// UpdateAccountPath computes a request path to the update action of Account.
func UpdateAccountPath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/account/%s", param0)
}

// update
func (c *Client) UpdateAccount(ctx context.Context, path string, payload *UpdateAccountPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateAccountRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateAccountRequest create the request corresponding to the update action endpoint of the Account resource.
func (c *Client) NewUpdateAccountRequest(ctx context.Context, path string, payload *UpdateAccountPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}