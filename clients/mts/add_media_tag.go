package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddMediaTagRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Tag                  string `position:"Query" name:"Tag"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
}

func (r AddMediaTagRequest) Invoke(client *sdk.Client) (response *AddMediaTagResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddMediaTagRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "AddMediaTag", "", "")

	resp := struct {
		*responses.BaseResponse
		AddMediaTagResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddMediaTagResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddMediaTagResponse struct {
	RequestId string
}
