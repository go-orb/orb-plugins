package mdns

import (
	"testing"
	"time"

	_ "github.com/go-micro/plugins/log/text"
	"go-micro.dev/v5/log"
	"go-micro.dev/v5/registry"
	"go-micro.dev/v5/types"
)

func TestMDNS(t *testing.T) {
	testData := []*registry.Service{
		{
			Name:    "test1",
			Version: "1.0.1",
			Nodes: []*registry.Node{
				{
					ID:      "test1-1",
					Address: "10.0.0.1:10001",
					Metadata: map[string]string{
						"foo": "bar",
					},
				},
			},
		},
		{
			Name:    "test2",
			Version: "1.0.2",
			Nodes: []*registry.Node{
				{
					ID:      "test2-1",
					Address: "10.0.0.2:10002",
					Metadata: map[string]string{
						"foo2": "bar2",
					},
				},
			},
		},
		{
			Name:    "test3",
			Version: "1.0.3",
			Nodes: []*registry.Node{
				{
					ID:      "test3-1",
					Address: "10.0.0.3:10003",
					Metadata: map[string]string{
						"foo3": "bar3",
					},
				},
			},
		},
		{
			Name:    "test4",
			Version: "1.0.4",
			Nodes: []*registry.Node{
				{
					ID:      "test4-1",
					Address: "[::]:10004",
					Metadata: map[string]string{
						"foo4": "bar4",
					},
				},
			},
		},
	}

	// new registry
	l, err := log.New(log.NewConfig())
	if err != nil {
		t.Fatal(err)
	}

	cfg, err := NewConfig(types.ServiceName("test.service"), nil)
	if err != nil {
		t.Fatal(err)
	}

	r := New(cfg, l)
	if err = r.Start(); err != nil {
		t.Fatal(err)
	}

	for _, service := range testData {
		// register service
		if err = r.Register(service); err != nil {
			t.Fatal(err)
		}

		// get registered service
		var s []*registry.Service
		s, err = r.GetService(service.Name)
		if err != nil {
			t.Fatal(err)
		}

		if len(s) != 1 {
			t.Fatalf("Expected one result for %s got %d", service.Name, len(s))
		}

		if s[0].Name != service.Name {
			t.Fatalf("Expected name %s got %s", service.Name, s[0].Name)
		}

		if s[0].Version != service.Version {
			t.Fatalf("Expected version %s got %s", service.Version, s[0].Version)
		}

		if len(s[0].Nodes) != 1 {
			t.Fatalf("Expected 1 node, got %d", len(s[0].Nodes))
		}

		node := s[0].Nodes[0]

		if node.ID != service.Nodes[0].ID {
			t.Fatalf("Expected node id %s got %s", service.Nodes[0].ID, node.ID)
		}

		if node.Address != service.Nodes[0].Address {
			t.Fatalf("Expected node address %s got %s", service.Nodes[0].Address, node.Address)
		}
	}

	services, err := r.ListServices()
	if err != nil {
		t.Fatal(err)
	}

	for _, service := range testData {
		var seen bool
		for _, s := range services {
			if s.Name == service.Name {
				seen = true
				break
			}
		}
		if !seen {
			t.Fatalf("Expected service %s got nothing", service.Name)
		}

		// deregister
		if err := r.Deregister(service); err != nil {
			t.Fatal(err)
		}

		time.Sleep(time.Millisecond * 5)

		// check its gone
		s, _ := r.GetService(service.Name) //nolint:errcheck
		if len(s) > 0 {
			t.Fatalf("Expected nothing got %+v", s[0])
		}
	}
}

