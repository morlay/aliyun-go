package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UnbindAliasRequest struct {
	requests.RpcRequest
	AliasName string `position:"Query" name:"AliasName"`
	AppKey    int64  `position:"Query" name:"AppKey"`
	DeviceId  string `position:"Query" name:"DeviceId"`
	UnbindAll string `position:"Query" name:"UnbindAll"`
}

func (req *UnbindAliasRequest) Invoke(client *sdk.Client) (resp *UnbindAliasResponse, err error) {
	req.InitWithApiInfo("Push", "2016-08-01", "UnbindAlias", "", "")
	resp = &UnbindAliasResponse{}
	err = client.DoAction(req, resp)
	return
}

type UnbindAliasResponse struct {
	responses.BaseResponse
	RequestId common.String
}
