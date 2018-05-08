package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	Config    DescribeClusterServiceConfigConfig
}

type DescribeClusterServiceConfigConfig struct {
	ServiceName      common.String
	ConfigVersion    common.String
	Applied          common.String
	CreateTime       common.String
	Author           common.String
	Comment          common.String
	ConfigValueList  DescribeClusterServiceConfigConfigValueList
	PropertyInfoList DescribeClusterServiceConfigPropertyInfoList
}

type DescribeClusterServiceConfigConfigValue struct {
	ConfigName          common.String
	AllowCustom         bool
	ConfigItemValueList DescribeClusterServiceConfigConfigItemValueList
}

type DescribeClusterServiceConfigConfigItemValue struct {
	ItemName    common.String
	Value       common.String
	IsCustom    bool
	Description common.String
}

type DescribeClusterServiceConfigPropertyInfo struct {
	Name                    common.String
	Value                   common.String
	Description             common.String
	FileName                common.String
	DisplayName             common.String
	ServiceName             common.String
	Component               common.String
	PropertyTypes           DescribeClusterServiceConfigPropertyTypeList
	PropertyValueAttributes DescribeClusterServiceConfigPropertyValueAttributes
	EffectWay               DescribeClusterServiceConfigEffectWay
}

type DescribeClusterServiceConfigPropertyValueAttributes struct {
	Type          common.String
	Maximum       common.String
	Mimimum       common.String
	Unit          common.String
	ReadOnly      bool
	Hidden        bool
	IncrememtStep common.String
	Entries       DescribeClusterServiceConfigValueEntryInfoList
}

type DescribeClusterServiceConfigValueEntryInfo struct {
	Value       common.String
	Label       common.String
	Description common.String
}

type DescribeClusterServiceConfigEffectWay struct {
	EffectType        common.String
	InvokeServiceName common.String
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

type DescribeClusterServiceConfigPropertyTypeList []common.String

func (list *DescribeClusterServiceConfigPropertyTypeList) UnmarshalJSON(data []byte) error {
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
