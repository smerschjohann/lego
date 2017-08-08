package godaddy

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	godaddyAPIKey    string
	godaddyAPISecret string
	godaddyDomain    string
)

func init() {
	godaddyAPIKey = os.Getenv("GODADDY_API_KEY")
	godaddyAPISecret = os.Getenv("GODADDY_API_SECRET")
	godaddyDomain = os.Getenv("GODADDY_DOMAIN")
}

func TestDNSProvider_Present(t *testing.T) {
	provider, err := NewDNSProvider()
	assert.NoError(t, err)

	err = provider.Present(godaddyDomain, "", "123d==")
	assert.NoError(t, err)
}

func TestDNSProvider_CleanUp(t *testing.T) {
	provider, err := NewDNSProvider()
	assert.NoError(t, err)

	err = provider.CleanUp(godaddyDomain, "", "123d==")
	assert.NoError(t, err)
}
