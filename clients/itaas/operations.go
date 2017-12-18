package itaas

import "encoding/json"

func (c *ItaasClient) GetIPSegmentsList(req *GetIPSegmentsListArgs) (resp *GetIPSegmentsListResponse, err error) {
	resp = &GetIPSegmentsListResponse{}
	err = c.Request("GetIPSegmentsList", req, resp)
	return
}

type GetIPSegmentsListIpsegmentInfo struct {
	CreateTimeL int64
	Ipsegment   string
	Memo        string
	ModifyTimeL int64
	Uuid        string
}

type GetIPSegmentsListErrorMessage struct {
	ErrorMessage string
}
type GetIPSegmentsListArgs struct {
	Sysfrom     string
	Operator    string
	Clientappid string
}
type GetIPSegmentsListResponse struct {
	RequestId string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	Data      GetIPSegmentsListIpsegmentInfoList
	ErrorList GetIPSegmentsListErrorMessageList
}

type GetIPSegmentsListIpsegmentInfoList []GetIPSegmentsListIpsegmentInfo

func (list *GetIPSegmentsListIpsegmentInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetIPSegmentsListIpsegmentInfo)
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

type GetIPSegmentsListErrorMessageList []GetIPSegmentsListErrorMessage

func (list *GetIPSegmentsListErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetIPSegmentsListErrorMessage)
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

func (c *ItaasClient) RemoveIPSegment(req *RemoveIPSegmentArgs) (resp *RemoveIPSegmentResponse, err error) {
	resp = &RemoveIPSegmentResponse{}
	err = c.Request("RemoveIPSegment", req, resp)
	return
}

type RemoveIPSegmentErrorMessage struct {
	ErrorMessage string
}
type RemoveIPSegmentArgs struct {
	Sysfrom     string
	Operator    string
	Clientappid string
	Uuid        string
}
type RemoveIPSegmentResponse struct {
	RequestId string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	ErrorList RemoveIPSegmentErrorMessageList
}

type RemoveIPSegmentErrorMessageList []RemoveIPSegmentErrorMessage

func (list *RemoveIPSegmentErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RemoveIPSegmentErrorMessage)
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

func (c *ItaasClient) AddIPSegment(req *AddIPSegmentArgs) (resp *AddIPSegmentResponse, err error) {
	resp = &AddIPSegmentResponse{}
	err = c.Request("AddIPSegment", req, resp)
	return
}

type AddIPSegmentErrorMessage struct {
	ErrorMessage string
}
type AddIPSegmentArgs struct {
	Sysfrom     string
	Operator    string
	Clientappid string
	Ipsegment   string
	Memo        string
}
type AddIPSegmentResponse struct {
	RequestId string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	ErrorList AddIPSegmentErrorMessageList
}

type AddIPSegmentErrorMessageList []AddIPSegmentErrorMessage

func (list *AddIPSegmentErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddIPSegmentErrorMessage)
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

func (c *ItaasClient) RemoveBoxCode(req *RemoveBoxCodeArgs) (resp *RemoveBoxCodeResponse, err error) {
	resp = &RemoveBoxCodeResponse{}
	err = c.Request("RemoveBoxCode", req, resp)
	return
}

type RemoveBoxCodeErrorMessage struct {
	ErrorMessage string
}
type RemoveBoxCodeArgs struct {
	Sysfrom     string
	Operator    string
	Clientappid string
	Code        string
}
type RemoveBoxCodeResponse struct {
	RequestId string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	ErrorList RemoveBoxCodeErrorMessageList
}

type RemoveBoxCodeErrorMessageList []RemoveBoxCodeErrorMessage

func (list *RemoveBoxCodeErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RemoveBoxCodeErrorMessage)
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

func (c *ItaasClient) UpdateRoomName(req *UpdateRoomNameArgs) (resp *UpdateRoomNameResponse, err error) {
	resp = &UpdateRoomNameResponse{}
	err = c.Request("UpdateRoomName", req, resp)
	return
}

type UpdateRoomNameErrorMessage struct {
	ErrorMessage string
}
type UpdateRoomNameArgs struct {
	Sysfrom     string
	Operator    string
	Clientappid string
	Screencode  string
	Drname      string
}
type UpdateRoomNameResponse struct {
	RequestId string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	ErrorList UpdateRoomNameErrorMessageList
}

