package configs

import (
	"fmt"
	"github.com/ForbiddenR/nacos/options"

	"github.com/spf13/viper"
	remote "github.com/yoyofxteam/nacos-viper-remote"
)

// InitViper reads all needed configs from nacos, putting them to local viper.
func InitViper(options *options.Options) (*viper.Viper, error) {
	err := remote.SetOptions(&remote.Option{
		Url:         options.Url,
		Port:        options.Port,
		NamespaceId: options.NamespaceId,
		GroupName:   options.GroupName,
		Config:      remote.Config{DataId: options.DataId},
		Auth:        &remote.Auth{Enable: true, User: options.Username, Password: options.Password},
	})
	if err != nil {
		return nil, err
	}
	remoteViper := viper.New()
	endpoint := fmt.Sprintf("%s:%d", options.Url, options.Port)
	err = remoteViper.AddRemoteProvider("nacos", endpoint, "")
	if err != nil {
		return nil, err
	}

	remoteViper.SetConfigType("properties")
	err = remoteViper.ReadRemoteConfig()
	if err != nil {
		return nil, err
	}
	return remoteViper, nil
}
