package ehpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type EditJobTemplateRequest struct {
	requests.RpcRequest
	StderrRedirectPath string `position:"Query" name:"StderrRedirectPath"`
	Variables          string `position:"Query" name:"Variables"`
	RunasUser          string `position:"Query" name:"RunasUser"`
	TemplateId         string `position:"Query" name:"TemplateId"`
	Priority           int    `position:"Query" name:"Priority"`
	CommandLine        string `position:"Query" name:"CommandLine"`
	ArrayRequest       string `position:"Query" name:"ArrayRequest"`
	PackagePath        string `position:"Query" name:"PackagePath"`
	ReRunnable         string `position:"Query" name:"ReRunnable"`
	Name               string `position:"Query" name:"Name"`
	StdoutRedirectPath string `position:"Query" name:"StdoutRedirectPath"`
}

func (req *EditJobTemplateRequest) Invoke(client *sdk.Client) (resp *EditJobTemplateResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "EditJobTemplate", "ehs", "")
	resp = &EditJobTemplateResponse{}
	err = client.DoAction(req, resp)
	return
}

type EditJobTemplateResponse struct {
	responses.BaseResponse
	RequestId  string
	TemplateId string
}
