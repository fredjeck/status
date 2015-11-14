package status_test

import (
	"github.com/fredjeck/status"
	"testing"
)

func TestStatuses(t *testing.T) {
	sw := status.NewWriter(80)
	sw.Pending(" Achievement unlocked")
	sw.MkSuccess().Done()

	sw.Pending(" Something is wrong here")
	sw.MkWarning().Done()

	sw.Pending(" Houston we have a problem")
	sw.MkFailure().Done()

	sw.Pending(" I will pass all the statuses and end failing")
	sw.MkSuccess()
	sw.MkWarning()
	sw.MkFailure().Done()

	sw.Pending(" I'll stay pending")
	sw.Done()

	sw.Pending(" This is a very very very very long status message that should normally be truncated to avoir overflows")
	sw.MkSuccess()
	sw.Done()
}
