package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"github.com/crooks/vmware_exporter/config"
	"github.com/vmware/govmomi/session/cache"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/soap"
)

var (
	cfg *config.Config
)

func newClient(ctx context.Context) (*vim25.Client, error) {
	u, err := soap.ParseURL(cfg.API.URL)
	u.User = url.UserPassword(cfg.API.UserID, cfg.API.Password)
	s := &cache.Session{
		URL: u,
		Insecure: cfg.API.Insecure,
	}

	c := new(vim25.Client)
	err = s.Login(ctx, c, nil)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func main() {
	var err error
	flags := config.ParseFlags()
	cfg, err = config.ParseConfig(flags.ConfigFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg.API.UserID)

	ctx := context.Background()
    c, err := newClient(ctx)
    if err != nil {
        log.Fatal(err)
    }

	m := view.NewManager(c)
	v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, []string{"HostSystem"}, true)
    if err != nil {
        log.Fatal(err)
    }
    defer v.Destroy(ctx)

	var hss []mo.HostSystem
	err = v.Retrieve(ctx, []string{"HostSystem"}, []string{"summary"}, &hss)
	if err != nil {
		log.Fatal(err)
	}
	for _, hs := range hss {
		fmt.Println(hs)
	}
}
