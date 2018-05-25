package pvtz

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddZoneRecordRequest struct {
	requests.RpcRequest
	Rr           string `position:"Query" name:"Rr"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	ZoneId       string `position:"Query" name:"ZoneId"`
	Lang         string `position:"Query" name:"Lang"`
	Type         string `position:"Query" name:"Type"`
	Priority     int    `position:"Query" name:"Priority"`
	Ttl          int    `position:"Query" name:"Ttl"`
	Value        string `position:"Query" name:"Value"`
}

func (req *AddZoneRecordRequest) Invoke(client *sdk.Client) (resp *AddZoneRecordResponse, err error) {
	req.InitWithApiInfo("pvtz", "2018-01-01", "AddZoneRecord", "pvtz", "")
	resp = &AddZoneRecordResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddZoneRecordResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	RecordId  common.Long
}
