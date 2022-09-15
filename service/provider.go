package service

import (
	"context"
	"fmt"
	"mime"
	"os"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

var stderr = os.Stderr

func New() provider.Provider {
	return &TerraformProvider{}
}

var _ provider.Provider = &TerraformProvider{}

type TerraformProvider struct {
	configured bool
}

// Provider schema struct
type providerData struct{}

// With the provider.Provider implementation
func (p *TerraformProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "http"
}

func (p *TerraformProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
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
func (p *TerraformProvider) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{},
	}, nil
}

// DataSources - Defines provider data sources
func (p *TerraformProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewDataSourceHttpRequest,
	}
}

// Resources - Defines provider resources
func (p *TerraformProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewResourceHttpRequest,
	}
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
