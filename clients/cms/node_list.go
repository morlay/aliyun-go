package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type NodeListRequest struct {
	HostName      string `position:"Query" name:"HostName"`
	InstanceIds   string `position:"Query" name:"InstanceIds"`
	PageSize      int    `position:"Query" name:"PageSize"`
	KeyWord       string `position:"Query" name:"KeyWord"`
	UserId        int64  `position:"Query" name:"UserId"`
	SerialNumbers string `position:"Query" name:"SerialNumbers"`
	PageNumber    int    `position:"Query" name:"PageNumber"`
	Status        string `position:"Query" name:"Status"`
}

func (r NodeListRequest) Invoke(client *sdk.Client) (response *NodeListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		NodeListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cms", "2017-03-01", "NodeList", "cms", "")

	resp := struct {
		*responses.BaseResponse
		NodeListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.NodeListResponse

	err = client.DoAction(&req, &resp)
	return
}

type NodeListResponse struct {
	ErrorCode    int
	ErrorMessage string
	Success      bool
	RequestId    string
	PageNumber   int
	PageSize     int
	PageTotal    int
	Total        int
	Nodes        NodeListNodeList
}

type NodeListNode struct {
	InstanceId       string
	SerialNumber     string
	HostName         string
	AliUid           int64
	OperatingSystem  string
	IpGroup          string
	Region           string
	TianjimonVersion string
	EipAddress       string
	EipId            string
	AliyunHost       bool
	NatIp            string
	NetworkType      string
}

type NodeListNodeList []NodeListNode

func (list *NodeListNodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]NodeListNode)
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
