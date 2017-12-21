package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type BindAliasRequest struct {
	requests.RpcRequest
	AliasName string `position:"Query" name:"AliasName"`
	AppKey    int64  `position:"Query" name:"AppKey"`
	DeviceId  string `position:"Query" name:"DeviceId"`
}

func (req *BindAliasRequest) Invoke(client *sdk.Client) (resp *BindAliasResponse, err error) {
	req.InitWithApiInfo("Push", "2016-08-01", "BindAlias", "", "")
	resp = &BindAliasResponse{}
	err = client.DoAction(req, resp)
	return
}

type BindAliasResponse struct {
	responses.BaseResponse
	RequestId string
}
