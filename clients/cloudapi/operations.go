package cloudapi

import (
	"encoding/json"

	"github.com/morlay/aliyun-go/core"
)

func (c *CloudapiClient) ModifySignature(req *ModifySignatureArgs) (resp *ModifySignatureResponse, err error) {
	resp = &ModifySignatureResponse{}
	err = c.Request("ModifySignature", req, resp)
	return
}

type ModifySignatureArgs struct {
	SignatureId     string
	SignatureName   string
	SignatureKey    string
	SignatureSecret string
}
type ModifySignatureResponse struct {
	RequestId     string
	SignatureId   string
	SignatureName string
}

func (c *CloudapiClient) DeleteApiGroup(req *DeleteApiGroupArgs) (resp *DeleteApiGroupResponse, err error) {
	resp = &DeleteApiGroupResponse{}
	err = c.Request("DeleteApiGroup", req, resp)
	return
}

type DeleteApiGroupArgs struct {
	GroupId string
}
type DeleteApiGroupResponse struct {
	RequestId string
}

func (c *CloudapiClient) SetDomain(req *SetDomainArgs) (resp *SetDomainResponse, err error) {
	resp = &SetDomainResponse{}
	err = c.Request("SetDomain", req, resp)
	return
}

type SetDomainArgs struct {
	GroupId               string
	DomainName            string
	CertificateName       string
	CertificateBody       string
	CertificatePrivateKey string
}
type SetDomainResponse struct {
	RequestId           string
	GroupId             string
	DomainName          string
	SubDomain           string
	DomainBindingStatus string
	DomainLegalStatus   string
	DomainRemark        string
}

func (c *CloudapiClient) DescribeSignatures(req *DescribeSignaturesArgs) (resp *DescribeSignaturesResponse, err error) {
	resp = &DescribeSignaturesResponse{}
	err = c.Request("DescribeSignatures", req, resp)
	return
}

type DescribeSignaturesSignatureInfo struct {
	RegionId        string
	SignatureId     string
	SignatureName   string
	SignatureKey    string
	SignatureSecret string
	CreatedTime     string
	ModifiedTime    string
}
type DescribeSignaturesArgs struct {
	SignatureId   string
	SignatureName string
	PageNumber    int
	PageSize      int
}
type DescribeSignaturesResponse struct {
	RequestId      string
	TotalCount     int
	PageSize       int
	PageNumber     int
	SignatureInfos DescribeSignaturesSignatureInfoList
}

type DescribeSignaturesSignatureInfoList []DescribeSignaturesSignatureInfo

func (list *DescribeSignaturesSignatureInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSignaturesSignatureInfo)
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

func (c *CloudapiClient) DescribeTrafficControls(req *DescribeTrafficControlsArgs) (resp *DescribeTrafficControlsResponse, err error) {
	resp = &DescribeTrafficControlsResponse{}
	err = c.Request("DescribeTrafficControls", req, resp)
	return
}

type DescribeTrafficControlsTrafficControl struct {
	TrafficControlId   string
	TrafficControlName string
	Description        string
	TrafficControlUnit string
	ApiDefault         int
	UserDefault        int
	AppDefault         int
	CreatedTime        string
	ModifiedTime       string
	SpecialPolicies    DescribeTrafficControlsSpecialPolicyList
}

type DescribeTrafficControlsSpecialPolicy struct {
	SpecialType string
	Specials    DescribeTrafficControlsSpecialList
}

type DescribeTrafficControlsSpecial struct {
	SpecialKey   string
	TrafficValue int
}
type DescribeTrafficControlsArgs struct {
	TrafficControlId   string
	GroupId            string
	ApiId              string
	StageName          string
	TrafficControlName string
	PageNumber         int
	PageSize           int
}
type DescribeTrafficControlsResponse struct {
	RequestId       string
	TotalCount      int
	PageSize        int
	PageNumber      int
	TrafficControls DescribeTrafficControlsTrafficControlList
}

type DescribeTrafficControlsSpecialPolicyList []DescribeTrafficControlsSpecialPolicy

func (list *DescribeTrafficControlsSpecialPolicyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTrafficControlsSpecialPolicy)
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

type DescribeTrafficControlsSpecialList []DescribeTrafficControlsSpecial

func (list *DescribeTrafficControlsSpecialList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTrafficControlsSpecial)
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

type DescribeTrafficControlsTrafficControlList []DescribeTrafficControlsTrafficControl

func (list *DescribeTrafficControlsTrafficControlList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTrafficControlsTrafficControl)
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

func (c *CloudapiClient) SwitchApi(req *SwitchApiArgs) (resp *SwitchApiResponse, err error) {
	resp = &SwitchApiResponse{}
	err = c.Request("SwitchApi", req, resp)
	return
}

type SwitchApiArgs struct {
	GroupId        string
	ApiId          string
	StageName      string
	Description    string
	HistoryVersion string
}
type SwitchApiResponse struct {
	RequestId string
}

func (c *CloudapiClient) DescribeApiLatencyData(req *DescribeApiLatencyDataArgs) (resp *DescribeApiLatencyDataResponse, err error) {
	resp = &DescribeApiLatencyDataResponse{}
	err = c.Request("DescribeApiLatencyData", req, resp)
	return
}

type DescribeApiLatencyDataMonitorItem struct {
	ItemTime  string
	ItemValue string
}
type DescribeApiLatencyDataArgs struct {
	ApiId     string
	GroupId   string
	StartTime string
	EndTime   string
}
type DescribeApiLatencyDataResponse struct {
	RequestId    string
	CallLatencys DescribeApiLatencyDataMonitorItemList
}

type DescribeApiLatencyDataMonitorItemList []DescribeApiLatencyDataMonitorItem

func (list *DescribeApiLatencyDataMonitorItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiLatencyDataMonitorItem)
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

func (c *CloudapiClient) SetTrafficControlApis(req *SetTrafficControlApisArgs) (resp *SetTrafficControlApisResponse, err error) {
	resp = &SetTrafficControlApisResponse{}
	err = c.Request("SetTrafficControlApis", req, resp)
	return
}

type SetTrafficControlApisArgs struct {
	TrafficControlId string
	GroupId          string
	ApiIds           string
	StageName        string
}
type SetTrafficControlApisResponse struct {
	RequestId string
}

func (c *CloudapiClient) CreateApp(req *CreateAppArgs) (resp *CreateAppResponse, err error) {
	resp = &CreateAppResponse{}
	err = c.Request("CreateApp", req, resp)
	return
}

type CreateAppArgs struct {
	AppName     string
	Description string
}
type CreateAppResponse struct {
	RequestId string
	AppId     int64
}

func (c *CloudapiClient) CreateTrafficControl(req *CreateTrafficControlArgs) (resp *CreateTrafficControlResponse, err error) {
	resp = &CreateTrafficControlResponse{}
	err = c.Request("CreateTrafficControl", req, resp)
	return
}

type CreateTrafficControlArgs struct {
	TrafficControlName string
	TrafficControlUnit string
	ApiDefault         int
	UserDefault        int
	AppDefault         int
	Description        string
}
type CreateTrafficControlResponse struct {
	RequestId        string
	TrafficControlId string
}

func (c *CloudapiClient) RemoveTrafficControlApis(req *RemoveTrafficControlApisArgs) (resp *RemoveTrafficControlApisResponse, err error) {
	resp = &RemoveTrafficControlApisResponse{}
	err = c.Request("RemoveTrafficControlApis", req, resp)
	return
}

type RemoveTrafficControlApisArgs struct {
	TrafficControlId string
	GroupId          string
	ApiIds           string
	StageName        string
}
type RemoveTrafficControlApisResponse struct {
	RequestId string
}