func TestEncoding(t *testing.T) {
	testData := []*mdnsTxt{
		{
			Version: "1.0.0",
			Metadata: map[string]string{
				"foo": "bar",
			},
			Endpoints: []*registry.Endpoint{
				{
					Name: "endpoint1",
					Request: &registry.Value{
						Name: "request",
						Type: "request",
					},
					Response: &registry.Value{
						Name: "response",
						Type: "response",
					},
					Metadata: map[string]string{
						"foo1": "bar1",
					},
				},
			},
		},
	}

	for _, d := range testData {
		encoded, err := encode(d)
		if err != nil {
			t.Fatal(err)
		}

		for _, txt := range encoded {
			if len(txt) > 255 {
				t.Fatalf("One of parts for txt is %d characters", len(txt))
			}
		}

		decoded, err := decode(encoded)
		if err != nil {
			t.Fatal(err)
		}

		if decoded.Version != d.Version {
			t.Fatalf("Expected version %s got %s", d.Version, decoded.Version)
		}

		if len(decoded.Endpoints) != len(d.Endpoints) {
			t.Fatalf("Expected %d endpoints, got %d", len(d.Endpoints), len(decoded.Endpoints))
		}

		for k, v := range d.Metadata {
			if val := decoded.Metadata[k]; val != v {
				t.Fatalf("Expected %s=%s got %s=%s", k, v, k, val)
			}
		}
	}
}

func TestWatcher(t *testing.T) {
	testData := []*registry.Service{
		{
			Name:    "test1",
			Version: "1.0.1",
			Nodes: []*registry.Node{
				{
					ID:      "test1-1",
					Address: "10.0.0.1:10001",
					Metadata: map[string]string{
						"foo": "bar",
					},
				},
			},
		},
		{
			Name:    "test2",
			Version: "1.0.2",
			Nodes: []*registry.Node{
				{
					ID:      "test2-1",
					Address: "10.0.0.2:10002",
					Metadata: map[string]string{
						"foo2": "bar2",
					},
				},
			},
		},
		{
			Name:    "test3",
			Version: "1.0.3",
			Nodes: []*registry.Node{
				{
					ID:      "test3-1",
					Address: "10.0.0.3:10003",
					Metadata: map[string]string{
						"foo3": "bar3",
					},
				},
			},
		},
		{
			Name:    "test4",
			Version: "1.0.4",
			Nodes: []*registry.Node{
				{
					ID:      "test4-1",
					Address: "[::]:10004",
					Metadata: map[string]string{
						"foo4": "bar4",
					},
				},
			},
		},
	}

	testFn := func(service, s *registry.Service) {
		if s == nil {
			t.Fatalf("Expected one result for %s got nil", service.Name)
		}

		if s.Name != service.Name {
			t.Fatalf("Expected name %s got %s", service.Name, s.Name)
		}

		if s.Version != service.Version {
			t.Fatalf("Expected version %s got %s", service.Version, s.Version)
		}

		if len(s.Nodes) != 1 {
			t.Fatalf("Expected 1 node, got %d", len(s.Nodes))
		}

		node := s.Nodes[0]

		if node.ID != service.Nodes[0].ID {
			t.Fatalf("Expected node id %s got %s", service.Nodes[0].ID, node.ID)
		}

		if node.Address != service.Nodes[0].Address {
			t.Fatalf("Expected node address %s got %s", service.Nodes[0].Address, node.Address)
		}
	}

	// New registry
	l, err := log.New(log.NewConfig())
	if err != nil {
		t.Fatal(err)
	}

	cfg, err := NewConfig(types.ServiceName("test.service"), nil)
	if err != nil {
		t.Fatal(err)
	}

	r := New(cfg, l)
	if err = r.Start(); err != nil {
		t.Fatal(err)
	}

	w, err := r.Watch()
	if err != nil {
		t.Fatal(err)
	}
	defer w.Stop()

	for _, service := range testData {
		// register service
		if err := r.Register(service); err != nil {
			t.Fatal(err)
		}

		for {
			res, err := w.Next()
			if err != nil {
				t.Fatal(err)
			}

			if res.Service.Name != service.Name {
				continue
			}

			if res.Action != "create" {
				t.Fatalf("Expected create event got %s for %s", res.Action, res.Service.Name)
			}

			testFn(service, res.Service)
			break
		}

		// deregister
		if err := r.Deregister(service); err != nil {
			t.Fatal(err)
		}

		for {
			res, err := w.Next()
			if err != nil {
				t.Fatal(err)
			}

			if res.Service.Name != service.Name {
				continue
			}

			if res.Action != "delete" {
				continue
			}

			testFn(service, res.Service)
			break
		}
	}
}
