package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListClusterServiceConfigRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	TagValue        string `position:"Query" name:"TagValue"`
	GroupId         int64  `position:"Query" name:"GroupId"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ConfigVersion   string `position:"Query" name:"ConfigVersion"`
}

func (req *ListClusterServiceConfigRequest) Invoke(client *sdk.Client) (resp *ListClusterServiceConfigResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListClusterServiceConfig", "", "")
	resp = &ListClusterServiceConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListClusterServiceConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
	Configs   ListClusterServiceConfigConfigList
}

type ListClusterServiceConfigConfig struct {
	ServiceName      common.String
	ConfigVersion    common.String
	Applied          common.String
	CreateTime       common.String
	Author           common.String
	Comment          common.String
	ConfigValueList  ListClusterServiceConfigConfigValueList
	PropertyInfoList ListClusterServiceConfigPropertyInfoList
}

type ListClusterServiceConfigConfigValue struct {
	ConfigName          common.String
	AllowCustom         bool
	ConfigItemValueList ListClusterServiceConfigConfigItemValueList
}

type ListClusterServiceConfigConfigItemValue struct {
	ItemName    common.String
	Value       common.String
	IsCustom    bool
	Description common.String
}

type ListClusterServiceConfigPropertyInfo struct {
	Name                    common.String
	Value                   common.String
	Description             common.String
	FileName                common.String
	DisplayName             common.String
	ServiceName             common.String
	Component               common.String
	PropertyTypes           ListClusterServiceConfigPropertyTypeList
	PropertyValueAttributes ListClusterServiceConfigPropertyValueAttributes
	EffectWay               ListClusterServiceConfigEffectWay
}

type ListClusterServiceConfigPropertyValueAttributes struct {
	Type          common.String
	Maximum       common.String
	Mimimum       common.String
	Unit          common.String
	ReadOnly      bool
	Hidden        bool
	IncrememtStep common.String
	Entries       ListClusterServiceConfigValueEntryInfoList
}

type ListClusterServiceConfigValueEntryInfo struct {
	Value       common.String
	Label       common.String
	Description common.String
}

type ListClusterServiceConfigEffectWay struct {
	EffectType        common.String
	InvokeServiceName common.String
}

type ListClusterServiceConfigConfigList []ListClusterServiceConfigConfig

func (list *ListClusterServiceConfigConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterServiceConfigConfig)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type ListClusterServiceConfigConfigValueList []ListClusterServiceConfigConfigValue

func (list *ListClusterServiceConfigConfigValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterServiceConfigConfigValue)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type ListClusterServiceConfigPropertyInfoList []ListClusterServiceConfigPropertyInfo

func (list *ListClusterServiceConfigPropertyInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterServiceConfigPropertyInfo)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type ListClusterServiceConfigConfigItemValueList []ListClusterServiceConfigConfigItemValue

func (list *ListClusterServiceConfigConfigItemValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterServiceConfigConfigItemValue)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type ListClusterServiceConfigPropertyTypeList []common.String

func (list *ListClusterServiceConfigPropertyTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type ListClusterServiceConfigValueEntryInfoList []ListClusterServiceConfigValueEntryInfo

func (list *ListClusterServiceConfigValueEntryInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterServiceConfigValueEntryInfo)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
