package app

import (
	"context"
	"flag"
	"k8s.io/client-go/rest"

	"github.com/spf13/cobra"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/klog/v2"

	"github.com/jwcesign/kcloud/cmd/apiserver/app/options"
	"github.com/jwcesign/kcloud/pkg/version"
)

func NewAPIServerCommand(ctx context.Context) *cobra.Command {
	opts := options.NewOptions()

	cmd := &cobra.Command{
		Use:  "kcloud-apiserver",
		Long: "kcloud-apiserver used for serve as the apiserver",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.ApplyAndValidate(); err != nil {
				return err
			}
			if err := run(ctx, opts); err != nil {
				return err
			}
		},
	}

	fss := cliflag.NamedFlagSets{}
	logFlagSet := fss.FlagSet("log")
	klog.InitFlags(flag.CommandLine)
	logFlagSet.AddGoFlagSet(flag.CommandLine)
	cmd.Flags().AddFlagSet(logFlagSet)

	return cmd
}

var (
	scheme = runtime.NewScheme()
)

func init() {
	_ = v1.AddToScheme(scheme)
}

func run(ctx context.Context, opts *options.Options) error {
	currentVersion := version.Get()
	klog.Infof("Start cloudpilot-apiserver, version: %s, commit: %s", currentVersion.GitVersion, currentVersion.GitCommit)

	cfg, err := rest.InClusterConfig()
	if err != nil {
		klog.Fatalf("Failed to get in-cluster config: %v", err)
	}
	setQPS(cfg, opts)

	<-ctx.Done()
	return nil
}

func setQPS(cfg *rest.Config, opts *options.Options) {
	cfg.QPS = opts.APIQPS
	cfg.Burst = opts.APIBurst
}
