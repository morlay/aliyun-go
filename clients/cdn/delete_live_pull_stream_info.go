package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteLivePullStreamInfoRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (r DeleteLivePullStreamInfoRequest) Invoke(client *sdk.Client) (response *DeleteLivePullStreamInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteLivePullStreamInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DeleteLivePullStreamInfo", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteLivePullStreamInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteLivePullStreamInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteLivePullStreamInfoResponse struct {
	RequestId string
}
