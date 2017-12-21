package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateGlobalAccelerationInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ServiceLocation      string `position:"Query" name:"ServiceLocation"`
	Bandwidth            string `position:"Query" name:"Bandwidth"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	InternetChargeType   string `position:"Query" name:"InternetChargeType"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r CreateGlobalAccelerationInstanceRequest) Invoke(client *sdk.Client) (response *CreateGlobalAccelerationInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateGlobalAccelerationInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "CreateGlobalAccelerationInstance", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		CreateGlobalAccelerationInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateGlobalAccelerationInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateGlobalAccelerationInstanceResponse struct {
	RequestId                    string
	GlobalAccelerationInstanceId string
	IpAddress                    string
}
