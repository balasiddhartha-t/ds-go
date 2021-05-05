//Creating a custom Set file in golang
package main

import "fmt"

type customSet map[string]struct{}

func (s customSet) add(value string) {
	s[value] = struct{}{}
}

func (s customSet) del(value string) {
	delete(s, value)
}

func (s customSet) has(value string) bool {
	_, ok := s[value]
	return ok
}

func main() {
	names := customSet{}
	names.add("siddhartha")
	names.add("bala")
	names.add("bala siddhartha")
	fmt.Println(names.has("bala siddhartha"))
	names.del("bala siddhartha")
	for v := range names {
		fmt.Println(v)
	}

}
