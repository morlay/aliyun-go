package ram

import "encoding/json"

func (c *RamClient) ChangePassword(req *ChangePasswordArgs) (resp *ChangePasswordResponse, err error) {
	resp = &ChangePasswordResponse{}
	err = c.Request("ChangePassword", req, resp)
	return
}

type ChangePasswordArgs struct {
	OldPassword string
	NewPassword string
}
type ChangePasswordResponse struct {
	RequestId string
}

func (c *RamClient) CreateGroup(req *CreateGroupArgs) (resp *CreateGroupResponse, err error) {
	resp = &CreateGroupResponse{}
	err = c.Request("CreateGroup", req, resp)
	return
}

type CreateGroupGroup struct {
	GroupName  string
	Comments   string
	CreateDate string
}
type CreateGroupArgs struct {
	Comments  string
	GroupName string
}
type CreateGroupResponse struct {
	RequestId string
	Group     CreateGroupGroup
}

func (c *RamClient) GetGroup(req *GetGroupArgs) (resp *GetGroupResponse, err error) {
	resp = &GetGroupResponse{}
	err = c.Request("GetGroup", req, resp)
	return
}

type GetGroupGroup struct {
	GroupName  string
	Comments   string
	CreateDate string
	UpdateDate string
}
type GetGroupArgs struct {
	GroupName string
}
type GetGroupResponse struct {
	RequestId string
	Group     GetGroupGroup
}

func (c *RamClient) CreateRole(req *CreateRoleArgs) (resp *CreateRoleResponse, err error) {
	resp = &CreateRoleResponse{}
	err = c.Request("CreateRole", req, resp)
	return
}

type CreateRoleRole struct {
	RoleId                   string
	RoleName                 string
	Arn                      string
	Description              string
	AssumeRolePolicyDocument string
	CreateDate               string
}
type CreateRoleArgs struct {
	RoleName                 string
	Description              string
	AssumeRolePolicyDocument string
}
type CreateRoleResponse struct {
	RequestId string
	Role      CreateRoleRole
}

func (c *RamClient) DetachPolicyFromUser(req *DetachPolicyFromUserArgs) (resp *DetachPolicyFromUserResponse, err error) {
	resp = &DetachPolicyFromUserResponse{}
	err = c.Request("DetachPolicyFromUser", req, resp)
	return
}

type DetachPolicyFromUserArgs struct {
	PolicyType string
	PolicyName string
	UserName   string
}
type DetachPolicyFromUserResponse struct {
	RequestId string
}

func (c *RamClient) DeleteRole(req *DeleteRoleArgs) (resp *DeleteRoleResponse, err error) {
	resp = &DeleteRoleResponse{}
	err = c.Request("DeleteRole", req, resp)
	return
}

type DeleteRoleArgs struct {
	RoleName string
}
type DeleteRoleResponse struct {
	RequestId string
}

func (c *RamClient) ListPolicyVersions(req *ListPolicyVersionsArgs) (resp *ListPolicyVersionsResponse, err error) {
	resp = &ListPolicyVersionsResponse{}
	err = c.Request("ListPolicyVersions", req, resp)
	return
}

type ListPolicyVersionsPolicyVersion struct {
	VersionId        string
	IsDefaultVersion bool
	PolicyDocument   string
	CreateDate       string
}
type ListPolicyVersionsArgs struct {
	PolicyType string
	PolicyName string
}
type ListPolicyVersionsResponse struct {
	RequestId      string
	PolicyVersions ListPolicyVersionsPolicyVersionList
}

type ListPolicyVersionsPolicyVersionList []ListPolicyVersionsPolicyVersion

func (list *ListPolicyVersionsPolicyVersionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPolicyVersionsPolicyVersion)
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

func (c *RamClient) DetachPolicyFromGroup(req *DetachPolicyFromGroupArgs) (resp *DetachPolicyFromGroupResponse, err error) {
	resp = &DetachPolicyFromGroupResponse{}
	err = c.Request("DetachPolicyFromGroup", req, resp)
	return
}

type DetachPolicyFromGroupArgs struct {
	PolicyType string
	PolicyName string
	GroupName  string
}
type DetachPolicyFromGroupResponse struct {
	RequestId string
}

func (c *RamClient) DeletePolicy(req *DeletePolicyArgs) (resp *DeletePolicyResponse, err error) {
	resp = &DeletePolicyResponse{}
	err = c.Request("DeletePolicy", req, resp)
	return
}

type DeletePolicyArgs struct {
	PolicyName string
}
type DeletePolicyResponse struct {
	RequestId string
}

func (c *RamClient) UpdateUser(req *UpdateUserArgs) (resp *UpdateUserResponse, err error) {
	resp = &UpdateUserResponse{}
	err = c.Request("UpdateUser", req, resp)
	return
}

