package cloudphoto

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetMeRequest struct {
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	FaceId    int64  `position:"Query" name:"FaceId"`
}

func (r SetMeRequest) Invoke(client *sdk.Client) (response *SetMeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetMeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "SetMe", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		SetMeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetMeResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetMeResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
}
