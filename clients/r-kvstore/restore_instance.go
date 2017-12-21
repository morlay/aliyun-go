
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type RestoreInstanceRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
InstanceId string `position:"Query" name:"InstanceId"`
SecurityToken string `position:"Query" name:"SecurityToken"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
BackupId string `position:"Query" name:"BackupId"`
OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (r RestoreInstanceRequest) Invoke(client *sdk.Client) (response *RestoreInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RestoreInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "RestoreInstance", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		RestoreInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.RestoreInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type RestoreInstanceResponse struct {
RequestId string
}

