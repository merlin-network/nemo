package types_test

import (
	"os"
	"testing"

	"github.com/merlin-network/nemo/app"
)

func TestMain(m *testing.M) {
	app.SetSDKConfig()
	os.Exit(m.Run())
}
