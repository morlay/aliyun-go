package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreatePhysicalConnectionNewRequest struct {
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
	InterfaceName                 string `position:"Query" name:"InterfaceName"`
	Type                          string `position:"Query" name:"Type"`
	OwnerId                       int64  `position:"Query" name:"OwnerId"`
	LineOperator                  string `position:"Query" name:"LineOperator"`
	Name                          string `position:"Query" name:"Name"`
	DeviceName                    string `position:"Query" name:"DeviceName"`
}

func (req *CreatePhysicalConnectionNewRequest) Invoke(client *sdk.Client) (resp *CreatePhysicalConnectionNewResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "CreatePhysicalConnectionNew", "vpc", "")
	resp = &CreatePhysicalConnectionNewResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreatePhysicalConnectionNewResponse struct {
	responses.BaseResponse
	RequestId            common.String
	PhysicalConnectionId common.String
}
