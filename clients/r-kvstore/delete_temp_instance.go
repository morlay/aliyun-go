
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type DeleteTempInstanceRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
TempInstanceId string `position:"Domain" name:"TempInstanceId"`
InstanceId string `position:"Domain" name:"InstanceId"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (r DeleteTempInstanceRequest) Invoke(client *sdk.Client) (response *DeleteTempInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteTempInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DeleteTempInstance", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		DeleteTempInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.DeleteTempInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteTempInstanceResponse struct {
RequestId string
}

