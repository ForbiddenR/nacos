package register

import (
	"errors"
	"github.com/ForbiddenR/nacos/options"
	"github.com/go-playground/validator/v10"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"strconv"
)

// NamingClient will keep this server alive in nacos until the server is terminated.
var NamingClient naming_client.INamingClient

// RegisterToNacos registers this server to nacos.
func RegisterToNacos(validate *validator.Validate, port, ip, name string) error {
	option, err := options.InitConfig(validate, nil)
	if err != nil {
		return err
	}
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(
			option.Url, option.Port),
	}

	cc := *constant.NewClientConfig(
		constant.WithUsername(option.Username),
		constant.WithPassword(option.Password),
		constant.WithNamespaceId(option.NamespaceId),
		constant.WithNotLoadCacheAtStart(true),
	)

	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		})
	if err != nil {
		panic("failed getting naming client of nacos " + err.Error())
	}

	ports, err := strconv.Atoi(port)
	if err != nil {
		return err
	}
	success, err := client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          ip,
		Port:        uint64(ports),
		Healthy:     true,
		ServiceName: name,
		GroupName:   option.GroupName,
		Ephemeral:   true,
		Enable:      true,
	})

	if !success {
		return errors.New("failed registering server to nacos")
	} else if err != nil {
		return err
	}

	NamingClient = client
	return nil
}

// Close closes the client.
func Close() {
	NamingClient.CloseClient()
}
