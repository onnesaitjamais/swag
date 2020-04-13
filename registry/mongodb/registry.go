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
	"github.com/arnumina/swag/runner"
	"github.com/arnumina/swag/util"
	"github.com/arnumina/swag/util/failure"
	"github.com/arnumina/swag/util/mongodb"
)

type registry struct {
	runner     *runner.Runner
	interval   int
	uri        string
	clMongo    *mongodb.Client
	coMutex    *mongo.Collection
	coRegistry *mongo.Collection
}

type document struct {
	ID      string `bson:"_id"`
	Service *_registry.Service
}

func (r *registry) build() (*registry, error) {
	client := mongodb.NewClient("swag.registry", r.uri)

	if err := client.Connect(); err != nil {
		return nil, err
	}

	swag := client.Database("swag")

	r.clMongo = client
	r.coMutex = swag.Collection("registry.mutex")
	r.coRegistry = swag.Collection("registry")

	return r, nil
}

// Interval AFAIRE
func (r *registry) Interval() int {
	return r.interval
}

func (r *registry) lock(owner string) error {
	ctx, cancel := context.WithTimeout(r.clMongo.Context(), 5*time.Second)
	defer cancel()

	for {
		result, err := r.coMutex.UpdateOne(
			ctx,
			bson.D{
				{Key: "_id", Value: "registry"},
				{Key: "locked", Value: false},
			},
			bson.M{"$set": bson.M{"locked": true, "owner": owner}},
		)
		if err != nil {
			return err
		}

		if result.ModifiedCount != 0 {
			return nil
		}

		time.Sleep(50 * time.Millisecond)
	}
}

func (r *registry) unlock(owner string) {
	_, err := r.coMutex.UpdateOne(
		r.clMongo.Context(),
		bson.D{
			{Key: "_id", Value: "registry"},
			{Key: "locked", Value: true},
			{Key: "owner", Value: owner},
		},
		bson.M{"$set": bson.M{"locked": false}},
	)
	if err != nil {
		f := failure.New(err).
			Set("component", "registry").
			Set("collection", r.coMutex.Name()).
			Set("document", "registry").
			Msg("Impossible to unlock the single document") ////////////////////////////////////////////////////////////

		r.runner.Logger().Error(f.Error())

		util.Alert(f)
	}
}

func (r *registry) find(ctx context.Context, filter interface{}) (_registry.Services, error) {
	cursor, err := r.coRegistry.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var ds []document

	if err = cursor.All(r.clMongo.Context(), &ds); err != nil {
		return nil, err
	}

	var s _registry.Services

	for _, d := range ds {
		s = append(s, d.Service)
	}

	return s, nil
}

// Preregister AFAIRE
func (r *registry) Preregister(id, name string, fn func([]int) (*_registry.Service, error)) error {
	owner := id + "@" + name

	if err := r.lock(owner); err != nil {
		return err
	}

	defer r.unlock(owner)

	services, err := r.find(r.clMongo.Context(), bson.M{"service.port": bson.M{"$gt": 0}})
	if err != nil {
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

	d := &document{
		ID:      service.ID,
		Service: service,
	}

	_, err = r.coRegistry.InsertOne(r.clMongo.Context(), d)

	return err
}

// Register AFAIRE
func (r *registry) Register(service *_registry.Service) error {
	_, err := r.coRegistry.UpdateOne(
		r.clMongo.Context(),
		bson.M{"_id": service.ID},
		bson.M{"$set": bson.M{"service": service}},
		options.Update().SetUpsert(true),
	)

	return err
}

// Deregister AFAIRE
func (r *registry) Deregister(id, _ string) error {
	_, err := r.coRegistry.DeleteOne(
		r.clMongo.Context(),
		bson.M{"_id": id},
	)

	return err
}

// Find AFAIRE
func (r *registry) Find(name string) (_registry.Services, error) {
	return r.find(r.clMongo.Context(), bson.M{"name": name})
}

// List AFAIRE
func (r *registry) List() (_registry.Services, error) {
	return r.find(r.clMongo.Context(), bson.D{})
}

// Close AFAIRE
func (r *registry) Close() error {
	if r.clMongo == nil {
		return nil
	}

	return r.clMongo.Disconnect()
}

/*
######################################################################################################## @(°_°)@ #######
*/
