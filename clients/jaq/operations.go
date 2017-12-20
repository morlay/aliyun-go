package jaq

import (
	"github.com/morlay/aliyun-go/core"
)

func (c *JaqClient) ActivityPrevention(req *ActivityPreventionArgs) (resp *ActivityPreventionResponse, err error) {
	resp = &ActivityPreventionResponse{}
	err = c.Request("ActivityPrevention", req, resp)
	return
}

type ActivityPreventionData struct {
	FnalDecision     int
	EventId          string
	UserId           string
	FinalScore       int
	FinalDesc        string
	Detail           string
	CaptchaCheckData string
}
type ActivityPreventionArgs struct {
	CallerName          string
	Ip                  string
	ProtocolVersion     string
	Source              int
	ActivityDescription string
	ActivityId          string
	Prize               string
	PrizeType           int
	PhoneNumber         string
	Email               string
	UserId              string
	IdType              int
	CurrentUrl          string
	Agent               string
	Cookie              string
	SessionId           string
	MacAddress          string
	Referer             string
	UserName            string
	CompanyName         string
	Address             string
	IDNumber            string
	BankCardNumber      string
	RegisterIp          string
	RegisterDate        int64
	ExtendData          string
	JsToken             string
	SDKToken            string
}
type ActivityPreventionResponse struct {
	ErrorCode int
	ErrorMsg  string
	Data      ActivityPreventionData
}

func (c *JaqClient) CheckAccountAndPasswordRisk(req *CheckAccountAndPasswordRiskArgs) (resp *CheckAccountAndPasswordRiskResponse, err error) {
	resp = &CheckAccountAndPasswordRiskResponse{}
	err = c.Request("CheckAccountAndPasswordRisk", req, resp)
	return
}

type CheckAccountAndPasswordRiskData struct {
	FnalDecision     int
	EventId          string
	UserId           string
	FinalScore       int
	FinalDesc        string
	Detail           string
	CaptchaCheckData string
}
type CheckAccountAndPasswordRiskArgs struct {
	CallerName   string
	AccountName  string
	PasswordHash string
}
type CheckAccountAndPasswordRiskResponse struct {
	ErrorCode int
	ErrorMsg  string
	Data      CheckAccountAndPasswordRiskData
}

func (c *JaqClient) AfsAppCheck(req *AfsAppCheckArgs) (resp *AfsAppCheckResponse, err error) {
	resp = &AfsAppCheckResponse{}
	err = c.Request("AfsAppCheck", req, resp)
	return
}

type AfsAppCheckData struct {
	SecondCheckResult int
}
type AfsAppCheckArgs struct {
	CallerName string
	Session    string
}
type AfsAppCheckResponse struct {
	ErrorCode int
	ErrorMsg  string
	Data      AfsAppCheckData
}

func (c *JaqClient) LoginPrevention(req *LoginPreventionArgs) (resp *LoginPreventionResponse, err error) {
	resp = &LoginPreventionResponse{}
	err = c.Request("LoginPrevention", req, resp)
	return
}

type LoginPreventionData struct {
	FnalDecision     int
	EventId          string
	UserId           string
	FinalScore       int
	FinalDesc        string
	Detail           string
	CaptchaCheckData string
}
type LoginPreventionArgs struct {
	CallerName      string
	Ip              string
	ProtocolVersion string
	Source          int
	PhoneNumber     string
	Email           string
	UserId          string
	IdType          int
	CurrentUrl      string
	Agent           string
	Cookie          string
	SessionId       string
	MacAddress      string
	Referer         string
	UserName        string
	CompanyName     string
	Address         string
	IDNumber        string
	BankCardNumber  string
	RegisterIp      string
	RegisterDate    int64
	AccountExist    int
	ExtendData      string
	JsToken         string
	SDKToken        string
	PasswordHash    string
	LoginType       int
	PasswordCorrect int
}
type LoginPreventionResponse struct {
	ErrorCode int
	ErrorMsg  string
	Data      LoginPreventionData
}

