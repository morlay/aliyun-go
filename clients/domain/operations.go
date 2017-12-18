package domain

import "encoding/json"

func (c *DomainClient) QueryBatchTaskList(req *QueryBatchTaskListArgs) (resp *QueryBatchTaskListResponse, err error) {
	resp = &QueryBatchTaskListResponse{}
	err = c.Request("QueryBatchTaskList", req, resp)
	return
}

type QueryBatchTaskListTaskInfo struct {
	TaskType   string
	TaskNum    int
	TaskStatus string
	CreateTime string
	Clientip   string
	TaskNo     string
}
type QueryBatchTaskListArgs struct {
	BeginCreateTime string
	EndCreateTime   string
	UserClientIp    string
	PageSize        int
	Lang            string
	PageNum         int
}
type QueryBatchTaskListResponse struct {
	RequestId      string
	TotalItemNum   int
	CurrentPageNum int
	TotalPageNum   int
	PageSize       int
	PrePage        bool
	NextPage       bool
	Data           QueryBatchTaskListTaskInfoList
}

type QueryBatchTaskListTaskInfoList []QueryBatchTaskListTaskInfo

func (list *QueryBatchTaskListTaskInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryBatchTaskListTaskInfo)
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

func (c *DomainClient) QueryFailReasonList(req *QueryFailReasonListArgs) (resp *QueryFailReasonListResponse, err error) {
	resp = &QueryFailReasonListResponse{}
	err = c.Request("QueryFailReasonList", req, resp)
	return
}

type QueryFailReasonListFailRecord struct {
	Date       string
	FailReason string
}
type QueryFailReasonListArgs struct {
	SaleId            string
	UserClientIp      string
	DomainName        string
	Lang              string
	ContactTemplateId int64
}
type QueryFailReasonListResponse struct {
	RequestId string
	Data      QueryFailReasonListFailRecordList
}

type QueryFailReasonListFailRecordList []QueryFailReasonListFailRecord

func (list *QueryFailReasonListFailRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryFailReasonListFailRecord)
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

func (c *DomainClient) SaveContactTemplate(req *SaveContactTemplateArgs) (resp *SaveContactTemplateResponse, err error) {
	resp = &SaveContactTemplateResponse{}
	err = c.Request("SaveContactTemplate", req, resp)
	return
}

type SaveContactTemplateArgs struct {
	CCompany          string
	DefaultTemplate   bool
	TelArea           string
	ECompany          string
	TelMain           string
	CName             string
	CProvince         string
	ECity             string
	CCity             string
	RegType           string
	EName             string
	TelExt            string
	CVenu             string
	EProvince         string
	PostalCode        string
	UserClientIp      string
	CCountry          string
	Lang              string
	EVenu             string
	Email             string
	ContactTemplateId int64
}
type SaveContactTemplateResponse struct {
	RequestId         string
	Success           bool
	ContactTemplateId int64
}

func (c *DomainClient) SaveTaskForSubmittingDomainNameCredentialByTemplateId(req *SaveTaskForSubmittingDomainNameCredentialByTemplateIdArgs) (resp *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse, err error) {
	resp = &SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse{}
	err = c.Request("SaveTaskForSubmittingDomainNameCredentialByTemplateId", req, resp)
	return
}

type SaveTaskForSubmittingDomainNameCredentialByTemplateIdArgs struct {
	SaleId            string
	UserClientIp      string
	DomainName        string
	Lang              string
	ContactTemplateId int64
}
type SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse struct {
	RequestId string
	Success   bool
	TaskNo    string
}

func (c *DomainClient) QueryContact(req *QueryContactArgs) (resp *QueryContactResponse, err error) {
	resp = &QueryContactResponse{}
	err = c.Request("QueryContact", req, resp)
	return
}

