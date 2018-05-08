package ehpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SubmitJobRequest struct {
	requests.RpcRequest
	StderrRedirectPath string `position:"Query" name:"StderrRedirectPath"`
	Variables          string `position:"Query" name:"Variables"`
	RunasUserPassword  string `position:"Query" name:"RunasUserPassword"`
	RunasUser          string `position:"Query" name:"RunasUser"`
	ClusterId          string `position:"Query" name:"ClusterId"`
	ReRunable          string `position:"Query" name:"ReRunable"`
	Priority           int    `position:"Query" name:"Priority"`
	CommandLine        string `position:"Query" name:"CommandLine"`
	ArrayRequest       string `position:"Query" name:"ArrayRequest"`
	PackagePath        string `position:"Query" name:"PackagePath"`
	Name               string `position:"Query" name:"Name"`
	StdoutRedirectPath string `position:"Query" name:"StdoutRedirectPath"`
}

func (req *SubmitJobRequest) Invoke(client *sdk.Client) (resp *SubmitJobResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "SubmitJob", "ehs", "")
	resp = &SubmitJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubmitJobResponse struct {
	responses.BaseResponse
	RequestId common.String
	JobId     common.String
}
