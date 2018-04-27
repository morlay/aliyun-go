package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterServiceConfigRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	TagValue        string `position:"Query" name:"TagValue"`
	GroupId         int64  `position:"Query" name:"GroupId"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ConfigVersion   string `position:"Query" name:"ConfigVersion"`
}

func (req *DescribeClusterServiceConfigRequest) Invoke(client *sdk.Client) (resp *DescribeClusterServiceConfigResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeClusterServiceConfig", "", "")
	resp = &DescribeClusterServiceConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClusterServiceConfigResponse struct {
	responses.BaseResponse
	RequestId string
	Config    DescribeClusterServiceConfigConfig
}

type DescribeClusterServiceConfigConfig struct {
	ServiceName      string
	ConfigVersion    string
	Applied          string
	CreateTime       string
	Author           string
	Comment          string
	ConfigValueList  DescribeClusterServiceConfigConfigValueList
	PropertyInfoList DescribeClusterServiceConfigPropertyInfoList
}

type DescribeClusterServiceConfigConfigValue struct {
	ConfigName          string
	AllowCustom         bool
	ConfigItemValueList DescribeClusterServiceConfigConfigItemValueList
}

type DescribeClusterServiceConfigConfigItemValue struct {
	ItemName    string
	Value       string
	IsCustom    bool
	Description string
}

type DescribeClusterServiceConfigPropertyInfo struct {
	Name                    string
	Value                   string
	Description             string
	FileName                string
	DisplayName             string
	ServiceName             string
	Component               string
	PropertyTypes           DescribeClusterServiceConfigPropertyTypeList
	PropertyValueAttributes DescribeClusterServiceConfigPropertyValueAttributes
	EffectWay               DescribeClusterServiceConfigEffectWay
}

type DescribeClusterServiceConfigPropertyValueAttributes struct {
	Type          string
	Maximum       string
	Mimimum       string
	Unit          string
	ReadOnly      bool
	Hidden        bool
	IncrememtStep string
	Entries       DescribeClusterServiceConfigValueEntryInfoList
}

type DescribeClusterServiceConfigValueEntryInfo struct {
	Value       string
	Label       string
	Description string
}

type DescribeClusterServiceConfigEffectWay struct {
	EffectType        string
	InvokeServiceName string
}

type DescribeClusterServiceConfigConfigValueList []DescribeClusterServiceConfigConfigValue

func (list *DescribeClusterServiceConfigConfigValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterServiceConfigConfigValue)
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

type DescribeClusterServiceConfigPropertyInfoList []DescribeClusterServiceConfigPropertyInfo

func (list *DescribeClusterServiceConfigPropertyInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterServiceConfigPropertyInfo)
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

type DescribeClusterServiceConfigConfigItemValueList []DescribeClusterServiceConfigConfigItemValue

func (list *DescribeClusterServiceConfigConfigItemValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterServiceConfigConfigItemValue)
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

type DescribeClusterServiceConfigPropertyTypeList []string

func (list *DescribeClusterServiceConfigPropertyTypeList) UnmarshalJSON(data []byte) error {
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

type DescribeClusterServiceConfigValueEntryInfoList []DescribeClusterServiceConfigValueEntryInfo

func (list *DescribeClusterServiceConfigValueEntryInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterServiceConfigValueEntryInfo)
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
