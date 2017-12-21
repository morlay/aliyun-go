package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateMediaCoverRequest struct {
	CoverURL             string `position:"Query" name:"CoverURL"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
}

func (r UpdateMediaCoverRequest) Invoke(client *sdk.Client) (response *UpdateMediaCoverResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdateMediaCoverRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "UpdateMediaCover", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdateMediaCoverResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateMediaCoverResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateMediaCoverResponse struct {
	RequestId string
}