type UpdateRoomNameErrorMessageList []UpdateRoomNameErrorMessage

func (list *UpdateRoomNameErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]UpdateRoomNameErrorMessage)
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

func (c *ItaasClient) RemoveRegisterBox(req *RemoveRegisterBoxArgs) (resp *RemoveRegisterBoxResponse, err error) {
	resp = &RemoveRegisterBoxResponse{}
	err = c.Request("RemoveRegisterBox", req, resp)
	return
}

type RemoveRegisterBoxErrorMessage struct {
	ErrorMessage string
}
type RemoveRegisterBoxArgs struct {
	Sysfrom     string
	Operator    string
	Clientappid string
	Drsessionid string
}
type RemoveRegisterBoxResponse struct {
	RequestId string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	ErrorList RemoveRegisterBoxErrorMessageList
}

type RemoveRegisterBoxErrorMessageList []RemoveRegisterBoxErrorMessage

func (list *RemoveRegisterBoxErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RemoveRegisterBoxErrorMessage)
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

func (c *ItaasClient) GetWelcomePageURI(req *GetWelcomePageURIArgs) (resp *GetWelcomePageURIResponse, err error) {
	resp = &GetWelcomePageURIResponse{}
	err = c.Request("GetWelcomePageURI", req, resp)
	return
}

type GetWelcomePageURIErrorMessage struct {
	ErrorMessage string
}
type GetWelcomePageURIArgs struct {
	Sysfrom     string
	Operator    string
	Clientappid string
}
type GetWelcomePageURIResponse struct {
	RequestId string
	Data      string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	ErrorList GetWelcomePageURIErrorMessageList
}

type GetWelcomePageURIErrorMessageList []GetWelcomePageURIErrorMessage

func (list *GetWelcomePageURIErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetWelcomePageURIErrorMessage)
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

func (c *ItaasClient) GetRegisterHistoryList(req *GetRegisterHistoryListArgs) (resp *GetRegisterHistoryListResponse, err error) {
	resp = &GetRegisterHistoryListResponse{}
	err = c.Request("GetRegisterHistoryList", req, resp)
	return
}

type GetRegisterHistoryListRegisterHistoryInfo struct {
	CreateTimeL  int64
	DrIp         string
	DrMac        string
	DrName       string
	Eventinfo    string
	Eventtype    int
	EventtypeTxt string
	Memo         string
	Screencode   string
}

type GetRegisterHistoryListErrorMessage struct {
	ErrorMessage string
}
type GetRegisterHistoryListArgs struct {
	Sysfrom     string
	Operator    string
	Clientappid string
}
type GetRegisterHistoryListResponse struct {
	RequestId string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	Data      GetRegisterHistoryListRegisterHistoryInfoList
	ErrorList GetRegisterHistoryListErrorMessageList
}

type GetRegisterHistoryListRegisterHistoryInfoList []GetRegisterHistoryListRegisterHistoryInfo

func (list *GetRegisterHistoryListRegisterHistoryInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetRegisterHistoryListRegisterHistoryInfo)
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

type GetRegisterHistoryListErrorMessageList []GetRegisterHistoryListErrorMessage

func (list *GetRegisterHistoryListErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetRegisterHistoryListErrorMessage)
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

func (c *ItaasClient) GetRegisterBoxList(req *GetRegisterBoxListArgs) (resp *GetRegisterBoxListResponse, err error) {
	resp = &GetRegisterBoxListResponse{}
	err = c.Request("GetRegisterBoxList", req, resp)
	return
}

type GetRegisterBoxListErrorMessage struct {
	ErrorMessage string
}

type GetRegisterBoxListData struct {
	ActivedNumber int
	BuyNumber     int
	BoxesList     GetRegisterBoxListBoxInfoList
}

type GetRegisterBoxListBoxInfo struct {
	CurVersion      string
	DrName          string
	DrSessionId     string
	DrStatus        string
	DrStatusTxt     string
	Ipaddress       string
	LastReportTimeL int64
	OnlineTimeL     int64
	Screencode      string
	SysVersion      string
}
type GetRegisterBoxListArgs struct {
	Sysfrom     string
	Operator    string
	Clientappid string
}
type GetRegisterBoxListResponse struct {
	RequestId string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	ErrorList GetRegisterBoxListErrorMessageList
	Data      GetRegisterBoxListData
}

