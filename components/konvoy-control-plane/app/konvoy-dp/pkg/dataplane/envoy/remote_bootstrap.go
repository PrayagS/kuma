package envoy

import (
	"bytes"
	"encoding/json"
	"fmt"
	konvoy_dp "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/config/app/konvoy-dp"
	util_proto "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/util/proto"
	xds_bootstrap "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/xds/bootstrap"
	envoy_bootstrap "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v2"
	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type remoteBootstrap struct {
	client *http.Client
}

func NewRemoteBootstrapGenerator(client *http.Client) BootstrapConfigFactoryFunc {
	rb := remoteBootstrap{client: client}
	return rb.Generate
}

func (b *remoteBootstrap) Generate(cfg konvoy_dp.Config) (proto.Message, error) {
	url := fmt.Sprintf("http://%s:%d/bootstrap", cfg.ControlPlane.BootstrapServer.Address, cfg.ControlPlane.BootstrapServer.Port)
	request := xds_bootstrap.BootstrapRequest{
		NodeId: cfg.Dataplane.Id,
	}
	jsonBytes, err := json.Marshal(request)
	if err != nil {
		return nil, errors.Wrap(err, "could not marshal request to json")
	}
	resp, err := b.client.Post(url, "application/json", bytes.NewReader(jsonBytes))
	if err != nil {
		return nil, errors.Wrap(err, "request to bootstrap server failed")
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 404 {
			return nil, errors.New("status: 404. Did you first applied Dataplane resource?")
		}
		return nil, errors.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	bootstrap := envoy_bootstrap.Bootstrap{}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "could not read the body of the response")
	}
	if err := util_proto.FromYAML(respBytes, &bootstrap); err != nil {
		return nil, errors.Wrap(err, "could not parse the bootstrap configuration")
	}

	return &bootstrap, nil
}