type UpdateUserUser struct {
	UserId      string
	UserName    string
	DisplayName string
	MobilePhone string
	Email       string
	Comments    string
	CreateDate  string
	UpdateDate  string
}
type UpdateUserArgs struct {
	NewUserName    string
	NewDisplayName string
	NewMobilePhone string
	NewComments    string
	NewEmail       string
	UserName       string
}
type UpdateUserResponse struct {
	RequestId string
	User      UpdateUserUser
}

func (c *RamClient) AttachPolicyToGroup(req *AttachPolicyToGroupArgs) (resp *AttachPolicyToGroupResponse, err error) {
	resp = &AttachPolicyToGroupResponse{}
	err = c.Request("AttachPolicyToGroup", req, resp)
	return
}

type AttachPolicyToGroupArgs struct {
	PolicyType string
	PolicyName string
	GroupName  string
}
type AttachPolicyToGroupResponse struct {
	RequestId string
}

func (c *RamClient) GetPasswordPolicy(req *GetPasswordPolicyArgs) (resp *GetPasswordPolicyResponse, err error) {
	resp = &GetPasswordPolicyResponse{}
	err = c.Request("GetPasswordPolicy", req, resp)
	return
}

type GetPasswordPolicyPasswordPolicy struct {
	MinimumPasswordLength      int
	RequireLowercaseCharacters bool
	RequireUppercaseCharacters bool
	RequireNumbers             bool
	RequireSymbols             bool
	HardExpiry                 bool
	MaxPasswordAge             int
	PasswordReusePrevention    int
	MaxLoginAttemps            int
}
type GetPasswordPolicyArgs struct {
}
type GetPasswordPolicyResponse struct {
	RequestId      string
	PasswordPolicy GetPasswordPolicyPasswordPolicy
}

func (c *RamClient) SetDefaultPolicyVersion(req *SetDefaultPolicyVersionArgs) (resp *SetDefaultPolicyVersionResponse, err error) {
	resp = &SetDefaultPolicyVersionResponse{}
	err = c.Request("SetDefaultPolicyVersion", req, resp)
	return
}

type SetDefaultPolicyVersionArgs struct {
	VersionId  string
	PolicyName string
}
type SetDefaultPolicyVersionResponse struct {
	RequestId string
}

func (c *RamClient) DeleteVirtualMFADevice(req *DeleteVirtualMFADeviceArgs) (resp *DeleteVirtualMFADeviceResponse, err error) {
	resp = &DeleteVirtualMFADeviceResponse{}
	err = c.Request("DeleteVirtualMFADevice", req, resp)
	return
}

type DeleteVirtualMFADeviceArgs struct {
	SerialNumber string
}
type DeleteVirtualMFADeviceResponse struct {
	RequestId string
}

func (c *RamClient) DeletePublicKey(req *DeletePublicKeyArgs) (resp *DeletePublicKeyResponse, err error) {
	resp = &DeletePublicKeyResponse{}
	err = c.Request("DeletePublicKey", req, resp)
	return
}

type DeletePublicKeyArgs struct {
	UserPublicKeyId string
	UserName        string
}
type DeletePublicKeyResponse struct {
	RequestId string
}

func (c *RamClient) CreateLoginProfile(req *CreateLoginProfileArgs) (resp *CreateLoginProfileResponse, err error) {
	resp = &CreateLoginProfileResponse{}
	err = c.Request("CreateLoginProfile", req, resp)
	return
}

type CreateLoginProfileLoginProfile struct {
	UserName              string
	PasswordResetRequired bool
	MFABindRequired       bool
	CreateDate            string
}
type CreateLoginProfileArgs struct {
	Password              string
	PasswordResetRequired bool
	MFABindRequired       bool
	UserName              string
}
type CreateLoginProfileResponse struct {
	RequestId    string
	LoginProfile CreateLoginProfileLoginProfile
}

func (c *RamClient) GetRole(req *GetRoleArgs) (resp *GetRoleResponse, err error) {
	resp = &GetRoleResponse{}
	err = c.Request("GetRole", req, resp)
	return
}

type GetRoleRole struct {
	RoleId                   string
	RoleName                 string
	Arn                      string
	Description              string
	AssumeRolePolicyDocument string
	CreateDate               string
	UpdateDate               string
}
type GetRoleArgs struct {
	RoleName string
}
type GetRoleResponse struct {
	RequestId string
	Role      GetRoleRole
}

func (c *RamClient) RemoveUserFromGroup(req *RemoveUserFromGroupArgs) (resp *RemoveUserFromGroupResponse, err error) {
	resp = &RemoveUserFromGroupResponse{}
	err = c.Request("RemoveUserFromGroup", req, resp)
	return
}

type RemoveUserFromGroupArgs struct {
	GroupName string
	UserName  string
}
type RemoveUserFromGroupResponse struct {
	RequestId string
}

func (c *RamClient) GetPublicKey(req *GetPublicKeyArgs) (resp *GetPublicKeyResponse, err error) {
	resp = &GetPublicKeyResponse{}
	err = c.Request("GetPublicKey", req, resp)
	return
}

