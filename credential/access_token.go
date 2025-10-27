package credential

import "context"

type AccessTokenHandle interface {
	GetAccessToken() (accessToken string, err error)
}

type AccessTokenContextHandle interface {
	AccessTokenHandle
	GetAccessTokenContext(ctx context.Context) (accessToken string, err error)
}
