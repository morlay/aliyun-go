package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeVpcAccessesRequest struct {
	requests.RpcRequest
}

func (req *DescribeVpcAccessesRequest) Invoke(client *sdk.Client) (resp *DescribeVpcAccessesResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeVpcAccesses", "apigateway", "")
	resp = &DescribeVpcAccessesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeVpcAccessesResponse struct {
	responses.BaseResponse
	RequestId           common.String
	TotalCount          common.Integer
	PageSize            common.Integer
	PageNumber          common.Integer
	VpcAccessAttributes DescribeVpcAccessesVpcAccessAttributeList
}

type DescribeVpcAccessesVpcAccessAttribute struct {
	VpcId       common.String
	InstanceId  common.String
	Port        common.Integer
	Name        common.String
	CreatedTime common.String
}

type DescribeVpcAccessesVpcAccessAttributeList []DescribeVpcAccessesVpcAccessAttribute

func (list *DescribeVpcAccessesVpcAccessAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVpcAccessesVpcAccessAttribute)
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
