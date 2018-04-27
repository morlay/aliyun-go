package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveSingleTaskForCancelingTransferInRequest struct {
	requests.RpcRequest
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

func (req *SaveSingleTaskForCancelingTransferInRequest) Invoke(client *sdk.Client) (resp *SaveSingleTaskForCancelingTransferInResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "SaveSingleTaskForCancelingTransferIn", "", "")
	resp = &SaveSingleTaskForCancelingTransferInResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveSingleTaskForCancelingTransferInResponse struct {
	responses.BaseResponse
	RequestId string
	TaskNo    string
}
