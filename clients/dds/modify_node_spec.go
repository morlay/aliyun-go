package dds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyNodeSpecRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	NodeStorage          int    `position:"Query" name:"NodeStorage"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	NodeClass            string `position:"Query" name:"NodeClass"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	NodeId               string `position:"Query" name:"NodeId"`
}

func (r ModifyNodeSpecRequest) Invoke(client *sdk.Client) (response *ModifyNodeSpecResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyNodeSpecRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Dds", "2015-12-01", "ModifyNodeSpec", "dds", "")

	resp := struct {
		*responses.BaseResponse
		ModifyNodeSpecResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyNodeSpecResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyNodeSpecResponse struct {
	RequestId string
	OrderId   string
}
