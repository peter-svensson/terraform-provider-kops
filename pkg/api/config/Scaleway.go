package config

type Scaleway struct {
	// AccessKey defines the Scaleway access key
	AccessKey string
	// SecretKey defines the Scaleway secret key
	SecretKey string
	// ProjectId defines the Scaleway project ID
	ProjectId string
	// Region defines the Scaleway region (e.g. fr-par, nl-ams)
	Region string
	// Zone defines the Scaleway zone (e.g. fr-par-1, nl-ams-1)
	Zone string
	// Profile defines the Scaleway profile to load from config file
	Profile string
	// KopsControllerImage overrides the kops-controller container image
	KopsControllerImage string
}
