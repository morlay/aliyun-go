package pvtz

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetZoneRecordStatusRequest struct {
	requests.RpcRequest
	RecordId     int64  `position:"Query" name:"RecordId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	Status       string `position:"Query" name:"Status"`
}

func (req *SetZoneRecordStatusRequest) Invoke(client *sdk.Client) (resp *SetZoneRecordStatusResponse, err error) {
	req.InitWithApiInfo("pvtz", "2018-01-01", "SetZoneRecordStatus", "pvtz", "")
	resp = &SetZoneRecordStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetZoneRecordStatusResponse struct {
	responses.BaseResponse
	RequestId common.String
	RecordId  common.Long
	Status    common.String
}
