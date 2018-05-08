package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type NodeListRequest struct {
	requests.RpcRequest
	HostName         string `position:"Query" name:"HostName"`
	InstanceIds      string `position:"Query" name:"InstanceIds"`
	InstanceRegionId string `position:"Query" name:"InstanceRegionId"`
	PageSize         int    `position:"Query" name:"PageSize"`
	KeyWord          string `position:"Query" name:"KeyWord"`
	UserId           int64  `position:"Query" name:"UserId"`
	PageNumber       int    `position:"Query" name:"PageNumber"`
	SerialNumbers    string `position:"Query" name:"SerialNumbers"`
	Status           string `position:"Query" name:"Status"`
}

func (req *NodeListRequest) Invoke(client *sdk.Client) (resp *NodeListResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "NodeList", "cms", "")
	resp = &NodeListResponse{}
	err = client.DoAction(req, resp)
	return
}

type NodeListResponse struct {
	responses.BaseResponse
	ErrorCode    common.Integer
	ErrorMessage common.String
	Success      bool
	RequestId    common.String
	PageNumber   common.Integer
	PageSize     common.Integer
	PageTotal    common.Integer
	Total        common.Integer
	Nodes        NodeListNodeList
}

type NodeListNode struct {
	InstanceId       common.String
	SerialNumber     common.String
	HostName         common.String
	AliUid           common.Long
	OperatingSystem  common.String
	IpGroup          common.String
	Region           common.String
	TianjimonVersion common.String
	EipAddress       common.String
	EipId            common.String
	AliyunHost       bool
	NatIp            common.String
	NetworkType      common.String
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
