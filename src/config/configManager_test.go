package config

import (
	"testing"

	types "github.com/accuknox/knoxAutoPolicy/src/types"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfigDB(t *testing.T) {
	cfg := LoadConfigDB()

	assert.NotEmpty(t, cfg.DBDriver, "db driver should not be empty")
	assert.NotEmpty(t, cfg.DBUser, "db user should not be empty")
	assert.NotEmpty(t, cfg.DBPass, "db pass should not be empty")
	assert.NotEmpty(t, cfg.DBName, "db name should not be empty")
	assert.NotEmpty(t, cfg.DBHost, "db host should not be empty")
	assert.NotEmpty(t, cfg.DBPort, "db host should not be empty")

	assert.NotEmpty(t, cfg.TableNetworkLog, "table networkf_flow should not be empty")
	assert.NotEmpty(t, cfg.TableNetworkPolicy, "table discovered_policy should not be empty")
	assert.NotEmpty(t, cfg.TableConfiguration, "table auto_policy_config should not be empty")
	assert.NotEmpty(t, cfg.TableSystemLog, "table auto_policy_config should not be empty")
}

func TestLoadDefaultConfig(t *testing.T) {
	LoadDefaultConfig()

	assert.NotEmpty(t, Cfg.OperationMode, "operation mode should not be empty")

	assert.NotEmpty(t, Cfg.NetworkLogFrom, "network log from should not be empty")
	assert.NotEmpty(t, Cfg.NetworkLogFile, "network log file should not be empty")

	assert.NotEmpty(t, Cfg.NetworkPolicyTo, "discovery policy to should not be empty")

	assert.NotEmpty(t, Cfg.NetPolicyTypes, "discovery policy types should not be empty")
	assert.NotEmpty(t, Cfg.NetPolicyRuleTypes, "discovery rule types should not be empty")
	assert.NotEmpty(t, Cfg.NetPolicyCIDRBits, "cidr bits should not be empty")

	assert.NotEmpty(t, Cfg.NetPolicyL3Level, "L3 aggregation level should not be empty")
	assert.NotEmpty(t, Cfg.NetPolicyL4Level, "L4 compression should not be empty")
	assert.NotEmpty(t, Cfg.NetPolicyL7Level, "L7 aggregation level should not be empty")
}

func TestSetLogFile(t *testing.T) {
	SetLogFile("test_log.log")

	assert.Equal(t, Cfg.NetworkLogFile, "test_log.log", "network log file should be \"test_log.log\"")
}

func TestManageConfiguration(t *testing.T) {
	newCfg := types.Configuration{}
	newCfg.ConfigName = "test_config"
	newCfg.NetPolicyCIDRBits = 32

	// add configuration
	err := AddConfiguration(newCfg)
	assert.NoError(t, err)

	// get configuration
	results, err := GetConfigurations(newCfg.ConfigName)
	assert.NoError(t, err)
	assert.Equal(t, results[0].ConfigName, "test_config")

	// apply configuration
	err = ApplyConfiguration(newCfg.ConfigName)
	assert.NoError(t, err)

	// update configuration
	upCfg := types.Configuration{}
	upCfg.ConfigName = "test_config"
	upCfg.NetPolicyCIDRBits = 24
	err = UpdateConfiguration("test_config", upCfg)

	results, err = GetConfigurations(newCfg.ConfigName)
	assert.NoError(t, err)
	assert.Equal(t, results[0].NetPolicyCIDRBits, upCfg.NetPolicyCIDRBits)

	// delete configuration
	err = DeleteConfiguration(newCfg.ConfigName)
	assert.NoError(t, err)
}