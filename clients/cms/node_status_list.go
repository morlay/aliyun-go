package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type NodeStatusListRequest struct {
	requests.RpcRequest
	InstanceIds string `position:"Query" name:"InstanceIds"`
}

func (req *NodeStatusListRequest) Invoke(client *sdk.Client) (resp *NodeStatusListResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "NodeStatusList", "cms", "")
	resp = &NodeStatusListResponse{}
	err = client.DoAction(req, resp)
	return
}

type NodeStatusListResponse struct {
	responses.BaseResponse
	ErrorCode      int
	ErrorMessage   string
	Success        bool
	RequestId      string
	NodeStatusList NodeStatusListNodeStatusList
}

type NodeStatusListNodeStatus struct {
	InstanceId  string
	AutoInstall bool
	Status      string
}

type NodeStatusListNodeStatusList []NodeStatusListNodeStatus

func (list *NodeStatusListNodeStatusList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]NodeStatusListNodeStatus)
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
