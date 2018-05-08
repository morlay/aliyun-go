package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeVpcZoneNosRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Region               string `position:"Query" name:"Region"`
}

func (req *DescribeVpcZoneNosRequest) Invoke(client *sdk.Client) (resp *DescribeVpcZoneNosResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeVpcZoneNos", "rds", "")
	resp = &DescribeVpcZoneNosResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeVpcZoneNosResponse struct {
	responses.BaseResponse
	RequestId common.String
	Items     DescribeVpcZoneNosVpcZoneIdList
}

type DescribeVpcZoneNosVpcZoneId struct {
	ZoneId    common.String
	Region    common.String
	SubDomain common.String
}

type DescribeVpcZoneNosVpcZoneIdList []DescribeVpcZoneNosVpcZoneId

func (list *DescribeVpcZoneNosVpcZoneIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVpcZoneNosVpcZoneId)
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