type QueryContactArgs struct {
	ContactType  string
	UserClientIp string
	DomainName   string
	Lang         string
}
type QueryContactResponse struct {
	RequestId  string
	CreateDate string
	UpdateDate string
	CName      string
	EName      string
	CCompany   string
	ECompany   string
	CCountry   string
	CProvince  string
	EProvince  string
	CCity      string
	ECity      string
	CVenu      string
	EVenu      string
	Email      string
	TelArea    string
	PostalCode string
	TelMain    string
	TelExt     string
	RegType    string
}

func (c *DomainClient) CheckDomain(req *CheckDomainArgs) (resp *CheckDomainResponse, err error) {
	resp = &CheckDomainResponse{}
	err = c.Request("CheckDomain", req, resp)
	return
}

type CheckDomainArgs struct {
	DomainName string
}
type CheckDomainResponse struct {
	RequestId   string
	Name        string
	Avail       int
	Reason      string
	FeeCurrency string
	FeePeriod   int
	FeeFee      string
	RmbFee      string
	FeeCommand  string
}

func (c *DomainClient) SaveTaskForModifyingDomainDns(req *SaveTaskForModifyingDomainDnsArgs) (resp *SaveTaskForModifyingDomainDnsResponse, err error) {
	resp = &SaveTaskForModifyingDomainDnsResponse{}
	err = c.Request("SaveTaskForModifyingDomainDns", req, resp)
	return
}

type SaveTaskForModifyingDomainDnsArgs struct {
	SaleId       string
	UserClientIp string
	DomainName   string
	Lang         string
	AliyunDns    bool
	DnsLists     SaveTaskForModifyingDomainDnsDnsListList
}
type SaveTaskForModifyingDomainDnsResponse struct {
	RequestId string
	TaskNo    string
}

type SaveTaskForModifyingDomainDnsDnsListList []string

