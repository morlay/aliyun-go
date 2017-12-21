package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type BindTagRequest struct {
	TagName   string `position:"Query" name:"TagName"`
	ClientKey string `position:"Query" name:"ClientKey"`
	AppKey    int64  `position:"Query" name:"AppKey"`
	KeyType   string `position:"Query" name:"KeyType"`
}

func (r BindTagRequest) Invoke(client *sdk.Client) (response *BindTagResponse, err error) {
	req := struct {
		*requests.RpcRequest
		BindTagRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Push", "2016-08-01", "BindTag", "", "")

	resp := struct {
		*responses.BaseResponse
		BindTagResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.BindTagResponse

	err = client.DoAction(&req, &resp)
	return
}

type BindTagResponse struct {
	RequestId string
}
