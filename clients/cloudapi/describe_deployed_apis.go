package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDeployedApisRequest struct {
	requests.RpcRequest
	GroupId    string `position:"Query" name:"GroupId"`
	StageName  string `position:"Query" name:"StageName"`
	ApiId      string `position:"Query" name:"ApiId"`
	ApiName    string `position:"Query" name:"ApiName"`
	PageNumber int    `position:"Query" name:"PageNumber"`
	PageSize   int    `position:"Query" name:"PageSize"`
}

func (req *DescribeDeployedApisRequest) Invoke(client *sdk.Client) (resp *DescribeDeployedApisResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeDeployedApis", "apigateway", "")
	resp = &DescribeDeployedApisResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDeployedApisResponse struct {
	responses.BaseResponse
	RequestId    common.String
	TotalCount   common.Integer
	PageSize     common.Integer
	PageNumber   common.Integer
	DeployedApis DescribeDeployedApisDeployedApiItemList
}

type DescribeDeployedApisDeployedApiItem struct {
	RegionId     common.String
	ApiId        common.String
	ApiName      common.String
	GroupId      common.String
	GroupName    common.String
	StageName    common.String
	Visibility   common.String
	Description  common.String
	DeployedTime common.String
}

type DescribeDeployedApisDeployedApiItemList []DescribeDeployedApisDeployedApiItem

func (list *DescribeDeployedApisDeployedApiItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDeployedApisDeployedApiItem)
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
