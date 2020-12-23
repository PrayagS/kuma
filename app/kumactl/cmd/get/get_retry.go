package get

import (
	"context"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/kumahq/kuma/app/kumactl/pkg/output"
	"github.com/kumahq/kuma/app/kumactl/pkg/output/printers"
	"github.com/kumahq/kuma/pkg/core/resources/apis/mesh"
	rest_types "github.com/kumahq/kuma/pkg/core/resources/model/rest"
	"github.com/kumahq/kuma/pkg/core/resources/store"
)

func newGetRetryCmd(pctx *getContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "retry NAME",
		Short: "Show a single retry resource",
		Long:  `Show a single retry resource.`,
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			rs, err := pctx.CurrentResourceStore()
			if err != nil {
				return err
			}
			name := args[0]
			currentMesh := pctx.CurrentMesh()
			retry := mesh.NewRetryResource()
			if err := rs.Get(
				context.Background(),
				retry,
				store.GetByKey(name, currentMesh),
			); err != nil {
				if store.IsResourceNotFound(err) {
					return errors.Errorf("No resources found in %s mesh", currentMesh)
				}
				return errors.Wrapf(err, "failed to get a retry in mesh %s", currentMesh)
			}
			retries := &mesh.RetryResourceList{
				Items: []*mesh.RetryResource{retry},
			}
			switch format := output.Format(pctx.args.outputFormat); format {
			case output.TableFormat:
				return printRetries(pctx.Now(), retries, cmd.OutOrStdout())
			default:
				printer, err := printers.NewGenericPrinter(format)
				if err != nil {
					return err
				}
				return printer.Print(
					rest_types.From.Resource(retry),
					cmd.OutOrStdout(),
				)
			}
		},
	}
	return cmd
}