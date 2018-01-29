package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeApiHistoriesRequest struct {
	requests.RpcRequest
	StageName  string `position:"Query" name:"StageName"`
	ApiName    string `position:"Query" name:"ApiName"`
	GroupId    string `position:"Query" name:"GroupId"`
	PageSize   string `position:"Query" name:"PageSize"`
	ApiId      string `position:"Query" name:"ApiId"`
	PageNumber string `position:"Query" name:"PageNumber"`
}

func (req *DescribeApiHistoriesRequest) Invoke(client *sdk.Client) (resp *DescribeApiHistoriesResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApiHistories", "apigateway", "")
	resp = &DescribeApiHistoriesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeApiHistoriesResponse struct {
	responses.BaseResponse
	RequestId   string
	TotalCount  int
	PageSize    int
	PageNumber  int
	ApiHisItems DescribeApiHistoriesApiHisItemList
}

type DescribeApiHistoriesApiHisItem struct {
	RegionId       string
	ApiId          string
	ApiName        string
	GroupId        string
	GroupName      string
	StageName      string
	HistoryVersion string
	Status         string
	Description    string
	DeployedTime   string
}

type DescribeApiHistoriesApiHisItemList []DescribeApiHistoriesApiHisItem

func (list *DescribeApiHistoriesApiHisItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiHistoriesApiHisItem)
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
