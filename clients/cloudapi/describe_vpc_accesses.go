package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeVpcAccessesRequest struct {
}

func (r DescribeVpcAccessesRequest) Invoke(client *sdk.Client) (response *DescribeVpcAccessesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeVpcAccessesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeVpcAccesses", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		DescribeVpcAccessesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeVpcAccessesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeVpcAccessesResponse struct {
	RequestId           string
	TotalCount          int
	PageSize            int
	PageNumber          int
	VpcAccessAttributes DescribeVpcAccessesVpcAccessAttributeList
}

type DescribeVpcAccessesVpcAccessAttribute struct {
	VpcId       string
	InstanceId  string
	Port        int
	Name        string
	CreatedTime string
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
