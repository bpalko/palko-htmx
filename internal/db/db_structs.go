package db

// Build represents a user's build.
type Build struct {
	ID     int64 // Unique identifier for the build
	Barrel string
	Grip   string
	Sight  string
}
