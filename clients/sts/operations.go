package sts

func (c *StsClient) GetCallerIdentity(req *GetCallerIdentityArgs) (resp *GetCallerIdentityResponse, err error) {
	resp = &GetCallerIdentityResponse{}
	err = c.Request("GetCallerIdentity", req, resp)
	return
}

type GetCallerIdentityArgs struct {
}
type GetCallerIdentityResponse struct {
	AccountId string
	UserId    string
	Arn       string
	RequestId string
}

func (c *StsClient) AssumeRole(req *AssumeRoleArgs) (resp *AssumeRoleResponse, err error) {
	resp = &AssumeRoleResponse{}
	err = c.Request("AssumeRole", req, resp)
	return
}

type AssumeRoleCredentials struct {
	SecurityToken   string
	AccessKeySecret string
	AccessKeyId     string
	Expiration      string
}

type AssumeRoleAssumedRoleUser struct {
	Arn           string
	AssumedRoleId string
}
type AssumeRoleArgs struct {
	RoleArn         string
	RoleSessionName string
	DurationSeconds int64
	Policy          string
}
type AssumeRoleResponse struct {
	RequestId       string
	Credentials     AssumeRoleCredentials
	AssumedRoleUser AssumeRoleAssumedRoleUser
}

func (c *StsClient) GenerateSessionAccessKey(req *GenerateSessionAccessKeyArgs) (resp *GenerateSessionAccessKeyResponse, err error) {
	resp = &GenerateSessionAccessKeyResponse{}
	err = c.Request("GenerateSessionAccessKey", req, resp)
	return
}

type GenerateSessionAccessKeySessionAccessKey struct {
	SessionAccessKeyId     string
	SessionAccessKeySecret string
	Expiration             string
}
type GenerateSessionAccessKeyArgs struct {
	DurationSeconds int64
}
type GenerateSessionAccessKeyResponse struct {
	RequestId        string
	SessionAccessKey GenerateSessionAccessKeySessionAccessKey
}
