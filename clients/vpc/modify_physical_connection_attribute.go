package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyPhysicalConnectionAttributeRequest struct {
	requests.RpcRequest
	RedundantPhysicalConnectionId string `position:"Query" name:"RedundantPhysicalConnectionId"`
	PeerLocation                  string `position:"Query" name:"PeerLocation"`
	ResourceOwnerId               int64  `position:"Query" name:"ResourceOwnerId"`
	PortType                      string `position:"Query" name:"PortType"`
	CircuitCode                   string `position:"Query" name:"CircuitCode"`
	Bandwidth                     int    `position:"Query" name:"Bandwidth"`
	ClientToken                   string `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount          string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                  string `position:"Query" name:"OwnerAccount"`
	Description                   string `position:"Query" name:"Description"`
	OwnerId                       int64  `position:"Query" name:"OwnerId"`
	LineOperator                  string `position:"Query" name:"LineOperator"`
	PhysicalConnectionId          string `position:"Query" name:"PhysicalConnectionId"`
	Name                          string `position:"Query" name:"Name"`
}

func (req *ModifyPhysicalConnectionAttributeRequest) Invoke(client *sdk.Client) (resp *ModifyPhysicalConnectionAttributeResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyPhysicalConnectionAttribute", "vpc", "")
	resp = &ModifyPhysicalConnectionAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyPhysicalConnectionAttributeResponse struct {
	responses.BaseResponse
	RequestId common.String
}
