// Copyright 2016 NewsAI. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package heroku provides constants for using OAuth2 to access Heroku.
package instagram // import "github.com/news-ai/oauth2/instagram"

import (
	"golang.org/x/oauth2"
)

// Endpoint is Heroku's OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://api.instagram.com/oauth/authorize",
	TokenURL: "https://api.instagram.com/oauth/access_token",
}
