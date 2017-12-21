
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type ModifyInstanceMinorVersionRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
InstanceId string `position:"Query" name:"InstanceId"`
SecurityToken string `position:"Query" name:"SecurityToken"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
Minorversion string `position:"Query" name:"Minorversion"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (r ModifyInstanceMinorVersionRequest) Invoke(client *sdk.Client) (response *ModifyInstanceMinorVersionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyInstanceMinorVersionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyInstanceMinorVersion", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		ModifyInstanceMinorVersionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.ModifyInstanceMinorVersionResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyInstanceMinorVersionResponse struct {
RequestId string
}

