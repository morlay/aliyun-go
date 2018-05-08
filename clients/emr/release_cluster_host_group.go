package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ReleaseClusterHostGroupRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	HostGroupId     string `position:"Query" name:"HostGroupId"`
	InstanceIdList  string `position:"Query" name:"InstanceIdList"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *ReleaseClusterHostGroupRequest) Invoke(client *sdk.Client) (resp *ReleaseClusterHostGroupResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ReleaseClusterHostGroup", "", "")
	resp = &ReleaseClusterHostGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReleaseClusterHostGroupResponse struct {
	responses.BaseResponse
	RequestId common.String
}
