package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddLivePullStreamInfoRequest struct {
	SourceUrl     string `position:"Query" name:"SourceUrl"`
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (r AddLivePullStreamInfoRequest) Invoke(client *sdk.Client) (response *AddLivePullStreamInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddLivePullStreamInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "AddLivePullStreamInfo", "", "")

	resp := struct {
		*responses.BaseResponse
		AddLivePullStreamInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddLivePullStreamInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddLivePullStreamInfoResponse struct {
	RequestId string
}