type GetRegisterBoxListBoxInfoList []GetRegisterBoxListBoxInfo

func (list *GetRegisterBoxListBoxInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetRegisterBoxListBoxInfo)
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

type GetRegisterBoxListErrorMessageList []GetRegisterBoxListErrorMessage

func (list *GetRegisterBoxListErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetRegisterBoxListErrorMessage)
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

func (c *ItaasClient) SetWelcomePageURI(req *SetWelcomePageURIArgs) (resp *SetWelcomePageURIResponse, err error) {
	resp = &SetWelcomePageURIResponse{}
	err = c.Request("SetWelcomePageURI", req, resp)
	return
}

type SetWelcomePageURIErrorMessage struct {
	ErrorMessage string
}
type SetWelcomePageURIArgs struct {
	Sysfrom     string
	Operator    string
	Clientappid string
	Welcomeuri  string
}
type SetWelcomePageURIResponse struct {
	RequestId string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	ErrorList SetWelcomePageURIErrorMessageList
}

type SetWelcomePageURIErrorMessageList []SetWelcomePageURIErrorMessage

func (list *SetWelcomePageURIErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SetWelcomePageURIErrorMessage)
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

func (c *ItaasClient) CreateBoxCode(req *CreateBoxCodeArgs) (resp *CreateBoxCodeResponse, err error) {
	resp = &CreateBoxCodeResponse{}
	err = c.Request("CreateBoxCode", req, resp)
	return
}

type CreateBoxCodeErrorMessage struct {
	ErrorMessage string
}

type CreateBoxCodeData struct {
	ClientAppid       string
	Code              string
	CurrentTimeMillis int64
	ExpiresIn         int
	ExpiresInUnit     string
}
type CreateBoxCodeArgs struct {
	Sysfrom     string
	Operator    string
	Clientappid string
}
type CreateBoxCodeResponse struct {
	RequestId string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	ErrorList CreateBoxCodeErrorMessageList
	Data      CreateBoxCodeData
}

type CreateBoxCodeErrorMessageList []CreateBoxCodeErrorMessage

func (list *CreateBoxCodeErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateBoxCodeErrorMessage)
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

func (c *ItaasClient) GetRegisterBoxNumber(req *GetRegisterBoxNumberArgs) (resp *GetRegisterBoxNumberResponse, err error) {
	resp = &GetRegisterBoxNumberResponse{}
	err = c.Request("GetRegisterBoxNumber", req, resp)
	return
}

type GetRegisterBoxNumberErrorMessage struct {
	ErrorMessage string
}

type GetRegisterBoxNumberData struct {
	ActivedNumber int
	BuyNumber     int
	BoxesList     GetRegisterBoxNumberBoxInfoList
}

type GetRegisterBoxNumberBoxInfo struct {
	CurVersion      string
	DrName          string
	DrSessionId     string
	DrStatus        string
	DrStatusTxt     string
	Ipaddress       string
	LastReportTimeL int64
	OnlineTimeL     int64
	Screencode      string
	SysVersion      string
}
type GetRegisterBoxNumberArgs struct {
	Sysfrom     string
	Operator    string
	Clientappid string
}
type GetRegisterBoxNumberResponse struct {
	RequestId string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	ErrorList GetRegisterBoxNumberErrorMessageList
	Data      GetRegisterBoxNumberData
}

type GetRegisterBoxNumberBoxInfoList []GetRegisterBoxNumberBoxInfo

func (list *GetRegisterBoxNumberBoxInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetRegisterBoxNumberBoxInfo)
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

type GetRegisterBoxNumberErrorMessageList []GetRegisterBoxNumberErrorMessage

func (list *GetRegisterBoxNumberErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetRegisterBoxNumberErrorMessage)
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

func (c *ItaasClient) CreateEnterprise(req *CreateEnterpriseArgs) (resp *CreateEnterpriseResponse, err error) {
	resp = &CreateEnterpriseResponse{}
	err = c.Request("CreateEnterprise", req, resp)
	return
}

type CreateEnterpriseErrorMessage struct {
	ErrorMessage string
}
type CreateEnterpriseArgs struct {
	Sysfrom     string
	Operator    string
	Clientappid string
	BoxNumber   int
	ServiceFlag bool
}
type CreateEnterpriseResponse struct {
	RequestId string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	ErrorList CreateEnterpriseErrorMessageList
}

