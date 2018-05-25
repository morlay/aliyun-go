package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListJobTemplatesRequest struct {
	requests.RpcRequest
	Name       string `position:"Query" name:"Name"`
	PageSize   int    `position:"Query" name:"PageSize"`
	PageNumber int    `position:"Query" name:"PageNumber"`
}

func (req *ListJobTemplatesRequest) Invoke(client *sdk.Client) (resp *ListJobTemplatesResponse, err error) {
	req.InitWithApiInfo("EHPC", "2018-04-12", "ListJobTemplates", "ehs", "")
	resp = &ListJobTemplatesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListJobTemplatesResponse struct {
	responses.BaseResponse
	RequestId  common.String
	TotalCount common.Integer
	PageNumber common.Integer
	PageSize   common.Integer
	Templates  ListJobTemplatesJobTemplatesList
}

type ListJobTemplatesJobTemplates struct {
	Id                 common.String
	Name               common.String
	CommandLine        common.String
	RunasUser          common.String
	Priority           common.Integer
	PackagePath        common.String
	StdoutRedirectPath common.String
	StderrRedirectPath common.String
	ReRunable          bool
	ArrayRequest       common.String
	Variables          common.String
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
