package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDrdsIpWhiteListRequest struct {
	requests.RpcRequest
	Mode           string `position:"Query" name:"Mode"`
	DbName         string `position:"Query" name:"DbName"`
	GroupAttribute string `position:"Query" name:"GroupAttribute"`
	IpWhiteList    string `position:"Query" name:"IpWhiteList"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	GroupName      string `position:"Query" name:"GroupName"`
}

func (req *ModifyDrdsIpWhiteListRequest) Invoke(client *sdk.Client) (resp *ModifyDrdsIpWhiteListResponse, err error) {
	req.InitWithApiInfo("Drds", "2017-10-16", "ModifyDrdsIpWhiteList", "", "")
	resp = &ModifyDrdsIpWhiteListResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyDrdsIpWhiteListResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
}
