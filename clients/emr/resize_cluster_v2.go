package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ResizeClusterV2Request struct {
	requests.RpcRequest
	HostGroups *ResizeClusterV2HostGroupList `position:"Query" type:"Repeated" name:"HostGroup"`
	ClusterId  string                        `position:"Query" name:"ClusterId"`
}

func (req *ResizeClusterV2Request) Invoke(client *sdk.Client) (resp *ResizeClusterV2Response, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ResizeClusterV2", "", "")
	resp = &ResizeClusterV2Response{}
	err = client.DoAction(req, resp)
	return
}

type ResizeClusterV2HostGroup struct {
	ClusterId       string `name:"ClusterId"`
	HostGroupId     string `name:"HostGroupId"`
	HostGroupName   string `name:"HostGroupName"`
	HostGroupType   string `name:"HostGroupType"`
	Comment         string `name:"Comment"`
	CreateType      string `name:"CreateType"`
	ChargeType      string `name:"ChargeType"`
	Period          int    `name:"Period"`
	NodeCount       int    `name:"NodeCount"`
	InstanceType    string `name:"InstanceType"`
	DiskType        string `name:"DiskType"`
	DiskCapacity    int    `name:"DiskCapacity"`
	DiskCount       int    `name:"DiskCount"`
	SysDiskType     string `name:"SysDiskType"`
	SysDiskCapacity int    `name:"SysDiskCapacity"`
	AutoRenew       string `name:"AutoRenew"`
	VswitchId       int    `name:"VswitchId"`
}

type ResizeClusterV2Response struct {
	responses.BaseResponse
	RequestId common.String
	ClusterId common.String
}

type ResizeClusterV2HostGroupList []ResizeClusterV2HostGroup

func (list *ResizeClusterV2HostGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ResizeClusterV2HostGroup)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
