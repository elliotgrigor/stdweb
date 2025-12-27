package internal

const (
	devAuthenticated = true
	devAuthorized    = true
)

func isAuthenticated() bool {
	if devMode {
		return devAuthenticated
	}
	// TODO: Implement authentication logic
	return false
}

func isAuthorised() bool {
	if devMode {
		return devAuthorized
	}
	// TODO: Implement authorisation logic
	return false
}
