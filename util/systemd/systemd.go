/*
#######
##        ____    _____ ____ _
##       (_-< |/|/ / _ `/ _ `/
##      /___/__,__/\_,_/\_, /
##                     /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package systemd

import (
	"time"

	"github.com/coreos/go-systemd/v22/daemon"
)

// WatchdogDelay AFAIRE
func WatchdogDelay() (time.Duration, error) {
	td, err := daemon.SdWatchdogEnabled(false)
	if err != nil || td == 0 {
		return td, err
	}

	return (td / 2) - 250*time.Millisecond, nil
}

func notify(state string) {
	_, _ = daemon.SdNotify(false, state)
}

// NotifyReady AFAIRE
func NotifyReady() {
	notify(daemon.SdNotifyReady)
}

// NotifyStopping AFAIRE
func NotifyStopping() {
	notify(daemon.SdNotifyStopping)
}

// NotifyWatchdog AFAIRE
func NotifyWatchdog() {
	notify(daemon.SdNotifyWatchdog)
}

/*
######################################################################################################## @(°_°)@ #######
*/
