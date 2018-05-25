package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeFCTriggerRequest struct {
	requests.RpcRequest
	TriggerARN string `position:"Query" name:"TriggerARN"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeFCTriggerRequest) Invoke(client *sdk.Client) (resp *DescribeFCTriggerResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeFCTrigger", "", "")
	resp = &DescribeFCTriggerResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeFCTriggerResponse struct {
	responses.BaseResponse
	RequestId common.String
	FCTrigger DescribeFCTriggerFCTrigger
}

type DescribeFCTriggerFCTrigger struct {
	EventMetaName    common.String
	EventMetaVersion common.String
	TriggerARN       common.String
	RoleARN          common.String
	SourceArn        common.String
	Notes            common.String
}
