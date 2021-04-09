package ads

import "github.com/secure-dns/lists/scripts"

func Build() {
	scripts.ReadLists("ads/ads.txt", scripts.GetLists("ads/lists.txt"))
}
