package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteRecycleBinRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceIds          string `position:"Query" name:"ResourceIds"`
}

func (r DeleteRecycleBinRequest) Invoke(client *sdk.Client) (response *DeleteRecycleBinResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteRecycleBinRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteRecycleBin", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DeleteRecycleBinResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteRecycleBinResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteRecycleBinResponse struct {
	RequestId string
}
