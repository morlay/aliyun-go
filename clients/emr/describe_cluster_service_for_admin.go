package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterServiceForAdminRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *DescribeClusterServiceForAdminRequest) Invoke(client *sdk.Client) (resp *DescribeClusterServiceForAdminResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeClusterServiceForAdmin", "", "")
	resp = &DescribeClusterServiceForAdminResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClusterServiceForAdminResponse struct {
	responses.BaseResponse
	RequestId   string
	ServiceInfo DescribeClusterServiceForAdminServiceInfo
}

type DescribeClusterServiceForAdminServiceInfo struct {
	ServiceName                  string
	ServiceVersion               string
	ServiceStatus                string
	NeedRestartInfo              string
	NeedRestartNum               int
	ServiceActionList            DescribeClusterServiceForAdminServiceActionList
	ClusterServiceSummaryList    DescribeClusterServiceForAdminClusterServiceSummaryList
	NeedRestartHostIdList        DescribeClusterServiceForAdminNeedRestartHostIdListList
	NeedRestartComponentNameList DescribeClusterServiceForAdminNeedRestartComponentNameListList
}

type DescribeClusterServiceForAdminServiceAction struct {
	ServiceName   string
	ComponentName string
	ActionName    string
	Command       string
	DisplayName   string
}

type DescribeClusterServiceForAdminClusterServiceSummary struct {
	Key         string
	DisplayName string
	Value       string
	Status      string
	Type        string
	AlertInfo   string
}

type DescribeClusterServiceForAdminServiceActionList []DescribeClusterServiceForAdminServiceAction

func (list *DescribeClusterServiceForAdminServiceActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterServiceForAdminServiceAction)
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

type DescribeClusterServiceForAdminClusterServiceSummaryList []DescribeClusterServiceForAdminClusterServiceSummary

func (list *DescribeClusterServiceForAdminClusterServiceSummaryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterServiceForAdminClusterServiceSummary)
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

type DescribeClusterServiceForAdminNeedRestartHostIdListList []string

func (list *DescribeClusterServiceForAdminNeedRestartHostIdListList) UnmarshalJSON(data []byte) error {
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

type DescribeClusterServiceForAdminNeedRestartComponentNameListList []string

func (list *DescribeClusterServiceForAdminNeedRestartComponentNameListList) UnmarshalJSON(data []byte) error {
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
