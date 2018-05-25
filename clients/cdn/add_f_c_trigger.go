package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddFCTriggerRequest struct {
	requests.RpcRequest
	Notes            string `position:"Query" name:"Notes"`
	EventMetaVersion string `position:"Query" name:"EventMetaVersion"`
	TriggerARN       string `position:"Query" name:"TriggerARN"`
	SourceARN        string `position:"Query" name:"SourceARN"`
	OwnerId          int64  `position:"Query" name:"OwnerId"`
	RoleARN          string `position:"Query" name:"RoleARN"`
	EventMetaName    string `position:"Query" name:"EventMetaName"`
}

func (req *AddFCTriggerRequest) Invoke(client *sdk.Client) (resp *AddFCTriggerResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "AddFCTrigger", "", "")
	resp = &AddFCTriggerResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddFCTriggerResponse struct {
	responses.BaseResponse
	RequestId common.String
}
