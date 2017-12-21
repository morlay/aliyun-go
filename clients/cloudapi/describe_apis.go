package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeApisRequest struct {
	requests.RpcRequest
	GroupId    string `position:"Query" name:"GroupId"`
	ApiId      string `position:"Query" name:"ApiId"`
	ApiName    string `position:"Query" name:"ApiName"`
	CatalogId  string `position:"Query" name:"CatalogId"`
	Visibility string `position:"Query" name:"Visibility"`
	PageSize   int    `position:"Query" name:"PageSize"`
	PageNumber int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeApisRequest) Invoke(client *sdk.Client) (resp *DescribeApisResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApis", "apigateway", "")
	resp = &DescribeApisResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeApisResponse struct {
	responses.BaseResponse
	RequestId   string
	TotalCount  int
	PageSize    int
	PageNumber  int
	ApiSummarys DescribeApisApiSummaryList
}

type DescribeApisApiSummary struct {
	RegionId     string
	GroupId      string
	GroupName    string
	ApiId        string
	ApiName      string
	Visibility   string
	Description  string
	CreatedTime  string
	ModifiedTime string
}

type DescribeApisApiSummaryList []DescribeApisApiSummary

func (list *DescribeApisApiSummaryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApisApiSummary)
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
