package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeIpControlsRequest struct {
	requests.RpcRequest
	IpControlId   string `position:"Query" name:"IpControlId"`
	IpControlName string `position:"Query" name:"IpControlName"`
	IpControlType string `position:"Query" name:"IpControlType"`
	PageSize      int    `position:"Query" name:"PageSize"`
	PageNumber    int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeIpControlsRequest) Invoke(client *sdk.Client) (resp *DescribeIpControlsResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeIpControls", "apigateway", "")
	resp = &DescribeIpControlsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeIpControlsResponse struct {
	responses.BaseResponse
	RequestId      common.String
	TotalCount     common.Integer
	PageSize       common.Integer
	PageNumber     common.Integer
	IpControlInfos DescribeIpControlsIpControlInfoList
}

type DescribeIpControlsIpControlInfo struct {
	IpControlId   common.String
	IpControlName common.String
	IpControlType common.String
	Description   common.String
	CreateTime    common.String
	ModifiedTime  common.String
	RegionId      common.String
}

type DescribeIpControlsIpControlInfoList []DescribeIpControlsIpControlInfo

func (list *DescribeIpControlsIpControlInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeIpControlsIpControlInfo)
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
