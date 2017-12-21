package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDrdsIpWhiteListRequest struct {
	Mode           string `position:"Query" name:"Mode"`
	DbName         string `position:"Query" name:"DbName"`
	GroupAttribute string `position:"Query" name:"GroupAttribute"`
	IpWhiteList    string `position:"Query" name:"IpWhiteList"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	GroupName      string `position:"Query" name:"GroupName"`
}

func (r ModifyDrdsIpWhiteListRequest) Invoke(client *sdk.Client) (response *ModifyDrdsIpWhiteListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyDrdsIpWhiteListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Drds", "2017-10-16", "ModifyDrdsIpWhiteList", "", "")

	resp := struct {
		*responses.BaseResponse
		ModifyDrdsIpWhiteListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyDrdsIpWhiteListResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyDrdsIpWhiteListResponse struct {
	RequestId string
	Success   bool
}
