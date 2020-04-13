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
	"container/list"
	"regexp"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_broker "github.com/arnumina/swag/component/broker"
	"github.com/arnumina/swag/runner"
	"github.com/arnumina/swag/util"
	"github.com/arnumina/swag/util/failure"
	"github.com/arnumina/swag/util/mongodb"
)

const (
	_maxPendingDocuments = 100
	_maxNoMsgTimeout     = 2
	_noMsgDelay          = 50
	_maxErrorTimeout     = 10
	_maxRetries          = 3
	_retryDelay          = 30
)

type broker struct {
	runner     *runner.Runner
	uri        string
	bindings   map[string][]*regexp.Regexp
	clMongo    *mongodb.Client
	dbSwag     *mongo.Database
	mutex      sync.Mutex
	pDocuments *list.List
}

type document struct {
	ID       string `bson:"_id"`
	Booked   bool
	Retries  int
	UseAfter time.Time
	Message  *_broker.Message
}

type pendingDocument struct {
	queue string
	doc   *document
}

func (b *broker) build() (*broker, error) {
	client := mongodb.NewClient("swag.broker", b.uri)

	if err := client.Connect(); err != nil {
		return nil, err
	}

	b.clMongo = client
	b.dbSwag = client.Database("swag")

	b.pDocuments = list.New()

	return b, nil
}

func (b *broker) insert(queue string, d *document) error {
	_, err := b.dbSwag.Collection("broker."+queue).InsertOne(
		b.clMongo.Context(),
		d,
	)
	if err != nil {
		b.runner.Logger().Warning( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"Unable to publish this message",
			"id", d.ID,
			"queue", queue,
			"reason", err.Error(),
		)

		return err
	}

	b.runner.Logger().Trace( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
		"Message published",
		"id", d.ID,
		"queue", queue,
		"elapsed", time.Since(d.Message.CreatedAt).String(),
	)

	return nil
}

func (b *broker) enqueue(queue string, d *document) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.pDocuments.Len() == 0 {
		if err := b.insert(queue, d); err == nil {
			return
		}
	}

	b.runner.Logger().Notice( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
		"The publication of this message is delayed",
		"id", d.ID,
		"queue", queue,
		"pending", b.pDocuments.Len(),
	)

	b.pDocuments.PushBack(
		&pendingDocument{
			queue: queue,
			doc:   d,
		},
	)

	for b.pDocuments.Len() > 0 {
		e := b.pDocuments.Front()
		pd := e.Value.(*pendingDocument)

		if err := b.insert(pd.queue, pd.doc); err != nil {
			if b.pDocuments.Len() <= _maxPendingDocuments {
				return
			}

			b.runner.Logger().Error( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
				"This message will be destroyed",
				"id", pd.doc.ID,
				"queue", pd.queue,
			)
		}

		b.pDocuments.Remove(e)
	}
}

// Publish AFAIRE
func (b *broker) Publish(event string, data interface{}) error {
	d := &document{
		ID:       util.NewUUID(),
		Booked:   false,
		Retries:  0,
		UseAfter: time.Now(),
		Message: &_broker.Message{
			Event:     event,
			Payload:   data,
			CreatedAt: time.Now(),
		},
	}

	b.runner.Logger().Trace( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
		"New message",
		"id", d.ID,
		"event", event,
	)

	for queue, events := range b.bindings {
		for _, re := range events {
			if re.MatchString(event) {
				b.enqueue(queue, d)
				break
			}
		}
	}

	return nil
}

func (b *broker) findOne(co *mongo.Collection) (*document, error) {
	var d document
	if err := co.FindOneAndUpdate(
		b.clMongo.Context(),
		bson.D{
			{Key: "booked", Value: false},
			{Key: "useafter", Value: bson.M{"$lte": time.Now()}},
		},
		bson.M{"$set": bson.M{"booked": true}},
		options.FindOneAndUpdate().SetSort(bson.D{{Key: "useafter", Value: 1}}),
	).Decode(&d); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		b.runner.Logger().Warning( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
			"Impossible to consume a message",
			"reason", err.Error(),
		)

		return nil, err
	}

	b.runner.Logger().Trace( //:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
		"Message consumed",
		"id", d.ID,
		"elapsed", time.Since(d.Message.CreatedAt).String(),
	)

	return &d, nil
}

func (b *broker) ack(co *mongo.Collection, d *document) {
	_, err := co.DeleteOne(b.clMongo.Context(), bson.M{"_id": d.ID})
	if err != nil {
		f := failure.New(err).
			Set("component", "broker").
			Set("collection", co.Name()).
			Set("id", d.ID).
			Msg("Cannot delete this message") //////////////////////////////////////////////////////////////////////////

		b.runner.Logger().Error(f.Error())

		util.Alert(f)
	}
}

func (b *broker) nack(co *mongo.Collection, d *document) {
	if d.Retries < _maxRetries {
		_, err := co.UpdateOne(
			b.clMongo.Context(),
			bson.M{"_id": d.ID},
			bson.M{"$set": bson.M{
				"booked":   false,
				"retries":  d.Retries + 1,
				"useafter": d.UseAfter.Add(_retryDelay * time.Second),
			}},
		)
		if err != nil {
			f := failure.New(err).
				Set("component", "broker").
				Set("collection", co.Name()).
				Set("id", d.ID).
				Msg("Unable to update this message for a new attempt") /////////////////////////////////////////////////

			b.runner.Logger().Error(f.Error())

			util.Alert(f)
		}
	} else {
		f := failure.New(nil).
			Set("component", "broker").
			Set("collection", co.Name()).
			Set("id", d.ID).
			Msg("Cannot process this message") /////////////////////////////////////////////////////////////////////////

		b.runner.Logger().Error(f.Error())

		util.Alert(f)
	}
}

func (b *broker) consume(td time.Duration, co *mongo.Collection, fn func(*_broker.Message) bool) time.Duration {
	d, err := b.findOne(co)
	if err != nil {
		return _maxErrorTimeout * time.Second
	}

	if d == nil {
		if td < _maxNoMsgTimeout*time.Second {
			td += _noMsgDelay * time.Millisecond
		}

		return td
	}

	if fn(d.Message) {
		b.ack(co, d)
	} else {
		b.nack(co, d)
	}

	return 0
}

// Subscribe AFAIRE
func (b *broker) Subscribe(queue string, fn func(*_broker.Message) bool) {
	collection := b.dbSwag.Collection("broker." + queue)

	b.runner.AddGroupFn(
		func(stop <-chan struct{}) error {
			delay := 1 * time.Second

		loop:
			for {
				select {
				case <-stop:
					break loop
				case <-time.After(delay):
					delay = b.consume(delay, collection, fn)
				}
			}

			return nil
		},
	)
}

// Close AFAIRE
func (b *broker) Close() error {
	if b.clMongo == nil {
		return nil
	}

	return b.clMongo.Disconnect()
}

/*
######################################################################################################## @(°_°)@ #######
*/
