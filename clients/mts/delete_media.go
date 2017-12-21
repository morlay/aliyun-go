package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteMediaRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MediaIds             string `position:"Query" name:"MediaIds"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteMediaRequest) Invoke(client *sdk.Client) (response *DeleteMediaResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteMediaRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "DeleteMedia", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteMediaResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteMediaResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteMediaResponse struct {
	RequestId string
}
