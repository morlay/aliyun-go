
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type ModifyCertificationRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
InstanceId string `position:"Query" name:"InstanceId"`
SecurityToken string `position:"Query" name:"SecurityToken"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
NoCertification string `position:"Query" name:"NoCertification"`
}

func (r ModifyCertificationRequest) Invoke(client *sdk.Client) (response *ModifyCertificationResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyCertificationRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyCertification", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		ModifyCertificationResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.ModifyCertificationResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyCertificationResponse struct {
RequestId string
}

