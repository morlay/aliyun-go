package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ProduceEditingProjectVideoRequest struct {
	CoverURL             string `position:"Query" name:"CoverURL"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Timeline             string `position:"Query" name:"Timeline"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Title                string `position:"Query" name:"Title"`
	ProjectId            string `position:"Query" name:"ProjectId"`
}

func (r ProduceEditingProjectVideoRequest) Invoke(client *sdk.Client) (response *ProduceEditingProjectVideoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ProduceEditingProjectVideoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "ProduceEditingProjectVideo", "", "")

	resp := struct {
		*responses.BaseResponse
		ProduceEditingProjectVideoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ProduceEditingProjectVideoResponse

	err = client.DoAction(&req, &resp)
	return
}

type ProduceEditingProjectVideoResponse struct {
	RequestId string
	MediaId   string
}
