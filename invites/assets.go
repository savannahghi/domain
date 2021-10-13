package invites

// TODO: similar to "Attachment" or "Upload" in our older systems

type Asset struct {
	ID string

	Name        string
	Description string
	ContentType string // TODO: RFC ___ mime types
	InternalURL string // TODO: an internal reference to a location on GCS/S3 etc
	PublicURL   string
}

type AssetUpload interface {
	// TODO: secured API
	// Upload fetches and caches a file from a publicly accessible URL.
	// For example: a video can be uploaded to a DropBox/Drive etc location
	// then a public URL is obtained and sent to this service.
	// This service will return a public URL that can then be embedded in content
	// and that will be served in a 'smart' way e.g with proper caching, metrics.
	Upload(
		name string,
		description string,
		contentType string,
		sourceURL string,
	) (*Asset, error)
}
