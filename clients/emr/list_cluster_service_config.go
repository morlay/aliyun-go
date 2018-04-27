package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId string
	Configs   ListClusterServiceConfigConfigList
}

type ListClusterServiceConfigConfig struct {
	ServiceName      string
	ConfigVersion    string
	Applied          string
	CreateTime       string
	Author           string
	Comment          string
	ConfigValueList  ListClusterServiceConfigConfigValueList
	PropertyInfoList ListClusterServiceConfigPropertyInfoList
}

type ListClusterServiceConfigConfigValue struct {
	ConfigName          string
	AllowCustom         bool
	ConfigItemValueList ListClusterServiceConfigConfigItemValueList
}

type ListClusterServiceConfigConfigItemValue struct {
	ItemName    string
	Value       string
	IsCustom    bool
	Description string
}

type ListClusterServiceConfigPropertyInfo struct {
	Name                    string
	Value                   string
	Description             string
	FileName                string
	DisplayName             string
	ServiceName             string
	Component               string
	PropertyTypes           ListClusterServiceConfigPropertyTypeList
	PropertyValueAttributes ListClusterServiceConfigPropertyValueAttributes
	EffectWay               ListClusterServiceConfigEffectWay
}

type ListClusterServiceConfigPropertyValueAttributes struct {
	Type          string
	Maximum       string
	Mimimum       string
	Unit          string
	ReadOnly      bool
	Hidden        bool
	IncrememtStep string
	Entries       ListClusterServiceConfigValueEntryInfoList
}

type ListClusterServiceConfigValueEntryInfo struct {
	Value       string
	Label       string
	Description string
}

type ListClusterServiceConfigEffectWay struct {
	EffectType        string
	InvokeServiceName string
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

type ListClusterServiceConfigPropertyTypeList []string

func (list *ListClusterServiceConfigPropertyTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
