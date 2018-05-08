package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeClusterServiceRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *DescribeClusterServiceRequest) Invoke(client *sdk.Client) (resp *DescribeClusterServiceResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeClusterService", "", "")
	resp = &DescribeClusterServiceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClusterServiceResponse struct {
	responses.BaseResponse
	RequestId   common.String
	ServiceInfo DescribeClusterServiceServiceInfo
}

type DescribeClusterServiceServiceInfo struct {
	ServiceName                  common.String
	ServiceVersion               common.String
	ServiceStatus                common.String
	NeedRestartInfo              common.String
	NeedRestartNum               common.Integer
	ServiceActionList            DescribeClusterServiceServiceActionList
	ClusterServiceSummaryList    DescribeClusterServiceClusterServiceSummaryList
	NeedRestartHostIdList        DescribeClusterServiceNeedRestartHostIdListList
	NeedRestartComponentNameList DescribeClusterServiceNeedRestartComponentNameListList
}

type DescribeClusterServiceServiceAction struct {
	ServiceName   common.String
	ComponentName common.String
	ActionName    common.String
	Command       common.String
	DisplayName   common.String
}

type DescribeClusterServiceClusterServiceSummary struct {
	Key         common.String
	DisplayName common.String
	Value       common.String
	Status      common.String
	Type        common.String
	AlertInfo   common.String
}

type DescribeClusterServiceServiceActionList []DescribeClusterServiceServiceAction

func (list *DescribeClusterServiceServiceActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterServiceServiceAction)
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

type DescribeClusterServiceClusterServiceSummaryList []DescribeClusterServiceClusterServiceSummary

func (list *DescribeClusterServiceClusterServiceSummaryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterServiceClusterServiceSummary)
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

type DescribeClusterServiceNeedRestartHostIdListList []common.String

func (list *DescribeClusterServiceNeedRestartHostIdListList) UnmarshalJSON(data []byte) error {
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

type DescribeClusterServiceNeedRestartComponentNameListList []common.String

func (list *DescribeClusterServiceNeedRestartComponentNameListList) UnmarshalJSON(data []byte) error {
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
