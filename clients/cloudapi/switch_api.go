package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SwitchApiRequest struct {
	GroupId        string `position:"Query" name:"GroupId"`
	ApiId          string `position:"Query" name:"ApiId"`
	StageName      string `position:"Query" name:"StageName"`
	Description    string `position:"Query" name:"Description"`
	HistoryVersion string `position:"Query" name:"HistoryVersion"`
}

func (r SwitchApiRequest) Invoke(client *sdk.Client) (response *SwitchApiResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SwitchApiRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "SwitchApi", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		SwitchApiResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SwitchApiResponse

	err = client.DoAction(&req, &resp)
	return
}

type SwitchApiResponse struct {
	RequestId string
}
