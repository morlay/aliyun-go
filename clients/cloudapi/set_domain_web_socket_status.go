package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetDomainWebSocketStatusRequest struct {
	requests.RpcRequest
	GroupId     string `position:"Query" name:"GroupId"`
	DomainName  string `position:"Query" name:"DomainName"`
	ActionValue string `position:"Query" name:"ActionValue"`
}

func (req *SetDomainWebSocketStatusRequest) Invoke(client *sdk.Client) (resp *SetDomainWebSocketStatusResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "SetDomainWebSocketStatus", "apigateway", "")
	resp = &SetDomainWebSocketStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetDomainWebSocketStatusResponse struct {
	responses.BaseResponse
	RequestId string
}
