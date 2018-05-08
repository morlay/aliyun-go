package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeParameterTemplatesRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	Engine               string `position:"Query" name:"Engine"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EngineVersion        string `position:"Query" name:"EngineVersion"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeParameterTemplatesRequest) Invoke(client *sdk.Client) (resp *DescribeParameterTemplatesResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeParameterTemplates", "rds", "")
	resp = &DescribeParameterTemplatesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeParameterTemplatesResponse struct {
	responses.BaseResponse
	RequestId      common.String
	Engine         common.String
	EngineVersion  common.String
	ParameterCount common.String
	Parameters     DescribeParameterTemplatesTemplateRecordList
}

type DescribeParameterTemplatesTemplateRecord struct {
	ParameterName        common.String
	ParameterValue       common.String
	ForceModify          common.String
	ForceRestart         common.String
	CheckingCode         common.String
	ParameterDescription common.String
}

type DescribeParameterTemplatesTemplateRecordList []DescribeParameterTemplatesTemplateRecord

func (list *DescribeParameterTemplatesTemplateRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeParameterTemplatesTemplateRecord)
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
