package outlook

import (
	"golang.org/x/oauth2"
)

// Endpoint is Outlook's OAuth 2.0 endpoint in live (production) environment.
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://login.microsoftonline.com/common/oauth2/v2.0/authorize",
	TokenURL: "https://login.microsoftonline.com/common/oauth2/v2.0/token",
}
