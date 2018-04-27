package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId   string
	ServiceInfo DescribeClusterServiceServiceInfo
}

type DescribeClusterServiceServiceInfo struct {
	ServiceName                  string
	ServiceVersion               string
	ServiceStatus                string
	NeedRestartInfo              string
	NeedRestartNum               int
	ServiceActionList            DescribeClusterServiceServiceActionList
	ClusterServiceSummaryList    DescribeClusterServiceClusterServiceSummaryList
	NeedRestartHostIdList        DescribeClusterServiceNeedRestartHostIdListList
	NeedRestartComponentNameList DescribeClusterServiceNeedRestartComponentNameListList
}

type DescribeClusterServiceServiceAction struct {
	ServiceName   string
	ComponentName string
	ActionName    string
	Command       string
	DisplayName   string
}

type DescribeClusterServiceClusterServiceSummary struct {
	Key         string
	DisplayName string
	Value       string
	Status      string
	Type        string
	AlertInfo   string
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

type DescribeClusterServiceNeedRestartHostIdListList []string

func (list *DescribeClusterServiceNeedRestartHostIdListList) UnmarshalJSON(data []byte) error {
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

type DescribeClusterServiceNeedRestartComponentNameListList []string

func (list *DescribeClusterServiceNeedRestartComponentNameListList) UnmarshalJSON(data []byte) error {
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
