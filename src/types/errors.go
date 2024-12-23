package types

import "errors"

var (
	ErrTokenNotFound   = errors.New("bot token is required in the environment variable 'TOKEN'")
	ErrGuildIDNotFound = errors.New("guild ID is required in the environment variable 'GUILDID'")
)
