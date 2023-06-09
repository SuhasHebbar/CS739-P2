package bench

import (
    "github.com/spf13/viper"
)

type Config struct {
    // IP address of all the replicas
    Replicas []string `mapstructure:"replicas"`

    // Number of keys to generate
    NumKeys int32 `mapstructure:"num_keys"`

    // one of [random, read_recent, read_modify_update, read_range]
    // read range reads 8 entries at once
    Mode string `mapstructure:"mode"`

    // Proportion of write operations
    WriteProp float32 `mapstructure:"write_prop"`

    // Contiguous range of keys to scan
    // Applicable only in mode read_range
    RangeScanNumKeys int32 `mapstructure:"range_scan_num_keys"`

    // Length of each key
    KeyLen int32 `mapstructure:"key_len"`

    // Length of each value
    ValLen int32 `mapstructure:"val_len"`

    // Populate all the keys in the domain to avoid read failures
    PopulateAllKeys bool `mapstructure:"populate_all_keys"`

    // toggle fast get
    EnableFastGet bool `mapstructure:"enable_fast_get"`

    // simple server client
    ConnectSimpleServer bool `mapstructure:"connect_simple_server"`
}

const (
    RANDOM = "random"
    READ_RECENT = "read_recent"
    READ_MODIFY_UPDATE = "read_modify_update"
    READ_RANGE = "read_range"
)

func GetConfig(confname string) *Config {
    // set viper defaults
    viper.SetDefault("num_keys", 1024)
    viper.SetDefault("mode", "random")
    viper.SetDefault("write_prop", 0.5)
    viper.SetDefault("range_scan_num_keys", 8)
    viper.SetDefault("key_len", 8)
    viper.SetDefault("val_len", 8)
    viper.SetDefault("populate_all_keys", true)
    viper.SetDefault("enable_fast_get", false)
    viper.SetDefault("connect_simple_server", false)

    viper.SetConfigName(confname) // config filename
    viper.SetConfigType("yaml") // yaml config
    viper.AddConfigPath(".") // look for bench_config.yaml in the current working directoy

    // read config options from file
    if err := viper.ReadInConfig(); err != nil {
        panic(err)
    }

    // Unmarshal config into our struct
    config := &Config{}
    viper.UnmarshalExact(config)
    return config
}
