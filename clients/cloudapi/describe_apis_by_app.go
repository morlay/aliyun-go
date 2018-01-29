package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeApisByAppRequest struct {
	requests.RpcRequest
	AppId      int64 `position:"Query" name:"AppId"`
	PageSize   int   `position:"Query" name:"PageSize"`
	PageNumber int   `position:"Query" name:"PageNumber"`
}

func (req *DescribeApisByAppRequest) Invoke(client *sdk.Client) (resp *DescribeApisByAppResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApisByApp", "apigateway", "")
	resp = &DescribeApisByAppResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeApisByAppResponse struct {
	responses.BaseResponse
	RequestId           string
	TotalCount          int
	PageSize            int
	PageNumber          int
	AppApiRelationInfos DescribeApisByAppAppApiRelationInfoList
}

type DescribeApisByAppAppApiRelationInfo struct {
	RegionId            string
	GroupId             string
	GroupName           string
	StageName           string
	Operator            string
	ApiId               string
	ApiName             string
	AuthorizationSource string
	Description         string
	CreatedTime         string
}

type DescribeApisByAppAppApiRelationInfoList []DescribeApisByAppAppApiRelationInfo

func (list *DescribeApisByAppAppApiRelationInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApisByAppAppApiRelationInfo)
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
