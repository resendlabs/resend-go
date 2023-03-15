<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/resendlabs/resend-go"
    "github.com/resendlabs/resend-go/pkg/models/shared"
    "github.com/resendlabs/resend-go/pkg/models/operations"
)

func main() {
    s := resend.New(
        WithSecurity(        shared.Security{
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    req := shared.Email{
        Bcc: "unde",
        Cc: "deserunt",
        From: "porro",
        HTML: "nulla",
        ReplyTo: "id",
        Subject: "vero",
        Text: "perspiciatis",
        To: "nulla",
    }

    ctx := context.Background()
    res, err := s.Email.SendEmail(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.SendEmailResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->