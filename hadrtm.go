package packagertm

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertDatagrapic(db *mongo.Database, pendapatan, laba string, keuntungan string, kerugian string) (InsertedID interface{}) {
	var  gp grapic
	gp.Pendapatan = pendapatan
	gp.Laba = laba
	gp.Keuntungan = keuntungan
	gp.Kerugian = kerugian
	return InsertOneDoc(db, "grapic", gp)
}

func GetDataPendapatan(pendapatan string, db *mongo.Database, col string) (data grapic) {
	pdp := db.Collection(col)
	filter := bson.M{"pendapatan": pendapatan}
	err := pdp.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("GetDataPendapatan: %v\n", err)
	}
	return data
}

func GetDataKeuntungan(keuntungan string, db *mongo.Database, col string) (data []grapic) {
	ktg := db.Collection(col)
	filter := bson.M{"keuntungan": keuntungan}
	cursor, err := ktg.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetDataKeuntungan: %v\n", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func DeleteDatagrapic(kerugian string, db *mongo.Database, col string) (data grapic) {
	krg := db.Collection(col)
	filter := bson.M{"kerugian": kerugian}
	err, _ := krg.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("DeleteDataGrapicKinerjaKeungan : %v\n", err)
	}
	fmt.Println("Succes Delete data")
	return data
}