package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type NodeProcessesRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (r NodeProcessesRequest) Invoke(client *sdk.Client) (response *NodeProcessesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		NodeProcessesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cms", "2017-03-01", "NodeProcesses", "cms", "")

	resp := struct {
		*responses.BaseResponse
		NodeProcessesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.NodeProcessesResponse

	err = client.DoAction(&req, &resp)
	return
}

type NodeProcessesResponse struct {
	ErrorCode     int
	ErrorMessage  string
	Success       bool
	RequestId     string
	NodeProcesses NodeProcessesNodeProcessList
}

type NodeProcessesNodeProcess struct {
	Id          int64
	Name        string
	InstanceId  string
	ProcessName string
	ProcessUser string
	Command     string
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
