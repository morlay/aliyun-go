
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type CreateTempInstanceRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
SnapshotId string `position:"Domain" name:"SnapshotId"`
InstanceId string `position:"Domain" name:"InstanceId"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (r CreateTempInstanceRequest) Invoke(client *sdk.Client) (response *CreateTempInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateTempInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "CreateTempInstance", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		CreateTempInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.CreateTempInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateTempInstanceResponse struct {
RequestId string
InstanceId string
SnapshotId string
TempInstanceId string
Status string
}

