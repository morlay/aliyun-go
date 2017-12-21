package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreatePhysicalConnectionRequest struct {
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
	UserCidr                      string `position:"Query" name:"UserCidr"`
}

func (r CreatePhysicalConnectionRequest) Invoke(client *sdk.Client) (response *CreatePhysicalConnectionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreatePhysicalConnectionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "CreatePhysicalConnection", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		CreatePhysicalConnectionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreatePhysicalConnectionResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreatePhysicalConnectionResponse struct {
	RequestId            string
	PhysicalConnectionId string
}
