package controllers

import (
	"context"
	"golang-mongodb/hero"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateHero(client *mongo.Client, updatedData bson.M, filter bson.M) int64 {
	collection := client.Database("civilact").Collection("heroes")
	atualizacao := bson.D{{Key: "$set", Value: updatedData}}
	updatedResult, err := collection.UpdateOne(context.TODO(), filter, atualizacao)
	if err != nil {
		log.Fatal("Error on updating one Hero", err)
	}
	return updatedResult.ModifiedCount
}

func RemoveOneHero(client *mongo.Client, filter bson.M) int64 {
	collection := client.Database("civilact").Collection("heroes")
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on deleting one Hero", err)
	}
	return deleteResult.DeletedCount
}

func InsertNewHero(client *mongo.Client, hero hero.Hero) interface{} {

	collection := client.Database("civilact").Collection("heroes")
	insertResult, err := collection.InsertOne(context.TODO(), hero)
	if err != nil {
		log.Fatalln("Error on inserting new Hero", err)
	}
	return insertResult.InsertedID
}

func ReturnOneHero(client *mongo.Client, filter bson.M) hero.Hero {
	var hero hero.Hero
	collection := client.Database("civilact").Collection("heroes")
	documentReturned := collection.FindOne(context.TODO(), filter)
	documentReturned.Decode(&hero)
	return hero
}

func ReturnAllHeroes(client *mongo.Client, filter bson.M) []*hero.Hero {
	var heroes []*hero.Hero
	collection := client.Database("civilact").Collection("heroes")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var hero hero.Hero
		err = cur.Decode(&hero)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		heroes = append(heroes, &hero)
	}
	return heroes
}
