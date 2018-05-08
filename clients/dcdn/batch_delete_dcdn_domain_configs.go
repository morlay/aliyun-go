package dcdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type BatchDeleteDcdnDomainConfigsRequest struct {
	requests.RpcRequest
	FunctionNames string `position:"Query" name:"FunctionNames"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainNames   string `position:"Query" name:"DomainNames"`
	OwnerAccount  string `position:"Query" name:"OwnerAccount"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *BatchDeleteDcdnDomainConfigsRequest) Invoke(client *sdk.Client) (resp *BatchDeleteDcdnDomainConfigsResponse, err error) {
	req.InitWithApiInfo("dcdn", "2018-01-15", "BatchDeleteDcdnDomainConfigs", "dcdn", "")
	resp = &BatchDeleteDcdnDomainConfigsResponse{}
	err = client.DoAction(req, resp)
	return
}

type BatchDeleteDcdnDomainConfigsResponse struct {
	responses.BaseResponse
	RequestId string
}
