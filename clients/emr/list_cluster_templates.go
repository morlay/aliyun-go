package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListClusterTemplatesRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	BizId           string `position:"Query" name:"BizId"`
	PageSize        int    `position:"Query" name:"PageSize"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
}

func (req *ListClusterTemplatesRequest) Invoke(client *sdk.Client) (resp *ListClusterTemplatesResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListClusterTemplates", "", "")
	resp = &ListClusterTemplatesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListClusterTemplatesResponse struct {
	responses.BaseResponse
	RequestId        common.String
	TotalCount       common.String
	PageNumber       common.String
	PageSize         common.String
	TemplateInfoList ListClusterTemplatesTemplateInfoList
}

type ListClusterTemplatesTemplateInfo struct {
	Id                     common.String
	TemplateName           common.String
	LogEnable              bool
	LogPath                common.String
	UserId                 common.String
	UserDefinedEmrEcsRole  common.String
	MasterNodeTotal        common.Integer
	VpcId                  common.String
	VSwitchId              common.String
	NetType                common.String
	IoOptimized            bool
	InstanceGeneration     common.String
	HighAvailabilityEnable bool
	EasEnable              bool
	BootstrapActionList    ListClusterTemplatesBootstrapActionList
	HostGroupList          ListClusterTemplatesHostGroupList
	SoftwareInfo           ListClusterTemplatesSoftwareInfo
}

type ListClusterTemplatesBootstrapAction struct {
	Name common.String
	Path common.String
	Arg  common.String
}

type ListClusterTemplatesHostGroup struct {
	HostGroupId   common.String
	HostGroupName common.String
	HostGroupType common.String
	ChargeType    common.String
	Period        common.String
	NodeCount     common.Integer
	InstanceType  common.String
	DiskType      common.String
	DiskCapacity  common.Integer
	DiskCount     common.Integer
}

type ListClusterTemplatesSoftwareInfo struct {
	EmrVer      common.String
	ClusterType common.String
	Softwares   ListClusterTemplatesSoftwareList
}

type ListClusterTemplatesSoftware struct {
	DisplayName common.String
	Name        common.String
	OnlyDisplay bool
	StartTpe    common.Integer
	Version     common.String
}

type ListClusterTemplatesTemplateInfoList []ListClusterTemplatesTemplateInfo

func (list *ListClusterTemplatesTemplateInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterTemplatesTemplateInfo)
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

type ListClusterTemplatesBootstrapActionList []ListClusterTemplatesBootstrapAction

func (list *ListClusterTemplatesBootstrapActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterTemplatesBootstrapAction)
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

type ListClusterTemplatesHostGroupList []ListClusterTemplatesHostGroup

func (list *ListClusterTemplatesHostGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterTemplatesHostGroup)
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

type ListClusterTemplatesSoftwareList []ListClusterTemplatesSoftware

func (list *ListClusterTemplatesSoftwareList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterTemplatesSoftware)
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
