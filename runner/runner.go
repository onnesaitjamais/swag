/*
#######
##        ____    _____ ____ _
##       (_-< |/|/ / _ `/ _ `/
##      /___/__,__/\_,_/\_, /
##                     /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package runner

import (
	"strconv"
	"time"

	"github.com/heptio/workgroup"

	"github.com/arnumina/swag/component"
	"github.com/arnumina/swag/util"
	"github.com/arnumina/swag/util/value"
)

// Runner AFAIRE
type Runner struct {
	id         string
	name       string
	version    string
	builtAt    time.Time
	startedAt  time.Time
	fqdn       string
	sdInstance string
	cfgValue   *value.Value

	group workgroup.Group

	broker   component.Broker
	config   component.Config
	logger   component.Logger
	registry component.Registry
	server   component.Server
}

// New AFAIRE
func New(name, version, builtAt string) (*Runner, error) {
	ts, err := strconv.ParseInt(builtAt, 0, 64)
	if err != nil {
		return nil, err
	}

	fqdn, err := util.FQDN()
	if err != nil {
		return nil, err
	}

	runner := &Runner{
		id:        util.NewUUID(),
		name:      name,
		version:   version,
		builtAt:   time.Unix(ts, 0),
		startedAt: time.Now(),
		fqdn:      fqdn,
	}

	return runner, nil
}

// ID AFAIRE
func (r *Runner) ID() string {
	return r.id
}

// Name AFAIRE
func (r *Runner) Name() string {
	return r.name
}

// Version AFAIRE
func (r *Runner) Version() string {
	return r.version
}

// BuiltAt AFAIRE
func (r *Runner) BuiltAt() time.Time {
	return r.builtAt
}

// StartedAt AFAIRE
func (r *Runner) StartedAt() time.Time {
	return r.startedAt
}

// FQDN AFAIRE
func (r *Runner) FQDN() string {
	return r.fqdn
}

// SdInstance AFAIRE
func (r *Runner) SdInstance() string {
	return r.sdInstance
}

// SetSdInstance AFAIRE
func (r *Runner) SetSdInstance(instance string) {
	r.sdInstance = instance
}

// CfgValue AFAIRE
func (r *Runner) CfgValue() *value.Value {
	return r.cfgValue
}

// SetCfgValue AFAIRE
func (r *Runner) SetCfgValue(cfg *value.Value) {
	r.cfgValue = cfg
}

// AddGroupFn AFAIRE
func (r *Runner) AddGroupFn(fn func(<-chan struct{}) error) {
	r.group.Add(fn)
}

// RunGroup AFAIRE
func (r *Runner) RunGroup() error {
	return r.group.Run()
}

// Broker AFAIRE
func (r *Runner) Broker() component.Broker {
	return r.broker
}

// SetBroker AFAIRE
func (r *Runner) SetBroker(broker component.Broker) {
	r.broker = broker
}

// Config AFAIRE
func (r *Runner) Config() component.Config {
	return r.config
}

// SetConfig AFAIRE
func (r *Runner) SetConfig(config component.Config) {
	r.config = config
}

// Logger AFAIRE
func (r *Runner) Logger() component.Logger {
	return r.logger
}

// SetLogger AFAIRE
func (r *Runner) SetLogger(logger component.Logger) {
	r.logger = logger
}

// Registry AFAIRE
func (r *Runner) Registry() component.Registry {
	return r.registry
}

// SetRegistry AFAIRE
func (r *Runner) SetRegistry(registry component.Registry) {
	r.registry = registry
}

// Server AFAIRE
func (r *Runner) Server() component.Server {
	return r.server
}

// SetServer AFAIRE
func (r *Runner) SetServer(server component.Server) {
	r.server = server
}

// ServiceCfg AFAIRE
func (r *Runner) ServiceCfg(name string) (*value.Value, error) {
	_, v, err := r.cfgValue.Get("services", name)
	return v, err
}

/*
######################################################################################################## @(°_°)@ #######
*/
