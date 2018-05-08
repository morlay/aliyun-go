package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DecryptBizIdRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	Type            string `position:"Query" name:"Type"`
}

func (req *DecryptBizIdRequest) Invoke(client *sdk.Client) (resp *DecryptBizIdResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DecryptBizId", "", "")
	resp = &DecryptBizIdResponse{}
	err = client.DoAction(req, resp)
	return
}

type DecryptBizIdResponse struct {
	responses.BaseResponse
	RequestId common.String
	BizId     common.String
}
