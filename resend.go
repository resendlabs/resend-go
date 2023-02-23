package resend

import (
	"github.com/resendlabs/resend-go/pkg/models/shared"
	"github.com/resendlabs/resend-go/pkg/utils"
	"net/http"
	"time"
)

var ServerList = []string{
	"https://api.resend.com",
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Resend struct {
	Email *email

	// Non-idiomatic field names below are to namespace fields from the fields names above to avoid name conflicts
	_defaultClient  HTTPClient
	_securityClient HTTPClient
	_security       *shared.Security
	_serverURL      string
	_language       string
	_sdkVersion     string
	_genVersion     string
}

type SDKOption func(*Resend)

func WithServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Resend) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk._serverURL = serverURL
	}
}

func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Resend) {
		sdk._defaultClient = client
	}
}

func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *Resend) {
		sdk._security = &security
	}
}

func New(opts ...SDKOption) *Resend {
	sdk := &Resend{
		_language:   "go",
		_sdkVersion: "1.1.2",
		_genVersion: "1.4.8",
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk._defaultClient == nil {
		sdk._defaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk._securityClient == nil {

		if sdk._security != nil {
			sdk._securityClient = utils.ConfigureSecurityClient(sdk._defaultClient, sdk._security)
		} else {
			sdk._securityClient = sdk._defaultClient
		}

	}

	if sdk._serverURL == "" {
		sdk._serverURL = ServerList[0]
	}

	sdk.Email = newEmail(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	return sdk
}
