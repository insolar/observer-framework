// +build unit integration

package observerframework

import (
	"testing"
)

func Test_for_workflow(t *testing.T) {
	ci_workflow()
	t.Log("I'm robot!")
}
