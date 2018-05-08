package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId   common.String
	ServiceInfo DescribeClusterServiceForAdminServiceInfo
}

type DescribeClusterServiceForAdminServiceInfo struct {
	ServiceName                  common.String
	ServiceVersion               common.String
	ServiceStatus                common.String
	NeedRestartInfo              common.String
	NeedRestartNum               common.Integer
	ServiceActionList            DescribeClusterServiceForAdminServiceActionList
	ClusterServiceSummaryList    DescribeClusterServiceForAdminClusterServiceSummaryList
	NeedRestartHostIdList        DescribeClusterServiceForAdminNeedRestartHostIdListList
	NeedRestartComponentNameList DescribeClusterServiceForAdminNeedRestartComponentNameListList
}

type DescribeClusterServiceForAdminServiceAction struct {
	ServiceName   common.String
	ComponentName common.String
	ActionName    common.String
	Command       common.String
	DisplayName   common.String
}

type DescribeClusterServiceForAdminClusterServiceSummary struct {
	Key         common.String
	DisplayName common.String
	Value       common.String
	Status      common.String
	Type        common.String
	AlertInfo   common.String
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

type DescribeClusterServiceForAdminNeedRestartHostIdListList []common.String

func (list *DescribeClusterServiceForAdminNeedRestartHostIdListList) UnmarshalJSON(data []byte) error {
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

type DescribeClusterServiceForAdminNeedRestartComponentNameListList []common.String

func (list *DescribeClusterServiceForAdminNeedRestartComponentNameListList) UnmarshalJSON(data []byte) error {
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