type GetPublicKeyPublicKey struct {
	PublicKeyId   string
	PublicKeySpec string
	Status        string
	CreateDate    string
}
type GetPublicKeyArgs struct {
	UserPublicKeyId string
	UserName        string
}
type GetPublicKeyResponse struct {
	RequestId string
	PublicKey GetPublicKeyPublicKey
}

func (c *RamClient) UploadPublicKey(req *UploadPublicKeyArgs) (resp *UploadPublicKeyResponse, err error) {
	resp = &UploadPublicKeyResponse{}
	err = c.Request("UploadPublicKey", req, resp)
	return
}

type UploadPublicKeyPublicKey struct {
	PublicKeyId   string
	PublicKeySpec string
	Status        string
	CreateDate    string
}
type UploadPublicKeyArgs struct {
	PublicKeySpec string
	UserName      string
}
type UploadPublicKeyResponse struct {
	RequestId string
	PublicKey UploadPublicKeyPublicKey
}

func (c *RamClient) UpdateGroup(req *UpdateGroupArgs) (resp *UpdateGroupResponse, err error) {
	resp = &UpdateGroupResponse{}
	err = c.Request("UpdateGroup", req, resp)
	return
}

type UpdateGroupGroup struct {
	GroupName  string
	Comments   string
	CreateDate string
	UpdateDate string
}
type UpdateGroupArgs struct {
	NewGroupName string
	NewComments  string
	GroupName    string
}
type UpdateGroupResponse struct {
	RequestId string
	Group     UpdateGroupGroup
}

func (c *RamClient) ListEntitiesForPolicy(req *ListEntitiesForPolicyArgs) (resp *ListEntitiesForPolicyResponse, err error) {
	resp = &ListEntitiesForPolicyResponse{}
	err = c.Request("ListEntitiesForPolicy", req, resp)
	return
}

type ListEntitiesForPolicyGroup struct {
	GroupName  string
	Comments   string
	AttachDate string
}

type ListEntitiesForPolicyUser struct {
	UserId      string
	UserName    string
	DisplayName string
	AttachDate  string
}

type ListEntitiesForPolicyRole struct {
	RoleId      string
	RoleName    string
	Arn         string
	Description string
	AttachDate  string
}
type ListEntitiesForPolicyArgs struct {
	PolicyType string
	PolicyName string
}
type ListEntitiesForPolicyResponse struct {
	RequestId string
	Groups    ListEntitiesForPolicyGroupList
	Users     ListEntitiesForPolicyUserList
	Roles     ListEntitiesForPolicyRoleList
}

type ListEntitiesForPolicyGroupList []ListEntitiesForPolicyGroup

func (list *ListEntitiesForPolicyGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListEntitiesForPolicyGroup)
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

type ListEntitiesForPolicyUserList []ListEntitiesForPolicyUser

func (list *ListEntitiesForPolicyUserList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListEntitiesForPolicyUser)
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

type ListEntitiesForPolicyRoleList []ListEntitiesForPolicyRole

func (list *ListEntitiesForPolicyRoleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListEntitiesForPolicyRole)
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

func (c *RamClient) UpdatePublicKey(req *UpdatePublicKeyArgs) (resp *UpdatePublicKeyResponse, err error) {
	resp = &UpdatePublicKeyResponse{}
	err = c.Request("UpdatePublicKey", req, resp)
	return
}

type UpdatePublicKeyArgs struct {
	UserPublicKeyId string
	UserName        string
	Status          string
}
type UpdatePublicKeyResponse struct {
	RequestId string
}

func (c *RamClient) ListAccessKeys(req *ListAccessKeysArgs) (resp *ListAccessKeysResponse, err error) {
	resp = &ListAccessKeysResponse{}
	err = c.Request("ListAccessKeys", req, resp)
	return
}

type ListAccessKeysAccessKey struct {
	AccessKeyId string
	Status      string
	CreateDate  string
}
type ListAccessKeysArgs struct {
	UserName string
}
type ListAccessKeysResponse struct {
	RequestId  string
	AccessKeys ListAccessKeysAccessKeyList
}

type ListAccessKeysAccessKeyList []ListAccessKeysAccessKey

func (list *ListAccessKeysAccessKeyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAccessKeysAccessKey)
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

func (c *RamClient) BindMFADevice(req *BindMFADeviceArgs) (resp *BindMFADeviceResponse, err error) {
	resp = &BindMFADeviceResponse{}
	err = c.Request("BindMFADevice", req, resp)
	return
}

type BindMFADeviceArgs struct {
	SerialNumber        string
	AuthenticationCode2 string
	AuthenticationCode1 string
	UserName            string
}
type BindMFADeviceResponse struct {
	RequestId string
}

func (c *RamClient) DetachPolicyFromRole(req *DetachPolicyFromRoleArgs) (resp *DetachPolicyFromRoleResponse, err error) {
	resp = &DetachPolicyFromRoleResponse{}
	err = c.Request("DetachPolicyFromRole", req, resp)
	return
}

type DetachPolicyFromRoleArgs struct {
	PolicyType string
	RoleName   string
	PolicyName string
}
type DetachPolicyFromRoleResponse struct {
	RequestId string
}

