package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.uber.org/multierr"

	"currency/internal/app/config"
	"currency/internal/app/delivery/rest"
	"currency/internal/app/repository"
	"currency/internal/app/service"
	"currency/pkg/process"
)

type Listener interface {
	Listen(addr string) error
}

const configPath = "configs/main.yaml"

func app() (err error){
	c := make(chan os.Signal)
	out := make(chan bool)

	signal.Notify(c,  syscall.SIGINT)

	go func() {
		for {
			select {
			case _, ok := <-c:
				if ok {
					out <- true
					close(c)

					return
				}
			}
		}
	}()

	logrus.Info("initialize config")
	conf, err := config.New(configPath)
	if err != nil {
		return err
	}

	logrus.Info("initialize database connection")
	db, err := sqlx.Open("postgres", conf.DB.Conn)
	if err != nil {
		return errors.Wrap(err, "connect to database")
	}
	defer func(err error) {
		multierr.AppendInto(&err, db.Close())
	}(err)

	logrus.Info("initialize repository")
	storage := repository.New(db)

	logrus.Info("initialize service usecases")
	usecase := service.New(storage, conf.API)

	logrus.Info("initialize process usecases")
	proc := process.New(time.Duration(conf.Process.Ticker) * time.Second)

	logrus.Info("add usecase process")
	proc.Add(usecase.Pair.UpdateAll)

	logrus.Info("initialize service rest core")
	api := rest.New(usecase)

	logrus.Info("initialize application context")
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		proc.Wait(ctx)
	}()

	go func() {
		logrus.Infof("listen and serving http on port %s", conf.App.Port)
		lerr := listenHttp(api, ":"+conf.App.Port)
		if lerr != nil {
			multierr.AppendInto(&err, lerr)
			return
		}
	}()

	<-out
	logrus.Info("shutdown service")
	cancel()
	close(out)
	time.Sleep(time.Second * 2)
	return nil
}

func listenHttp(l Listener, port string) error {
	return l.Listen(port)
}