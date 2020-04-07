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

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

// Client AFAIRE
type Client struct {
	options *options.ClientOptions
	ctx     context.Context
	client  *mongo.Client
}

// NewClient AFAIRE
func NewClient(app, uri string) *Client {
	options := options.Client().
		SetAppName(app).
		SetConnectTimeout(5 * time.Second).
		SetReadConcern(readconcern.Majority()).
		SetServerSelectionTimeout(5 * time.Second).
		SetWriteConcern(
			writeconcern.New(
				writeconcern.J(true),
				writeconcern.WMajority(),
				writeconcern.WTimeout(2*time.Second),
			),
		).
		ApplyURI(uri)

	ctx := context.Background()

	return &Client{
		options: options,
		ctx:     ctx,
	}
}

// Options AFAIRE
func (c *Client) Options() *options.ClientOptions {
	return c.options
}

// Context AFAIRE
func (c *Client) Context() context.Context {
	return c.ctx
}

// Connect AFAIRE
func (c *Client) Connect() error {
	client, err := mongo.Connect(c.ctx, c.options)
	if err != nil {
		return err
	}

	err = client.Ping(c.ctx, nil)
	if err != nil {
		return err
	}

	c.client = client

	return nil
}

// Client AFAIRE
func (c *Client) Client() *mongo.Client {
	return c.client
}

// Database AFAIRE
func (c *Client) Database(name string, opts ...*options.DatabaseOptions) *mongo.Database {
	return c.client.Database(name, opts...)
}

// Disconnect AFAIRE
func (c *Client) Disconnect() error {
	return c.client.Disconnect(c.ctx)
}

/*
######################################################################################################## @(°_°)@ #######
*/
