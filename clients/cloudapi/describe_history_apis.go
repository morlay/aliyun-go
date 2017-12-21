package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeHistoryApisRequest struct {
	requests.RpcRequest
	GroupId    string `position:"Query" name:"GroupId"`
	StageName  string `position:"Query" name:"StageName"`
	ApiId      string `position:"Query" name:"ApiId"`
	ApiName    string `position:"Query" name:"ApiName"`
	PageSize   string `position:"Query" name:"PageSize"`
	PageNumber string `position:"Query" name:"PageNumber"`
}

func (req *DescribeHistoryApisRequest) Invoke(client *sdk.Client) (resp *DescribeHistoryApisResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeHistoryApis", "apigateway", "")
	resp = &DescribeHistoryApisResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeHistoryApisResponse struct {
	responses.BaseResponse
	RequestId   string
	TotalCount  int
	PageSize    int
	PageNumber  int
	ApiHisItems DescribeHistoryApisApiHisItemList
}

type DescribeHistoryApisApiHisItem struct {
	RegionId       string
	ApiId          string
	ApiName        string
	GroupId        string
	GroupName      string
	StageName      string
	HistoryVersion string
	Status         DescribeHistoryApisStatus
	Description    string
	DeployedTime   string
}

type DescribeHistoryApisStatus struct {
	StringValue string
}

type DescribeHistoryApisApiHisItemList []DescribeHistoryApisApiHisItem

func (list *DescribeHistoryApisApiHisItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeHistoryApisApiHisItem)
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
