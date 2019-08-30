package main

import (
	flag "github.com/jessevdk/go-flags"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/klog"
	"os"
	"strconv"
	"time"

	"github.com/kubeedge/kubeedge/cloud/pkg/apis/devices/v1alpha1"

	"github.com/subpathdev/kubeedge-database/database"
	"github.com/subpathdev/kubeedge-database/kubernetes"
)

var opts struct {
	Server     string `short:"s" long:"server" required:"no" description:"address of the kubernetes api server"`
	ConfigPath string `short:"c" long:"config" required:"no" description:"path of the kubernetes config"`
	Database   string `short:"d" long:"database" required:"yes" description:"name of the database"`
	Address    string `short:"a" long:"address" required:"yes" description:"address of the db server"`
	Port       int    `short:"p" long:"port" required:"no" default:"5432" description:"port of the database connection"`
	SslMode    string `short:"m" long:"sslmode" required:"no" default:"required" description:"set the ssl mode which should be used to connect to the database"`
	Schema     string `short:"S" long:"schema" default:"public" required:"no" description:"is the used database schema to store the device data"`
	Timescale  bool   `short:"t" long:"timescale" required:"no" description:"you can enable this option, when we should use the timescale plugin"`
}

var env struct {
	User     string
	Password string
}

func main() {
	klog.InitFlags(nil)

	_, err := flag.Parse(&opts)
	if err != nil {
		pr := flag.WroteHelp(err)
		args := []string{
			"-h",
		}
		if !pr {
			_, err := flag.ParseArgs(&opts, args)
			if err != nil {
				klog.Error(err)
				os.Exit(2)
			}
			os.Exit(1)
		} else {
			os.Exit(0)
		}
	}

	env.User = os.Getenv("DATABASE-USER")
	env.Password = os.Getenv("DATABASE-PASSWORD")

	events := make(chan watch.Event, 10)

	db, err := database.NewDatabase(opts.Database, env.User, opts.Address, env.Password, opts.SslMode, opts.Schema, opts.Port, opts.Timescale)
	if err != nil {
		klog.Errorf("could not open database connection; err: %v\n", err)
		os.Exit(1)
	}

	if err := kubernetes.Init(opts.Server, opts.ConfigPath, events); err != nil {
		klog.Errorf("could not start api connection successfully; err: %v\n", err)

	}
	handleKube(events, db)
}

func handleKube(event chan watch.Event, db database.Database) {
	for {
		ev := <-event
		dev, ok := ev.Object.(*v1alpha1.Device)
		if !ok {
			klog.Errorf("could not convert event data to KubeEdge v1alpha1.Device")
			continue
		}
		switch ev.Type {
		case watch.Deleted:
			klog.Infof("handle delete event")
			if err := db.Delete(dev.Name, dev.Namespace); err != nil {
				klog.Errorf("could not marked as non active err: %v\n", err)
				continue
			}
		case watch.Added:
			klog.Infof("handle added event")
			for _, twin := range dev.Status.Twins {
				var tt time.Time
				if val, ok := twin.Reported.Metadata["timestamp"]; ok {
					tim, err := strconv.ParseInt(val, 10, 64)
					if err != nil {
						tt = time.Now()
					} else {
						tt = time.Unix(tim/1000, (tim%1000)*1000)
					}
				} else {
					tt = time.Now()
				}
				err := db.Insert(tt, twin.Reported.Value, dev.Name, dev.Namespace, twin.PropertyName, true)
				if err != nil {
					klog.Errorf("could not insert data; err: %v\n")
				}
			}
		case watch.Modified:
			klog.Infof("handle modified event")
			for _, twin := range dev.Status.Twins {
				var tt time.Time
				if val, ok := twin.Reported.Metadata["timestamp"]; ok {
					tim, err := strconv.ParseInt(val, 10, 64)
					if err != nil {
						tt = time.Now()
					} else {
						tt = time.Unix(tim/1000, (tim%1000)*1000)
					}
				} else {
					tt = time.Now()
				}
				err := db.Insert(tt, twin.Reported.Value, dev.Name, dev.Namespace, twin.PropertyName, true)
				if err != nil {
					klog.Errorf("could not insert data; err: %v\n")
				}
			}
		}
	}
}
