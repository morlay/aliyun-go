package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UpdateFCTriggerRequest struct {
	requests.RpcRequest
	Notes      string `position:"Query" name:"Notes"`
	TriggerARN string `position:"Query" name:"TriggerARN"`
	SourceARN  string `position:"Query" name:"SourceARN"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	RoleARN    string `position:"Query" name:"RoleARN"`
}

func (req *UpdateFCTriggerRequest) Invoke(client *sdk.Client) (resp *UpdateFCTriggerResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "UpdateFCTrigger", "", "")
	resp = &UpdateFCTriggerResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateFCTriggerResponse struct {
	responses.BaseResponse
	RequestId common.String
}
