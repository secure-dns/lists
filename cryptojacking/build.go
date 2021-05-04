package cryptojacking

import "github.com/secure-dns/lists/scripts"

func Build() {
	scripts.ReadLists("lists/cryptojacking", scripts.GetLists("cryptojacking/lists.txt"))
}
