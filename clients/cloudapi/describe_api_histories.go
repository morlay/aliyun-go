package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId   common.String
	TotalCount  common.Integer
	PageSize    common.Integer
	PageNumber  common.Integer
	ApiHisItems DescribeApiHistoriesApiHisItemList
}

type DescribeApiHistoriesApiHisItem struct {
	RegionId       common.String
	ApiId          common.String
	ApiName        common.String
	GroupId        common.String
	GroupName      common.String
	StageName      common.String
	HistoryVersion common.String
	Status         common.String
	Description    common.String
	DeployedTime   common.String
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
