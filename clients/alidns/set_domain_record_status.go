package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetDomainRecordStatusRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	RecordId     string `position:"Query" name:"RecordId"`
	Status       string `position:"Query" name:"Status"`
}

func (r SetDomainRecordStatusRequest) Invoke(client *sdk.Client) (response *SetDomainRecordStatusResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetDomainRecordStatusRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Alidns", "2015-01-09", "SetDomainRecordStatus", "", "")

	resp := struct {
		*responses.BaseResponse
		SetDomainRecordStatusResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetDomainRecordStatusResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetDomainRecordStatusResponse struct {
	RequestId string
	RecordId  string
	Status    string
}
