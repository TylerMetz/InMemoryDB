package main

import (
	"errors"
	"fmt"
)

type InMemoryDB struct {
	data          map[string]int
	transactionOn bool
	transaction   map[string]int
}

func (db *InMemoryDB) beginTransaction() {
	if db.transactionOn {
		fmt.Println("Transaction already in progress.")
	} else {
		db.transactionOn = true
		db.transaction = make(map[string]int)
		fmt.Println("Transaction started.")
	}
}

func (db *InMemoryDB) put(key string, value int) error {
	if !db.transactionOn {
		return errors.New("Transaction not in progress")
	}
	db.transaction[key] = value
	fmt.Printf("Key '%s' set to %d in transaction.\n", key, value)
	return nil
}

func (db *InMemoryDB) get(key string) (int, error) {
	if val, ok := db.transaction[key]; ok {
		return val, nil
	}
	if val, ok := db.data[key]; ok {
		return val, nil
	}
	return 0, errors.New("Key not found")
}

func (db *InMemoryDB) commit() {
	if db.transactionOn {
		for key, val := range db.transaction {
			db.data[key] = val
		}
		db.transactionOn = false
		db.transaction = nil
		fmt.Println("Transaction committed.")
	} else {
		fmt.Println("No transaction in progress to commit.")
	}
}

func (db *InMemoryDB) rollback() {
	if db.transactionOn {
		db.transactionOn = false
		db.transaction = nil
		fmt.Println("Transaction rolled back.")
	} else {
		fmt.Println("No transaction in progress to rollback.")
	}
}

func main() {
	inmemoryDB := InMemoryDB{data: make(map[string]int)}

	// Example usage as per provided instructions
	val, err := inmemoryDB.get("A")
	fmt.Println("get(\"A\"):", val, err) // Should return null

	err = inmemoryDB.put("A", 5) // Set value of A to 5 within the transaction
	fmt.Println("put(\"A\", 5):", err)

	inmemoryDB.beginTransaction()
	fmt.Println("beginTransaction():") // Starts a new transaction

	err = inmemoryDB.put("A", 5) // Set value of A to 5 within the transaction
	fmt.Println("put(\"A\", 5):", err)

	val, err = inmemoryDB.get("A") // Should return null, because updates to A are not committed yet
	fmt.Println("get(\"A\"):", val, err)

	err = inmemoryDB.put("A", 6) // Update A’s value to 6 within the transaction
	fmt.Println("put(\"A\", 6):", err)

	inmemoryDB.commit() // Commits the open transaction
	fmt.Println("commit():")

	val, err = inmemoryDB.get("A") // Should return 6, the last value of A that was committed
	fmt.Println("get(\"A\"):", val, err)

	inmemoryDB.commit() // Commits the not-open transaction

	inmemoryDB.rollback()

	val, err = inmemoryDB.get("B") // Should return null because B does not exist in the database
	fmt.Println("get(\"B\"):", val, err)

	inmemoryDB.beginTransaction() // Starts a new transaction
	fmt.Println("beginTransaction():")

	err = inmemoryDB.put("B", 10) // Set key B’s value to 10 within the transaction
	fmt.Println("put(\"B\", 10):", err)

	inmemoryDB.rollback() // Rollback the transaction - revert any changes made to B
	fmt.Println("rollback():")

	val, err = inmemoryDB.get("B") // Should return null because changes to B were rolled back
	fmt.Println("get(\"B\"):", val, err)
}
