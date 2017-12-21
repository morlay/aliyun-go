
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type VerifyPasswordRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
Password string `position:"Query" name:"Password"`
InstanceId string `position:"Query" name:"InstanceId"`
SecurityToken string `position:"Query" name:"SecurityToken"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (r VerifyPasswordRequest) Invoke(client *sdk.Client) (response *VerifyPasswordResponse, err error) {
	req := struct {
		*requests.RpcRequest
		VerifyPasswordRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "VerifyPassword", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		VerifyPasswordResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.VerifyPasswordResponse

	err = client.DoAction(&req, &resp)
	return
}

type VerifyPasswordResponse struct {
RequestId string
}

