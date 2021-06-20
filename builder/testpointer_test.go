package builder

import (
	"testing"
)

func TestTalk(t *testing.T) {
	shout(admin{})
	shout(&admin{})
}
