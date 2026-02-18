// Package tests provides test configuration and utilities
package tests

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/rsl6/loyalty-client/client"
	"github.com/rsl6/loyalty-client/mock"
)

// TestConfig holds test configuration
type TestConfig struct {
	// UseMockServer determines whether to use mock server or real server
	UseMockServer bool
	// RealServerURL is the URL of the real server (used when UseMockServer is false)
	RealServerURL string
}

// GetTestConfig returns test configuration from environment variables
func GetTestConfig() *TestConfig {
	useMock := os.Getenv("USE_MOCK_SERVER") != "false"
	realURL := os.Getenv("REAL_SERVER_URL")
	if realURL == "" {
		realURL = "http://localhost:8080"
	}

	return &TestConfig{
		UseMockServer: useMock,
		RealServerURL: realURL,
	}
}

// TestContext holds test context including client and mock server
type TestContext struct {
	Client     *client.Client
	MockServer *mock.Server
	HTTPServer *httptest.Server
	Config     *TestConfig
}

// SetupTestContext creates a new test context
func SetupTestContext(t *testing.T) *TestContext {
	t.Helper()

	config := GetTestConfig()
	ctx := &TestContext{
		Config: config,
	}

	if config.UseMockServer {
		// Create mock server
		ctx.MockServer = mock.NewServer()
		ctx.HTTPServer = httptest.NewServer(ctx.MockServer)

		// Create client pointing to mock server
		ctx.Client = client.NewClient(&client.Config{
			BaseURL: ctx.HTTPServer.URL,
		})
	} else {
		// Create client pointing to real server
		ctx.Client = client.NewClient(&client.Config{
			BaseURL: config.RealServerURL,
		})
	}

	return ctx
}

// Cleanup cleans up test resources
func (ctx *TestContext) Cleanup() {
	if ctx.HTTPServer != nil {
		ctx.HTTPServer.Close()
	}
}

// SkipIfMock skips test if using mock server
func (ctx *TestContext) SkipIfMock(t *testing.T, reason string) {
	if ctx.Config.UseMockServer {
		t.Skipf("Skipping test with mock server: %s", reason)
	}
}

// SkipIfReal skips test if using real server
func (ctx *TestContext) SkipIfReal(t *testing.T, reason string) {
	if !ctx.Config.UseMockServer {
		t.Skipf("Skipping test with real server: %s", reason)
	}
}
