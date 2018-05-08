package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribePurchasedApisRequest struct {
	requests.RpcRequest
	StageName  string `position:"Query" name:"StageName"`
	ApiName    string `position:"Query" name:"ApiName"`
	Visibility string `position:"Query" name:"Visibility"`
	GroupId    string `position:"Query" name:"GroupId"`
	PageSize   int    `position:"Query" name:"PageSize"`
	ApiId      string `position:"Query" name:"ApiId"`
	PageNumber int    `position:"Query" name:"PageNumber"`
}

func (req *DescribePurchasedApisRequest) Invoke(client *sdk.Client) (resp *DescribePurchasedApisResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribePurchasedApis", "apigateway", "")
	resp = &DescribePurchasedApisResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribePurchasedApisResponse struct {
	responses.BaseResponse
	RequestId     common.String
	TotalCount    common.Integer
	PageSize      common.Integer
	PageNumber    common.Integer
	PurchasedApis DescribePurchasedApisPurchasedApiList
}

type DescribePurchasedApisPurchasedApi struct {
	RegionId      common.String
	GroupId       common.String
	GroupName     common.String
	ApiId         common.String
	ApiName       common.String
	StageName     common.String
	Description   common.String
	PurchasedTime common.String
}

type DescribePurchasedApisPurchasedApiList []DescribePurchasedApisPurchasedApi

func (list *DescribePurchasedApisPurchasedApiList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePurchasedApisPurchasedApi)
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
