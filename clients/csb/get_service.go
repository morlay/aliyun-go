package csb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	Code      int
	Message   string
	RequestId string
	Data      GetServiceData
}

type GetServiceData struct {
	Service GetServiceService
}

type GetServiceService struct {
	AccessParamsJSON       string
	Active                 bool
	Alias                  string
	AllVisiable            bool
	ApproveUserId          string
	CasTargets             string
	ConsumeTypesJSON       string
	CreateTime             int64
	CsbId                  int64
	ErrDefJSON             string
	Id                     int64
	InterfaceName          string
	IpBlackStr             string
	IpWhiteStr             string
	ModelVersion           string
	ModifiedTime           int64
	OldVersion             string
	OpenRestfulPath        string
	OttFlag                bool
	OwnerId                string
	PolicyHandler          string
	PrincipalName          string
	ProjectId              int64
	ProjectName            string
	ProvideType            string
	RouteConfJson          string
	SSL                    bool
	Scope                  string
	ServiceName            string
	ServiceOpenRestfulPath string
	ServiceProviderType    string
	ServiceVersion         string
	SkipAuth               bool
	StatisticName          string
	Status                 int
	UserId                 string
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
	Id                int64
	OldVersion        string
	OttFlag           bool
	SSL               bool
	Scope             string
	ServiceVersion    string
	SkipAuth          bool
	StatisticName     string
	Status            int
	ValidConsumeTypes bool
	ValidProvideType  bool
}

type GetServiceVisiableGroup struct {
	Id           int64
	GroupId      int64
	UserId       string
	ServiceId    int64
	CreateTime   int64
	ModifiedTime int64
	Status       int
}

type GetServiceRouteConf struct {
	ServiceRouteStrategy string
	ImportConf           GetServiceImportConf
	ImportConfs          GetServiceImportConfs
}

type GetServiceImportConf struct {
	AccessEndpointJSON string
	ProvideType        string
	InputParameterMap  GetServiceInputParameterList
	OutputParameterMap GetServiceOutputParameterList
}

type GetServiceInputParameter struct {
	CatType      int
	Depth        int
	ExtType      int
	MapStyle     int
	Optional     bool
	OriginalName string
	ParamType    string
	PassMethod   string
	TargetName   string
}

type GetServiceOutputParameter struct {
	CatType      int
	Depth        int
	ExtType      int
	MapStyle     int
	Optional     bool
	OriginalName string
	ParamType    string
	PassMethod   string
	TargetName   string
}

type GetServiceImportConfs struct {
	AccessEndpointJSON  string
	ProvideType         string
	InputParameterMap1  GetServiceInputParameter3List
	OutputParameterMap2 GetServiceOutputParameter4List
}

type GetServiceInputParameter3 struct {
	CatType      int
	Depth        int
	ExtType      int
	MapStyle     int
	Optional     bool
	OriginalName string
	ParamType    string
	PassMethod   string
	TargetName   string
}

type GetServiceOutputParameter4 struct {
	CatType      int
	Depth        int
	ExtType      int
	MapStyle     int
	Optional     bool
	OriginalName string
	ParamType    string
	PassMethod   string
	TargetName   string
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

type GetServiceCasServTargetList []string

func (list *GetServiceCasServTargetList) UnmarshalJSON(data []byte) error {
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

type GetServiceConsumeTypeList []string

func (list *GetServiceConsumeTypeList) UnmarshalJSON(data []byte) error {
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
