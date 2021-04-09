package cryptojacking

import "github.com/secure-dns/lists/scripts"

func Build() {
	scripts.ReadLists("cryptojacking/cryptojacking.txt", scripts.GetLists("cryptojacking/lists.txt"))
}
