package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListClusterNodeForAdminRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64                                  `position:"Query" name:"ResourceOwnerId"`
	StatusLists     *ListClusterNodeForAdminStatusListList `position:"Query" type:"Repeated" name:"StatusList"`
	PageSize        int                                    `position:"Query" name:"PageSize"`
	ClusterId       string                                 `position:"Query" name:"ClusterId"`
	UserId          string                                 `position:"Query" name:"UserId"`
	PageNumber      int                                    `position:"Query" name:"PageNumber"`
}

func (req *ListClusterNodeForAdminRequest) Invoke(client *sdk.Client) (resp *ListClusterNodeForAdminResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListClusterNodeForAdmin", "", "")
	resp = &ListClusterNodeForAdminResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListClusterNodeForAdminResponse struct {
	responses.BaseResponse
	RequestId       common.String
	ClusterNodeList ListClusterNodeForAdminClusterNodeList
}

type ListClusterNodeForAdminClusterNode struct {
	ClusterId       common.String
	CpuCore         common.String
	Daemons         common.String
	DiskDevices     common.String
	DiskInfo        common.String
	DiskType        common.String
	GmtCreate       common.String
	GmtModified     common.String
	HostName        common.String
	Id              common.String
	ImageId         common.String
	InnerIpAddress  common.String
	InstanceId      common.String
	InstanceType    common.String
	IsMaster        common.String
	MemCapacity     common.String
	Payment         common.String
	PublicIpAddress common.String
	RegionId        common.String
	SecurityGroupId common.String
	SerialNumber    common.String
	Status          common.String
	UserId          common.String
	ZoneId          common.String
}

type ListClusterNodeForAdminStatusListList []string

func (list *ListClusterNodeForAdminStatusListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type ListClusterNodeForAdminClusterNodeList []ListClusterNodeForAdminClusterNode

func (list *ListClusterNodeForAdminClusterNodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterNodeForAdminClusterNode)
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
