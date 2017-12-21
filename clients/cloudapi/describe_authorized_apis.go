package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeAuthorizedApisRequest struct {
	AppId      int64 `position:"Query" name:"AppId"`
	PageNumber int   `position:"Query" name:"PageNumber"`
	PageSize   int   `position:"Query" name:"PageSize"`
}

func (r DescribeAuthorizedApisRequest) Invoke(client *sdk.Client) (response *DescribeAuthorizedApisResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeAuthorizedApisRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeAuthorizedApis", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		DescribeAuthorizedApisResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeAuthorizedApisResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeAuthorizedApisResponse struct {
	RequestId      string
	TotalCount     int
	PageSize       int
	PageNumber     int
	AuthorizedApis DescribeAuthorizedApisAuthorizedApiList
}

type DescribeAuthorizedApisAuthorizedApi struct {
	RegionId            string
	GroupId             string
	GroupName           string
	StageName           string
	Operator            string
	ApiId               string
	ApiName             string
	AuthorizationSource string
	Description         string
	AuthorizedTime      string
}

type DescribeAuthorizedApisAuthorizedApiList []DescribeAuthorizedApisAuthorizedApi

func (list *DescribeAuthorizedApisAuthorizedApiList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAuthorizedApisAuthorizedApi)
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
