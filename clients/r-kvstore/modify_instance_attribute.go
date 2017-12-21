
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type ModifyInstanceAttributeRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
InstanceId string `position:"Query" name:"InstanceId"`
InstanceName string `position:"Query" name:"InstanceName"`
SecurityToken string `position:"Query" name:"SecurityToken"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
NewPassword string `position:"Query" name:"NewPassword"`
}

func (r ModifyInstanceAttributeRequest) Invoke(client *sdk.Client) (response *ModifyInstanceAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyInstanceAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyInstanceAttribute", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		ModifyInstanceAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.ModifyInstanceAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyInstanceAttributeResponse struct {
RequestId string
}

