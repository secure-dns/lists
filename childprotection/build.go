package childprotection

import "github.com/secure-dns/lists/scripts"

func Build() {
	scripts.ReadLists("childprotection/childprotection.txt", scripts.GetLists("childprotection/lists.txt"))
}