func (list *SaveTaskForModifyingDomainDnsDnsListList) UnmarshalJSON(data []byte) error {
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

func (c *DomainClient) SaveTaskForUpdatingContactByTemplateId(req *SaveTaskForUpdatingContactByTemplateIdArgs) (resp *SaveTaskForUpdatingContactByTemplateIdResponse, err error) {
	resp = &SaveTaskForUpdatingContactByTemplateIdResponse{}
	err = c.Request("SaveTaskForUpdatingContactByTemplateId", req, resp)
	return
}

type SaveTaskForUpdatingContactByTemplateIdArgs struct {
	SaleId            string
	ContactType       string
	UserClientIp      string
	DomainName        string
	AddTransferLock   bool
	Lang              string
	ContactTemplateId int64
}
type SaveTaskForUpdatingContactByTemplateIdResponse struct {
	RequestId string
	Success   bool
	TaskNo    string
}

func (c *DomainClient) QueryDomainList(req *QueryDomainListArgs) (resp *QueryDomainListResponse, err error) {
	resp = &QueryDomainListResponse{}
	err = c.Request("QueryDomainList", req, resp)
	return
}

type QueryDomainListDomain struct {
	DomainName        string
	SaleId            string
	DeadDate          string
	RegDate           string
	DomainAuditStatus string
	DomainRegType     string
	GroupId           string
	DomainType        string
	DomainStatus      string
	DeadDateStatus    string
	ProductId         string
	DeadDateLong      int64
	RegDateLong       int64
	Remark            string
	Premium           bool
}
type QueryDomainListArgs struct {
	ProductDomainType string
	RegStartDate      int64
	OrderKeyType      string
	GroupId           string
	DeadEndDate       int64
	DomainName        string
	StartDate         string
	PageNum           int
	OrderByType       string
	RegEndDate        int64
	EndDate           string
	DomainType        string
	DeadStartDate     int64
	UserClientIp      string
	PageSize          int
	Lang              string
	QueryType         string
}
type QueryDomainListResponse struct {
	RequestId      string
	TotalItemNum   int
	CurrentPageNum int
	TotalPageNum   int
	PageSize       int
	PrePage        bool
	NextPage       bool
	Data           QueryDomainListDomainList
}

type QueryDomainListDomainList []QueryDomainListDomain

func (list *QueryDomainListDomainList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryDomainListDomain)
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

func (c *DomainClient) GetWhoisInfo(req *GetWhoisInfoArgs) (resp *GetWhoisInfoResponse, err error) {
	resp = &GetWhoisInfoResponse{}
	err = c.Request("GetWhoisInfo", req, resp)
	return
}

type GetWhoisInfoDomainStatus struct {
	Status string
	Desc   string
	Tip    string
}
type GetWhoisInfoArgs struct {
	DomainName string
}
type GetWhoisInfoResponse struct {
	RequestId                  string
	ReferralURL                string
	DomainName                 string
	Registrar                  string
	RegistrarWhoisServer       string
	StatusList                 string
	FormatCreationDate         string
	FormatExpirationDate       string
	FormatUpdatedDate          string
	NameServerList             string
	UpdatedDate                string
	CreationDate               string
	ExpirationDate             string
	RegistryDomainID           string
	RegistrarURL               string
	RegistrarIANAID            string
	RegistrarAbuseContactEmail string
	RegistrarAbuseContactPhone string
	RegistryRegistrantId       string
	RegistrantName             string
	RegistrantOrganization     string
	RegistrantStreet           string
	RegistrantCity             string
	RegistrantStateProvince    string
	RegistrantPostalCode       string
	RegistrantCountry          string
	RegistrantPhone            string
	RegistrantPhoneExt         string
	RegistrantFax              string
	RegistrantFaxExt           string
	RegistrantEmail            string
	RegistryAdminID            string
	RegistryTechID             string
	AdminName                  string
	AdminOrganization          string
	AdminStreet                string
	AdminCity                  string
	AdminStateProvince         string
	AdminPostalCode            string
	AdminCountry               string
	AdminPhone                 string
	AdminPhoneExt              string
	AdminFax                   string
	AdminFaxExt                string
	AdminEmail                 string
	TechName                   string
	TechOrganization           string
	TechStreet                 string
	TechCity                   string
	TechStateProvince          string
	TechPostalCode             string
	TechCountry                string
	TechPhone                  string
	TechPhoneExt               string
	TechFax                    string
	TechFaxExt                 string
	TechEmail                  string
	OriginalInfo               string
	CacheUpdatedDate           string
	WhoisProtected             string
	DomainStatusList           GetWhoisInfoDomainStatusList
}

type GetWhoisInfoDomainStatusList []GetWhoisInfoDomainStatus

func (list *GetWhoisInfoDomainStatusList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetWhoisInfoDomainStatus)
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

func (c *DomainClient) QueryContactTemplate(req *QueryContactTemplateArgs) (resp *QueryContactTemplateResponse, err error) {
	resp = &QueryContactTemplateResponse{}
	err = c.Request("QueryContactTemplate", req, resp)
	return
}

type QueryContactTemplateContactTemplate struct {
	ContactTemplateId int64
	CreateTime        string
	UpdateTime        string
	UserId            string
	RegType           string
	DefaultTemplate   bool
	AuditStatus       string
	CName             string
	EName             string
	CCompany          string
	ECompany          string
	CCountry          string
	CProvince         string
	EProvince         string
	CCity             string
	ECity             string
	CVenu             string
	EVenu             string
	Email             string
	TelArea           string
	PostalCode        string
	TelMain           string
	TelExt            string
}
type QueryContactTemplateArgs struct {
	CCompany          string
	AuditStatus       string
	DefaultTemplate   bool
	ECompany          string
	UserClientIp      string
	PageSize          int
	Lang              string
	PageNum           int
	ContactTemplateId int64
	RegType           string
}
type QueryContactTemplateResponse struct {
	RequestId        string
	TotalItemNum     int
	CurrentPageNum   int
	TotalPageNum     int
	PageSize         int
	PrePage          bool
	NextPage         bool
	ContactTemplates QueryContactTemplateContactTemplateList
}

type QueryContactTemplateContactTemplateList []QueryContactTemplateContactTemplate

func (list *QueryContactTemplateContactTemplateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryContactTemplateContactTemplate)
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

func (c *DomainClient) QueryDomainBySaleId(req *QueryDomainBySaleIdArgs) (resp *QueryDomainBySaleIdResponse, err error) {
	resp = &QueryDomainBySaleIdResponse{}
	err = c.Request("QueryDomainBySaleId", req, resp)
	return
}

type QueryDomainBySaleIdArgs struct {
	SaleId       string
	UserClientIp string
	Lang         string
}
type QueryDomainBySaleIdResponse struct {
	UserId               string
	DomainName           string
	SaleId               string
	CreationDate         string
	ExpirationDate       string
	DomainRegType        string
	EnglishHolder        string
	ChineseHolder        string
	EnglishContactPerson string
	ChineseContactPerson string
	HolderEmail          string
	TransferOutStatus    string
	SafetyLock           string
	TransferLock         string
	WhoisProtected       bool
	Premium              bool
	Remark               string
	DnsList              QueryDomainBySaleIdDnsListList
}

type QueryDomainBySaleIdDnsListList []string

func (list *QueryDomainBySaleIdDnsListList) UnmarshalJSON(data []byte) error {
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

func (c *DomainClient) QueryBatchTaskDetailList(req *QueryBatchTaskDetailListArgs) (resp *QueryBatchTaskDetailListResponse, err error) {
	resp = &QueryBatchTaskDetailListResponse{}
	err = c.Request("QueryBatchTaskDetailList", req, resp)
	return
}

type QueryBatchTaskDetailListTaskDetail struct {
	TaskNo     string
	TaskType   string
	DomainName string
	TaskStatus string
	UpdateTime string
	TryCount   int
	ErrorMsg   string
}
type QueryBatchTaskDetailListArgs struct {
	TaskStatus   int
	SaleId       string
	UserClientIp string
	TaskNo       string
	DomainName   string
	PageSize     int
	Lang         string
	PageNum      int
}
type QueryBatchTaskDetailListResponse struct {
	RequestId      string
	TotalItemNum   int
	CurrentPageNum int
	TotalPageNum   int
	PageSize       int
	PrePage        bool
	NextPage       bool
	Data           QueryBatchTaskDetailListTaskDetailList
}

type QueryBatchTaskDetailListTaskDetailList []QueryBatchTaskDetailListTaskDetail

func (list *QueryBatchTaskDetailListTaskDetailList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryBatchTaskDetailListTaskDetail)
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

func (c *DomainClient) DeleteContactTemplate(req *DeleteContactTemplateArgs) (resp *DeleteContactTemplateResponse, err error) {
	resp = &DeleteContactTemplateResponse{}
	err = c.Request("DeleteContactTemplate", req, resp)
	return
}

type DeleteContactTemplateArgs struct {
	UserClientIp      string
	Lang              string
	ContactTemplateId int64
}
type DeleteContactTemplateResponse struct {
	RequestId string
	Success   bool
}

func (c *DomainClient) SaveTaskForSubmittingDomainNameCredential(req *SaveTaskForSubmittingDomainNameCredentialArgs) (resp *SaveTaskForSubmittingDomainNameCredentialResponse, err error) {
	resp = &SaveTaskForSubmittingDomainNameCredentialResponse{}
	err = c.Request("SaveTaskForSubmittingDomainNameCredential", req, resp)
	return
}

type SaveTaskForSubmittingDomainNameCredentialArgs struct {
	CredentialNo string
	SaleId       string
	Credential   string
	UserClientIp string
	DomainName   string
	Lang         string
}
type SaveTaskForSubmittingDomainNameCredentialResponse struct {
	RequestId string
	TaskNo    string
}

func (c *DomainClient) SaveContactTemplateCredential(req *SaveContactTemplateCredentialArgs) (resp *SaveContactTemplateCredentialResponse, err error) {
	resp = &SaveContactTemplateCredentialResponse{}
	err = c.Request("SaveContactTemplateCredential", req, resp)
	return
}

type SaveContactTemplateCredentialArgs struct {
	CredentialNo      string
	Credential        string
	UserClientIp      string
	Lang              string
	ContactTemplateId int64
}
type SaveContactTemplateCredentialResponse struct {
	RequestId string
}

func (c *DomainClient) WhoisProtection(req *WhoisProtectionArgs) (resp *WhoisProtectionResponse, err error) {
	resp = &WhoisProtectionResponse{}
	err = c.Request("WhoisProtection", req, resp)
	return
}

type WhoisProtectionArgs struct {
	WhoisProtect bool
	DataSource   int
	UserClientIp string
	DataContent  string
	Lang         string
}
type WhoisProtectionResponse struct {
	RequestId string
	Result    int
}

func (c *DomainClient) QueryOrder(req *QueryOrderArgs) (resp *QueryOrderResponse, err error) {
	resp = &QueryOrderResponse{}
	err = c.Request("QueryOrder", req, resp)
	return
}

type QueryOrderSubOrderResult struct {
	TrackID      string
	OrderID      string
	SaleID       string
	UserID       string
	ClassID      string
	ProductName  string
	RelatedName  string
	ActionType   string
	PeriodUnit   int
	PeriodNum    int
	Amount       int
	OrderDate    string
	CheckFlag    bool
	CheckDate    string
	BizStatus    int
	UpdateDate   string
	DeadDate     string
	ValidFlag    bool
	Money        string
	ParentSaleID string
}
type QueryOrderArgs struct {
	OrderID string
}
type QueryOrderResponse struct {
	RequestId     string
	OrderID       string
	UserID        string
	Money         string
	OrderDate     string
	CheckFlag     bool
	CheckDate     string
	ValidFlag     bool
	CheckType     bool
	OrderProducts QueryOrderSubOrderResultList
}

type QueryOrderSubOrderResultList []QueryOrderSubOrderResult

func (list *QueryOrderSubOrderResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryOrderSubOrderResult)
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

func (c *DomainClient) OSuborderDomain(req *OSuborderDomainArgs) (resp *OSuborderDomainResponse, err error) {
	resp = &OSuborderDomainResponse{}
	err = c.Request("OSuborderDomain", req, resp)
	return
}

type OSuborderDomainObject struct {
	CommodityKey  string
	CommodityCode string
	Amount        int64
	SettleTime    string
}
type OSuborderDomainArgs struct {
	EndDate   string
	PageSize  int
	Type      string
	StartDate string
	PageNum   int
}
type OSuborderDomainResponse struct {
	RequestId      string
	TotalItemNum   int
	CurrentPageNum int
	PageSize       int
	Data           OSuborderDomainObjectList
}

type OSuborderDomainObjectList []OSuborderDomainObject

func (list *OSuborderDomainObjectList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OSuborderDomainObject)
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

func (c *DomainClient) CreateOrder(req *CreateOrderArgs) (resp *CreateOrderResponse, err error) {
	resp = &CreateOrderResponse{}
	err = c.Request("CreateOrder", req, resp)
	return
}

type CreateOrderSubOrderParam struct {
	SaleID           string
	RelatedName      string
	Action           string
	Period           int
	DomainTemplateID string
}
type CreateOrderArgs struct {
	SubOrderParams CreateOrderSubOrderParamList
}
type CreateOrderResponse struct {
	RequestId string
	OrderID   string
}

type CreateOrderSubOrderParamList []CreateOrderSubOrderParam

func (list *CreateOrderSubOrderParamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateOrderSubOrderParam)
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
