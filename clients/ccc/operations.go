package ccc

import (
	"encoding/json"

	"github.com/morlay/aliyun-go/core"
)

func (c *CccClient) RequestLoginInfo(req *RequestLoginInfoArgs) (resp *RequestLoginInfoResponse, err error) {
	resp = &RequestLoginInfoResponse{}
	err = c.Request("RequestLoginInfo", req, resp)
	return
}

type RequestLoginInfoLoginInfo struct {
	UserName       string
	DisplayName    string
	Region         string
	WebRtcUrl      string
	AgentServerUrl string
	Extension      string
	TenantId       string
	Signature      string
	SignData       string
}
type RequestLoginInfoArgs struct {
	InstanceId string
}
type RequestLoginInfoResponse struct {
	RequestId      string
	Success        core.Bool
	Code           string
	Message        string
	HttpStatusCode int
	LoginInfo      RequestLoginInfoLoginInfo
}

func (c *CccClient) AddPhoneNumber(req *AddPhoneNumberArgs) (resp *AddPhoneNumberResponse, err error) {
	resp = &AddPhoneNumberResponse{}
	err = c.Request("AddPhoneNumber", req, resp)
	return
}

type AddPhoneNumberPhoneNumber struct {
	PhoneNumberId          string
	InstanceId             string
	Number                 string
	PhoneNumberDescription string
	TestOnly               core.Bool
	RemainingTime          int
	AllowOutbound          core.Bool
	Usage                  string
	Trunks                 int
	ContactFlow            AddPhoneNumberContactFlow
}

type AddPhoneNumberContactFlow struct {
	ContactFlowId          string
	InstanceId             string
	ContactFlowName        string
	ContactFlowDescription string
	Type                   string
}
type AddPhoneNumberArgs struct {
	ContactFlowId string
	InstanceId    string
	Usage         string
	PhoneNumber   string
}
type AddPhoneNumberResponse struct {
	RequestId      string
	Success        core.Bool
	Code           string
	Message        string
	HttpStatusCode int
	PhoneNumber    AddPhoneNumberPhoneNumber
}

func (c *CccClient) ListPhoneNumbers(req *ListPhoneNumbersArgs) (resp *ListPhoneNumbersResponse, err error) {
	resp = &ListPhoneNumbersResponse{}
	err = c.Request("ListPhoneNumbers", req, resp)
	return
}

type ListPhoneNumbersPhoneNumber struct {
	PhoneNumberId          string
	InstanceId             string
	Number                 string
	PhoneNumberDescription string
	TestOnly               core.Bool
	RemainingTime          int
	AllowOutbound          core.Bool
	Usage                  string
	Trunks                 int
	ContactFlow            ListPhoneNumbersContactFlow
}

type ListPhoneNumbersContactFlow struct {
	ContactFlowId          string
	InstanceId             string
	ContactFlowName        string
	ContactFlowDescription string
	Type                   string
}
type ListPhoneNumbersArgs struct {
	OutboundOnly core.Bool
	InstanceId   string
}
type ListPhoneNumbersResponse struct {
	RequestId      string
	Success        core.Bool
	Code           string
	Message        string
	HttpStatusCode int
	PhoneNumbers   ListPhoneNumbersPhoneNumberList
}

type ListPhoneNumbersPhoneNumberList []ListPhoneNumbersPhoneNumber

func (list *ListPhoneNumbersPhoneNumberList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPhoneNumbersPhoneNumber)
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

func (c *CccClient) CreateSkillGroup(req *CreateSkillGroupArgs) (resp *CreateSkillGroupResponse, err error) {
	resp = &CreateSkillGroupResponse{}
	err = c.Request("CreateSkillGroup", req, resp)
	return
}

type CreateSkillGroupArgs struct {
	SkillLevels            CreateSkillGroupSkillLevelList
	InstanceId             string
	OutboundPhoneNumberIds CreateSkillGroupOutboundPhoneNumberIdList
	Name                   string
	Description            string
	UserIds                CreateSkillGroupUserIdList
}
type CreateSkillGroupResponse struct {
	RequestId      string
	Success        core.Bool
	Code           string
	Message        string
	HttpStatusCode int
	SkillGroupId   string
}

type CreateSkillGroupSkillLevelList []int

func (list *CreateSkillGroupSkillLevelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int)
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

