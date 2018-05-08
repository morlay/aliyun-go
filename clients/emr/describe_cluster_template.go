package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId    common.String
	TemplateInfo DescribeClusterTemplateTemplateInfo
}

type DescribeClusterTemplateTemplateInfo struct {
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
	BootstrapActionList    DescribeClusterTemplateBootstrapActionList
	HostGroupList          DescribeClusterTemplateHostGroupList
}

type DescribeClusterTemplateBootstrapAction struct {
	Name common.String
	Path common.String
	Arg  common.String
}

type DescribeClusterTemplateHostGroup struct {
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
