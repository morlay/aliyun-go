package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type KillFlowJobRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	JobInstanceId   string `position:"Query" name:"JobInstanceId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
}

func (req *KillFlowJobRequest) Invoke(client *sdk.Client) (resp *KillFlowJobResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "KillFlowJob", "", "")
	resp = &KillFlowJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type KillFlowJobResponse struct {
	responses.BaseResponse
	RequestId common.String
	Data      bool
}
