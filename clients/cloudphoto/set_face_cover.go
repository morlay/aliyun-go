package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetFaceCoverRequest struct {
	LibraryId string `position:"Query" name:"LibraryId"`
	PhotoId   int64  `position:"Query" name:"PhotoId"`
	StoreName string `position:"Query" name:"StoreName"`
	FaceId    int64  `position:"Query" name:"FaceId"`
}

func (r SetFaceCoverRequest) Invoke(client *sdk.Client) (response *SetFaceCoverResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetFaceCoverRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "SetFaceCover", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		SetFaceCoverResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetFaceCoverResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetFaceCoverResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
}
