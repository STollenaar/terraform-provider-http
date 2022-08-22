package service

import "github.com/hashicorp/terraform-plugin-framework/types"

type httpmodel struct {
	ID              types.String `tfsdk:"id"`
	URL             types.String `tfsdk:"url"`
	Method          types.String `tfsdk:"method"`
	RequestHeaders  types.Map    `tfsdk:"request_headers"`
	RequestBody     types.String `tfsdk:"request_body"`
	ResponseHeaders types.Map    `tfsdk:"response_headers"`
	ResponseBody    types.String `tfsdk:"response_body"`
	Body            types.String `tfsdk:"body"`
	StatusCode      types.Int64  `tfsdk:"status_code"`
}
