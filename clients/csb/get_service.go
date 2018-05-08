package csb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetServiceRequest struct {
	requests.RpcRequest
	CsbId     int64 `position:"Query" name:"CsbId"`
	ServiceId int64 `position:"Query" name:"ServiceId"`
}

func (req *GetServiceRequest) Invoke(client *sdk.Client) (resp *GetServiceResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "GetService", "CSB", "")
	resp = &GetServiceResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetServiceResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Message   common.String
	RequestId common.String
	Data      GetServiceData
}

type GetServiceData struct {
	Service GetServiceService
}

type GetServiceService struct {
	AccessParamsJSON       common.String
	Active                 bool
	Alias                  common.String
	AllVisiable            bool
	ApproveUserId          common.String
	CasTargets             common.String
	ConsumeTypesJSON       common.String
	CreateTime             common.Long
	CsbId                  common.Long
	ErrDefJSON             common.String
	Id                     common.Long
	InterfaceName          common.String
	IpBlackStr             common.String
	IpWhiteStr             common.String
	ModelVersion           common.String
	ModifiedTime           common.Long
	OldVersion             common.String
	OpenRestfulPath        common.String
	OttFlag                bool
	OwnerId                common.String
	PolicyHandler          common.String
	PrincipalName          common.String
	ProjectId              common.Long
	ProjectName            common.String
	ProvideType            common.String
	RouteConfJson          common.String
	SSL                    bool
	Scope                  common.String
	ServiceName            common.String
	ServiceOpenRestfulPath common.String
	ServiceProviderType    common.String
	ServiceVersion         common.String
	SkipAuth               bool
	StatisticName          common.String
	Status                 common.Integer
	UserId                 common.String
	ValidConsumeTypes      bool
	ValidProvideType       bool
	ServiceVersionsList    GetServiceServiceVersionList
	VisiableGroupList      GetServiceVisiableGroupList
	CasServTargets         GetServiceCasServTargetList
	ConsumeTypes           GetServiceConsumeTypeList
	RouteConf              GetServiceRouteConf
}

type GetServiceServiceVersion struct {
	Active            bool
	AllVisiable       bool
	Id                common.Long
	OldVersion        common.String
	OttFlag           bool
	SSL               bool
	Scope             common.String
	ServiceVersion    common.String
	SkipAuth          bool
	StatisticName     common.String
	Status            common.Integer
	ValidConsumeTypes bool
	ValidProvideType  bool
}

type GetServiceVisiableGroup struct {
	Id           common.Long
	GroupId      common.Long
	UserId       common.String
	ServiceId    common.Long
	CreateTime   common.Long
	ModifiedTime common.Long
	Status       common.Integer
}

type GetServiceRouteConf struct {
	ServiceRouteStrategy common.String
	ImportConf           GetServiceImportConf
	ImportConfs          GetServiceImportConfs
}

type GetServiceImportConf struct {
	AccessEndpointJSON common.String
	ProvideType        common.String
	InputParameterMap  GetServiceInputParameterList
	OutputParameterMap GetServiceOutputParameterList
}

type GetServiceInputParameter struct {
	CatType      common.Integer
	Depth        common.Integer
	ExtType      common.Integer
	MapStyle     common.Integer
	Optional     bool
	OriginalName common.String
	ParamType    common.String
	PassMethod   common.String
	TargetName   common.String
}

type GetServiceOutputParameter struct {
	CatType      common.Integer
	Depth        common.Integer
	ExtType      common.Integer
	MapStyle     common.Integer
	Optional     bool
	OriginalName common.String
	ParamType    common.String
	PassMethod   common.String
	TargetName   common.String
}

type GetServiceImportConfs struct {
	AccessEndpointJSON  common.String
	ProvideType         common.String
	InputParameterMap1  GetServiceInputParameter3List
	OutputParameterMap2 GetServiceOutputParameter4List
}

type GetServiceInputParameter3 struct {
	CatType      common.Integer
	Depth        common.Integer
	ExtType      common.Integer
	MapStyle     common.Integer
	Optional     bool
	OriginalName common.String
	ParamType    common.String
	PassMethod   common.String
	TargetName   common.String
}

type GetServiceOutputParameter4 struct {
	CatType      common.Integer
	Depth        common.Integer
	ExtType      common.Integer
	MapStyle     common.Integer
	Optional     bool
	OriginalName common.String
	ParamType    common.String
	PassMethod   common.String
	TargetName   common.String
}

type GetServiceServiceVersionList []GetServiceServiceVersion

func (list *GetServiceServiceVersionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetServiceServiceVersion)
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

type GetServiceVisiableGroupList []GetServiceVisiableGroup

func (list *GetServiceVisiableGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetServiceVisiableGroup)
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

type GetServiceCasServTargetList []common.String

func (list *GetServiceCasServTargetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type GetServiceConsumeTypeList []common.String

func (list *GetServiceConsumeTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type GetServiceInputParameterList []GetServiceInputParameter

func (list *GetServiceInputParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetServiceInputParameter)
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

type GetServiceOutputParameterList []GetServiceOutputParameter

func (list *GetServiceOutputParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetServiceOutputParameter)
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

type GetServiceInputParameter3List []GetServiceInputParameter3

func (list *GetServiceInputParameter3List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetServiceInputParameter3)
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

type GetServiceOutputParameter4List []GetServiceOutputParameter4

func (list *GetServiceOutputParameter4List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetServiceOutputParameter4)
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
