package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeApisByIpControlRequest struct {
	requests.RpcRequest
	IpControlId string `position:"Query" name:"IpControlId"`
	PageSize    int    `position:"Query" name:"PageSize"`
	PageNumber  int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeApisByIpControlRequest) Invoke(client *sdk.Client) (resp *DescribeApisByIpControlResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApisByIpControl", "apigateway", "")
	resp = &DescribeApisByIpControlResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeApisByIpControlResponse struct {
	responses.BaseResponse
	RequestId  common.String
	TotalCount common.Integer
	PageSize   common.Integer
	PageNumber common.Integer
	ApiInfos   DescribeApisByIpControlApiInfoList
}

type DescribeApisByIpControlApiInfo struct {
	RegionId    common.String
	GroupId     common.String
	GroupName   common.String
	StageName   common.String
	ApiId       common.String
	ApiName     common.String
	Description common.String
	Visibility  common.String
	BoundTime   common.String
}

type DescribeApisByIpControlApiInfoList []DescribeApisByIpControlApiInfo

func (list *DescribeApisByIpControlApiInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApisByIpControlApiInfo)
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
