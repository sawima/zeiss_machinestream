package main

import (
	"encoding/json"
	"log"

	"github.com/hashicorp/go-memdb"
)

var ramdb *memdb.MemDB

func init() {
	ramdb = ramdbInit()
}

func ramdbInit() *memdb.MemDB {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"machine": &memdb.TableSchema{
				Name: "machine",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "MachineID"},
					},
				},
			},
		},
	}
	// Create a new data base
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}
	return db
}

func updateNewRecord(wsmsg []byte) {
	wsRecord := &WSMessage{}
	if err := json.Unmarshal(wsmsg, wsRecord); err != nil {
		panic(err)
	}
	machine := wsRecord.Payload
	txn := ramdb.Txn(true)
	existItem, err := txn.First("machine", "id", machine.MachineID)
	if err != nil {
		log.Println("query exist item error")
		panic(err)
	}
	if existItem != nil {
		txn.Delete("machine", existItem)
	}
	if err = txn.Insert("machine", machine); err != nil {
		panic(err)
	}
	txn.Commit()
}

func machineRecords() []Machine {
	txn := ramdb.Txn(false)
	defer txn.Abort()
	machines, err := txn.Get("machine", "id")
	if err != nil {
		panic(err)
	}
	records := make([]Machine, 0)
	for obj := machines.Next(); obj != nil; obj = machines.Next() {
		newMachine := obj.(Machine)
		records = append(records, newMachine)
	}
	return records
}
