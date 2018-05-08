package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteDomainRecordRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	RecordId     string `position:"Query" name:"RecordId"`
}

func (req *DeleteDomainRecordRequest) Invoke(client *sdk.Client) (resp *DeleteDomainRecordResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "DeleteDomainRecord", "", "")
	resp = &DeleteDomainRecordResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteDomainRecordResponse struct {
	responses.BaseResponse
	RequestId common.String
	RecordId  common.String
}
