package app

import (
	"context"
	"flag"

	"github.com/spf13/cobra"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/klog"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"

	"github.com/jwcesign/kloud/cmd/analyzer/app/options"
	"github.com/jwcesign/kloud/pkg/apiserver/router"
	"github.com/jwcesign/kloud/pkg/controller"
	"github.com/jwcesign/kloud/pkg/version"
)

func NewAnalyzer(ctx context.Context) *cobra.Command {
	opts := options.NewOptions()

	cmd := &cobra.Command{
		Use:  "kloud-analyzer",
		Long: "kloud-analyzer is used to estimate the Kubernetes cost when migrating to a specific public cloud provider",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.ApplyAndValidate(); err != nil {
				return err
			}
			return run(ctx, opts)
		},
	}

	fss := cliflag.NamedFlagSets{}
	logFlagSet := fss.FlagSet("log")
	klog.InitFlags(flag.CommandLine)
	logFlagSet.AddGoFlagSet(flag.CommandLine)
	cmd.Flags().AddFlagSet(logFlagSet)

	return cmd
}

func run(ctx context.Context, opts *options.Options) error {
	klog.Infof("Start kloud-analyzer, version: %s, commit: %s...", version.Get().GitVersion, version.Get().GitCommit)

	mgr, err := manager.New(config.GetConfigOrDie(),
		manager.Options{Metrics: metricsserver.Options{BindAddress: "0"}})
	if err != nil {
		klog.Fatalf("Failed to create controller manager: %v", err)
	}

	if err := controller.SetupController(mgr); err != nil {
		klog.Fatalf("Failed to setup controller: %v", err)
	}
	go func() {
		if err := mgr.Start(ctx); err != nil {
			klog.Fatalf("Failed to start controller manager: %v", err)
		}
	}()

	serverRouter := router.NewAnalyzerAPIServer()
	go func() {
		if err := serverRouter.Run(":8080"); err != nil {
			klog.Fatalf("Failed to start server: %v", err)
		}
	}()

	<-ctx.Done()
	return nil
}
