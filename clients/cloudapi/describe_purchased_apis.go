package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId     string
	TotalCount    int
	PageSize      int
	PageNumber    int
	PurchasedApis DescribePurchasedApisPurchasedApiList
}

type DescribePurchasedApisPurchasedApi struct {
	RegionId      string
	GroupId       string
	GroupName     string
	ApiId         string
	ApiName       string
	StageName     string
	Description   string
	PurchasedTime string
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
