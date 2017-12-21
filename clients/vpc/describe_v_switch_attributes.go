package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeVSwitchAttributesRequest struct {
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeVSwitchAttributesRequest) Invoke(client *sdk.Client) (response *DescribeVSwitchAttributesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeVSwitchAttributesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeVSwitchAttributes", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		DescribeVSwitchAttributesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeVSwitchAttributesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeVSwitchAttributesResponse struct {
	RequestId               string
	VSwitchId               string
	VpcId                   string
	Status                  string
	CidrBlock               string
	ZoneId                  string
	AvailableIpAddressCount int64
	Description             string
	VSwitchName             string
	CreationTime            string
	IsDefault               bool
	CloudResources          DescribeVSwitchAttributesCloudResourceSetTypeList
}

type DescribeVSwitchAttributesCloudResourceSetType struct {
	ResourceType  string
	ResourceCount int
}

type DescribeVSwitchAttributesCloudResourceSetTypeList []DescribeVSwitchAttributesCloudResourceSetType

func (list *DescribeVSwitchAttributesCloudResourceSetTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVSwitchAttributesCloudResourceSetType)
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
