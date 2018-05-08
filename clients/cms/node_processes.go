package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type NodeProcessesRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (req *NodeProcessesRequest) Invoke(client *sdk.Client) (resp *NodeProcessesResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "NodeProcesses", "cms", "")
	resp = &NodeProcessesResponse{}
	err = client.DoAction(req, resp)
	return
}

type NodeProcessesResponse struct {
	responses.BaseResponse
	ErrorCode     common.Integer
	ErrorMessage  common.String
	Success       bool
	RequestId     common.String
	NodeProcesses NodeProcessesNodeProcessList
}

type NodeProcessesNodeProcess struct {
	Id          common.Long
	Name        common.String
	InstanceId  common.String
	ProcessName common.String
	ProcessUser common.String
	Command     common.String
}

type NodeProcessesNodeProcessList []NodeProcessesNodeProcess

func (list *NodeProcessesNodeProcessList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]NodeProcessesNodeProcess)
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
