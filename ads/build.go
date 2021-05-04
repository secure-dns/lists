package ads

import "github.com/secure-dns/lists/scripts"

func Build() {
	scripts.ReadLists("lists/ads", scripts.GetLists("ads/lists.txt"))
}
