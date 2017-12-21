package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId string
	Success   bool
	Data      DescribeReadOnlyAccountData
}

type DescribeReadOnlyAccountData struct {
	DbName         string
	DrdsInstanceId string
	AccountName    string
}
