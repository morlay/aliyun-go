package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeApisByTrafficControlRequest struct {
	requests.RpcRequest
	TrafficControlId string `position:"Query" name:"TrafficControlId"`
	PageSize         int    `position:"Query" name:"PageSize"`
	PageNumber       int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeApisByTrafficControlRequest) Invoke(client *sdk.Client) (resp *DescribeApisByTrafficControlResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApisByTrafficControl", "apigateway", "")
	resp = &DescribeApisByTrafficControlResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeApisByTrafficControlResponse struct {
	responses.BaseResponse
	RequestId  common.String
	TotalCount common.Integer
	PageSize   common.Integer
	PageNumber common.Integer
	ApiInfos   DescribeApisByTrafficControlApiInfoList
}

type DescribeApisByTrafficControlApiInfo struct {
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

type DescribeApisByTrafficControlApiInfoList []DescribeApisByTrafficControlApiInfo

func (list *DescribeApisByTrafficControlApiInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApisByTrafficControlApiInfo)
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
