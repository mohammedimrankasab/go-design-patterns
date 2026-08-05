package builder

import "time"

// HTTPClient represents a configured HTTP client.
type HTTPClient struct {
	BaseURL     string
	Timeout     time.Duration
	Retries     int
	Headers     map[string]string
	EnableTrace bool
}

// ClientBuilder builds HTTPClient instances.
type ClientBuilder struct {
	client HTTPClient
}

func NewClientBuilder() *ClientBuilder {
	return &ClientBuilder{
		client: HTTPClient{
			Timeout: 30 * time.Second,
			Retries: 3,
			Headers: make(map[string]string),
		},
	}
}

// WithBaseURL sets the service endpoint.
func (b *ClientBuilder) WithBaseURL(url string) *ClientBuilder {
	b.client.BaseURL = url
	return b
}

// WithTimeout configures request timeout.
func (b *ClientBuilder) WithTimeout(timeout time.Duration) *ClientBuilder {
	b.client.Timeout = timeout
	return b
}

// WithRetries configures retry attempts.
func (b *ClientBuilder) WithRetries(retries int) *ClientBuilder {
	b.client.Retries = retries
	return b
}

// WithHeader adds a custom HTTP header.
func (b *ClientBuilder) WithHeader(key, value string) *ClientBuilder {
	b.client.Headers[key] = value
	return b
}

// EnableTracing enables request tracing.
func (b *ClientBuilder) EnableTracing() *ClientBuilder {
	b.client.EnableTrace = true
	return b
}

// Build constructs the HTTPClient instance.
func (b *ClientBuilder) Build() HTTPClient {
	return b.client
}
