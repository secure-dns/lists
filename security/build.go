package security

import (
	"github.com/secure-dns/lists/scripts"
)

func Build() {
	scripts.ReadLists("security/security.txt", scripts.GetLists("security/lists.txt"))
}
