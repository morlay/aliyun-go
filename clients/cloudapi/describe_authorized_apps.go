package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeAuthorizedAppsRequest struct {
	requests.RpcRequest
	GroupId    string `position:"Query" name:"GroupId"`
	StageName  string `position:"Query" name:"StageName"`
	ApiId      string `position:"Query" name:"ApiId"`
	PageNumber int    `position:"Query" name:"PageNumber"`
	PageSize   int    `position:"Query" name:"PageSize"`
}

func (req *DescribeAuthorizedAppsRequest) Invoke(client *sdk.Client) (resp *DescribeAuthorizedAppsResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeAuthorizedApps", "apigateway", "")
	resp = &DescribeAuthorizedAppsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeAuthorizedAppsResponse struct {
	responses.BaseResponse
	RequestId      common.String
	TotalCount     common.Integer
	PageSize       common.Integer
	PageNumber     common.Integer
	AuthorizedApps DescribeAuthorizedAppsAuthorizedAppList
}

type DescribeAuthorizedAppsAuthorizedApp struct {
	StageName           common.String
	AppId               common.Long
	AppName             common.String
	Operator            common.String
	AuthorizationSource common.String
	Description         common.String
	AuthorizedTime      common.String
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
