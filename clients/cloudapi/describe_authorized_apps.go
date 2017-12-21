package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeAuthorizedAppsRequest struct {
	GroupId    string `position:"Query" name:"GroupId"`
	StageName  string `position:"Query" name:"StageName"`
	ApiId      string `position:"Query" name:"ApiId"`
	PageNumber int    `position:"Query" name:"PageNumber"`
	PageSize   int    `position:"Query" name:"PageSize"`
}

func (r DescribeAuthorizedAppsRequest) Invoke(client *sdk.Client) (response *DescribeAuthorizedAppsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeAuthorizedAppsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeAuthorizedApps", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		DescribeAuthorizedAppsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeAuthorizedAppsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeAuthorizedAppsResponse struct {
	RequestId      string
	TotalCount     int
	PageSize       int
	PageNumber     int
	AuthorizedApps DescribeAuthorizedAppsAuthorizedAppList
}

type DescribeAuthorizedAppsAuthorizedApp struct {
	StageName           string
	AppId               int64
	AppName             string
	Operator            string
	AuthorizationSource string
	Description         string
	AuthorizedTime      string
}

type DescribeAuthorizedAppsAuthorizedAppList []DescribeAuthorizedAppsAuthorizedApp

func (list *DescribeAuthorizedAppsAuthorizedAppList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAuthorizedAppsAuthorizedApp)
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