func (c *RamClient) GetSecurityPreference(req *GetSecurityPreferenceArgs) (resp *GetSecurityPreferenceResponse, err error) {
	resp = &GetSecurityPreferenceResponse{}
	err = c.Request("GetSecurityPreference", req, resp)
	return
}

type GetSecurityPreferenceSecurityPreference struct {
	LoginProfilePreference GetSecurityPreferenceLoginProfilePreference
	AccessKeyPreference    GetSecurityPreferenceAccessKeyPreference
	PublicKeyPreference    GetSecurityPreferencePublicKeyPreference
	MFAPreference          GetSecurityPreferenceMFAPreference
}

type GetSecurityPreferenceLoginProfilePreference struct {
	EnableSaveMFATicket       bool
	AllowUserToChangePassword bool
}

type GetSecurityPreferenceAccessKeyPreference struct {
	AllowUserToManageAccessKeys bool
}

type GetSecurityPreferencePublicKeyPreference struct {
	AllowUserToManagePublicKeys bool
}

type GetSecurityPreferenceMFAPreference struct {
	AllowUserToManageMFADevices bool
}
type GetSecurityPreferenceArgs struct {
}
type GetSecurityPreferenceResponse struct {
	RequestId          string
	SecurityPreference GetSecurityPreferenceSecurityPreference
}

func (c *RamClient) DeleteUser(req *DeleteUserArgs) (resp *DeleteUserResponse, err error) {
	resp = &DeleteUserResponse{}
	err = c.Request("DeleteUser", req, resp)
	return
}

type DeleteUserArgs struct {
	UserName string
}
type DeleteUserResponse struct {
	RequestId string
}

func (c *RamClient) CreatePolicyVersion(req *CreatePolicyVersionArgs) (resp *CreatePolicyVersionResponse, err error) {
	resp = &CreatePolicyVersionResponse{}
	err = c.Request("CreatePolicyVersion", req, resp)
	return
}

type CreatePolicyVersionPolicyVersion struct {
	VersionId        string
	IsDefaultVersion bool
	PolicyDocument   string
	CreateDate       string
}
type CreatePolicyVersionArgs struct {
	SetAsDefault   bool
	PolicyName     string
	PolicyDocument string
}
type CreatePolicyVersionResponse struct {
	RequestId     string
	PolicyVersion CreatePolicyVersionPolicyVersion
}

func (c *RamClient) UpdateLoginProfile(req *UpdateLoginProfileArgs) (resp *UpdateLoginProfileResponse, err error) {
	resp = &UpdateLoginProfileResponse{}
	err = c.Request("UpdateLoginProfile", req, resp)
	return
}

type UpdateLoginProfileArgs struct {
	Password              string
	PasswordResetRequired bool
	MFABindRequired       bool
	UserName              string
}
type UpdateLoginProfileResponse struct {
	RequestId string
}

func (c *RamClient) CreateUser(req *CreateUserArgs) (resp *CreateUserResponse, err error) {
	resp = &CreateUserResponse{}
	err = c.Request("CreateUser", req, resp)
	return
}

type CreateUserUser struct {
	UserId      string
	UserName    string
	DisplayName string
	MobilePhone string
	Email       string
	Comments    string
	CreateDate  string
}
type CreateUserArgs struct {
	Comments    string
	DisplayName string
	MobilePhone string
	Email       string
	UserName    string
}
type CreateUserResponse struct {
	RequestId string
	User      CreateUserUser
}

func (c *RamClient) ListUsers(req *ListUsersArgs) (resp *ListUsersResponse, err error) {
	resp = &ListUsersResponse{}
	err = c.Request("ListUsers", req, resp)
	return
}

