package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyFlowProjectClusterSettingRequest struct {
	requests.RpcRequest
	UserLists       *ModifyFlowProjectClusterSettingUserListList  `position:"Query" type:"Repeated" name:"UserList"`
	ResourceOwnerId int64                                         `position:"Query" name:"ResourceOwnerId"`
	QueueLists      *ModifyFlowProjectClusterSettingQueueListList `position:"Query" type:"Repeated" name:"QueueList"`
	HostLists       *ModifyFlowProjectClusterSettingHostListList  `position:"Query" type:"Repeated" name:"HostList"`
	ClusterId       string                                        `position:"Query" name:"ClusterId"`
	DefaultQueue    string                                        `position:"Query" name:"DefaultQueue"`
	ProjectId       string                                        `position:"Query" name:"ProjectId"`
	DefaultUser     string                                        `position:"Query" name:"DefaultUser"`
}

func (req *ModifyFlowProjectClusterSettingRequest) Invoke(client *sdk.Client) (resp *ModifyFlowProjectClusterSettingResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ModifyFlowProjectClusterSetting", "", "")
	resp = &ModifyFlowProjectClusterSettingResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyFlowProjectClusterSettingResponse struct {
	responses.BaseResponse
	RequestId string
	Data      bool
}

type ModifyFlowProjectClusterSettingUserListList []string

func (list *ModifyFlowProjectClusterSettingUserListList) UnmarshalJSON(data []byte) error {
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

type ModifyFlowProjectClusterSettingQueueListList []string

func (list *ModifyFlowProjectClusterSettingQueueListList) UnmarshalJSON(data []byte) error {
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

type ModifyFlowProjectClusterSettingHostListList []string

func (list *ModifyFlowProjectClusterSettingHostListList) UnmarshalJSON(data []byte) error {
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
