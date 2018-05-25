package pvtz

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeZonesRequest struct {
	requests.RpcRequest
	PageSize     int    `position:"Query" name:"PageSize"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	Keyword      string `position:"Query" name:"Keyword"`
	PageNumber   int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeZonesRequest) Invoke(client *sdk.Client) (resp *DescribeZonesResponse, err error) {
	req.InitWithApiInfo("pvtz", "2018-01-01", "DescribeZones", "pvtz", "")
	resp = &DescribeZonesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeZonesResponse struct {
	responses.BaseResponse
	RequestId  common.String
	TotalItems common.Integer
	TotalPages common.Integer
	PageSize   common.Integer
	PageNumber common.Integer
	Zones      DescribeZonesZoneList
}

type DescribeZonesZone struct {
	ZoneId          common.String
	ZoneName        common.String
	Remark          common.String
	RecordCount     common.Integer
	CreateTime      common.String
	CreateTimestamp common.Long
	UpdateTime      common.String
	UpdateTimestamp common.Long
	IsPtr           bool
}

type DescribeZonesZoneList []DescribeZonesZone

func (list *DescribeZonesZoneList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeZonesZone)
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
