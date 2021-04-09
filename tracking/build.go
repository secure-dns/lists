package tracking

import (
	"github.com/secure-dns/lists/scripts"
)

func Build() {
	scripts.ReadLists("tracking/tracking.txt", scripts.GetLists("tracking/lists.txt"))
}
