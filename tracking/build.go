package tracking

import (
	"github.com/secure-dns/lists/scripts"
)

func Build() {
	scripts.ReadLists("lists/tracking", scripts.GetLists("tracking/lists.txt"))
}
