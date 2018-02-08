package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListJobTemplatesRequest struct {
	requests.RpcRequest
	Name       string `position:"Query" name:"Name"`
	PageSize   int    `position:"Query" name:"PageSize"`
	PageNumber int    `position:"Query" name:"PageNumber"`
}

func (req *ListJobTemplatesRequest) Invoke(client *sdk.Client) (resp *ListJobTemplatesResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "ListJobTemplates", "ehs", "")
	resp = &ListJobTemplatesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListJobTemplatesResponse struct {
	responses.BaseResponse
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	Templates  ListJobTemplatesJobTemplatesList
}

type ListJobTemplatesJobTemplates struct {
	Id                 string
	Name               string
	CommandLine        string
	RunasUser          string
	Priority           int
	_package           string
	StdoutRedirectPath string
	StderrRedirectPath string
	ReRunable          bool
	ArrayRequest       string
	Variables          string
}

type ListJobTemplatesJobTemplatesList []ListJobTemplatesJobTemplates

func (list *ListJobTemplatesJobTemplatesList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobTemplatesJobTemplates)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
