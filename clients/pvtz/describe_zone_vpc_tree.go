package pvtz

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeZoneVpcTreeRequest struct {
	requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

func (req *DescribeZoneVpcTreeRequest) Invoke(client *sdk.Client) (resp *DescribeZoneVpcTreeResponse, err error) {
	req.InitWithApiInfo("pvtz", "2018-01-01", "DescribeZoneVpcTree", "pvtz", "")
	resp = &DescribeZoneVpcTreeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeZoneVpcTreeResponse struct {
	responses.BaseResponse
	RequestId common.String
	Zones     DescribeZoneVpcTreeZoneList
}

type DescribeZoneVpcTreeZone struct {
	ZoneId          common.String
	ZoneName        common.String
	Remark          common.String
	RecordCount     common.Integer
	CreateTime      common.String
	CreateTimestamp common.Long
	UpdateTime      common.String
	UpdateTimestamp common.Long
	IsPtr           bool
	Vpcs            DescribeZoneVpcTreeVpcList
}

type DescribeZoneVpcTreeVpc struct {
	RegionId   common.String
	RegionName common.String
	VpcId      common.String
	VpcName    common.String
}

type DescribeZoneVpcTreeZoneList []DescribeZoneVpcTreeZone

func (list *DescribeZoneVpcTreeZoneList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeZoneVpcTreeZone)
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

type DescribeZoneVpcTreeVpcList []DescribeZoneVpcTreeVpc

func (list *DescribeZoneVpcTreeVpcList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeZoneVpcTreeVpc)
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
