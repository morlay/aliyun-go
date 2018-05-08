package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UnbindTagRequest struct {
	requests.RpcRequest
	TagName   string `position:"Query" name:"TagName"`
	ClientKey string `position:"Query" name:"ClientKey"`
	AppKey    int64  `position:"Query" name:"AppKey"`
	KeyType   string `position:"Query" name:"KeyType"`
}

func (req *UnbindTagRequest) Invoke(client *sdk.Client) (resp *UnbindTagResponse, err error) {
	req.InitWithApiInfo("Push", "2016-08-01", "UnbindTag", "", "")
	resp = &UnbindTagResponse{}
	err = client.DoAction(req, resp)
	return
}

type UnbindTagResponse struct {
	responses.BaseResponse
	RequestId common.String
}
