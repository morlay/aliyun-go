package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId        string
	TotalCount       string
	PageNumber       string
	PageSize         string
	TemplateInfoList ListClusterTemplatesTemplateInfoList
}

type ListClusterTemplatesTemplateInfo struct {
	Id                     string
	TemplateName           string
	LogEnable              bool
	LogPath                string
	UserId                 string
	UserDefinedEmrEcsRole  string
	MasterNodeTotal        int
	VpcId                  string
	VSwitchId              string
	NetType                string
	IoOptimized            bool
	InstanceGeneration     string
	HighAvailabilityEnable bool
	EasEnable              bool
	BootstrapActionList    ListClusterTemplatesBootstrapActionList
	HostGroupList          ListClusterTemplatesHostGroupList
	SoftwareInfo           ListClusterTemplatesSoftwareInfo
}

type ListClusterTemplatesBootstrapAction struct {
	Name string
	Path string
	Arg  string
}

type ListClusterTemplatesHostGroup struct {
	HostGroupId   string
	HostGroupName string
	HostGroupType string
	ChargeType    string
	Period        string
	NodeCount     int
	InstanceType  string
	DiskType      string
	DiskCapacity  int
	DiskCount     int
}

type ListClusterTemplatesSoftwareInfo struct {
	EmrVer      string
	ClusterType string
	Softwares   ListClusterTemplatesSoftwareList
}

type ListClusterTemplatesSoftware struct {
	DisplayName string
	Name        string
	OnlyDisplay bool
	StartTpe    int
	Version     string
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
