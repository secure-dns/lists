package childprotection

import "github.com/secure-dns/lists/scripts"

func Build() {
	scripts.ReadLists("lists/childprotection", scripts.GetLists("childprotection/lists.txt"))
}
