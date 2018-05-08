package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SubmitFlowJobRequest struct {
	requests.RpcRequest
	JobId           string `position:"Query" name:"JobId"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	HostName        string `position:"Query" name:"HostName"`
	Conf            string `position:"Query" name:"Conf"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
}

func (req *SubmitFlowJobRequest) Invoke(client *sdk.Client) (resp *SubmitFlowJobResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "SubmitFlowJob", "", "")
	resp = &SubmitFlowJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubmitFlowJobResponse struct {
	responses.BaseResponse
	RequestId common.String
	Id        common.String
}
