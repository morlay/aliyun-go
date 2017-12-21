package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type NodeStatusListRequest struct {
	InstanceIds string `position:"Query" name:"InstanceIds"`
}

func (r NodeStatusListRequest) Invoke(client *sdk.Client) (response *NodeStatusListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		NodeStatusListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cms", "2017-03-01", "NodeStatusList", "cms", "")

	resp := struct {
		*responses.BaseResponse
		NodeStatusListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.NodeStatusListResponse

	err = client.DoAction(&req, &resp)
	return
}

type NodeStatusListResponse struct {
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
