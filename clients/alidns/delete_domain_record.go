package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteDomainRecordRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	RecordId     string `position:"Query" name:"RecordId"`
}

func (r DeleteDomainRecordRequest) Invoke(client *sdk.Client) (response *DeleteDomainRecordResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteDomainRecordRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Alidns", "2015-01-09", "DeleteDomainRecord", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteDomainRecordResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteDomainRecordResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteDomainRecordResponse struct {
	RequestId string
	RecordId  string
}
