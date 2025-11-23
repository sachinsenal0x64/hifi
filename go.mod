// Hifi: avoid using external libraries where possible, except for the one below, because I don't want to reinvent the wheel.

module hifi

go 1.25.3

require (
	github.com/alexedwards/argon2id v1.0.0
	github.com/go-co-op/gocron/v2 v2.18.0
	github.com/valkey-io/valkey-go v1.0.68
)

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/jonboulle/clockwork v0.5.0 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
)
