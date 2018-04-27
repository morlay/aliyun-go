package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListFlowProjectClusterSettingRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int    `position:"Query" name:"PageSize"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
}

func (req *ListFlowProjectClusterSettingRequest) Invoke(client *sdk.Client) (resp *ListFlowProjectClusterSettingResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListFlowProjectClusterSetting", "", "")
	resp = &ListFlowProjectClusterSettingResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListFlowProjectClusterSettingResponse struct {
	responses.BaseResponse
	RequestId       string
	PageNumber      int
	PageSize        int
	Total           int
	ClusterSettings ListFlowProjectClusterSettingClusterSettingList
}

type ListFlowProjectClusterSettingClusterSetting struct {
	GmtCreate    int64
	GmtModified  int64
	ProjectId    string
	ClusterId    string
	DefaultUser  string
	DefaultQueue string
	UserList     ListFlowProjectClusterSettingUserListList
	QueueList    ListFlowProjectClusterSettingQueueListList
	HostList     ListFlowProjectClusterSettingHostListList
}

type ListFlowProjectClusterSettingClusterSettingList []ListFlowProjectClusterSettingClusterSetting

func (list *ListFlowProjectClusterSettingClusterSettingList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFlowProjectClusterSettingClusterSetting)
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

type ListFlowProjectClusterSettingUserListList []string

func (list *ListFlowProjectClusterSettingUserListList) UnmarshalJSON(data []byte) error {
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

type ListFlowProjectClusterSettingQueueListList []string

func (list *ListFlowProjectClusterSettingQueueListList) UnmarshalJSON(data []byte) error {
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

type ListFlowProjectClusterSettingHostListList []string

func (list *ListFlowProjectClusterSettingHostListList) UnmarshalJSON(data []byte) error {
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
