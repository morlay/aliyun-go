package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreatePhysicalConnectionNewRequest struct {
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
	UserCidr                      string `position:"Query" name:"UserCidr"`
}

func (r CreatePhysicalConnectionNewRequest) Invoke(client *sdk.Client) (response *CreatePhysicalConnectionNewResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreatePhysicalConnectionNewRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "CreatePhysicalConnectionNew", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		CreatePhysicalConnectionNewResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreatePhysicalConnectionNewResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreatePhysicalConnectionNewResponse struct {
	RequestId            string
	PhysicalConnectionId string
}
