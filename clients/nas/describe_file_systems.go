package nas

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId   common.String
	TotalCount  common.Integer
	PageSize    common.Integer
	PageNumber  common.Integer
	FileSystems DescribeFileSystemsFileSystemList
}

type DescribeFileSystemsFileSystem struct {
	FileSystemId common.String
	Destription  common.String
	CreateTime   common.String
	RegionId     common.String
	ProtocolType common.String
	StorageType  common.String
	MeteredSize  common.Long
	MountTargets DescribeFileSystemsMountTargetList
	Packages     DescribeFileSystems_PackageList
}

type DescribeFileSystemsMountTarget struct {
	MountTargetDomain common.String
}

type DescribeFileSystems_Package struct {
	PackageId common.String
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
