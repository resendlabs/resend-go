# Resend Go SDK

Resend is the email platform for developers. Learn more on our [docsite](https://resend.com/docs/api-reference/concepts) 

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/resendlabs/resend-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
    "github.com/resendlabs/resend-go"
    "github.com/resendlabs/resend-go/pkg/models/shared"
    "github.com/resendlabs/resend-go/pkg/models/operations"
)

func main() {
    opts := []sdk.SDKOption{
        sdk.WithSecurity(
            shared.Security{
                BearerAuth: shared.SchemeBearerAuth{
                    Authorization: "Bearer YOUR_BEARER_TOKEN_HERE",
                },
            }
        ),
    }

    s := sdk.New(opts...)
    
    req := operations.SendEmailRequest{
        Request: shared.Email{
            Bcc: "sit",
            Cc: "voluptas",
            From: "culpa",
            HTML: "expedita",
            React: "consequuntur",
            ReplyTo: "dolor",
            Subject: "expedita",
            Text: "voluptas",
            To: "fugit",
        },
    }
    
    res, err := s.Emails.SendEmail(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.SendEmailResponse != nil {
        // handle response
    }
```
<!-- End SDK Example Usage -->

## Authentication

To authenticate you need to add an Authorization header with the contents of the header being Bearer re_123456789 where re_123456789 is your [API Key](https://resend.com/login?redirectedFrom=%2Fapi-keys).

```bash
Authorization: Bearer re_123
```

<!-- Start SDK Available Operations -->
## SDK Available Operations

### emails

* `SendEmail` - Send an email

<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
