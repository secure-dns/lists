package security

import (
	"github.com/secure-dns/lists/scripts"
)

func Build() {
	scripts.ReadLists("lists/security", scripts.GetLists("security/lists.txt"))
}
