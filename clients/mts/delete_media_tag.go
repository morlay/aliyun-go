package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteMediaTagRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Tag                  string `position:"Query" name:"Tag"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
}

func (r DeleteMediaTagRequest) Invoke(client *sdk.Client) (response *DeleteMediaTagResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteMediaTagRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "DeleteMediaTag", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteMediaTagResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteMediaTagResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteMediaTagResponse struct {
	RequestId string
}
