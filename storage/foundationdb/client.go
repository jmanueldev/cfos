package storage

import "github.com/apple/foundationdb/bindings/go/src/fdb"

func Open() fdb.Database {

    fdb.MustAPIVersion(710)

    return fdb.MustOpenDefault()
}