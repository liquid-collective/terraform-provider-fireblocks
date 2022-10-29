package sdk

import (
	"context"

	kilnapp "github.com/kilnfi/go-utils/app"
	"github.com/liquid-collective/terraform-provider-fireblocks/pkg/ethereum"
	fireblocksclient "github.com/liquid-collective/terraform-provider-fireblocks/pkg/fireblocks/client"
	fireblocksclienthttp "github.com/liquid-collective/terraform-provider-fireblocks/pkg/fireblocks/client/http"
)

type SDK struct {
	fireblocksclient.Client
	cfg *Config
	abi ethereum.ABI
}

func New(cfg *Config) (*SDK, error) {
	client, err := fireblocksclienthttp.New(cfg.Fireblocks)
	if err != nil {
		return nil, err
	}

	return &SDK{
		cfg:    cfg,
		Client: client,
	}, nil
}

func (sdk *SDK) Init(ctx context.Context) error {
	if sdk.cfg.ABIPath != "" {
		abi, err := ethereum.LoadABI(sdk.cfg.ABIPath)
		if err != nil {
			return err
		}
		sdk.abi = abi
	}

	if init, ok := sdk.Client.(kilnapp.Initializable); ok {
		return init.Init(ctx)
	}

	return nil
}
