package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeVpcZoneNosRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Region               string `position:"Query" name:"Region"`
}

func (r DescribeVpcZoneNosRequest) Invoke(client *sdk.Client) (response *DescribeVpcZoneNosResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeVpcZoneNosRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeVpcZoneNos", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeVpcZoneNosResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeVpcZoneNosResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeVpcZoneNosResponse struct {
	RequestId string
	Items     DescribeVpcZoneNosVpcZoneIdList
}

type DescribeVpcZoneNosVpcZoneId struct {
	ZoneId    string
	Region    string
	SubDomain string
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
