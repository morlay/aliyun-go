package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeFlowProjectClusterSettingRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
}

func (req *DescribeFlowProjectClusterSettingRequest) Invoke(client *sdk.Client) (resp *DescribeFlowProjectClusterSettingResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeFlowProjectClusterSetting", "", "")
	resp = &DescribeFlowProjectClusterSettingResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeFlowProjectClusterSettingResponse struct {
	responses.BaseResponse
	RequestId    string
	GmtCreate    int64
	GmtModified  int64
	ProjectId    string
	ClusterId    string
	DefaultUser  string
	DefaultQueue string
	UserList     DescribeFlowProjectClusterSettingUserListList
	QueueList    DescribeFlowProjectClusterSettingQueueListList
	HostList     DescribeFlowProjectClusterSettingHostListList
}

type DescribeFlowProjectClusterSettingUserListList []string

func (list *DescribeFlowProjectClusterSettingUserListList) UnmarshalJSON(data []byte) error {
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

type DescribeFlowProjectClusterSettingQueueListList []string

func (list *DescribeFlowProjectClusterSettingQueueListList) UnmarshalJSON(data []byte) error {
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

type DescribeFlowProjectClusterSettingHostListList []string

func (list *DescribeFlowProjectClusterSettingHostListList) UnmarshalJSON(data []byte) error {
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
