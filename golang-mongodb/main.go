package main

import (
	"context"
	"golang-mongodb/controllers"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb+srv://root:root@blog.7t70c.mongodb.net/test?retryWrites=true&w=majority")

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func main() {

	c := GetClient()
	err := c.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}

	// insert hero
	// var hero hero.Hero
	// hero = Hero{Name: "T", Alias: "W", Signed: true}
	// insertedID := controllers.InsertNewHero(c, hero)
	// log.Println(insertedID)

	//return one hero
	// var hero Hero
	// hero = ReturnOneHero(c, bson.M{"alias": "Doctor Strange"})
	// log.Println(hero.Name, hero.Alias, hero.Signed)

	//delete one hero
	// heroesRemoved := RemoveOneHero(c, bson.M{"alias": "Doctor Strange"})
	// log.Println("Heroes removed count:", heroesRemove
	// hero = ReturnOneHero(c, bson.M{"alias": "Doctor Strange"})
	// log.Println("Is Hero empty?", hero == Hero{ })

	//update data
	// UpdatedHero := UpdateHero(c, bson.M{"alias": "Doctor TOl"}, bson.M{"alias": "Doctor Strange"})
	// log.Println("Heroes update count:", UpdatedHero)

	//call all hero
	heroes := controllers.ReturnAllHeroes(c, bson.M{})
	for _, hero := range heroes {
		log.Println(hero.Name, hero.Alias, hero.Signed)
	}
}