func (c *CloudapiClient) DescribeVpcAccesses(req *DescribeVpcAccessesArgs) (resp *DescribeVpcAccessesResponse, err error) {
	resp = &DescribeVpcAccessesResponse{}
	err = c.Request("DescribeVpcAccesses", req, resp)
	return
}

type DescribeVpcAccessesVpcAccessAttribute struct {
	VpcId       string
	InstanceId  string
	Port        int
	Name        string
	CreatedTime string
}
type DescribeVpcAccessesArgs struct {
}
type DescribeVpcAccessesResponse struct {
	RequestId           string
	TotalCount          int
	PageSize            int
	PageNumber          int
	VpcAccessAttributes DescribeVpcAccessesVpcAccessAttributeList
}

type DescribeVpcAccessesVpcAccessAttributeList []DescribeVpcAccessesVpcAccessAttribute

func (list *DescribeVpcAccessesVpcAccessAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVpcAccessesVpcAccessAttribute)
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

func (c *CloudapiClient) DescribeApisByTrafficControl(req *DescribeApisByTrafficControlArgs) (resp *DescribeApisByTrafficControlResponse, err error) {
	resp = &DescribeApisByTrafficControlResponse{}
	err = c.Request("DescribeApisByTrafficControl", req, resp)
	return
}

type DescribeApisByTrafficControlApiInfo struct {
	RegionId    string
	GroupId     string
	GroupName   string
	StageName   string
	ApiId       string
	ApiName     string
	Description string
	Visibility  string
	BoundTime   string
}
type DescribeApisByTrafficControlArgs struct {
	TrafficControlId string
	PageSize         int
	PageNumber       int
}
type DescribeApisByTrafficControlResponse struct {
	RequestId  string
	TotalCount int
	PageSize   int
	PageNumber int
	ApiInfos   DescribeApisByTrafficControlApiInfoList
}

type DescribeApisByTrafficControlApiInfoList []DescribeApisByTrafficControlApiInfo

func (list *DescribeApisByTrafficControlApiInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApisByTrafficControlApiInfo)
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

func (c *CloudapiClient) DescribeApiTrafficControls(req *DescribeApiTrafficControlsArgs) (resp *DescribeApiTrafficControlsResponse, err error) {
	resp = &DescribeApiTrafficControlsResponse{}
	err = c.Request("DescribeApiTrafficControls", req, resp)
	return
}

type DescribeApiTrafficControlsApiTrafficControlItem struct {
	ApiId              string
	ApiName            string
	TrafficControlId   string
	TrafficControlName string
	BoundTime          string
}
type DescribeApiTrafficControlsArgs struct {
	StageName  string
	GroupId    string
	ApiIds     string
	PageNumber int
	PageSize   int
}
type DescribeApiTrafficControlsResponse struct {
	RequestId          string
	TotalCount         int
	PageSize           int
	PageNumber         int
	ApiTrafficControls DescribeApiTrafficControlsApiTrafficControlItemList
}

type DescribeApiTrafficControlsApiTrafficControlItemList []DescribeApiTrafficControlsApiTrafficControlItem

func (list *DescribeApiTrafficControlsApiTrafficControlItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiTrafficControlsApiTrafficControlItem)
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

func (c *CloudapiClient) DescribeSignaturesByApi(req *DescribeSignaturesByApiArgs) (resp *DescribeSignaturesByApiResponse, err error) {
	resp = &DescribeSignaturesByApiResponse{}
	err = c.Request("DescribeSignaturesByApi", req, resp)
	return
}

type DescribeSignaturesByApiSignatureItem struct {
	SignatureId   string
	SignatureName string
	BoundTime     string
}
type DescribeSignaturesByApiArgs struct {
	GroupId   string
	ApiId     string
	StageName string
}
type DescribeSignaturesByApiResponse struct {
	RequestId  string
	Signatures DescribeSignaturesByApiSignatureItemList
}

type DescribeSignaturesByApiSignatureItemList []DescribeSignaturesByApiSignatureItem

func (list *DescribeSignaturesByApiSignatureItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSignaturesByApiSignatureItem)
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

func (c *CloudapiClient) DescribeApiDoc(req *DescribeApiDocArgs) (resp *DescribeApiDocResponse, err error) {
	resp = &DescribeApiDocResponse{}
	err = c.Request("DescribeApiDoc", req, resp)
	return
}

type DescribeApiDocArgs struct {
	GroupId   string
	StageName string
	ApiId     string
}
type DescribeApiDocResponse struct {
	RequestId         string
	RegionId          string
	GroupId           string
	GroupName         string
	StageName         string
	ApiId             string
	ApiName           string
	Description       string
	Visibility        string
	AuthType          string
	ResultType        string
	ResultSample      string
	FailResultSample  string
	DeployedTime      string
	ErrorCodeSamples  DescribeApiDocErrorCodeSampleList
	RequestParameters DescribeApiDocRequestParameterList
	RequestConfig     DescribeApiRequestConfig
}

type DescribeApiDocErrorCodeSampleList []DescribeApiErrorCodeSample

func (list *DescribeApiDocErrorCodeSampleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiErrorCodeSample)
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

type DescribeApiDocRequestParameterList []DescribeApiRequestParameter

func (list *DescribeApiDocRequestParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiRequestParameter)
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

func (c *CloudapiClient) DeleteSignature(req *DeleteSignatureArgs) (resp *DeleteSignatureResponse, err error) {
	resp = &DeleteSignatureResponse{}
	err = c.Request("DeleteSignature", req, resp)
	return
}

type DeleteSignatureArgs struct {
	SignatureId string
}
type DeleteSignatureResponse struct {
	RequestId string
}

func (c *CloudapiClient) RemoveApisAuthorities(req *RemoveApisAuthoritiesArgs) (resp *RemoveApisAuthoritiesResponse, err error) {
	resp = &RemoveApisAuthoritiesResponse{}
	err = c.Request("RemoveApisAuthorities", req, resp)
	return
}

type RemoveApisAuthoritiesArgs struct {
	GroupId     string
	AppId       int64
	StageName   string
	ApiIds      string
	Description string
}
type RemoveApisAuthoritiesResponse struct {
	RequestId string
}

func (c *CloudapiClient) AddTrafficSpecialControl(req *AddTrafficSpecialControlArgs) (resp *AddTrafficSpecialControlResponse, err error) {
	resp = &AddTrafficSpecialControlResponse{}
	err = c.Request("AddTrafficSpecialControl", req, resp)
	return
}

type AddTrafficSpecialControlArgs struct {
	TrafficControlId string
	SpecialType      string
	SpecialKey       string
	TrafficValue     int
}
type AddTrafficSpecialControlResponse struct {
	RequestId string
}

func (c *CloudapiClient) CreateApiGroup(req *CreateApiGroupArgs) (resp *CreateApiGroupResponse, err error) {
	resp = &CreateApiGroupResponse{}
	err = c.Request("CreateApiGroup", req, resp)
	return
}

type CreateApiGroupArgs struct {
	GroupName   string
	Description string
}
type CreateApiGroupResponse struct {
	RequestId   string
	GroupId     string
	GroupName   string
	SubDomain   string
	Description string
}

func (c *CloudapiClient) DeleteAllTrafficSpecialControl(req *DeleteAllTrafficSpecialControlArgs) (resp *DeleteAllTrafficSpecialControlResponse, err error) {
	resp = &DeleteAllTrafficSpecialControlResponse{}
	err = c.Request("DeleteAllTrafficSpecialControl", req, resp)
	return
}

type DeleteAllTrafficSpecialControlArgs struct {
	TrafficControlId string
}
type DeleteAllTrafficSpecialControlResponse struct {
	RequestId string
}

func (c *CloudapiClient) DescribeTrafficControlsByApi(req *DescribeTrafficControlsByApiArgs) (resp *DescribeTrafficControlsByApiResponse, err error) {
	resp = &DescribeTrafficControlsByApiResponse{}
	err = c.Request("DescribeTrafficControlsByApi", req, resp)
	return
}

