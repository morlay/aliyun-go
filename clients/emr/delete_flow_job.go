package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteFlowJobRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	ProjectId       string `position:"Query" name:"ProjectId"`
}

func (req *DeleteFlowJobRequest) Invoke(client *sdk.Client) (resp *DeleteFlowJobResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DeleteFlowJob", "", "")
	resp = &DeleteFlowJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteFlowJobResponse struct {
	responses.BaseResponse
	RequestId string
	Data      bool
}
