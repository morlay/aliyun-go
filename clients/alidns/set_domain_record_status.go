package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetDomainRecordStatusRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	RecordId     string `position:"Query" name:"RecordId"`
	Status       string `position:"Query" name:"Status"`
}

func (req *SetDomainRecordStatusRequest) Invoke(client *sdk.Client) (resp *SetDomainRecordStatusResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "SetDomainRecordStatus", "", "")
	resp = &SetDomainRecordStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetDomainRecordStatusResponse struct {
	responses.BaseResponse
	RequestId string
	RecordId  string
	Status    string
}
