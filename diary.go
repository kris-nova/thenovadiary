package thenovadiary

import (
	"sync"
	"time"

	"github.com/kris-nova/photoprism-client-go"

	"github.com/kris-nova/logger"
)

type DiaryConfig struct {
	Name           string // Unique identifier for logs
	TwitterUser    string
	TwitterPass    string
	PhotoprismPass string
	PhotoprismUser string
	PhotoprismConn string
	validated      bool
}

type Diary struct {
	photoprismClient *photoprism.Client
	config           *DiaryConfig
	lock             sync.Mutex
	cache            *Cache
}

func New(cfg *DiaryConfig) *Diary {
	return &Diary{
		cache:  NewCache(cfg.Name),
		config: cfg,
	}
}

func (d *Diary) Service() error {
	logger.Always("Starting service...")
	run := true
	delta, err := d.cache.Recover()
	logger.Info("Delta found: %d", delta)
	if err != nil {
		logger.Info("Unable to recover cache %s, starting with empty cache: %v", d.cache.path.Name(), err)
	} else {
		logger.Info("Successful cache recovery from %s", d.cache.path.Name())
	}
	logger.Debug("Service loop...")
	for run {
		//d.lock.Lock()
		// ----------------------------------

		//logger.Debug("...")
		// TODO @kris-nova please remove this
		time.Sleep(2 * time.Second)

		// Daily Tweet Here
		err := d.SendDailyTweet()
		if err != nil {
			logger.Critical("Service Loop Error: SendDailyTweet: %v", err)
		}

		// Listen for /nova on twitter

		// Check for container updates

		// compile ZFS kernel modules

		// compile Falco kernel modules

		// ----------------------------------
		//d.lock.Unlock()
		d.cache.Persist()
	}
	return nil

}
