
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type ModifyInstanceMaintainTimeRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
InstanceId string `position:"Query" name:"InstanceId"`
SecurityToken string `position:"Query" name:"SecurityToken"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
MaintainStartTime string `position:"Query" name:"MaintainStartTime"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
MaintainEndTime string `position:"Query" name:"MaintainEndTime"`
}

func (r ModifyInstanceMaintainTimeRequest) Invoke(client *sdk.Client) (response *ModifyInstanceMaintainTimeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyInstanceMaintainTimeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyInstanceMaintainTime", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		ModifyInstanceMaintainTimeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.ModifyInstanceMaintainTimeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyInstanceMaintainTimeResponse struct {
RequestId string
}