type CreateSkillGroupOutboundPhoneNumberIdList []string

func (list *CreateSkillGroupOutboundPhoneNumberIdList) UnmarshalJSON(data []byte) error {
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

type CreateSkillGroupUserIdList []string

func (list *CreateSkillGroupUserIdList) UnmarshalJSON(data []byte) error {
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

func (c *CccClient) CreateUser(req *CreateUserArgs) (resp *CreateUserResponse, err error) {
	resp = &CreateUserResponse{}
	err = c.Request("CreateUser", req, resp)
	return
}

type CreateUserArgs struct {
	SkillLevels   CreateUserSkillLevelList
	InstanceId    string
	LoginName     string
	Phone         string
	RoleIds       CreateUserRoleIdList
	DisplayName   string
	SkillGroupIds CreateUserSkillGroupIdList
	Email         string
}
type CreateUserResponse struct {
	RequestId      string
	Success        core.Bool
	Code           string
	Message        string
	HttpStatusCode int
	UserId         string
}

type CreateUserSkillLevelList []int

func (list *CreateUserSkillLevelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int)
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

type CreateUserRoleIdList []string

func (list *CreateUserRoleIdList) UnmarshalJSON(data []byte) error {
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

type CreateUserSkillGroupIdList []string

func (list *CreateUserSkillGroupIdList) UnmarshalJSON(data []byte) error {
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

func (c *CccClient) ListUsers(req *ListUsersArgs) (resp *ListUsersResponse, err error) {
	resp = &ListUsersResponse{}
	err = c.Request("ListUsers", req, resp)
	return
}

type ListUsersUsers struct {
	TotalCount int
	PageNumber int
	PageSize   int
	List       ListUsersUserList
}

type ListUsersUser struct {
	UserId      string
	RamId       string
	InstanceId  string
	Primary     core.Bool
	Roles       ListUsersRoleList
	SkillLevels ListUsersSkillLevelList
	Detail      ListUsersDetail
}

type ListUsersRole struct {
	RoleId          string
	InstanceId      string
	RoleName        string
	RoleDescription string
}

type ListUsersSkillLevel struct {
	SkillLevelId string
	Level        int
	Skill        ListUsersSkill
}

type ListUsersSkill struct {
	SkillGroupId          string
	InstanceId            string
	SkillGroupName        string
	SkillGroupDescription string
}

type ListUsersDetail struct {
	LoginName   string
	DisplayName string
	Phone       string
	Email       string
	Department  string
}
type ListUsersArgs struct {
	InstanceId string
	PageSize   int
	PageNumber int
}
type ListUsersResponse struct {
	RequestId      string
	Success        core.Bool
	Code           string
	Message        string
	HttpStatusCode int
	Users          ListUsersUsers
}

type ListUsersUserList []ListUsersUser

func (list *ListUsersUserList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListUsersUser)
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

type ListUsersRoleList []ListUsersRole

func (list *ListUsersRoleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListUsersRole)
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

type ListUsersSkillLevelList []ListUsersSkillLevel

func (list *ListUsersSkillLevelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListUsersSkillLevel)
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

func (c *CccClient) ListRoles(req *ListRolesArgs) (resp *ListRolesResponse, err error) {
	resp = &ListRolesResponse{}
	err = c.Request("ListRoles", req, resp)
	return
}

type ListRolesRole struct {
	RoleId          string
	InstanceId      string
	RoleName        string
	RoleDescription string
}
type ListRolesArgs struct {
	InstanceId string
}
type ListRolesResponse struct {
	RequestId      string
	Success        core.Bool
	Code           string
	Message        string
	HttpStatusCode int
	Roles          ListRolesRoleList
}

type ListRolesRoleList []ListRolesRole

func (list *ListRolesRoleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListRolesRole)
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

func (c *CccClient) ModifyPhoneNumber(req *ModifyPhoneNumberArgs) (resp *ModifyPhoneNumberResponse, err error) {
	resp = &ModifyPhoneNumberResponse{}
	err = c.Request("ModifyPhoneNumber", req, resp)
	return
}

type ModifyPhoneNumberPhoneNumber struct {
	PhoneNumberId          string
	InstanceId             string
	Number                 string
	PhoneNumberDescription string
	TestOnly               core.Bool
	RemainingTime          int
	AllowOutbound          core.Bool
	Usage                  string
	Trunks                 int
	ContactFlow            ModifyPhoneNumberContactFlow
}

type ModifyPhoneNumberContactFlow struct {
	ContactFlowId          string
	InstanceId             string
	ContactFlowName        string
	ContactFlowDescription string
	Type                   string
}
type ModifyPhoneNumberArgs struct {
	ContactFlowId string
	InstanceId    string
	PhoneNumberId string
	Usage         string
}
type ModifyPhoneNumberResponse struct {
	RequestId      string
	Success        core.Bool
	Code           string
	Message        string
	HttpStatusCode int
	PhoneNumber    ModifyPhoneNumberPhoneNumber
}

func (c *CccClient) ListContactFlows(req *ListContactFlowsArgs) (resp *ListContactFlowsResponse, err error) {
	resp = &ListContactFlowsResponse{}
	err = c.Request("ListContactFlows", req, resp)
	return
}

type ListContactFlowsContactFlow struct {
	ContactFlowId          string
	InstanceId             string
	ContactFlowName        string
	ContactFlowDescription string
	Type                   string
	AppliedVersion         string
	Versions               ListContactFlowsContactFlowVersionList
	PhoneNumbers           ListContactFlowsPhoneNumberList
}

type ListContactFlowsContactFlowVersion struct {
	ContactFlowVersionId          string
	Version                       string
	ContactFlowVersionDescription string
	LastModified                  string
	LastModifiedBy                string
	LockedBy                      string
	Status                        string
}

type ListContactFlowsPhoneNumber struct {
	PhoneNumberId          string
	InstanceId             string
	Number                 string
	PhoneNumberDescription string
	TestOnly               core.Bool
	RemainingTime          int
	AllowOutbound          core.Bool
	Usage                  string
	Trunks                 int
}
type ListContactFlowsArgs struct {
	InstanceId string
}
type ListContactFlowsResponse struct {
	RequestId      string
	Success        core.Bool
	Code           string
	Message        string
	HttpStatusCode int
	ContactFlows   ListContactFlowsContactFlowList
}

type ListContactFlowsContactFlowVersionList []ListContactFlowsContactFlowVersion

func (list *ListContactFlowsContactFlowVersionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListContactFlowsContactFlowVersion)
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

type ListContactFlowsPhoneNumberList []ListContactFlowsPhoneNumber

func (list *ListContactFlowsPhoneNumberList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListContactFlowsPhoneNumber)
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

type ListContactFlowsContactFlowList []ListContactFlowsContactFlow

func (list *ListContactFlowsContactFlowList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListContactFlowsContactFlow)
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

func (c *CccClient) GetUser(req *GetUserArgs) (resp *GetUserResponse, err error) {
	resp = &GetUserResponse{}
	err = c.Request("GetUser", req, resp)
	return
}

type GetUserUser struct {
	UserId      string
	RamId       string
	InstanceId  string
	Roles       GetUserRoleList
	SkillLevels GetUserSkillLevelList
	Detail      GetUserDetail
}

type GetUserRole struct {
	RoleId          string
	InstanceId      string
	RoleName        string
	RoleDescription string
}

type GetUserSkillLevel struct {
	SkillLevelId string
	Level        int
	Skill        GetUserSkill
}

type GetUserSkill struct {
	SkillGroupId          string
	InstanceId            string
	SkillGroupName        string
	SkillGroupDescription string
}

type GetUserDetail struct {
	LoginName   string
	DisplayName string
	Phone       string
	Email       string
	Department  string
}
type GetUserArgs struct {
	InstanceId string
	UserId     string
}
type GetUserResponse struct {
	RequestId      string
	Success        core.Bool
	Code           string
	Message        string
	HttpStatusCode int
	User           GetUserUser
}

type GetUserRoleList []GetUserRole

func (list *GetUserRoleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetUserRole)
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

type GetUserSkillLevelList []GetUserSkillLevel

func (list *GetUserSkillLevelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetUserSkillLevel)
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

func (c *CccClient) ModifySkillGroup(req *ModifySkillGroupArgs) (resp *ModifySkillGroupResponse, err error) {
	resp = &ModifySkillGroupResponse{}
	err = c.Request("ModifySkillGroup", req, resp)
	return
}

type ModifySkillGroupArgs struct {
	SkillLevels            ModifySkillGroupSkillLevelList
	InstanceId             string
	OutboundPhoneNumberIds ModifySkillGroupOutboundPhoneNumberIdList
	SkillGroupId           string
	Name                   string
	Description            string
	UserIds                ModifySkillGroupUserIdList
}
type ModifySkillGroupResponse struct {
	RequestId      string
	Success        core.Bool
	Code           string
	Message        string
	HttpStatusCode int
}

type ModifySkillGroupSkillLevelList []int

func (list *ModifySkillGroupSkillLevelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int)
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

type ModifySkillGroupOutboundPhoneNumberIdList []string

func (list *ModifySkillGroupOutboundPhoneNumberIdList) UnmarshalJSON(data []byte) error {
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

type ModifySkillGroupUserIdList []string

func (list *ModifySkillGroupUserIdList) UnmarshalJSON(data []byte) error {
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

func (c *CccClient) GetServiceExtensions(req *GetServiceExtensionsArgs) (resp *GetServiceExtensionsResponse, err error) {
	resp = &GetServiceExtensionsResponse{}
	err = c.Request("GetServiceExtensions", req, resp)
	return
}

type GetServiceExtensionsServiceExtension struct {
	Name   string
	Number string
}
type GetServiceExtensionsArgs struct {
	ServiceType string
	InstanceId  string
}
type GetServiceExtensionsResponse struct {
	RequestId         string
	Success           core.Bool
	Code              string
	Message           string
	HttpStatusCode    int
	ServiceExtensions GetServiceExtensionsServiceExtensionList
}

type GetServiceExtensionsServiceExtensionList []GetServiceExtensionsServiceExtension

func (list *GetServiceExtensionsServiceExtensionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetServiceExtensionsServiceExtension)
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

func (c *CccClient) DeleteSkillGroup(req *DeleteSkillGroupArgs) (resp *DeleteSkillGroupResponse, err error) {
	resp = &DeleteSkillGroupResponse{}
	err = c.Request("DeleteSkillGroup", req, resp)
	return
}

type DeleteSkillGroupArgs struct {
	InstanceId   string
	SkillGroupId string
}
type DeleteSkillGroupResponse struct {
	RequestId      string
	Success        core.Bool
	Code           string
	Message        string
	HttpStatusCode int
}

func (c *CccClient) GetConfig(req *GetConfigArgs) (resp *GetConfigResponse, err error) {
	resp = &GetConfigResponse{}
	err = c.Request("GetConfig", req, resp)
	return
}

type GetConfigConfigItem struct {
	Name  string
	Value string
}
type GetConfigArgs struct {
	InstanceId string
	Name       string
	ObjectType string
	ObjectId   string
}
type GetConfigResponse struct {
	RequestId      string
	Success        core.Bool
	Code           string
	Message        string
	HttpStatusCode int
	ConfigItem     GetConfigConfigItem
}

func (c *CccClient) ListSkillGroups(req *ListSkillGroupsArgs) (resp *ListSkillGroupsResponse, err error) {
	resp = &ListSkillGroupsResponse{}
	err = c.Request("ListSkillGroups", req, resp)
	return
}

type ListSkillGroupsSkillGroup struct {
	SkillGroupId          string
	InstanceId            string
	SkillGroupName        string
	AccSkillGroupName     string
	AccQueueName          string
	SkillGroupDescription string
	UserCount             int
	OutboundPhoneNumbers  ListSkillGroupsPhoneNumberList
}

type ListSkillGroupsPhoneNumber struct {
	PhoneNumberId          string
	InstanceId             string
	Number                 string
	PhoneNumberDescription string
	TestOnly               core.Bool
	RemainingTime          int
	AllowOutbound          core.Bool
	Usage                  string
	Trunks                 int
}
type ListSkillGroupsArgs struct {
	InstanceId string
}
type ListSkillGroupsResponse struct {
	RequestId      string
	Success        core.Bool
	Code           string
	Message        string
	HttpStatusCode int
	SkillGroups    ListSkillGroupsSkillGroupList
}

type ListSkillGroupsPhoneNumberList []ListSkillGroupsPhoneNumber

func (list *ListSkillGroupsPhoneNumberList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListSkillGroupsPhoneNumber)
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

type ListSkillGroupsSkillGroupList []ListSkillGroupsSkillGroup

func (list *ListSkillGroupsSkillGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListSkillGroupsSkillGroup)
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

func (c *CccClient) ListSkillGroupsOfUser(req *ListSkillGroupsOfUserArgs) (resp *ListSkillGroupsOfUserResponse, err error) {
	resp = &ListSkillGroupsOfUserResponse{}
	err = c.Request("ListSkillGroupsOfUser", req, resp)
	return
}

type ListSkillGroupsOfUserSkillLevel struct {
	SkillLevelId string
	Level        int
	Skill        ListSkillGroupsOfUserSkill
}

type ListSkillGroupsOfUserSkill struct {
	SkillGroupId          string
	InstanceId            string
	SkillGroupName        string
	SkillGroupDescription string
	OutboundPhoneNumbers  ListSkillGroupsOfUserPhoneNumberList
}

type ListSkillGroupsOfUserPhoneNumber struct {
	PhoneNumberId          string
	InstanceId             string
	Number                 string
	PhoneNumberDescription string
	TestOnly               core.Bool
	RemainingTime          int
	AllowOutbound          core.Bool
	Usage                  string
	Trunks                 int
}
type ListSkillGroupsOfUserArgs struct {
	InstanceId string
	UserId     string
}
type ListSkillGroupsOfUserResponse struct {
	RequestId      string
	Success        core.Bool
	Code           string
	Message        string
	HttpStatusCode int
	SkillLevels    ListSkillGroupsOfUserSkillLevelList
}

type ListSkillGroupsOfUserPhoneNumberList []ListSkillGroupsOfUserPhoneNumber

func (list *ListSkillGroupsOfUserPhoneNumberList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListSkillGroupsOfUserPhoneNumber)
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

type ListSkillGroupsOfUserSkillLevelList []ListSkillGroupsOfUserSkillLevel

func (list *ListSkillGroupsOfUserSkillLevelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListSkillGroupsOfUserSkillLevel)
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

func (c *CccClient) ModifyUser(req *ModifyUserArgs) (resp *ModifyUserResponse, err error) {
	resp = &ModifyUserResponse{}
	err = c.Request("ModifyUser", req, resp)
	return
}

type ModifyUserArgs struct {
	SkillLevels   ModifyUserSkillLevelList
	InstanceId    string
	Phone         string
	RoleIds       ModifyUserRoleIdList
	DisplayName   string
	SkillGroupIds ModifyUserSkillGroupIdList
	UserId        string
	Email         string
}
type ModifyUserResponse struct {
	RequestId      string
	Success        core.Bool
	Code           string
	Message        string
	HttpStatusCode int
}

type ModifyUserSkillLevelList []int

func (list *ModifyUserSkillLevelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int)
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

type ModifyUserRoleIdList []string

func (list *ModifyUserRoleIdList) UnmarshalJSON(data []byte) error {
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

type ModifyUserSkillGroupIdList []string

func (list *ModifyUserSkillGroupIdList) UnmarshalJSON(data []byte) error {
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

func (c *CccClient) RefreshToken(req *RefreshTokenArgs) (resp *RefreshTokenResponse, err error) {
	resp = &RefreshTokenResponse{}
	err = c.Request("RefreshToken", req, resp)
	return
}

type RefreshTokenToken struct {
	Signature string
	SignData  string
}
type RefreshTokenArgs struct {
	InstanceId string
}
type RefreshTokenResponse struct {
	RequestId      string
	Success        core.Bool
	Code           string
	Message        string
	HttpStatusCode int
	Token          RefreshTokenToken
}

func (c *CccClient) RemovePhoneNumber(req *RemovePhoneNumberArgs) (resp *RemovePhoneNumberResponse, err error) {
	resp = &RemovePhoneNumberResponse{}
	err = c.Request("RemovePhoneNumber", req, resp)
	return
}

type RemovePhoneNumberArgs struct {
	InstanceId    string
	PhoneNumberId string
}
type RemovePhoneNumberResponse struct {
	RequestId      string
	Success        core.Bool
	Code           string
	Message        string
	HttpStatusCode int
}

func (c *CccClient) AssignUsers(req *AssignUsersArgs) (resp *AssignUsersResponse, err error) {
	resp = &AssignUsersResponse{}
	err = c.Request("AssignUsers", req, resp)
	return
}

type AssignUsersArgs struct {
	UserRamIds    AssignUsersUserRamIdList
	SkillLevels   AssignUsersSkillLevelList
	InstanceId    string
	RoleIds       AssignUsersRoleIdList
	SkillGroupIds AssignUsersSkillGroupIdList
}
type AssignUsersResponse struct {
	RequestId      string
	Success        core.Bool
	Code           string
	Message        string
	HttpStatusCode int
}

type AssignUsersUserRamIdList []string

func (list *AssignUsersUserRamIdList) UnmarshalJSON(data []byte) error {
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

type AssignUsersSkillLevelList []int

func (list *AssignUsersSkillLevelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int)
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

type AssignUsersRoleIdList []string

func (list *AssignUsersRoleIdList) UnmarshalJSON(data []byte) error {
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

type AssignUsersSkillGroupIdList []string

func (list *AssignUsersSkillGroupIdList) UnmarshalJSON(data []byte) error {
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

func (c *CccClient) RemoveUsers(req *RemoveUsersArgs) (resp *RemoveUsersResponse, err error) {
	resp = &RemoveUsersResponse{}
	err = c.Request("RemoveUsers", req, resp)
	return
}

type RemoveUsersArgs struct {
	InstanceId string
	UserIds    RemoveUsersUserIdList
}
type RemoveUsersResponse struct {
	RequestId      string
	Success        core.Bool
	Code           string
	Message        string
	HttpStatusCode int
}

type RemoveUsersUserIdList []string

func (list *RemoveUsersUserIdList) UnmarshalJSON(data []byte) error {
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

func (c *CccClient) ListUsersOfSkillGroup(req *ListUsersOfSkillGroupArgs) (resp *ListUsersOfSkillGroupResponse, err error) {
	resp = &ListUsersOfSkillGroupResponse{}
	err = c.Request("ListUsersOfSkillGroup", req, resp)
	return
}

type ListUsersOfSkillGroupUsers struct {
	TotalCount int
	PageNumber int
	PageSize   int
	List       ListUsersOfSkillGroupUserList
}

type ListUsersOfSkillGroupUser struct {
	UserId      string
	RamId       string
	InstanceId  string
	Roles       ListUsersOfSkillGroupRoleList
	SkillLevels ListUsersOfSkillGroupSkillLevelList
	Detail      ListUsersOfSkillGroupDetail
}

type ListUsersOfSkillGroupRole struct {
	RoleId          string
	InstanceId      string
	RoleName        string
	RoleDescription string
	UserCount       int
	Privileges      ListUsersOfSkillGroupPrivilegeList
}

type ListUsersOfSkillGroupPrivilege struct {
	PrivilegeId          string
	PrivilegeName        string
	PrivilegeDescription string
}

type ListUsersOfSkillGroupSkillLevel struct {
	SkillLevelId string
	Level        int
	Skill        ListUsersOfSkillGroupSkill
}

type ListUsersOfSkillGroupSkill struct {
	SkillGroupId          string
	InstanceId            string
	SkillGroupName        string
	SkillGroupDescription string
}

type ListUsersOfSkillGroupDetail struct {
	LoginName   string
	DisplayName string
	Phone       string
	Email       string
	Department  string
}
type ListUsersOfSkillGroupArgs struct {
	InstanceId   string
	SkillGroupId string
	PageSize     int
	PageNumber   int
}
type ListUsersOfSkillGroupResponse struct {
	RequestId      string
	Success        core.Bool
	Code           string
	Message        string
	HttpStatusCode int
	Users          ListUsersOfSkillGroupUsers
}

type ListUsersOfSkillGroupUserList []ListUsersOfSkillGroupUser

func (list *ListUsersOfSkillGroupUserList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListUsersOfSkillGroupUser)
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

type ListUsersOfSkillGroupRoleList []ListUsersOfSkillGroupRole

func (list *ListUsersOfSkillGroupRoleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListUsersOfSkillGroupRole)
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

type ListUsersOfSkillGroupSkillLevelList []ListUsersOfSkillGroupSkillLevel

func (list *ListUsersOfSkillGroupSkillLevelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListUsersOfSkillGroupSkillLevel)
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

type ListUsersOfSkillGroupPrivilegeList []ListUsersOfSkillGroupPrivilege

func (list *ListUsersOfSkillGroupPrivilegeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListUsersOfSkillGroupPrivilege)
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
