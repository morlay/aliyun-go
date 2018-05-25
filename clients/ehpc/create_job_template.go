package ehpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateJobTemplateRequest struct {
	requests.RpcRequest
	StderrRedirectPath string `position:"Query" name:"StderrRedirectPath"`
	Variables          string `position:"Query" name:"Variables"`
	RunasUser          string `position:"Query" name:"RunasUser"`
	ReRunable          string `position:"Query" name:"ReRunable"`
	Priority           int    `position:"Query" name:"Priority"`
	CommandLine        string `position:"Query" name:"CommandLine"`
	ArrayRequest       string `position:"Query" name:"ArrayRequest"`
	PackagePath        string `position:"Query" name:"PackagePath"`
	Name               string `position:"Query" name:"Name"`
	StdoutRedirectPath string `position:"Query" name:"StdoutRedirectPath"`
}

func (req *CreateJobTemplateRequest) Invoke(client *sdk.Client) (resp *CreateJobTemplateResponse, err error) {
	req.InitWithApiInfo("EHPC", "2018-04-12", "CreateJobTemplate", "ehs", "")
	resp = &CreateJobTemplateResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateJobTemplateResponse struct {
	responses.BaseResponse
	RequestId  common.String
	TemplateId common.String
}