type DescribeTrafficControlsByApiTrafficControlItem struct {
	TrafficControlItemId   string
	TrafficControlItemName string
	BoundTime              string
}
type DescribeTrafficControlsByApiArgs struct {
	GroupId   string
	ApiId     string
	StageName string
}
type DescribeTrafficControlsByApiResponse struct {
	RequestId           string
	TrafficControlItems DescribeTrafficControlsByApiTrafficControlItemList
}

type DescribeTrafficControlsByApiTrafficControlItemList []DescribeTrafficControlsByApiTrafficControlItem

func (list *DescribeTrafficControlsByApiTrafficControlItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTrafficControlsByApiTrafficControlItem)
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

func (c *CloudapiClient) SetVpcAccess(req *SetVpcAccessArgs) (resp *SetVpcAccessResponse, err error) {
	resp = &SetVpcAccessResponse{}
	err = c.Request("SetVpcAccess", req, resp)
	return
}

type SetVpcAccessArgs struct {
	VpcId      string
	InstanceId string
	Port       int
	Name       string
}
type SetVpcAccessResponse struct {
	RequestId string
}

func (c *CloudapiClient) ModifyApiGroup(req *ModifyApiGroupArgs) (resp *ModifyApiGroupResponse, err error) {
	resp = &ModifyApiGroupResponse{}
	err = c.Request("ModifyApiGroup", req, resp)
	return
}

type ModifyApiGroupArgs struct {
	GroupId     string
	GroupName   string
	Description string
}
type ModifyApiGroupResponse struct {
	RequestId   string
	GroupId     string
	GroupName   string
	SubDomain   string
	Description string
}

func (c *CloudapiClient) CreateApiStageVariable(req *CreateApiStageVariableArgs) (resp *CreateApiStageVariableResponse, err error) {
	resp = &CreateApiStageVariableResponse{}
	err = c.Request("CreateApiStageVariable", req, resp)
	return
}

type CreateApiStageVariableArgs struct {
	GroupId       string
	StageId       string
	VariableName  string
	VariableValue string
}
type CreateApiStageVariableResponse struct {
	RequestId string
}

func (c *CloudapiClient) ModifyTrafficControl(req *ModifyTrafficControlArgs) (resp *ModifyTrafficControlResponse, err error) {
	resp = &ModifyTrafficControlResponse{}
	err = c.Request("ModifyTrafficControl", req, resp)
	return
}

type ModifyTrafficControlArgs struct {
	TrafficControlId   string
	TrafficControlName string
	TrafficControlUnit string
	ApiDefault         int
	UserDefault        int
	AppDefault         int
	Description        string
}
type ModifyTrafficControlResponse struct {
	RequestId string
}

func (c *CloudapiClient) AddCatalogRelations(req *AddCatalogRelationsArgs) (resp *AddCatalogRelationsResponse, err error) {
	resp = &AddCatalogRelationsResponse{}
	err = c.Request("AddCatalogRelations", req, resp)
	return
}

type AddCatalogRelationsArgs struct {
	CatalogId string
	ApiIds    string
	ApiList   AddCatalogRelationsApiListList
}
type AddCatalogRelationsResponse struct {
	RequestId string
}

type AddCatalogRelationsApiListList []string

