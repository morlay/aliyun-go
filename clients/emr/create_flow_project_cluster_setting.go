package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateFlowProjectClusterSettingRequest struct {
	requests.RpcRequest
	UserLists       *CreateFlowProjectClusterSettingUserListList  `position:"Query" type:"Repeated" name:"UserList"`
	ResourceOwnerId int64                                         `position:"Query" name:"ResourceOwnerId"`
	QueueLists      *CreateFlowProjectClusterSettingQueueListList `position:"Query" type:"Repeated" name:"QueueList"`
	HostLists       *CreateFlowProjectClusterSettingHostListList  `position:"Query" type:"Repeated" name:"HostList"`
	ClusterId       string                                        `position:"Query" name:"ClusterId"`
	DefaultQueue    string                                        `position:"Query" name:"DefaultQueue"`
	ProjectId       string                                        `position:"Query" name:"ProjectId"`
	DefaultUser     string                                        `position:"Query" name:"DefaultUser"`
}

func (req *CreateFlowProjectClusterSettingRequest) Invoke(client *sdk.Client) (resp *CreateFlowProjectClusterSettingResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "CreateFlowProjectClusterSetting", "", "")
	resp = &CreateFlowProjectClusterSettingResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateFlowProjectClusterSettingResponse struct {
	responses.BaseResponse
	RequestId string
	Data      bool
}

type CreateFlowProjectClusterSettingUserListList []string

func (list *CreateFlowProjectClusterSettingUserListList) UnmarshalJSON(data []byte) error {
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

type CreateFlowProjectClusterSettingQueueListList []string

func (list *CreateFlowProjectClusterSettingQueueListList) UnmarshalJSON(data []byte) error {
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

type CreateFlowProjectClusterSettingHostListList []string

func (list *CreateFlowProjectClusterSettingHostListList) UnmarshalJSON(data []byte) error {
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
