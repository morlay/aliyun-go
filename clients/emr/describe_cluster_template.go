package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterTemplateRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
}

func (req *DescribeClusterTemplateRequest) Invoke(client *sdk.Client) (resp *DescribeClusterTemplateResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeClusterTemplate", "", "")
	resp = &DescribeClusterTemplateResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClusterTemplateResponse struct {
	responses.BaseResponse
	RequestId    string
	TemplateInfo DescribeClusterTemplateTemplateInfo
}

type DescribeClusterTemplateTemplateInfo struct {
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
	BootstrapActionList    DescribeClusterTemplateBootstrapActionList
	HostGroupList          DescribeClusterTemplateHostGroupList
}

type DescribeClusterTemplateBootstrapAction struct {
	Name string
	Path string
	Arg  string
}

type DescribeClusterTemplateHostGroup struct {
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

type DescribeClusterTemplateBootstrapActionList []DescribeClusterTemplateBootstrapAction

func (list *DescribeClusterTemplateBootstrapActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterTemplateBootstrapAction)
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

type DescribeClusterTemplateHostGroupList []DescribeClusterTemplateHostGroup

func (list *DescribeClusterTemplateHostGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterTemplateHostGroup)
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
