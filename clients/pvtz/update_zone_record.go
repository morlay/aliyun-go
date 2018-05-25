package pvtz

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UpdateZoneRecordRequest struct {
	requests.RpcRequest
	Rr           string `position:"Query" name:"Rr"`
	RecordId     int64  `position:"Query" name:"RecordId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	Type         string `position:"Query" name:"Type"`
	Priority     int    `position:"Query" name:"Priority"`
	Ttl          int    `position:"Query" name:"Ttl"`
	Value        string `position:"Query" name:"Value"`
}

func (req *UpdateZoneRecordRequest) Invoke(client *sdk.Client) (resp *UpdateZoneRecordResponse, err error) {
	req.InitWithApiInfo("pvtz", "2018-01-01", "UpdateZoneRecord", "pvtz", "")
	resp = &UpdateZoneRecordResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateZoneRecordResponse struct {
	responses.BaseResponse
	RequestId common.String
	RecordId  common.Long
}
