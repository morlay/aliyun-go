
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type DeleteInstanceRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
InstanceId string `position:"Query" name:"InstanceId"`
SecurityToken string `position:"Query" name:"SecurityToken"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (r DeleteInstanceRequest) Invoke(client *sdk.Client) (response *DeleteInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DeleteInstance", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		DeleteInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.DeleteInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteInstanceResponse struct {
RequestId string
}

