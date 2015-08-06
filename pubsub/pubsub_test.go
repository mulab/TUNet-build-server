package pubsub_test

import (
	"testing"

	"github.com/mulab/TUNet-build-server/pubsub"
)

func TestPubsub(t *testing.T) {
	//	compile time assertion
	//	var _ pubsub.Publisher = (*pubsub.Pubsub)(nil)
	//	run time check, both is ok
	var p interface{} = &pubsub.Pubsub{}
	if _, ok := p.(pubsub.MessageQueue); !ok {
		t.Fatal("Pubsub does not implement MessageQueue")
	}
}
