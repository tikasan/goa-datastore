// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "appengine": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/tikasan/goa-datastore/design
// --out=$(GOPATH)/src/github.com/tikasan/goa-datastore
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
	"strconv"
)

// CreateAccountContext provides the Account create action context.
type CreateAccountContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateAccountPayload
}

// NewCreateAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the Account controller create action.
func NewCreateAccountContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreateAccountContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreateAccountContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createAccountPayload is the Account create action payload.
type createAccountPayload struct {
	// name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createAccountPayload) Validate() (err error) {
	if payload.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	return
}

// Publicize creates CreateAccountPayload from createAccountPayload
func (payload *createAccountPayload) Publicize() *CreateAccountPayload {
	var pub CreateAccountPayload
	if payload.Name != nil {
		pub.Name = *payload.Name
	}
	return &pub
}

// CreateAccountPayload is the Account create action payload.
type CreateAccountPayload struct {
	// name
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateAccountPayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateAccountContext) OK(r *Account) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.account+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreateAccountContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CreateAccountContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// DeleteAccountContext provides the Account delete action context.
type DeleteAccountContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID      int
	Payload *DeleteAccountPayload
}

// NewDeleteAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the Account controller delete action.
func NewDeleteAccountContext(ctx context.Context, r *http.Request, service *goa.Service) (*DeleteAccountContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := DeleteAccountContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// deleteAccountPayload is the Account delete action payload.
type deleteAccountPayload struct {
	// name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *deleteAccountPayload) Validate() (err error) {
	if payload.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	return
}

// Publicize creates DeleteAccountPayload from deleteAccountPayload
func (payload *deleteAccountPayload) Publicize() *DeleteAccountPayload {
	var pub DeleteAccountPayload
	if payload.Name != nil {
		pub.Name = *payload.Name
	}
	return &pub
}

// DeleteAccountPayload is the Account delete action payload.
type DeleteAccountPayload struct {
	// name
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate runs the validation rules defined in the design.
func (payload *DeleteAccountPayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *DeleteAccountContext) OK(r *Account) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.account+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *DeleteAccountContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *DeleteAccountContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// ListAccountContext provides the Account list action context.
type ListAccountContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the Account controller list action.
func NewListAccountContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListAccountContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListAccountContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListAccountContext) OK(r AccountCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.account+json; type=collection")
	if r == nil {
		r = AccountCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ListAccountContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListAccountContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// ShowAccountContext provides the Account show action context.
type ShowAccountContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID   int
	Name string
}

// NewShowAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the Account controller show action.
func NewShowAccountContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowAccountContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowAccountContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	paramName := req.Params["name"]
	if len(paramName) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("name"))
	} else {
		rawName := paramName[0]
		rctx.Name = rawName
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowAccountContext) OK(r *Account) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.account+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ShowAccountContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowAccountContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// UpdateAccountContext provides the Account update action context.
type UpdateAccountContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID      int
	Payload *UpdateAccountPayload
}

// NewUpdateAccountContext parses the incoming request URL and body, performs validations and creates the
// context used by the Account controller update action.
func NewUpdateAccountContext(ctx context.Context, r *http.Request, service *goa.Service) (*UpdateAccountContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := UpdateAccountContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// updateAccountPayload is the Account update action payload.
type updateAccountPayload struct {
	// name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *updateAccountPayload) Validate() (err error) {
	if payload.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	return
}

// Publicize creates UpdateAccountPayload from updateAccountPayload
func (payload *updateAccountPayload) Publicize() *UpdateAccountPayload {
	var pub UpdateAccountPayload
	if payload.Name != nil {
		pub.Name = *payload.Name
	}
	return &pub
}

// UpdateAccountPayload is the Account update action payload.
type UpdateAccountPayload struct {
	// name
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate runs the validation rules defined in the design.
func (payload *UpdateAccountPayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *UpdateAccountContext) OK(r *Account) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.account+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *UpdateAccountContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *UpdateAccountContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}