func (c *JaqClient) BbsPrevention(req *BbsPreventionArgs) (resp *BbsPreventionResponse, err error) {
	resp = &BbsPreventionResponse{}
	err = c.Request("BbsPrevention", req, resp)
	return
}

type BbsPreventionData struct {
	FnalDecision     int
	EventId          string
	UserId           string
	FinalScore       int
	FinalDesc        string
	Detail           string
	CaptchaCheckData string
}
type BbsPreventionArgs struct {
	CallerName      string
	Ip              string
	ProtocolVersion string
	Source          int
	PhoneNumber     string
	Email           string
	UserId          string
	IdType          int
	CurrentUrl      string
	Agent           string
	Cookie          string
	SessionId       string
	MacAddress      string
	Referer         string
	NickName        string
	CompanyName     string
	Address         string
	JsToken         string
	SDKToken        string
	ExtendData      string
}
type BbsPreventionResponse struct {
	ErrorCode int
	ErrorMsg  string
	Data      BbsPreventionData
}

func (c *JaqClient) OtherPrevention(req *OtherPreventionArgs) (resp *OtherPreventionResponse, err error) {
	resp = &OtherPreventionResponse{}
	err = c.Request("OtherPrevention", req, resp)
	return
}

type OtherPreventionData struct {
	FnalDecision     int
	EventId          string
	UserId           string
	FinalScore       int
	FinalDesc        string
	Detail           string
	CaptchaCheckData string
}
type OtherPreventionArgs struct {
	CallerName      string
	Ip              string
	ProtocolVersion string
	Source          int
	PhoneNumber     string
	Email           string
	UserId          string
	IdType          int
	CurrentUrl      string
	Agent           string
	Cookie          string
	SessionId       string
	MacAddress      string
	Referer         string
	UserName        string
	CompanyName     string
	Address         string
	IDNumber        string
	BankCardNumber  string
	RegisterIp      string
	RegisterDate    int64
	LoginIp         string
	LoginDate       int64
	ExtendData      string
	PasswordHash    string
	JsToken         string
	SDKToken        string
}
type OtherPreventionResponse struct {
	ErrorCode int
	ErrorMsg  string
	Data      OtherPreventionData
}

func (c *JaqClient) SpamRegisterPrevention(req *SpamRegisterPreventionArgs) (resp *SpamRegisterPreventionResponse, err error) {
	resp = &SpamRegisterPreventionResponse{}
	err = c.Request("SpamRegisterPrevention", req, resp)
	return
}

type SpamRegisterPreventionData struct {
	FnalDecision     int
	EventId          string
	UserId           string
	FinalScore       int
	FinalDesc        string
	Detail           string
	CaptchaCheckData string
}
type SpamRegisterPreventionArgs struct {
	CallerName      string
	Ip              string
	ProtocolVersion string
	Source          int
	PhoneNumber     string
	Email           string
	UserId          string
	IdType          int
	CurrentUrl      string
	Agent           string
	Cookie          string
	SessionId       string
	MacAddress      string
	Referer         string
	NickName        string
	CompanyName     string
	Address         string
	IDNumber        string
	BankCardNumber  string
	JsToken         string
	SDKToken        string
	ExtendData      string
}
type SpamRegisterPreventionResponse struct {
	ErrorCode int
	ErrorMsg  string
	Data      SpamRegisterPreventionData
}

func (c *JaqClient) AfsCheck(req *AfsCheckArgs) (resp *AfsCheckResponse, err error) {
	resp = &AfsCheckResponse{}
	err = c.Request("AfsCheck", req, resp)
	return
}

type AfsCheckArgs struct {
	CallerName string
	Session    string
	Token      string
	Sig        string
	Platform   int
	Scene      string
}
type AfsCheckResponse struct {
	ErrorCode int
	ErrorMsg  string
	Data      core.Bool
}
