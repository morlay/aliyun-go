package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterServiceConfigHistoryRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ConfigVersion   string `position:"Query" name:"ConfigVersion"`
}

func (req *DescribeClusterServiceConfigHistoryRequest) Invoke(client *sdk.Client) (resp *DescribeClusterServiceConfigHistoryResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeClusterServiceConfigHistory", "", "")
	resp = &DescribeClusterServiceConfigHistoryResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClusterServiceConfigHistoryResponse struct {
	responses.BaseResponse
	RequestId string
	Config    DescribeClusterServiceConfigHistoryConfig
}

type DescribeClusterServiceConfigHistoryConfig struct {
	ServiceName     string
	ConfigVersion   string
	Applied         bool
	CreateTime      string
	Author          string
	Comment         string
	ConfigValueList DescribeClusterServiceConfigHistoryConfigValueList
}

type DescribeClusterServiceConfigHistoryConfigValue struct {
	ConfigName          string
	ConfigItemValueList DescribeClusterServiceConfigHistoryConfigItemValueList
}

type DescribeClusterServiceConfigHistoryConfigItemValue struct {
	ItemName   string
	Value      string
	OldValue   string
	ChangeType string
}

type DescribeClusterServiceConfigHistoryConfigValueList []DescribeClusterServiceConfigHistoryConfigValue

func (list *DescribeClusterServiceConfigHistoryConfigValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterServiceConfigHistoryConfigValue)
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

type DescribeClusterServiceConfigHistoryConfigItemValueList []DescribeClusterServiceConfigHistoryConfigItemValue

func (list *DescribeClusterServiceConfigHistoryConfigItemValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterServiceConfigHistoryConfigItemValue)
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
