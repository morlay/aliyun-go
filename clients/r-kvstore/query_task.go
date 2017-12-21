
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type QueryTaskRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
InstanceId string `position:"Domain" name:"InstanceId"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (r QueryTaskRequest) Invoke(client *sdk.Client) (response *QueryTaskResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryTaskRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "QueryTask", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		QueryTaskResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.QueryTaskResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryTaskResponse struct {
RequestId string
Action string
Progress int
}

