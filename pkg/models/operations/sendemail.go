package operations

import (
	"github.com/resendlabs/resend-go/pkg/models/shared"
	"net/http"
)

type SendEmailResponse struct {
	ContentType       string
	SendEmailResponse *shared.SendEmailResponse
	StatusCode        int
	RawResponse       *http.Response
}
