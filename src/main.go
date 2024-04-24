package main

type InMemoryDB struct {
	data map[string]string	
}


func main() {
	db := InMemoryDB{data: make(map[string]string)}
	db.data["key"] = "value"

	//iterate through the map and print out values
	for key, value := range db.data {
		println(key, value)
	}
}