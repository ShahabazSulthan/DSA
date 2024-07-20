package main

import "fmt"

func main() {
	Hash_Table := make(map[string]int)
	Hash_Table["A"] = 1
	Hash_Table["B"] = 2
	Hash_Table["C"] = 3
	fmt.Println(Hash_Table)
	fmt.Println(Hash_Table["A"])
	delete(Hash_Table,"A")
	fmt.Println(Hash_Table)

	hashTable()

	structMap()

}