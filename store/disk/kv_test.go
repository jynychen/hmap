package disk

import (
	"path/filepath"
	"testing"
)

func TestKV(t *testing.T) {
	testOperations := TestOperations{
		Set:    true,
		Get:    true,
		Scan:   true,
		Delete: true,
	}
	var db DB
	// bbolt
	dbpath, _ := utiltestGetPath(t)
	dbpath = filepath.Join(dbpath, "boltdb")
	db, err := OpenBoltDBB(dbpath)
	if err != nil {
		t.Error(err)
	}
	db.(*BBoltDB).BucketName = "test"
	utiltestOperations(t, db, 100, testOperations)
	utiltestRemoveDb(t, db, dbpath)

	// pogreb
	dbpath, _ = utiltestGetPath(t)
	db, err = OpenPogrebDB(dbpath)
	if err != nil {
		t.Error(err)
	}
	utiltestOperations(t, db, 100, testOperations)
	utiltestRemoveDb(t, db, dbpath)

	// leveldb
	dbpath, _ = utiltestGetPath(t)
	db, err = OpenLevelDB(dbpath)
	if err != nil {
		t.Error(err)
	}
	utiltestOperations(t, db, 100, testOperations)
	utiltestRemoveDb(t, db, dbpath)

	// filedb
	testOperations.Get = false
	testOperations.Delete = false
	dbpath, _ = utiltestGetPath(t)
	dbpath = filepath.Join(dbpath, "filedb")
	db, err = OpenFileDB(dbpath)
	if err != nil {
		t.Error(err)
	}
	utiltestOperations(t, db, 100, testOperations)
	utiltestRemoveDb(t, db, dbpath)
}
