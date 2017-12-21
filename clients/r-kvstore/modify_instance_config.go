
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type ModifyInstanceConfigRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
InstanceId string `position:"Query" name:"InstanceId"`
SecurityToken string `position:"Query" name:"SecurityToken"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
Config string `position:"Query" name:"Config"`
}

func (r ModifyInstanceConfigRequest) Invoke(client *sdk.Client) (response *ModifyInstanceConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyInstanceConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyInstanceConfig", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		ModifyInstanceConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.ModifyInstanceConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyInstanceConfigResponse struct {
RequestId string
}

