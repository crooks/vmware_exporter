package main

import (
	"context"
	"fmt"
	"log"
	"github.com/crooks/vmware_exporter/config"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/view"
)

func main() {
	flags := config.ParseFlags()
	cfg, err := config.ParseConfig(flags.ConfigFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg.API.UserID)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	m := view.NewManager(ctx)
	v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, []string{"HostSystem"}, true)
    if err != nil {
        return err
    }
    defer v.Destroy(ctx)

	var hss []mo.HostSystem
	err = v.Retrieve(ctx, []string{"HostSystem"}, []string{"summary"}, &hss)
	if err != nil {
		return err
	}
	for _, hs := range hss {
		fmt.Println(hs)
	}
}