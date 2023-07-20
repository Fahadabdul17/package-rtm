package packagertm

import (
	"fmt"
	"os"
	"testing"

	"github.com/aiteung/atdb"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "grapic",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func TestInsertGrapicKinerjaKeuangan(t *testing.T) {
	Pendapatan := "90%"
	Keuntungan := "18%"
	Laba := "76%"
	Kerugian := "5%"
	hasil := InsertDatagrapic(MongoConn, Pendapatan, Keuntungan, Laba, Kerugian)
	fmt.Println(hasil)

}

func TestGetDataPendapatan(t *testing.T) {
	Pendapatan := "90%"
	hasil := GetDataPendapatan(Pendapatan, MongoConn, "grapic")
	fmt.Println(hasil)

}

func TestGetDataKeuntungan(t *testing.T) {
	Keuntungan := "18%"
	hasil := GetDataKeuntungan(Keuntungan, MongoConn, "grapic")
	fmt.Println(hasil)

}

func TestDeleteDataGrapicKinerjaKeuangan(t *testing.T) {
	Kerugian := "5%"
	hasil := DeleteDatagrapic(Kerugian, MongoConn, "grapic")
	fmt.Println(hasil)

}

