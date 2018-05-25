package pvtz

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeZoneInfoRequest struct {
	requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	ZoneId       string `position:"Query" name:"ZoneId"`
	Lang         string `position:"Query" name:"Lang"`
}

func (req *DescribeZoneInfoRequest) Invoke(client *sdk.Client) (resp *DescribeZoneInfoResponse, err error) {
	req.InitWithApiInfo("pvtz", "2018-01-01", "DescribeZoneInfo", "pvtz", "")
	resp = &DescribeZoneInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeZoneInfoResponse struct {
	responses.BaseResponse
	RequestId       common.String
	ZoneId          common.String
	ZoneName        common.String
	Remark          common.String
	RecordCount     common.Integer
	CreateTime      common.String
	CreateTimestamp common.Long
	UpdateTime      common.String
	UpdateTimestamp common.Long
	IsPtr           bool
	BindVpcs        DescribeZoneInfoVpcList
}

type DescribeZoneInfoVpc struct {
	ReionId    common.String
	VpcId      common.String
	VpcName    common.String
	RegionName common.String
}

type DescribeZoneInfoVpcList []DescribeZoneInfoVpc

func (list *DescribeZoneInfoVpcList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeZoneInfoVpc)
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