type CreateEnterpriseErrorMessageList []CreateEnterpriseErrorMessage

func (list *CreateEnterpriseErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateEnterpriseErrorMessage)
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

func (c *ItaasClient) GetEnterpriseConfig(req *GetEnterpriseConfigArgs) (resp *GetEnterpriseConfigResponse, err error) {
	resp = &GetEnterpriseConfigResponse{}
	err = c.Request("GetEnterpriseConfig", req, resp)
	return
}

type GetEnterpriseConfigErrorMessage struct {
	ErrorMessage string
}

type GetEnterpriseConfigData struct {
	AuthorizationNeedAccessToken bool
	DrMeetingQrUrl               string
	DrWelcomeUrl                 string
	ShareMboxNubmer              int
	ShareNeedInternet            bool
	ShareServiceFlag             bool
}
type GetEnterpriseConfigArgs struct {
	Sysfrom     string
	Operator    string
	Clientappid string
}
type GetEnterpriseConfigResponse struct {
	RequestId string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	ErrorList GetEnterpriseConfigErrorMessageList
	Data      GetEnterpriseConfigData
}

type GetEnterpriseConfigErrorMessageList []GetEnterpriseConfigErrorMessage

func (list *GetEnterpriseConfigErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetEnterpriseConfigErrorMessage)
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

func (c *ItaasClient) UpdateIPSegment(req *UpdateIPSegmentArgs) (resp *UpdateIPSegmentResponse, err error) {
	resp = &UpdateIPSegmentResponse{}
	err = c.Request("UpdateIPSegment", req, resp)
	return
}

type UpdateIPSegmentErrorMessage struct {
	ErrorMessage string
}
type UpdateIPSegmentArgs struct {
	Sysfrom     string
	Operator    string
	Clientappid string
	Uuid        string
	Ipsegment   string
	Memo        string
}
type UpdateIPSegmentResponse struct {
	RequestId string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	ErrorList UpdateIPSegmentErrorMessageList
}

type UpdateIPSegmentErrorMessageList []UpdateIPSegmentErrorMessage

func (list *UpdateIPSegmentErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]UpdateIPSegmentErrorMessage)
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

func (c *ItaasClient) GetBoxCodeList(req *GetBoxCodeListArgs) (resp *GetBoxCodeListResponse, err error) {
	resp = &GetBoxCodeListResponse{}
	err = c.Request("GetBoxCodeList", req, resp)
	return
}

type GetBoxCodeListBoxCodeInfo struct {
	BeginTime  int64
	BoxInfo    string
	Code       string
	EndTime    int64
	ModifyTime int64
	Operator   string
	Screencode string
	Status     int
	StatusTxt  string
}

type GetBoxCodeListErrorMessage struct {
	ErrorMessage string
}
type GetBoxCodeListArgs struct {
	Sysfrom     string
	Operator    string
	Clientappid string
}
type GetBoxCodeListResponse struct {
	RequestId string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	Data      GetBoxCodeListBoxCodeInfoList
	ErrorList GetBoxCodeListErrorMessageList
}

type GetBoxCodeListBoxCodeInfoList []GetBoxCodeListBoxCodeInfo

func (list *GetBoxCodeListBoxCodeInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetBoxCodeListBoxCodeInfo)
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

type GetBoxCodeListErrorMessageList []GetBoxCodeListErrorMessage

func (list *GetBoxCodeListErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetBoxCodeListErrorMessage)
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

func (c *ItaasClient) UpdateEnterpriseConfig(req *UpdateEnterpriseConfigArgs) (resp *UpdateEnterpriseConfigResponse, err error) {
	resp = &UpdateEnterpriseConfigResponse{}
	err = c.Request("UpdateEnterpriseConfig", req, resp)
	return
}

type UpdateEnterpriseConfigErrorMessage struct {
	ErrorMessage string
}
type UpdateEnterpriseConfigArgs struct {
	Sysfrom     string
	Operator    string
	Clientappid string
	ConfigKey   string
	ConfigValue string
	Memo        string
}
type UpdateEnterpriseConfigResponse struct {
	RequestId string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	ErrorList UpdateEnterpriseConfigErrorMessageList
}

type UpdateEnterpriseConfigErrorMessageList []UpdateEnterpriseConfigErrorMessage

func (list *UpdateEnterpriseConfigErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]UpdateEnterpriseConfigErrorMessage)
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
