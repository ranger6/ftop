// ftop converts a single argument (fully qualified domain name) to a filesystem path.
// usage: ftop fqdn
package main

import (
	"fmt"
	"os"
	"strings"
)

// ftop implements the underlying string to string conversion between fqdn's and paths
func ftop(fqdn string) string {
	var words = strings.Split(fqdn, ".")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, "/")
}

// main expects a single argument (fqdn) and prints the resulting path
func main() {
	fmt.Println(ftop(os.Args[1]))
}
