package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId       common.String
	PageNumber      common.Integer
	PageSize        common.Integer
	Total           common.Integer
	ClusterSettings ListFlowProjectClusterSettingClusterSettingList
}

type ListFlowProjectClusterSettingClusterSetting struct {
	GmtCreate    common.Long
	GmtModified  common.Long
	ProjectId    common.String
	ClusterId    common.String
	DefaultUser  common.String
	DefaultQueue common.String
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

type ListFlowProjectClusterSettingUserListList []common.String

func (list *ListFlowProjectClusterSettingUserListList) UnmarshalJSON(data []byte) error {
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

type ListFlowProjectClusterSettingQueueListList []common.String

func (list *ListFlowProjectClusterSettingQueueListList) UnmarshalJSON(data []byte) error {
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

type ListFlowProjectClusterSettingHostListList []common.String

func (list *ListFlowProjectClusterSettingHostListList) UnmarshalJSON(data []byte) error {
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
