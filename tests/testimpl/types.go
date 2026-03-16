package testimpl

import "github.com/launchbynttdata/lcaf-component-terratest/types"

type ThisTFModuleConfig struct {
	types.GenericTFModuleConfig
	// CloudWatch Event Bus module has no additional test config beyond GenericTFModuleConfig.
}
