package nas

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeFileSystemsRequest struct {
	requests.RpcRequest
	PageSize     int    `position:"Query" name:"PageSize"`
	PageNumber   int    `position:"Query" name:"PageNumber"`
	FileSystemId string `position:"Query" name:"FileSystemId"`
}

func (req *DescribeFileSystemsRequest) Invoke(client *sdk.Client) (resp *DescribeFileSystemsResponse, err error) {
	req.InitWithApiInfo("NAS", "2017-06-26", "DescribeFileSystems", "nas", "")
	resp = &DescribeFileSystemsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeFileSystemsResponse struct {
	responses.BaseResponse
	RequestId   string
	TotalCount  int
	PageSize    int
	PageNumber  int
	FileSystems DescribeFileSystemsFileSystemList
}

type DescribeFileSystemsFileSystem struct {
	FileSystemId string
	Destription  string
	CreateTime   string
	RegionId     string
	ProtocolType string
	StorageType  string
	MeteredSize  int64
	MountTargets DescribeFileSystemsMountTargetList
	Packages     DescribeFileSystems_PackageList
}

type DescribeFileSystemsMountTarget struct {
	MountTargetDomain string
}

type DescribeFileSystems_Package struct {
	PackageId string
}

type DescribeFileSystemsFileSystemList []DescribeFileSystemsFileSystem

func (list *DescribeFileSystemsFileSystemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeFileSystemsFileSystem)
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

type DescribeFileSystemsMountTargetList []DescribeFileSystemsMountTarget

func (list *DescribeFileSystemsMountTargetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeFileSystemsMountTarget)
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

type DescribeFileSystems_PackageList []DescribeFileSystems_Package

func (list *DescribeFileSystems_PackageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeFileSystems_Package)
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
