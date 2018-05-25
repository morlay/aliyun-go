package pvtz

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteZoneRecordRequest struct {
	requests.RpcRequest
	RecordId     int64  `position:"Query" name:"RecordId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

func (req *DeleteZoneRecordRequest) Invoke(client *sdk.Client) (resp *DeleteZoneRecordResponse, err error) {
	req.InitWithApiInfo("pvtz", "2018-01-01", "DeleteZoneRecord", "pvtz", "")
	resp = &DeleteZoneRecordResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteZoneRecordResponse struct {
	responses.BaseResponse
	RequestId common.String
	RecordId  common.Long
}
