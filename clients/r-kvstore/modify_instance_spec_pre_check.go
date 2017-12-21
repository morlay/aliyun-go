
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type ModifyInstanceSpecPreCheckRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
InstanceId string `position:"Query" name:"InstanceId"`
SecurityToken string `position:"Query" name:"SecurityToken"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
TargetInstanceClass string `position:"Query" name:"TargetInstanceClass"`
}

func (r ModifyInstanceSpecPreCheckRequest) Invoke(client *sdk.Client) (response *ModifyInstanceSpecPreCheckResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyInstanceSpecPreCheckRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyInstanceSpecPreCheck", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		ModifyInstanceSpecPreCheckResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.ModifyInstanceSpecPreCheckResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyInstanceSpecPreCheckResponse struct {
RequestId string
IsAllowModify bool
DisableCommands string
}

