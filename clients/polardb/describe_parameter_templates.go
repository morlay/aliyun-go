package polardb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeParameterTemplatesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeParameterTemplatesRequest) Invoke(client *sdk.Client) (response *DescribeParameterTemplatesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeParameterTemplatesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("polardb", "2017-08-01", "DescribeParameterTemplates", "polardb", "")

	resp := struct {
		*responses.BaseResponse
		DescribeParameterTemplatesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeParameterTemplatesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeParameterTemplatesResponse struct {
	RequestId      string
	Engine         string
	DBType         string
	DBVersion      string
	ParameterCount string
	Parameters     DescribeParameterTemplatesTemplateRecordList
}

type DescribeParameterTemplatesTemplateRecord struct {
	ParameterName        string
	ParameterValue       string
	ForceModify          string
	ForceRestart         string
	CheckingCode         string
	ParameterDescription string
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
