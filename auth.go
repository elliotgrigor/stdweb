package main

const (
	devAuthN = true
	devAuthZ = true
)

func isAuthenticated() bool {
	if devMode {
		return devAuthN
	}
	// TODO: Implement authentication logic
	return false
}

func isAuthorised() bool {
	if devMode {
		return devAuthZ
	}
	// TODO: Implement authorisation logic
	return false
}
