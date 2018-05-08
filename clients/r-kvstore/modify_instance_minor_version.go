package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyInstanceMinorVersionRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Minorversion         string `position:"Query" name:"Minorversion"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyInstanceMinorVersionRequest) Invoke(client *sdk.Client) (resp *ModifyInstanceMinorVersionResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyInstanceMinorVersion", "redisa", "")
	resp = &ModifyInstanceMinorVersionResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyInstanceMinorVersionResponse struct {
	responses.BaseResponse
	RequestId common.String
}