type ListUsersUser struct {
	UserId      string
	UserName    string
	DisplayName string
	MobilePhone string
	Email       string
	Comments    string
	CreateDate  string
	UpdateDate  string
}
type ListUsersArgs struct {
	Marker   string
	MaxItems int
}
type ListUsersResponse struct {
	RequestId   string
	IsTruncated bool
	Marker      string
	Users       ListUsersUserList
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

func (c *RamClient) GetPolicy(req *GetPolicyArgs) (resp *GetPolicyResponse, err error) {
	resp = &GetPolicyResponse{}
	err = c.Request("GetPolicy", req, resp)
	return
}

type GetPolicyPolicy struct {
	PolicyName      string
	PolicyType      string
	Description     string
	DefaultVersion  string
	PolicyDocument  string
	CreateDate      string
	UpdateDate      string
	AttachmentCount int
}
type GetPolicyArgs struct {
	PolicyType string
	PolicyName string
}
type GetPolicyResponse struct {
	RequestId string
	Policy    GetPolicyPolicy
}

func (c *RamClient) ListRoles(req *ListRolesArgs) (resp *ListRolesResponse, err error) {
	resp = &ListRolesResponse{}
	err = c.Request("ListRoles", req, resp)
	return
}

type ListRolesRole struct {
	RoleId      string
	RoleName    string
	Arn         string
	Description string
	CreateDate  string
	UpdateDate  string
}
type ListRolesArgs struct {
	Marker   string
	MaxItems int
}
type ListRolesResponse struct {
	RequestId   string
	IsTruncated bool
	Marker      string
	Roles       ListRolesRoleList
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

func (c *RamClient) DeletePolicyVersion(req *DeletePolicyVersionArgs) (resp *DeletePolicyVersionResponse, err error) {
	resp = &DeletePolicyVersionResponse{}
	err = c.Request("DeletePolicyVersion", req, resp)
	return
}

type DeletePolicyVersionArgs struct {
	VersionId  string
	PolicyName string
}
type DeletePolicyVersionResponse struct {
	RequestId string
}

func (c *RamClient) UpdateRole(req *UpdateRoleArgs) (resp *UpdateRoleResponse, err error) {
	resp = &UpdateRoleResponse{}
	err = c.Request("UpdateRole", req, resp)
	return
}

type UpdateRoleRole struct {
	RoleId                   string
	RoleName                 string
	Arn                      string
	Description              string
	AssumeRolePolicyDocument string
	CreateDate               string
	UpdateDate               string
}
type UpdateRoleArgs struct {
	NewAssumeRolePolicyDocument string
	RoleName                    string
}
type UpdateRoleResponse struct {
	RequestId string
	Role      UpdateRoleRole
}

func (c *RamClient) GetUser(req *GetUserArgs) (resp *GetUserResponse, err error) {
	resp = &GetUserResponse{}
	err = c.Request("GetUser", req, resp)
	return
}

type GetUserUser struct {
	UserId        string
	UserName      string
	DisplayName   string
	MobilePhone   string
	Email         string
	Comments      string
	CreateDate    string
	UpdateDate    string
	LastLoginDate string
}
type GetUserArgs struct {
	UserName string
}
type GetUserResponse struct {
	RequestId string
	User      GetUserUser
}

func (c *RamClient) UnbindMFADevice(req *UnbindMFADeviceArgs) (resp *UnbindMFADeviceResponse, err error) {
	resp = &UnbindMFADeviceResponse{}
	err = c.Request("UnbindMFADevice", req, resp)
	return
}

type UnbindMFADeviceMFADevice struct {
	SerialNumber string
}
type UnbindMFADeviceArgs struct {
	UserName string
}
type UnbindMFADeviceResponse struct {
	RequestId string
	MFADevice UnbindMFADeviceMFADevice
}

func (c *RamClient) CreatePolicy(req *CreatePolicyArgs) (resp *CreatePolicyResponse, err error) {
	resp = &CreatePolicyResponse{}
	err = c.Request("CreatePolicy", req, resp)
	return
}

type CreatePolicyPolicy struct {
	PolicyName     string
	PolicyType     string
	Description    string
	DefaultVersion string
	CreateDate     string
}
type CreatePolicyArgs struct {
	Description    string
	PolicyName     string
	PolicyDocument string
}
type CreatePolicyResponse struct {
	RequestId string
	Policy    CreatePolicyPolicy
}

func (c *RamClient) AddUserToGroup(req *AddUserToGroupArgs) (resp *AddUserToGroupResponse, err error) {
	resp = &AddUserToGroupResponse{}
	err = c.Request("AddUserToGroup", req, resp)
	return
}

type AddUserToGroupArgs struct {
	GroupName string
	UserName  string
}
type AddUserToGroupResponse struct {
	RequestId string
}

func (c *RamClient) ListGroupsForUser(req *ListGroupsForUserArgs) (resp *ListGroupsForUserResponse, err error) {
	resp = &ListGroupsForUserResponse{}
	err = c.Request("ListGroupsForUser", req, resp)
	return
}

type ListGroupsForUserGroup struct {
	GroupName string
	Comments  string
	JoinDate  string
}
type ListGroupsForUserArgs struct {
	UserName string
}
type ListGroupsForUserResponse struct {
	RequestId string
	Groups    ListGroupsForUserGroupList
}

type ListGroupsForUserGroupList []ListGroupsForUserGroup

func (list *ListGroupsForUserGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListGroupsForUserGroup)
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

func (c *RamClient) AttachPolicyToUser(req *AttachPolicyToUserArgs) (resp *AttachPolicyToUserResponse, err error) {
	resp = &AttachPolicyToUserResponse{}
	err = c.Request("AttachPolicyToUser", req, resp)
	return
}

type AttachPolicyToUserArgs struct {
	PolicyType string
	PolicyName string
	UserName   string
}
type AttachPolicyToUserResponse struct {
	RequestId string
}

func (c *RamClient) ListGroups(req *ListGroupsArgs) (resp *ListGroupsResponse, err error) {
	resp = &ListGroupsResponse{}
	err = c.Request("ListGroups", req, resp)
	return
}

type ListGroupsGroup struct {
	GroupName  string
	Comments   string
	CreateDate string
	UpdateDate string
}
type ListGroupsArgs struct {
	Marker   string
	MaxItems int
}
type ListGroupsResponse struct {
	RequestId   string
	IsTruncated bool
	Marker      string
	Groups      ListGroupsGroupList
}

type ListGroupsGroupList []ListGroupsGroup

func (list *ListGroupsGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListGroupsGroup)
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

func (c *RamClient) GetUserMFAInfo(req *GetUserMFAInfoArgs) (resp *GetUserMFAInfoResponse, err error) {
	resp = &GetUserMFAInfoResponse{}
	err = c.Request("GetUserMFAInfo", req, resp)
	return
}

type GetUserMFAInfoMFADevice struct {
	SerialNumber string
}
type GetUserMFAInfoArgs struct {
	UserName string
}
type GetUserMFAInfoResponse struct {
	RequestId string
	MFADevice GetUserMFAInfoMFADevice
}

func (c *RamClient) ListPoliciesForUser(req *ListPoliciesForUserArgs) (resp *ListPoliciesForUserResponse, err error) {
	resp = &ListPoliciesForUserResponse{}
	err = c.Request("ListPoliciesForUser", req, resp)
	return
}

type ListPoliciesForUserPolicy struct {
	PolicyName     string
	PolicyType     string
	Description    string
	DefaultVersion string
	AttachDate     string
}
type ListPoliciesForUserArgs struct {
	UserName string
}
type ListPoliciesForUserResponse struct {
	RequestId string
	Policies  ListPoliciesForUserPolicyList
}

type ListPoliciesForUserPolicyList []ListPoliciesForUserPolicy

func (list *ListPoliciesForUserPolicyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPoliciesForUserPolicy)
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

func (c *RamClient) ListUsersForGroup(req *ListUsersForGroupArgs) (resp *ListUsersForGroupResponse, err error) {
	resp = &ListUsersForGroupResponse{}
	err = c.Request("ListUsersForGroup", req, resp)
	return
}

type ListUsersForGroupUser struct {
	UserName    string
	DisplayName string
	JoinDate    string
}
type ListUsersForGroupArgs struct {
	GroupName string
}
type ListUsersForGroupResponse struct {
	RequestId string
	Users     ListUsersForGroupUserList
}

type ListUsersForGroupUserList []ListUsersForGroupUser

func (list *ListUsersForGroupUserList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListUsersForGroupUser)
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

func (c *RamClient) CreateAccessKey(req *CreateAccessKeyArgs) (resp *CreateAccessKeyResponse, err error) {
	resp = &CreateAccessKeyResponse{}
	err = c.Request("CreateAccessKey", req, resp)
	return
}

type CreateAccessKeyAccessKey struct {
	AccessKeyId     string
	AccessKeySecret string
	Status          string
	CreateDate      string
}
type CreateAccessKeyArgs struct {
	UserName string
}
type CreateAccessKeyResponse struct {
	RequestId string
	AccessKey CreateAccessKeyAccessKey
}

func (c *RamClient) ListVirtualMFADevices(req *ListVirtualMFADevicesArgs) (resp *ListVirtualMFADevicesResponse, err error) {
	resp = &ListVirtualMFADevicesResponse{}
	err = c.Request("ListVirtualMFADevices", req, resp)
	return
}

type ListVirtualMFADevicesVirtualMFADevice struct {
	SerialNumber string
	ActivateDate string
	User         ListVirtualMFADevicesUser
}

type ListVirtualMFADevicesUser struct {
	UserId      string
	UserName    string
	DisplayName string
}
type ListVirtualMFADevicesArgs struct {
}
type ListVirtualMFADevicesResponse struct {
	RequestId         string
	VirtualMFADevices ListVirtualMFADevicesVirtualMFADeviceList
}

type ListVirtualMFADevicesVirtualMFADeviceList []ListVirtualMFADevicesVirtualMFADevice

func (list *ListVirtualMFADevicesVirtualMFADeviceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListVirtualMFADevicesVirtualMFADevice)
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

func (c *RamClient) UpdateAccessKey(req *UpdateAccessKeyArgs) (resp *UpdateAccessKeyResponse, err error) {
	resp = &UpdateAccessKeyResponse{}
	err = c.Request("UpdateAccessKey", req, resp)
	return
}

type UpdateAccessKeyArgs struct {
	UserAccessKeyId string
	UserName        string
	Status          string
}
type UpdateAccessKeyResponse struct {
	RequestId string
}

func (c *RamClient) ListPolicies(req *ListPoliciesArgs) (resp *ListPoliciesResponse, err error) {
	resp = &ListPoliciesResponse{}
	err = c.Request("ListPolicies", req, resp)
	return
}

type ListPoliciesPolicy struct {
	PolicyName      string
	PolicyType      string
	Description     string
	DefaultVersion  string
	CreateDate      string
	UpdateDate      string
	AttachmentCount int
}
type ListPoliciesArgs struct {
	PolicyType string
	Marker     string
	MaxItems   int
}
type ListPoliciesResponse struct {
	RequestId   string
	IsTruncated bool
	Marker      string
	Policies    ListPoliciesPolicyList
}

type ListPoliciesPolicyList []ListPoliciesPolicy

func (list *ListPoliciesPolicyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPoliciesPolicy)
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

func (c *RamClient) CreateVirtualMFADevice(req *CreateVirtualMFADeviceArgs) (resp *CreateVirtualMFADeviceResponse, err error) {
	resp = &CreateVirtualMFADeviceResponse{}
	err = c.Request("CreateVirtualMFADevice", req, resp)
	return
}

type CreateVirtualMFADeviceVirtualMFADevice struct {
	SerialNumber     string
	Base32StringSeed string
	QRCodePNG        string
}
type CreateVirtualMFADeviceArgs struct {
	VirtualMFADeviceName string
}
type CreateVirtualMFADeviceResponse struct {
	RequestId        string
	VirtualMFADevice CreateVirtualMFADeviceVirtualMFADevice
}

func (c *RamClient) ListPoliciesForGroup(req *ListPoliciesForGroupArgs) (resp *ListPoliciesForGroupResponse, err error) {
	resp = &ListPoliciesForGroupResponse{}
	err = c.Request("ListPoliciesForGroup", req, resp)
	return
}

type ListPoliciesForGroupPolicy struct {
	PolicyName     string
	PolicyType     string
	Description    string
	DefaultVersion string
	AttachDate     string
}
type ListPoliciesForGroupArgs struct {
	GroupName string
}
type ListPoliciesForGroupResponse struct {
	RequestId string
	Policies  ListPoliciesForGroupPolicyList
}

type ListPoliciesForGroupPolicyList []ListPoliciesForGroupPolicy

func (list *ListPoliciesForGroupPolicyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPoliciesForGroupPolicy)
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

func (c *RamClient) DeleteLoginProfile(req *DeleteLoginProfileArgs) (resp *DeleteLoginProfileResponse, err error) {
	resp = &DeleteLoginProfileResponse{}
	err = c.Request("DeleteLoginProfile", req, resp)
	return
}

type DeleteLoginProfileArgs struct {
	UserName string
}
type DeleteLoginProfileResponse struct {
	RequestId string
}

func (c *RamClient) AttachPolicyToRole(req *AttachPolicyToRoleArgs) (resp *AttachPolicyToRoleResponse, err error) {
	resp = &AttachPolicyToRoleResponse{}
	err = c.Request("AttachPolicyToRole", req, resp)
	return
}

type AttachPolicyToRoleArgs struct {
	PolicyType string
	RoleName   string
	PolicyName string
}
type AttachPolicyToRoleResponse struct {
	RequestId string
}

func (c *RamClient) SetPasswordPolicy(req *SetPasswordPolicyArgs) (resp *SetPasswordPolicyResponse, err error) {
	resp = &SetPasswordPolicyResponse{}
	err = c.Request("SetPasswordPolicy", req, resp)
	return
}

type SetPasswordPolicyPasswordPolicy struct {
	MinimumPasswordLength      int
	RequireLowercaseCharacters bool
	RequireUppercaseCharacters bool
	RequireNumbers             bool
	RequireSymbols             bool
	HardExpiry                 bool
	MaxPasswordAge             int
	PasswordReusePrevention    int
	MaxLoginAttemps            int
}
type SetPasswordPolicyArgs struct {
	RequireNumbers             bool
	PasswordReusePrevention    int
	RequireUppercaseCharacters bool
	MaxPasswordAge             int
	MaxLoginAttemps            int
	HardExpiry                 bool
	MinimumPasswordLength      int
	RequireLowercaseCharacters bool
	RequireSymbols             bool
}
type SetPasswordPolicyResponse struct {
	RequestId      string
	PasswordPolicy SetPasswordPolicyPasswordPolicy
}

func (c *RamClient) DeleteGroup(req *DeleteGroupArgs) (resp *DeleteGroupResponse, err error) {
	resp = &DeleteGroupResponse{}
	err = c.Request("DeleteGroup", req, resp)
	return
}

type DeleteGroupArgs struct {
	GroupName string
}
type DeleteGroupResponse struct {
	RequestId string
}

func (c *RamClient) ListPublicKeys(req *ListPublicKeysArgs) (resp *ListPublicKeysResponse, err error) {
	resp = &ListPublicKeysResponse{}
	err = c.Request("ListPublicKeys", req, resp)
	return
}

type ListPublicKeysPublicKey struct {
	PublicKeyId string
	Status      string
	CreateDate  string
}
type ListPublicKeysArgs struct {
	UserName string
}
type ListPublicKeysResponse struct {
	RequestId  string
	PublicKeys ListPublicKeysPublicKeyList
}

type ListPublicKeysPublicKeyList []ListPublicKeysPublicKey

func (list *ListPublicKeysPublicKeyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPublicKeysPublicKey)
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

func (c *RamClient) GetPolicyVersion(req *GetPolicyVersionArgs) (resp *GetPolicyVersionResponse, err error) {
	resp = &GetPolicyVersionResponse{}
	err = c.Request("GetPolicyVersion", req, resp)
	return
}

type GetPolicyVersionPolicyVersion struct {
	VersionId        string
	IsDefaultVersion bool
	PolicyDocument   string
	CreateDate       string
}
type GetPolicyVersionArgs struct {
	VersionId  string
	PolicyType string
	PolicyName string
}
type GetPolicyVersionResponse struct {
	RequestId     string
	PolicyVersion GetPolicyVersionPolicyVersion
}

func (c *RamClient) DeleteAccessKey(req *DeleteAccessKeyArgs) (resp *DeleteAccessKeyResponse, err error) {
	resp = &DeleteAccessKeyResponse{}
	err = c.Request("DeleteAccessKey", req, resp)
	return
}

type DeleteAccessKeyArgs struct {
	UserAccessKeyId string
	UserName        string
}
type DeleteAccessKeyResponse struct {
	RequestId string
}

func (c *RamClient) SetAccountAlias(req *SetAccountAliasArgs) (resp *SetAccountAliasResponse, err error) {
	resp = &SetAccountAliasResponse{}
	err = c.Request("SetAccountAlias", req, resp)
	return
}

type SetAccountAliasArgs struct {
	AccountAlias string
}
type SetAccountAliasResponse struct {
	RequestId string
}

func (c *RamClient) SetSecurityPreference(req *SetSecurityPreferenceArgs) (resp *SetSecurityPreferenceResponse, err error) {
	resp = &SetSecurityPreferenceResponse{}
	err = c.Request("SetSecurityPreference", req, resp)
	return
}

type SetSecurityPreferenceSecurityPreference struct {
	LoginProfilePreference SetSecurityPreferenceLoginProfilePreference
	AccessKeyPreference    SetSecurityPreferenceAccessKeyPreference
	PublicKeyPreference    SetSecurityPreferencePublicKeyPreference
	MFAPreference          SetSecurityPreferenceMFAPreference
}

type SetSecurityPreferenceLoginProfilePreference struct {
	EnableSaveMFATicket       bool
	AllowUserToChangePassword bool
}

type SetSecurityPreferenceAccessKeyPreference struct {
	AllowUserToManageAccessKeys bool
}

type SetSecurityPreferencePublicKeyPreference struct {
	AllowUserToManagePublicKeys bool
}

type SetSecurityPreferenceMFAPreference struct {
	AllowUserToManageMFADevices bool
}
type SetSecurityPreferenceArgs struct {
	AllowUserToManageAccessKeys bool
	AllowUserToManageMFADevices bool
	AllowUserToManagePublicKeys bool
	EnableSaveMFATicket         bool
	AllowUserToChangePassword   bool
}
type SetSecurityPreferenceResponse struct {
	RequestId          string
	SecurityPreference SetSecurityPreferenceSecurityPreference
}

func (c *RamClient) GetLoginProfile(req *GetLoginProfileArgs) (resp *GetLoginProfileResponse, err error) {
	resp = &GetLoginProfileResponse{}
	err = c.Request("GetLoginProfile", req, resp)
	return
}

type GetLoginProfileLoginProfile struct {
	UserName              string
	PasswordResetRequired bool
	MFABindRequired       bool
	CreateDate            string
}
type GetLoginProfileArgs struct {
	UserName string
}
type GetLoginProfileResponse struct {
	RequestId    string
	LoginProfile GetLoginProfileLoginProfile
}

func (c *RamClient) GetAccountAlias(req *GetAccountAliasArgs) (resp *GetAccountAliasResponse, err error) {
	resp = &GetAccountAliasResponse{}
	err = c.Request("GetAccountAlias", req, resp)
	return
}

type GetAccountAliasArgs struct {
}
type GetAccountAliasResponse struct {
	RequestId    string
	AccountAlias string
}

func (c *RamClient) ClearAccountAlias(req *ClearAccountAliasArgs) (resp *ClearAccountAliasResponse, err error) {
	resp = &ClearAccountAliasResponse{}
	err = c.Request("ClearAccountAlias", req, resp)
	return
}

type ClearAccountAliasArgs struct {
}
type ClearAccountAliasResponse struct {
	RequestId string
}

func (c *RamClient) ListPoliciesForRole(req *ListPoliciesForRoleArgs) (resp *ListPoliciesForRoleResponse, err error) {
	resp = &ListPoliciesForRoleResponse{}
	err = c.Request("ListPoliciesForRole", req, resp)
	return
}

type ListPoliciesForRolePolicy struct {
	PolicyName     string
	PolicyType     string
	Description    string
	DefaultVersion string
	AttachDate     string
}
type ListPoliciesForRoleArgs struct {
	RoleName string
}
type ListPoliciesForRoleResponse struct {
	RequestId string
	Policies  ListPoliciesForRolePolicyList
}

type ListPoliciesForRolePolicyList []ListPoliciesForRolePolicy

func (list *ListPoliciesForRolePolicyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPoliciesForRolePolicy)
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
