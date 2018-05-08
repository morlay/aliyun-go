package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeReadOnlyAccountRequest struct {
	requests.RpcRequest
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (req *DescribeReadOnlyAccountRequest) Invoke(client *sdk.Client) (resp *DescribeReadOnlyAccountResponse, err error) {
	req.InitWithApiInfo("Drds", "2017-10-16", "DescribeReadOnlyAccount", "", "")
	resp = &DescribeReadOnlyAccountResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeReadOnlyAccountResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Data      DescribeReadOnlyAccountData
}

type DescribeReadOnlyAccountData struct {
	DbName         common.String
	DrdsInstanceId common.String
	AccountName    common.String
}
