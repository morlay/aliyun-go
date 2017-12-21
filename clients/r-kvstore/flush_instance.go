
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type FlushInstanceRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
InstanceId string `position:"Query" name:"InstanceId"`
SecurityToken string `position:"Query" name:"SecurityToken"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (r FlushInstanceRequest) Invoke(client *sdk.Client) (response *FlushInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		FlushInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "FlushInstance", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		FlushInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.FlushInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type FlushInstanceResponse struct {
RequestId string
}

