package models

// OauthUser user requesting access
type OauthUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// OauthAccessToken jwt access token for authentication
type OauthAccessToken struct {
	Token string `json:"token"`
}

// Response message for the request
type Response struct {
	Message string `json:"message"`
}

/*
type OauthClient struct {
	ID          string `json:"id"`
	Key         string `json:"key"`
	Secret      string `json:"secret"`
	RedirectURI string `json:"redirect"`
}

// Token ...
type Token struct {
	ClientID  string `json:"clientID"`
	UserID    string `json:"userID"`
	Client    *OauthClient
	User      *OauthUser
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expireAt"`
	Scope     string    `json:"scope"`
}

// OauthAuthorizationCode ...
type OauthAuthorizationCode struct {
	ClientID    string `json:"clientID"`
	UserID      string `json:"userID"`
	Client      *OauthClient
	User        *OauthUser
	Code        string    `json:"code"`
	RedirectURI string    `json:"redirect"`
	ExpiresAt   time.Time `json:"expireAt"`
	Scope       string    `json:"scope"`
}
*/
