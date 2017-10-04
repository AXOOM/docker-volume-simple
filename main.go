package main

import (
	"os"
	"path/filepath"
    "io/ioutil"
	"github.com/docker/go-plugins-helpers/volume"
)

func main() {
	volume.NewHandler(&simpleDriver{
		root: os.Getenv("ROOT"),
		scope: os.Getenv("SCOPE"),
	}).ServeUnix("/run/docker/plugins/plugin.sock", 0)
}

type simpleDriver struct {
	root      string
	scope     string
}

func (d *simpleDriver) Create(r *volume.CreateRequest) error {
	return os.MkdirAll(filepath.Join(d.root, r.Name), os.ModePerm)
}

func (d *simpleDriver) Remove(r *volume.RemoveRequest) error {
	return os.RemoveAll(filepath.Join(d.root, r.Name))
}

func (d *simpleDriver) Path(r *volume.PathRequest) (*volume.PathResponse, error) {
	return &volume.PathResponse{
		Mountpoint: filepath.Join(d.root, r.Name),
	}, nil
}

func (d *simpleDriver) Mount(r *volume.MountRequest) (*volume.MountResponse, error) {
	return &volume.MountResponse{
		Mountpoint: filepath.Join(d.root, r.Name),
	}, nil
}

func (d *simpleDriver) Unmount(r *volume.UnmountRequest) error {
	return nil
}

func (d *simpleDriver) Get(r *volume.GetRequest) (*volume.GetResponse, error) {
	path := filepath.Join(d.root, r.Name)
	stat, error := os.Stat(path);

	if error == nil && stat.IsDir() {
		return &volume.GetResponse{
			Volume: &volume.Volume{
				Name: r.Name,
				Mountpoint: path,
			},
		}, nil
	} else {
		return nil, error
	}
}

func (d *simpleDriver) List() (*volume.ListResponse, error) {
	list, err := ioutil.ReadDir(d.root)
	if err != nil {
		return nil, err
	}

	var vols []*volume.Volume
	for _, e := range list {
		if (e.IsDir()) {
			vols = append(vols, &volume.Volume{
				Name: e.Name(),
				Mountpoint: filepath.Join(d.root, e.Name()),
			})
		}
	}

	return &volume.ListResponse{
		Volumes: vols,
	}, nil
}

func (d *simpleDriver) Capabilities() *volume.CapabilitiesResponse {
	return &volume.CapabilitiesResponse{
		Capabilities: volume.Capability{
			Scope: d.scope,
		},
	}
}
