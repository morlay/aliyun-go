package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreatePhysicalConnectionRequest struct {
	requests.RpcRequest
	AccessPointId                 string `position:"Query" name:"AccessPointId"`
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
	Type                          string `position:"Query" name:"Type"`
	OwnerId                       int64  `position:"Query" name:"OwnerId"`
	LineOperator                  string `position:"Query" name:"LineOperator"`
	Name                          string `position:"Query" name:"Name"`
}

func (req *CreatePhysicalConnectionRequest) Invoke(client *sdk.Client) (resp *CreatePhysicalConnectionResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "CreatePhysicalConnection", "vpc", "")
	resp = &CreatePhysicalConnectionResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreatePhysicalConnectionResponse struct {
	responses.BaseResponse
	RequestId            string
	PhysicalConnectionId string
}
