package types

import "time"

type TokenPair struct {
	AccessToken struct {
		Token     string    `json:"token"`
		ExpiresAt time.Time `json:"expiresAt"`
	} `json:"accessToken"`
	RefreshToken struct {
		Token     string    `json:"token"`
		ExpiresAt time.Time `json:"expiresAt"`
	} `json:"refreshToken"`
}
