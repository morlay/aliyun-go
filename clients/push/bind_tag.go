package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type BindTagRequest struct {
	requests.RpcRequest
	TagName   string `position:"Query" name:"TagName"`
	ClientKey string `position:"Query" name:"ClientKey"`
	AppKey    int64  `position:"Query" name:"AppKey"`
	KeyType   string `position:"Query" name:"KeyType"`
}

func (req *BindTagRequest) Invoke(client *sdk.Client) (resp *BindTagResponse, err error) {
	req.InitWithApiInfo("Push", "2016-08-01", "BindTag", "", "")
	resp = &BindTagResponse{}
	err = client.DoAction(req, resp)
	return
}

type BindTagResponse struct {
	responses.BaseResponse
	RequestId common.String
}
