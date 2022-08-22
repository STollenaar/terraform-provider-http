package service

import (
	"context"
	"fmt"
	"mime"
	"os"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

var stderr = os.Stderr

func New() provider.Provider {
	return &terraformProvider{}
}

type terraformProvider struct {
	configured bool
}

// Provider schema struct
type providerData struct{}

func (p *terraformProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	fmt.Fprintf(stderr, "[DEBUG]- Already encountered an error")
	var config providerData
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		fmt.Fprint(stderr, "[DEBUG]- Already encountered an error")
		return
	}

	p.configured = true
}

// GetSchema -
func (p *terraformProvider) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{},
	}, nil
}

// GetDataSources - Defines provider data sources
func (p *terraformProvider) GetDataSources(_ context.Context) (map[string]provider.DataSourceType, diag.Diagnostics) {
	return map[string]provider.DataSourceType{
		"http_request": &DataSourceHttpRequestType{},
	}, nil
}

// GetResources - Defines provider resources
func (p *terraformProvider) GetResources(_ context.Context) (map[string]provider.ResourceType, diag.Diagnostics) {
	return map[string]provider.ResourceType{
		"http_request": &ResourceHttpRequestType{},
	}, nil
}

func isContentTypeText(contentType string) bool {

	parsedType, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		return false
	}

	allowedContentTypes := []*regexp.Regexp{
		regexp.MustCompile("^text/.+"),
		regexp.MustCompile("^application/json$"),
		regexp.MustCompile(`^application/samlmetadata\+xml`),
	}

	for _, r := range allowedContentTypes {
		if r.MatchString(parsedType) {
			charset := strings.ToLower(params["charset"])
			return charset == "" || charset == "utf-8" || charset == "us-ascii"
		}
	}

	return false
}
