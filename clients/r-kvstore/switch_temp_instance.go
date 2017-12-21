
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type SwitchTempInstanceRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
TempInstanceId string `position:"Domain" name:"TempInstanceId"`
InstanceId string `position:"Domain" name:"InstanceId"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (r SwitchTempInstanceRequest) Invoke(client *sdk.Client) (response *SwitchTempInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SwitchTempInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "SwitchTempInstance", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		SwitchTempInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.SwitchTempInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type SwitchTempInstanceResponse struct {
RequestId string
}

