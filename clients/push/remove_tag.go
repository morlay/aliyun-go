package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RemoveTagRequest struct {
	requests.RpcRequest
	TagName string `position:"Query" name:"TagName"`
	AppKey  int64  `position:"Query" name:"AppKey"`
}

func (req *RemoveTagRequest) Invoke(client *sdk.Client) (resp *RemoveTagResponse, err error) {
	req.InitWithApiInfo("Push", "2016-08-01", "RemoveTag", "", "")
	resp = &RemoveTagResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemoveTagResponse struct {
	responses.BaseResponse
	RequestId common.String
}
