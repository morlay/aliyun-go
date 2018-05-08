package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId   common.String
	TotalCount  common.Integer
	PageSize    common.Integer
	PageNumber  common.Integer
	ApiHisItems DescribeHistoryApisApiHisItemList
}

type DescribeHistoryApisApiHisItem struct {
	RegionId       common.String
	ApiId          common.String
	ApiName        common.String
	GroupId        common.String
	GroupName      common.String
	StageName      common.String
	HistoryVersion common.String
	Status         DescribeHistoryApisStatus
	Description    common.String
	DeployedTime   common.String
}

type DescribeHistoryApisStatus struct {
	StringValue common.String
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
