package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeAvailableInstanceTypeRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *DescribeAvailableInstanceTypeRequest) Invoke(client *sdk.Client) (resp *DescribeAvailableInstanceTypeResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeAvailableInstanceType", "", "")
	resp = &DescribeAvailableInstanceTypeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeAvailableInstanceTypeResponse struct {
	responses.BaseResponse
	RequestId                    common.String
	EmrSupportedInstanceTypeList DescribeAvailableInstanceTypeEmrSupportInstanceTypeList
}

type DescribeAvailableInstanceTypeEmrSupportInstanceType struct {
	ClusterType             common.String
	NodeTypeSupportInfoList DescribeAvailableInstanceTypeClusterNodeTypeSupportInfoList
}

type DescribeAvailableInstanceTypeClusterNodeTypeSupportInfo struct {
	ClusterNodeType         common.String
	SupportInstanceTypeList DescribeAvailableInstanceTypeSupportInstanceTypeListList
}

type DescribeAvailableInstanceTypeEmrSupportInstanceTypeList []DescribeAvailableInstanceTypeEmrSupportInstanceType

func (list *DescribeAvailableInstanceTypeEmrSupportInstanceTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAvailableInstanceTypeEmrSupportInstanceType)
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

type DescribeAvailableInstanceTypeClusterNodeTypeSupportInfoList []DescribeAvailableInstanceTypeClusterNodeTypeSupportInfo

func (list *DescribeAvailableInstanceTypeClusterNodeTypeSupportInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAvailableInstanceTypeClusterNodeTypeSupportInfo)
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

type DescribeAvailableInstanceTypeSupportInstanceTypeListList []common.String

func (list *DescribeAvailableInstanceTypeSupportInstanceTypeListList) UnmarshalJSON(data []byte) error {
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
