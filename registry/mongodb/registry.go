/*
#######
##        ____    _____ ____ _
##       (_-< |/|/ / _ `/ _ `/
##      /___/__,__/\_,_/\_, /
##                     /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_registry "github.com/arnumina/swag/component/registry"
	"github.com/arnumina/swag/util"
	"github.com/arnumina/swag/util/mongodb"
)

type registry struct {
	interval  int
	uri       string
	cMongo    *mongodb.Client
	cMutex    *mongo.Collection
	cRegistry *mongo.Collection
}

func (r *registry) build() (*registry, error) {
	client := mongodb.NewClient("swag.registry", r.uri)

	if err := client.Connect(); err != nil {
		return nil, err
	}

	swag := client.Database("swag")

	r.cMongo = client
	r.cMutex = swag.Collection("registry.mutex")
	r.cRegistry = swag.Collection("registry")

	return r, nil
}

// Interval AFAIRE
func (r *registry) Interval() int {
	return r.interval
}

func (r *registry) lock(owner string) error {
	ctx, cancel := context.WithTimeout(r.cMongo.Context(), 5*time.Second)
	defer cancel()

	for {
		result := r.cMutex.FindOneAndUpdate(
			ctx,
			bson.D{
				{Key: "_id", Value: "registry"},
				{Key: "locked", Value: false},
			},
			bson.M{
				"$set": bson.M{"locked": true, "owner": owner},
			},
		)

		err := result.Err()
		if err == nil {
			return nil
		}

		if err != mongo.ErrNoDocuments {
			return err
		}

		time.Sleep(50 * time.Millisecond)
	}
}

func (r *registry) unlock(owner string) {
	_, err := r.cMutex.UpdateOne(
		r.cMongo.Context(),
		bson.D{
			{Key: "_id", Value: "registry"},
			{Key: "locked", Value: true},
			{Key: "owner", Value: owner},
		},
		bson.M{
			"$set": bson.M{"locked": false},
		},
	)
	if err != nil {
		util.Alert(err)
	}
}

// Preregister AFAIRE
func (r *registry) Preregister(id, name string, fn func([]int) (*_registry.Service, error)) error {
	owner := name + "/" + id

	if err := r.lock(owner); err != nil {
		return err
	}

	defer r.unlock(owner)

	cursor, err := r.cRegistry.Find(r.cMongo.Context(), bson.D{})
	if err != nil {
		return err
	}

	var services _registry.Services

	if err = cursor.All(r.cMongo.Context(), &services); err != nil {
		return err
	}

	used := []int{}

	for _, s := range services {
		used = append(used, s.Port)
	}

	service, err := fn(used)
	if err != nil {
		return err
	}

	_, err = r.cRegistry.InsertOne(r.cMongo.Context(), service)

	return err
}

// Register AFAIRE
func (r *registry) Register(service *_registry.Service) error {
	_, err := r.cRegistry.UpdateOne(
		r.cMongo.Context(),
		bson.M{"_id": service.ID},
		bson.M{"$set": service},
		options.Update().SetUpsert(true),
	)

	return err
}

// Deregister AFAIRE
func (r *registry) Deregister(id, _ string) error {
	_, err := r.cRegistry.DeleteOne(
		r.cMongo.Context(),
		bson.M{"_id": id},
	)

	return err
}

// Get AFAIRE
func (r *registry) Get(name string) (_registry.Services, error) {
	cursor, err := r.cRegistry.Find(
		r.cMongo.Context(),
		bson.D{
			{Key: "name", Value: name},
		},
	)
	if err != nil {
		return nil, err
	}

	var services _registry.Services

	if err = cursor.All(r.cMongo.Context(), &services); err != nil {
		return nil, err
	}

	return services, nil
}

// List AFAIRE
func (r *registry) List() (_registry.Services, error) {
	cursor, err := r.cRegistry.Find(
		r.cMongo.Context(),
		bson.D{},
	)
	if err != nil {
		return nil, err
	}

	var services _registry.Services

	if err = cursor.All(r.cMongo.Context(), &services); err != nil {
		return nil, err
	}

	return services, nil
}

// Close AFAIRE
func (r *registry) Close() error {
	if r.cMongo == nil {
		return nil
	}

	return r.cMongo.Disconnect()
}

/*
######################################################################################################## @(°_°)@ #######
*/