func (list *AddCatalogRelationsApiListList) UnmarshalJSON(data []byte) error {
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

func (c *CloudapiClient) CreateApi(req *CreateApiArgs) (resp *CreateApiResponse, err error) {
	resp = &CreateApiResponse{}
	err = c.Request("CreateApi", req, resp)
	return
}

type CreateApiArgs struct {
	GroupId              string
	ApiName              string
	Visibility           string
	Description          string
	AuthType             string
	RequestConfig        string
	ServiceConfig        string
	RequestParameters    string
	SystemParameters     string
	ConstantParameters   string
	ServiceParameters    string
	ServiceParametersMap string
	ResultType           string
	ResultSample         string
	FailResultSample     string
	ErrorCodeSamples     string
	OpenIdConnectConfig  string
}
type CreateApiResponse struct {
	RequestId string
	ApiId     string
}

func (c *CloudapiClient) CreateSignature(req *CreateSignatureArgs) (resp *CreateSignatureResponse, err error) {
	resp = &CreateSignatureResponse{}
	err = c.Request("CreateSignature", req, resp)
	return
}

type CreateSignatureArgs struct {
	SignatureName   string
	SignatureKey    string
	SignatureSecret string
}
type CreateSignatureResponse struct {
	RequestId     string
	SignatureId   string
	SignatureName string
}

func (c *CloudapiClient) RemoveCatalogRelations(req *RemoveCatalogRelationsArgs) (resp *RemoveCatalogRelationsResponse, err error) {
	resp = &RemoveCatalogRelationsResponse{}
	err = c.Request("RemoveCatalogRelations", req, resp)
	return
}

type RemoveCatalogRelationsArgs struct {
	CatalogId string
}
type RemoveCatalogRelationsResponse struct {
	RequestId string
}

func (c *CloudapiClient) DescribeApiGroups(req *DescribeApiGroupsArgs) (resp *DescribeApiGroupsResponse, err error) {
	resp = &DescribeApiGroupsResponse{}
	err = c.Request("DescribeApiGroups", req, resp)
	return
}

type DescribeApiGroupsApiGroupAttribute struct {
	GroupId       string
	GroupName     string
	SubDomain     string
	Description   string
	CreatedTime   string
	ModifiedTime  string
	RegionId      string
	TrafficLimit  int
	BillingStatus DescribeApiGroupsBillingStatus
	IllegalStatus DescribeApiGroupsIllegalStatus
}

type DescribeApiGroupsBillingStatus struct {
	StringValue string
}

type DescribeApiGroupsIllegalStatus struct {
	StringValue string
}
type DescribeApiGroupsArgs struct {
	GroupId    string
	GroupName  string
	PageNumber int
	PageSize   int
}
type DescribeApiGroupsResponse struct {
	RequestId          string
	TotalCount         int
	PageSize           int
	PageNumber         int
	ApiGroupAttributes DescribeApiGroupsApiGroupAttributeList
}

type DescribeApiGroupsApiGroupAttributeList []DescribeApiGroupsApiGroupAttribute

func (list *DescribeApiGroupsApiGroupAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiGroupsApiGroupAttribute)
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

func (c *CloudapiClient) DeleteDomainCertificate(req *DeleteDomainCertificateArgs) (resp *DeleteDomainCertificateResponse, err error) {
	resp = &DeleteDomainCertificateResponse{}
	err = c.Request("DeleteDomainCertificate", req, resp)
	return
}

type DeleteDomainCertificateArgs struct {
	GroupId       string
	DomainName    string
	CertificateId string
}
type DeleteDomainCertificateResponse struct {
	RequestId string
}

func (c *CloudapiClient) DescribeApiHistory(req *DescribeApiHistoryArgs) (resp *DescribeApiHistoryResponse, err error) {
	resp = &DescribeApiHistoryResponse{}
	err = c.Request("DescribeApiHistory", req, resp)
	return
}

type DescribeApiHistoryArgs struct {
	GroupId        string
	ApiId          string
	StageName      string
	HistoryVersion string
}
type DescribeApiHistoryResponse struct {
	RequestId              string
	RegionId               string
	GroupId                string
	GroupName              string
	StageName              string
	ApiId                  string
	ApiName                string
	Description            string
	HistoryVersion         string
	Status                 string
	Visibility             string
	AuthType               string
	ResultType             string
	ResultSample           string
	FailResultSample       string
	DeployedTime           string
	ErrorCodeSamples       DescribeApiHistoryErrorCodeSampleList
	SystemParameters       DescribeApiHistorySystemParameterList
	CustomSystemParameters DescribeApiHistorySystemParameterList
	ConstantParameters     DescribeApiHistoryConstantParameterList
	RequestParameters      DescribeApiHistoryRequestParameterList
	ServiceParameters      DescribeApiHistoryServiceParameterList
	ServiceParametersMap   DescribeApiHistoryServiceParameterMapList
	RequestConfig          DescribeApiRequestConfig
	ServiceConfig          DescribeApiServiceConfig
	OpenIdConnectConfig    DescribeApiOpenIdConnectConfig
}

type DescribeApiHistoryErrorCodeSampleList []DescribeApiErrorCodeSample

func (list *DescribeApiHistoryErrorCodeSampleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiErrorCodeSample)
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

type DescribeApiHistorySystemParameterList []DescribeApiSystemParameter

func (list *DescribeApiHistorySystemParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiSystemParameter)
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

type DescribeApiHistoryConstantParameterList []DescribeApiConstantParameter

func (list *DescribeApiHistoryConstantParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiConstantParameter)
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

type DescribeApiHistoryRequestParameterList []DescribeApiRequestParameter

func (list *DescribeApiHistoryRequestParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiRequestParameter)
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

type DescribeApiHistoryServiceParameterList []DescribeApiServiceParameter

func (list *DescribeApiHistoryServiceParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiServiceParameter)
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

type DescribeApiHistoryServiceParameterMapList []DescribeApiServiceParameterMap

func (list *DescribeApiHistoryServiceParameterMapList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiServiceParameterMap)
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

func (c *CloudapiClient) DescribeApps(req *DescribeAppsArgs) (resp *DescribeAppsResponse, err error) {
	resp = &DescribeAppsResponse{}
	err = c.Request("DescribeApps", req, resp)
	return
}

type DescribeAppsAppItem struct {
	AppId       int64
	AppName     string
	Description string
}
type DescribeAppsArgs struct {
	AppId      int64
	AppOwner   string
	PageNumber int
	PageSize   int
}
type DescribeAppsResponse struct {
	RequestId  string
	TotalCount int
	PageSize   int
	PageNumber int
	Apps       DescribeAppsAppItemList
}

type DescribeAppsAppItemList []DescribeAppsAppItem

func (list *DescribeAppsAppItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAppsAppItem)
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

func (c *CloudapiClient) DeleteDomain(req *DeleteDomainArgs) (resp *DeleteDomainResponse, err error) {
	resp = &DeleteDomainResponse{}
	err = c.Request("DeleteDomain", req, resp)
	return
}

type DeleteDomainArgs struct {
	GroupId    string
	DomainName string
}
type DeleteDomainResponse struct {
	RequestId string
}

func (c *CloudapiClient) DescribeApis(req *DescribeApisArgs) (resp *DescribeApisResponse, err error) {
	resp = &DescribeApisResponse{}
	err = c.Request("DescribeApis", req, resp)
	return
}

type DescribeApisApiSummary struct {
	RegionId     string
	GroupId      string
	GroupName    string
	ApiId        string
	ApiName      string
	Visibility   string
	Description  string
	CreatedTime  string
	ModifiedTime string
}
type DescribeApisArgs struct {
	GroupId    string
	ApiId      string
	ApiName    string
	CatalogId  string
	Visibility string
	PageSize   int
	PageNumber int
}
type DescribeApisResponse struct {
	RequestId   string
	TotalCount  int
	PageSize    int
	PageNumber  int
	ApiSummarys DescribeApisApiSummaryList
}

type DescribeApisApiSummaryList []DescribeApisApiSummary

func (list *DescribeApisApiSummaryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApisApiSummary)
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

func (c *CloudapiClient) RemoveCatalogRelation(req *RemoveCatalogRelationArgs) (resp *RemoveCatalogRelationResponse, err error) {
	resp = &RemoveCatalogRelationResponse{}
	err = c.Request("RemoveCatalogRelation", req, resp)
	return
}

type RemoveCatalogRelationArgs struct {
	CatalogId string
	ApiId     string
}
type RemoveCatalogRelationResponse struct {
	RequestId string
}

func (c *CloudapiClient) RemoveAppsAuthorities(req *RemoveAppsAuthoritiesArgs) (resp *RemoveAppsAuthoritiesResponse, err error) {
	resp = &RemoveAppsAuthoritiesResponse{}
	err = c.Request("RemoveAppsAuthorities", req, resp)
	return
}

type RemoveAppsAuthoritiesArgs struct {
	GroupId   string
	ApiId     string
	StageName string
	AppIds    string
}
type RemoveAppsAuthoritiesResponse struct {
	RequestId string
}

func (c *CloudapiClient) DeleteTrafficControl(req *DeleteTrafficControlArgs) (resp *DeleteTrafficControlResponse, err error) {
	resp = &DeleteTrafficControlResponse{}
	err = c.Request("DeleteTrafficControl", req, resp)
	return
}

type DeleteTrafficControlArgs struct {
	TrafficControlId string
}
type DeleteTrafficControlResponse struct {
	RequestId string
}

func (c *CloudapiClient) RemoveVpcAccess(req *RemoveVpcAccessArgs) (resp *RemoveVpcAccessResponse, err error) {
	resp = &RemoveVpcAccessResponse{}
	err = c.Request("RemoveVpcAccess", req, resp)
	return
}

type RemoveVpcAccessArgs struct {
	VpcId      string
	InstanceId string
	Port       int
}
type RemoveVpcAccessResponse struct {
	RequestId string
}

func (c *CloudapiClient) ReactivateDomain(req *ReactivateDomainArgs) (resp *ReactivateDomainResponse, err error) {
	resp = &ReactivateDomainResponse{}
	err = c.Request("ReactivateDomain", req, resp)
	return
}

type ReactivateDomainArgs struct {
	GroupId    string
	DomainName string
}
type ReactivateDomainResponse struct {
	RequestId string
}

func (c *CloudapiClient) RemoveSignatureApis(req *RemoveSignatureApisArgs) (resp *RemoveSignatureApisResponse, err error) {
	resp = &RemoveSignatureApisResponse{}
	err = c.Request("RemoveSignatureApis", req, resp)
	return
}

type RemoveSignatureApisArgs struct {
	SignatureId string
	GroupId     string
	ApiIds      string
	StageName   string
}
type RemoveSignatureApisResponse struct {
	RequestId string
}

func (c *CloudapiClient) DescribeDeployedApis(req *DescribeDeployedApisArgs) (resp *DescribeDeployedApisResponse, err error) {
	resp = &DescribeDeployedApisResponse{}
	err = c.Request("DescribeDeployedApis", req, resp)
	return
}

type DescribeDeployedApisDeployedApiItem struct {
	RegionId     string
	ApiId        string
	ApiName      string
	GroupId      string
	GroupName    string
	StageName    string
	Visibility   string
	Description  string
	DeployedTime string
}
type DescribeDeployedApisArgs struct {
	GroupId    string
	StageName  string
	ApiId      string
	ApiName    string
	PageNumber int
	PageSize   int
}
type DescribeDeployedApisResponse struct {
	RequestId    string
	TotalCount   int
	PageSize     int
	PageNumber   int
	DeployedApis DescribeDeployedApisDeployedApiItemList
}

type DescribeDeployedApisDeployedApiItemList []DescribeDeployedApisDeployedApiItem

func (list *DescribeDeployedApisDeployedApiItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDeployedApisDeployedApiItem)
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

func (c *CloudapiClient) DescribeApisBySignature(req *DescribeApisBySignatureArgs) (resp *DescribeApisBySignatureResponse, err error) {
	resp = &DescribeApisBySignatureResponse{}
	err = c.Request("DescribeApisBySignature", req, resp)
	return
}

type DescribeApisBySignatureApiInfo struct {
	RegionId    string
	GroupId     string
	GroupName   string
	StageName   string
	ApiId       string
	ApiName     string
	Description string
	Visibility  string
	BoundTime   string
}
type DescribeApisBySignatureArgs struct {
	SignatureId string
	PageSize    int
	PageNumber  int
}
type DescribeApisBySignatureResponse struct {
	RequestId  string
	TotalCount int
	PageSize   int
	PageNumber int
	ApiInfos   DescribeApisBySignatureApiInfoList
}

type DescribeApisBySignatureApiInfoList []DescribeApisBySignatureApiInfo

func (list *DescribeApisBySignatureApiInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApisBySignatureApiInfo)
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

func (c *CloudapiClient) DescribeApiSignatures(req *DescribeApiSignaturesArgs) (resp *DescribeApiSignaturesResponse, err error) {
	resp = &DescribeApiSignaturesResponse{}
	err = c.Request("DescribeApiSignatures", req, resp)
	return
}

type DescribeApiSignaturesApiSignatureItem struct {
	ApiId         string
	ApiName       string
	SignatureId   string
	SignatureName string
	BoundTime     string
}
type DescribeApiSignaturesArgs struct {
	StageName  string
	GroupId    string
	ApiIds     string
	PageNumber int
	PageSize   int
}
type DescribeApiSignaturesResponse struct {
	RequestId     string
	TotalCount    int
	PageSize      int
	PageNumber    int
	ApiSignatures DescribeApiSignaturesApiSignatureItemList
}

type DescribeApiSignaturesApiSignatureItemList []DescribeApiSignaturesApiSignatureItem

func (list *DescribeApiSignaturesApiSignatureItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiSignaturesApiSignatureItem)
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

func (c *CloudapiClient) DescribeApiErrorData(req *DescribeApiErrorDataArgs) (resp *DescribeApiErrorDataResponse, err error) {
	resp = &DescribeApiErrorDataResponse{}
	err = c.Request("DescribeApiErrorData", req, resp)
	return
}

type DescribeApiErrorDataMonitorItem struct {
	ItemTime  string
	ItemValue string
}
type DescribeApiErrorDataArgs struct {
	ApiId     string
	GroupId   string
	StartTime string
	EndTime   string
}
type DescribeApiErrorDataResponse struct {
	RequestId    string
	ClientErrors DescribeApiErrorDataMonitorItemList
	ServerErrors DescribeApiErrorDataMonitorItemList
}

type DescribeApiErrorDataMonitorItemList []DescribeApiErrorDataMonitorItem

func (list *DescribeApiErrorDataMonitorItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiErrorDataMonitorItem)
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

func (c *CloudapiClient) DescribeDeployedApi(req *DescribeDeployedApiArgs) (resp *DescribeDeployedApiResponse, err error) {
	resp = &DescribeDeployedApiResponse{}
	err = c.Request("DescribeDeployedApi", req, resp)
	return
}

type DescribeDeployedApiArgs struct {
	GroupId   string
	ApiId     string
	StageName string
}
type DescribeDeployedApiResponse struct {
	RequestId              string
	RegionId               string
	GroupId                string
	GroupName              string
	StageName              string
	ApiId                  string
	ApiName                string
	Description            string
	Visibility             string
	AuthType               string
	ResultType             string
	ResultSample           string
	FailResultSample       string
	DeployedTime           string
	ErrorCodeSamples       DescribeDeployedApiErrorCodeSampleList
	SystemParameters       DescribeDeployedApiSystemParameterList
	CustomSystemParameters DescribeDeployedApiSystemParameterList
	ConstantParameters     DescribeDeployedApiConstantParameterList
	RequestParameters      DescribeDeployedApiRequestParameterList
	ServiceParameters      DescribeDeployedApiServiceParameterList
	ServiceParametersMap   DescribeDeployedApiServiceParameterMapList
	RequestConfig          DescribeApiRequestConfig
	ServiceConfig          DescribeApiServiceConfig
	OpenIdConnectConfig    DescribeApiOpenIdConnectConfig
}

type DescribeDeployedApiErrorCodeSampleList []DescribeApiErrorCodeSample

func (list *DescribeDeployedApiErrorCodeSampleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiErrorCodeSample)
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

type DescribeDeployedApiSystemParameterList []DescribeApiSystemParameter

func (list *DescribeDeployedApiSystemParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiSystemParameter)
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

type DescribeDeployedApiConstantParameterList []DescribeApiConstantParameter

func (list *DescribeDeployedApiConstantParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiConstantParameter)
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

type DescribeDeployedApiRequestParameterList []DescribeApiRequestParameter

func (list *DescribeDeployedApiRequestParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiRequestParameter)
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

type DescribeDeployedApiServiceParameterList []DescribeApiServiceParameter

func (list *DescribeDeployedApiServiceParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiServiceParameter)
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

type DescribeDeployedApiServiceParameterMapList []DescribeApiServiceParameterMap

func (list *DescribeDeployedApiServiceParameterMapList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiServiceParameterMap)
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

func (c *CloudapiClient) DescribeDomain(req *DescribeDomainArgs) (resp *DescribeDomainResponse, err error) {
	resp = &DescribeDomainResponse{}
	err = c.Request("DescribeDomain", req, resp)
	return
}

type DescribeDomainArgs struct {
	GroupId    string
	DomainName string
}
type DescribeDomainResponse struct {
	RequestId             string
	GroupId               string
	DomainName            string
	SubDomain             string
	CertificateId         string
	CertificateName       string
	CertificateBody       string
	CertificatePrivateKey string
	DomainBindingStatus   string
	DomainCNAMEStatus     string
	DomainLegalStatus     string
	DomainRemark          string
}

func (c *CloudapiClient) DescribeApiTrafficData(req *DescribeApiTrafficDataArgs) (resp *DescribeApiTrafficDataResponse, err error) {
	resp = &DescribeApiTrafficDataResponse{}
	err = c.Request("DescribeApiTrafficData", req, resp)
	return
}

type DescribeApiTrafficDataMonitorItem struct {
	ItemTime  string
	ItemValue string
}
type DescribeApiTrafficDataArgs struct {
	ApiId     string
	GroupId   string
	StartTime string
	EndTime   string
}
type DescribeApiTrafficDataResponse struct {
	RequestId     string
	CallUploads   DescribeApiTrafficDataMonitorItemList
	CallDownloads DescribeApiTrafficDataMonitorItemList
}

type DescribeApiTrafficDataMonitorItemList []DescribeApiTrafficDataMonitorItem

func (list *DescribeApiTrafficDataMonitorItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiTrafficDataMonitorItem)
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

func (c *CloudapiClient) AbolishApi(req *AbolishApiArgs) (resp *AbolishApiResponse, err error) {
	resp = &AbolishApiResponse{}
	err = c.Request("AbolishApi", req, resp)
	return
}

type AbolishApiArgs struct {
	GroupId   string
	ApiId     string
	StageName string
}
type AbolishApiResponse struct {
	RequestId string
}

func (c *CloudapiClient) DescribeSystemParameters(req *DescribeSystemParametersArgs) (resp *DescribeSystemParametersResponse, err error) {
	resp = &DescribeSystemParametersResponse{}
	err = c.Request("DescribeSystemParameters", req, resp)
	return
}

type DescribeSystemParametersSystemParamItem struct {
	ParamName   string
	ParamType   string
	DemoValue   string
	Description string
}
type DescribeSystemParametersArgs struct {
}
type DescribeSystemParametersResponse struct {
	RequestId    string
	SystemParams DescribeSystemParametersSystemParamItemList
}

type DescribeSystemParametersSystemParamItemList []DescribeSystemParametersSystemParamItem

func (list *DescribeSystemParametersSystemParamItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSystemParametersSystemParamItem)
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

func (c *CloudapiClient) AddCatalogRelation(req *AddCatalogRelationArgs) (resp *AddCatalogRelationResponse, err error) {
	resp = &AddCatalogRelationResponse{}
	err = c.Request("AddCatalogRelation", req, resp)
	return
}

type AddCatalogRelationArgs struct {
	CatalogId string
	ApiId     string
}
type AddCatalogRelationResponse struct {
	RequestId string
}

func (c *CloudapiClient) DeleteTrafficSpecialControl(req *DeleteTrafficSpecialControlArgs) (resp *DeleteTrafficSpecialControlResponse, err error) {
	resp = &DeleteTrafficSpecialControlResponse{}
	err = c.Request("DeleteTrafficSpecialControl", req, resp)
	return
}

type DeleteTrafficSpecialControlArgs struct {
	TrafficControlId string
	SpecialType      string
	SpecialKey       string
}
type DeleteTrafficSpecialControlResponse struct {
	RequestId string
}

func (c *CloudapiClient) DescribeAppAttributes(req *DescribeAppAttributesArgs) (resp *DescribeAppAttributesResponse, err error) {
	resp = &DescribeAppAttributesResponse{}
	err = c.Request("DescribeAppAttributes", req, resp)
	return
}

type DescribeAppAttributesAppAttribute struct {
	AppId        int64
	AppName      string
	Description  string
	CreatedTime  string
	ModifiedTime string
}
type DescribeAppAttributesArgs struct {
	AppId      int64
	PageNumber int
	PageSize   int
}
type DescribeAppAttributesResponse struct {
	RequestId  string
	TotalCount int
	PageSize   int
	PageNumber int
	Apps       DescribeAppAttributesAppAttributeList
}

type DescribeAppAttributesAppAttributeList []DescribeAppAttributesAppAttribute

func (list *DescribeAppAttributesAppAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAppAttributesAppAttribute)
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

func (c *CloudapiClient) ModifyApi(req *ModifyApiArgs) (resp *ModifyApiResponse, err error) {
	resp = &ModifyApiResponse{}
	err = c.Request("ModifyApi", req, resp)
	return
}

type ModifyApiArgs struct {
	GroupId              string
	ApiId                string
	ApiName              string
	Description          string
	Visibility           string
	AuthType             string
	RequestConfig        string
	ServiceConfig        string
	RequestParameters    string
	SystemParameters     string
	ConstantParameters   string
	ServiceParameters    string
	ServiceParametersMap string
	ResultType           string
	ResultSample         string
	FailResultSample     string
	ErrorCodeSamples     string
	OpenIdConnectConfig  string
}
type ModifyApiResponse struct {
	RequestId string
}

func (c *CloudapiClient) SetSignatureApis(req *SetSignatureApisArgs) (resp *SetSignatureApisResponse, err error) {
	resp = &SetSignatureApisResponse{}
	err = c.Request("SetSignatureApis", req, resp)
	return
}

type SetSignatureApisArgs struct {
	SignatureId string
	GroupId     string
	ApiIds      string
	StageName   string
}
type SetSignatureApisResponse struct {
	RequestId string
}

func (c *CloudapiClient) CreateCatalog(req *CreateCatalogArgs) (resp *CreateCatalogResponse, err error) {
	resp = &CreateCatalogResponse{}
	err = c.Request("CreateCatalog", req, resp)
	return
}

type CreateCatalogArgs struct {
	CatalogName string
	Description string
	ParentId    string
}
type CreateCatalogResponse struct {
	RequestId string
	CatalogId string
}

func (c *CloudapiClient) SetApisAuthorities(req *SetApisAuthoritiesArgs) (resp *SetApisAuthoritiesResponse, err error) {
	resp = &SetApisAuthoritiesResponse{}
	err = c.Request("SetApisAuthorities", req, resp)
	return
}

type SetApisAuthoritiesArgs struct {
	GroupId     string
	AppId       int64
	StageName   string
	ApiIds      string
	Description string
}
type SetApisAuthoritiesResponse struct {
	RequestId string
}

func (c *CloudapiClient) ClearCatalogRelations(req *ClearCatalogRelationsArgs) (resp *ClearCatalogRelationsResponse, err error) {
	resp = &ClearCatalogRelationsResponse{}
	err = c.Request("ClearCatalogRelations", req, resp)
	return
}

type ClearCatalogRelationsArgs struct {
	CatalogId string
}
type ClearCatalogRelationsResponse struct {
	RequestId string
}

func (c *CloudapiClient) DeleteCatalog(req *DeleteCatalogArgs) (resp *DeleteCatalogResponse, err error) {
	resp = &DeleteCatalogResponse{}
	err = c.Request("DeleteCatalog", req, resp)
	return
}

type DeleteCatalogArgs struct {
	CatalogId string
}
type DeleteCatalogResponse struct {
	RequestId string
}

func (c *CloudapiClient) ResetAppKeySecret(req *ResetAppKeySecretArgs) (resp *ResetAppKeySecretResponse, err error) {
	resp = &ResetAppKeySecretResponse{}
	err = c.Request("ResetAppKeySecret", req, resp)
	return
}

type ResetAppKeySecretArgs struct {
	AppKey string
}
type ResetAppKeySecretResponse struct {
	RequestId string
}

func (c *CloudapiClient) DeleteApi(req *DeleteApiArgs) (resp *DeleteApiResponse, err error) {
	resp = &DeleteApiResponse{}
	err = c.Request("DeleteApi", req, resp)
	return
}

type DeleteApiArgs struct {
	GroupId string
	ApiId   string
}
type DeleteApiResponse struct {
	RequestId string
}

func (c *CloudapiClient) SetDomainCertificate(req *SetDomainCertificateArgs) (resp *SetDomainCertificateResponse, err error) {
	resp = &SetDomainCertificateResponse{}
	err = c.Request("SetDomainCertificate", req, resp)
	return
}

type SetDomainCertificateArgs struct {
	GroupId               string
	DomainName            string
	CertificateName       string
	CertificateBody       string
	CertificatePrivateKey string
}
type SetDomainCertificateResponse struct {
	RequestId string
}

func (c *CloudapiClient) DescribeDomainsResolution(req *DescribeDomainsResolutionArgs) (resp *DescribeDomainsResolutionResponse, err error) {
	resp = &DescribeDomainsResolutionResponse{}
	err = c.Request("DescribeDomainsResolution", req, resp)
	return
}

type DescribeDomainsResolutionDomainResolution struct {
	DomainName             string
	DomainResolutionStatus string
}
type DescribeDomainsResolutionArgs struct {
	GroupId     string
	DomainNames string
}
type DescribeDomainsResolutionResponse struct {
	RequestId         string
	GroupId           string
	DomainResolutions DescribeDomainsResolutionDomainResolutionList
}

type DescribeDomainsResolutionDomainResolutionList []DescribeDomainsResolutionDomainResolution

func (list *DescribeDomainsResolutionDomainResolutionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainsResolutionDomainResolution)
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

func (c *CloudapiClient) DescribeApiQpsData(req *DescribeApiQpsDataArgs) (resp *DescribeApiQpsDataResponse, err error) {
	resp = &DescribeApiQpsDataResponse{}
	err = c.Request("DescribeApiQpsData", req, resp)
	return
}

type DescribeApiQpsDataMonitorItem struct {
	ItemTime  string
	ItemValue string
}
type DescribeApiQpsDataArgs struct {
	ApiId     string
	GroupId   string
	StartTime string
	EndTime   string
}
type DescribeApiQpsDataResponse struct {
	RequestId     string
	CallSuccesses DescribeApiQpsDataMonitorItemList
	CallFails     DescribeApiQpsDataMonitorItemList
}

type DescribeApiQpsDataMonitorItemList []DescribeApiQpsDataMonitorItem

func (list *DescribeApiQpsDataMonitorItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiQpsDataMonitorItem)
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

func (c *CloudapiClient) DescribeAppSecurity(req *DescribeAppSecurityArgs) (resp *DescribeAppSecurityResponse, err error) {
	resp = &DescribeAppSecurityResponse{}
	err = c.Request("DescribeAppSecurity", req, resp)
	return
}

type DescribeAppSecurityArgs struct {
	AppId int64
}
type DescribeAppSecurityResponse struct {
	RequestId    string
	AppKey       string
	AppSecret    string
	CreatedTime  string
	ModifiedTime string
}

func (c *CloudapiClient) DescribeCatalogs(req *DescribeCatalogsArgs) (resp *DescribeCatalogsResponse, err error) {
	resp = &DescribeCatalogsResponse{}
	err = c.Request("DescribeCatalogs", req, resp)
	return
}

type DescribeCatalogsCatalogAttribute struct {
	CatalogId    string
	CatalogName  string
	Description  string
	ParentId     string
	CreatedTime  string
	ModifiedTime string
	RegionId     string
}
type DescribeCatalogsArgs struct {
}
type DescribeCatalogsResponse struct {
	RequestId         string
	TotalCount        int
	PageSize          int
	PageNumber        int
	CatalogAttributes DescribeCatalogsCatalogAttributeList
}

type DescribeCatalogsCatalogAttributeList []DescribeCatalogsCatalogAttribute

func (list *DescribeCatalogsCatalogAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCatalogsCatalogAttribute)
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

func (c *CloudapiClient) DescribeHistoryApis(req *DescribeHistoryApisArgs) (resp *DescribeHistoryApisResponse, err error) {
	resp = &DescribeHistoryApisResponse{}
	err = c.Request("DescribeHistoryApis", req, resp)
	return
}

type DescribeHistoryApisApiHisItem struct {
	RegionId       string
	ApiId          string
	ApiName        string
	GroupId        string
	GroupName      string
	StageName      string
	HistoryVersion string
	Status         DescribeHistoryApisStatus
	Description    string
	DeployedTime   string
}

type DescribeHistoryApisStatus struct {
	StringValue string
}
type DescribeHistoryApisArgs struct {
	GroupId    string
	StageName  string
	ApiId      string
	ApiName    string
	PageSize   string
	PageNumber string
}
type DescribeHistoryApisResponse struct {
	RequestId   string
	TotalCount  int
	PageSize    int
	PageNumber  int
	ApiHisItems DescribeHistoryApisApiHisItemList
}

type DescribeHistoryApisApiHisItemList []DescribeHistoryApisApiHisItem

func (list *DescribeHistoryApisApiHisItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeHistoryApisApiHisItem)
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

func (c *CloudapiClient) DescribeApi(req *DescribeApiArgs) (resp *DescribeApiResponse, err error) {
	resp = &DescribeApiResponse{}
	err = c.Request("DescribeApi", req, resp)
	return
}

type DescribeApiSystemParameter struct {
	ParameterName        string
	ServiceParameterName string
	Location             string
	DemoValue            string
	Description          string
}

type DescribeApiConstantParameter struct {
	ServiceParameterName string
	ConstantValue        string
	Location             string
	Description          string
}

type DescribeApiRequestParameter struct {
	ApiParameterName  string
	Location          string
	ParameterType     string
	Required          string
	DefaultValue      string
	DemoValue         string
	MaxValue          int64
	MinValue          int64
	MaxLength         int64
	MinLength         int64
	RegularExpression string
	JsonScheme        string
	EnumValue         string
	DocShow           string
	DocOrder          int
	Description       string
}

type DescribeApiServiceParameter struct {
	ServiceParameterName string
	Location             string
	ParameterType        string
}

type DescribeApiServiceParameterMap struct {
	ServiceParameterName string
	RequestParameterName string
}

type DescribeApiDeployedInfo struct {
	StageName        string
	EffectiveVersion string
	DeployedStatus   string
}

type DescribeApiErrorCodeSample struct {
	Code        string
	Message     string
	Description string
}

type DescribeApiRequestConfig struct {
	RequestProtocol     string
	RequestHttpMethod   string
	RequestPath         string
	BodyFormat          string
	PostBodyDescription string
	RequestMode         string
}

type DescribeApiServiceConfig struct {
	ServiceProtocol   string
	ServiceAddress    string
	ServiceHttpMethod string
	ServicePath       string
	ServiceTimeout    int
	Mock              string
	MockResult        string
	ServiceVpcEnable  core.Bool
	VpcConfig         DescribeApiVpcConfig
}

type DescribeApiOpenIdConnectConfig struct {
	OpenIdApiType    string
	IdTokenParamName string
	PublicKeyId      string
	PublicKey        string
}

type DescribeApiVpcConfig struct {
	VpcId      string
	InstanceId string
	Port       int
}
type DescribeApiArgs struct {
	GroupId string
	ApiId   string
}
type DescribeApiResponse struct {
	RequestId              string
	RegionId               string
	ApiId                  string
	ApiName                string
	GroupId                string
	GroupName              string
	Visibility             string
	AuthType               string
	ResultType             string
	ResultSample           string
	FailResultSample       string
	CreatedTime            string
	ModifiedTime           string
	Description            string
	SystemParameters       DescribeApiSystemParameterList
	CustomSystemParameters DescribeApiSystemParameterList
	ConstantParameters     DescribeApiConstantParameterList
	RequestParameters      DescribeApiRequestParameterList
	ServiceParameters      DescribeApiServiceParameterList
	ServiceParametersMap   DescribeApiServiceParameterMapList
	ErrorCodeSamples       DescribeApiErrorCodeSampleList
	DeployedInfos          DescribeApiDeployedInfoList
	RequestConfig          DescribeApiRequestConfig
	ServiceConfig          DescribeApiServiceConfig
	OpenIdConnectConfig    DescribeApiOpenIdConnectConfig
}

type DescribeApiSystemParameterList []DescribeApiSystemParameter

func (list *DescribeApiSystemParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiSystemParameter)
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

type DescribeApiConstantParameterList []DescribeApiConstantParameter

func (list *DescribeApiConstantParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiConstantParameter)
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

type DescribeApiRequestParameterList []DescribeApiRequestParameter

func (list *DescribeApiRequestParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiRequestParameter)
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

type DescribeApiServiceParameterList []DescribeApiServiceParameter

func (list *DescribeApiServiceParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiServiceParameter)
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

type DescribeApiServiceParameterMapList []DescribeApiServiceParameterMap

func (list *DescribeApiServiceParameterMapList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiServiceParameterMap)
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

type DescribeApiErrorCodeSampleList []DescribeApiErrorCodeSample

func (list *DescribeApiErrorCodeSampleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiErrorCodeSample)
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

type DescribeApiDeployedInfoList []DescribeApiDeployedInfo

func (list *DescribeApiDeployedInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiDeployedInfo)
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

func (c *CloudapiClient) DescribeCatalog(req *DescribeCatalogArgs) (resp *DescribeCatalogResponse, err error) {
	resp = &DescribeCatalogResponse{}
	err = c.Request("DescribeCatalog", req, resp)
	return
}

type DescribeCatalogArgs struct {
	CatalogId string
}
type DescribeCatalogResponse struct {
	RequestId    string
	CatalogId    string
	CatalogName  string
	Description  string
	ParentId     string
	CreatedTime  string
	ModifiedTime string
	RegionId     string
	ApiIds       DescribeCatalogApiIdList
}

type DescribeCatalogApiIdList []string

func (list *DescribeCatalogApiIdList) UnmarshalJSON(data []byte) error {
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

func (c *CloudapiClient) ModifyCatalog(req *ModifyCatalogArgs) (resp *ModifyCatalogResponse, err error) {
	resp = &ModifyCatalogResponse{}
	err = c.Request("ModifyCatalog", req, resp)
	return
}

type ModifyCatalogArgs struct {
	CatalogId   string
	CatalogName string
	Description string
}
type ModifyCatalogResponse struct {
	RequestId string
}

func (c *CloudapiClient) DescribeAuthorizedApis(req *DescribeAuthorizedApisArgs) (resp *DescribeAuthorizedApisResponse, err error) {
	resp = &DescribeAuthorizedApisResponse{}
	err = c.Request("DescribeAuthorizedApis", req, resp)
	return
}

type DescribeAuthorizedApisAuthorizedApi struct {
	RegionId            string
	GroupId             string
	GroupName           string
	StageName           string
	Operator            string
	ApiId               string
	ApiName             string
	AuthorizationSource string
	Description         string
	AuthorizedTime      string
}
type DescribeAuthorizedApisArgs struct {
	AppId      int64
	PageNumber int
	PageSize   int
}
type DescribeAuthorizedApisResponse struct {
	RequestId      string
	TotalCount     int
	PageSize       int
	PageNumber     int
	AuthorizedApis DescribeAuthorizedApisAuthorizedApiList
}

type DescribeAuthorizedApisAuthorizedApiList []DescribeAuthorizedApisAuthorizedApi

func (list *DescribeAuthorizedApisAuthorizedApiList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAuthorizedApisAuthorizedApi)
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

func (c *CloudapiClient) DeleteApiStageVariable(req *DeleteApiStageVariableArgs) (resp *DeleteApiStageVariableResponse, err error) {
	resp = &DeleteApiStageVariableResponse{}
	err = c.Request("DeleteApiStageVariable", req, resp)
	return
}

type DeleteApiStageVariableArgs struct {
	GroupId      string
	StageId      string
	VariableName string
}
type DeleteApiStageVariableResponse struct {
	RequestId string
}

func (c *CloudapiClient) DescribeApiGroup(req *DescribeApiGroupArgs) (resp *DescribeApiGroupResponse, err error) {
	resp = &DescribeApiGroupResponse{}
	err = c.Request("DescribeApiGroup", req, resp)
	return
}

type DescribeApiGroupDomainItem struct {
	DomainName          string
	CertificateId       string
	CertificateName     string
	DomainCNAMEStatus   string
	DomainBindingStatus string
	DomainLegalStatus   string
	DomainRemark        string
}

type DescribeApiGroupStageInfo struct {
	StageId     string
	StageName   string
	Description string
}
type DescribeApiGroupArgs struct {
	GroupId string
}
type DescribeApiGroupResponse struct {
	RequestId     string
	GroupId       string
	GroupName     string
	SubDomain     string
	Description   string
	CreatedTime   string
	ModifiedTime  string
	RegionId      string
	Status        string
	BillingStatus string
	IllegalStatus string
	TrafficLimit  int
	CustomDomains DescribeApiGroupDomainItemList
	StageItems    DescribeApiGroupStageInfoList
}

type DescribeApiGroupDomainItemList []DescribeApiGroupDomainItem

func (list *DescribeApiGroupDomainItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiGroupDomainItem)
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

type DescribeApiGroupStageInfoList []DescribeApiGroupStageInfo

func (list *DescribeApiGroupStageInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiGroupStageInfo)
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

func (c *CloudapiClient) ModifyApp(req *ModifyAppArgs) (resp *ModifyAppResponse, err error) {
	resp = &ModifyAppResponse{}
	err = c.Request("ModifyApp", req, resp)
	return
}

type ModifyAppArgs struct {
	AppId       int64
	AppName     string
	Description string
}
type ModifyAppResponse struct {
	RequestId string
}

func (c *CloudapiClient) DescribeRegions(req *DescribeRegionsArgs) (resp *DescribeRegionsResponse, err error) {
	resp = &DescribeRegionsResponse{}
	err = c.Request("DescribeRegions", req, resp)
	return
}

type DescribeRegionsRegion struct {
	RegionId  string
	LocalName string
	EndPoint  string
}
type DescribeRegionsArgs struct {
}
type DescribeRegionsResponse struct {
	RequestId string
	Regions   DescribeRegionsRegionList
}

type DescribeRegionsRegionList []DescribeRegionsRegion

func (list *DescribeRegionsRegionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRegionsRegion)
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

func (c *CloudapiClient) SetAppsAuthorities(req *SetAppsAuthoritiesArgs) (resp *SetAppsAuthoritiesResponse, err error) {
	resp = &SetAppsAuthoritiesResponse{}
	err = c.Request("SetAppsAuthorities", req, resp)
	return
}

type SetAppsAuthoritiesArgs struct {
	GroupId     string
	ApiId       string
	StageName   string
	AppIds      string
	Description string
}
type SetAppsAuthoritiesResponse struct {
	RequestId string
}

func (c *CloudapiClient) ResetAppSecret(req *ResetAppSecretArgs) (resp *ResetAppSecretResponse, err error) {
	resp = &ResetAppSecretResponse{}
	err = c.Request("ResetAppSecret", req, resp)
	return
}

type ResetAppSecretArgs struct {
	AppKey string
}
type ResetAppSecretResponse struct {
	RequestId string
}

func (c *CloudapiClient) DescribeAuthorizedApps(req *DescribeAuthorizedAppsArgs) (resp *DescribeAuthorizedAppsResponse, err error) {
	resp = &DescribeAuthorizedAppsResponse{}
	err = c.Request("DescribeAuthorizedApps", req, resp)
	return
}

type DescribeAuthorizedAppsAuthorizedApp struct {
	StageName           string
	AppId               int64
	AppName             string
	Operator            string
	AuthorizationSource string
	Description         string
	AuthorizedTime      string
}
type DescribeAuthorizedAppsArgs struct {
	GroupId    string
	StageName  string
	ApiId      string
	PageNumber int
	PageSize   int
}
type DescribeAuthorizedAppsResponse struct {
	RequestId      string
	TotalCount     int
	PageSize       int
	PageNumber     int
	AuthorizedApps DescribeAuthorizedAppsAuthorizedAppList
}

type DescribeAuthorizedAppsAuthorizedAppList []DescribeAuthorizedAppsAuthorizedApp

func (list *DescribeAuthorizedAppsAuthorizedAppList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAuthorizedAppsAuthorizedApp)
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

func (c *CloudapiClient) DescribeApiStage(req *DescribeApiStageArgs) (resp *DescribeApiStageResponse, err error) {
	resp = &DescribeApiStageResponse{}
	err = c.Request("DescribeApiStage", req, resp)
	return
}

type DescribeApiStageVariableItem struct {
	VariableName  string
	VariableValue string
}
type DescribeApiStageArgs struct {
	GroupId string
	StageId string
}
type DescribeApiStageResponse struct {
	RequestId    string
	GroupId      string
	StageId      string
	StageName    string
	Description  string
	CreatedTime  string
	ModifiedTime string
	Variables    DescribeApiStageVariableItemList
}

type DescribeApiStageVariableItemList []DescribeApiStageVariableItem

func (list *DescribeApiStageVariableItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiStageVariableItem)
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

func (c *CloudapiClient) DeleteApp(req *DeleteAppArgs) (resp *DeleteAppResponse, err error) {
	resp = &DeleteAppResponse{}
	err = c.Request("DeleteApp", req, resp)
	return
}

type DeleteAppArgs struct {
	AppId int64
}
type DeleteAppResponse struct {
	RequestId string
}

func (c *CloudapiClient) DeployApi(req *DeployApiArgs) (resp *DeployApiResponse, err error) {
	resp = &DeployApiResponse{}
	err = c.Request("DeployApi", req, resp)
	return
}

type DeployApiArgs struct {
	GroupId     string
	ApiId       string
	StageName   string
	Description string
}
type DeployApiResponse struct {
	RequestId string
}
