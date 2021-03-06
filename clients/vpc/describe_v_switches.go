package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeVSwitchesRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	VpcId                string `position:"Query" name:"VpcId"`
	VSwitchName          string `position:"Query" name:"VSwitchName"`
	PageSize             int    `position:"Query" name:"PageSize"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	IsDefault            string `position:"Query" name:"IsDefault"`
}

func (req *DescribeVSwitchesRequest) Invoke(client *sdk.Client) (resp *DescribeVSwitchesResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeVSwitches", "vpc", "")
	resp = &DescribeVSwitchesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeVSwitchesResponse struct {
	responses.BaseResponse
	RequestId  common.String
	TotalCount common.Integer
	PageNumber common.Integer
	PageSize   common.Integer
	VSwitches  DescribeVSwitchesVSwitchList
}

type DescribeVSwitchesVSwitch struct {
	VSwitchId               common.String
	VpcId                   common.String
	Status                  common.String
	CidrBlock               common.String
	ZoneId                  common.String
	AvailableIpAddressCount common.Long
	Description             common.String
	VSwitchName             common.String
	CreationTime            common.String
	IsDefault               bool
}

type DescribeVSwitchesVSwitchList []DescribeVSwitchesVSwitch

func (list *DescribeVSwitchesVSwitchList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVSwitchesVSwitch)
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
