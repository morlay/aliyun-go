package pvtz

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeZoneRecordsRequest struct {
	requests.RpcRequest
	PageSize     int    `position:"Query" name:"PageSize"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	ZoneId       string `position:"Query" name:"ZoneId"`
	Tag          string `position:"Query" name:"Tag"`
	Keyword      string `position:"Query" name:"Keyword"`
	Lang         string `position:"Query" name:"Lang"`
	PageNumber   int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeZoneRecordsRequest) Invoke(client *sdk.Client) (resp *DescribeZoneRecordsResponse, err error) {
	req.InitWithApiInfo("pvtz", "2018-01-01", "DescribeZoneRecords", "pvtz", "")
	resp = &DescribeZoneRecordsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeZoneRecordsResponse struct {
	responses.BaseResponse
	RequestId  common.String
	TotalItems common.Integer
	TotalPages common.Integer
	PageSize   common.Integer
	PageNumber common.Integer
	Records    DescribeZoneRecordsRecordList
}

type DescribeZoneRecordsRecord struct {
	RecordId common.Long
	Rr       common.String
	Type     common.String
	Ttl      common.Integer
	Priority common.Integer
	Value    common.String
	Status   common.String
}

type DescribeZoneRecordsRecordList []DescribeZoneRecordsRecord

func (list *DescribeZoneRecordsRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeZoneRecordsRecord)
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
