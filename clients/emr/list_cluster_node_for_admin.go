package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId       string
	ClusterNodeList ListClusterNodeForAdminClusterNodeList
}

type ListClusterNodeForAdminClusterNode struct {
	ClusterId       string
	CpuCore         string
	Daemons         string
	DiskDevices     string
	DiskInfo        string
	DiskType        string
	GmtCreate       string
	GmtModified     string
	HostName        string
	Id              string
	ImageId         string
	InnerIpAddress  string
	InstanceId      string
	InstanceType    string
	IsMaster        string
	MemCapacity     string
	Payment         string
	PublicIpAddress string
	RegionId        string
	SecurityGroupId string
	SerialNumber    string
	Status          string
	UserId          string
	ZoneId          string
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
