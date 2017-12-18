package cloudauth

import "encoding/json"

func (c *CloudauthClient) CompareFaces(req *CompareFacesArgs) (resp *CompareFacesResponse, err error) {
	resp = &CompareFacesResponse{}
	err = c.Request("CompareFaces", req, resp)
	return
}

type CompareFacesData struct {
	SimilarityScore      float32
	ConfidenceThresholds string
}
type CompareFacesArgs struct {
	SourceImageType  string
	ResourceOwnerId  int64
	TargetImageType  string
	SourceImageValue string
	TargetImageValue string
}
type CompareFacesResponse struct {
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      CompareFacesData
}

func (c *CloudauthClient) GetVerifyToken(req *GetVerifyTokenArgs) (resp *GetVerifyTokenResponse, err error) {
	resp = &GetVerifyTokenResponse{}
	err = c.Request("GetVerifyToken", req, resp)
	return
}

type GetVerifyTokenData struct {
	VerifyToken GetVerifyTokenVerifyToken
	StsToken    GetVerifyTokenStsToken
}

type GetVerifyTokenVerifyToken struct {
	Token           string
	DurationSeconds int
}

type GetVerifyTokenStsToken struct {
	AccessKeyId     string
	AccessKeySecret string
	Expiration      string
	EndPoint        string
	BucketName      string
	Path            string
	Token           string
}
type GetVerifyTokenArgs struct {
	UserData        string
	ResourceOwnerId int64
	Biz             string
	Binding         string
	TicketId        string
}
type GetVerifyTokenResponse struct {
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      GetVerifyTokenData
}

func (c *CloudauthClient) GetMaterials(req *GetMaterialsArgs) (resp *GetMaterialsResponse, err error) {
	resp = &GetMaterialsResponse{}
	err = c.Request("GetMaterials", req, resp)
	return
}

type GetMaterialsData struct {
	Name                 string
	IdentificationNumber string
	IdCardType           string
	IdCardExpiry         string
	Address              string
	Sex                  string
	IdCardFrontPic       string
	IdCardBackPic        string
	FacePic              string
}
type GetMaterialsArgs struct {
	ResourceOwnerId int64
	Biz             string
	TicketId        string
}
type GetMaterialsResponse struct {
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      GetMaterialsData
}

func (c *CloudauthClient) SubmitMaterials(req *SubmitMaterialsArgs) (resp *SubmitMaterialsResponse, err error) {
	resp = &SubmitMaterialsResponse{}
	err = c.Request("SubmitMaterials", req, resp)
	return
}

type SubmitMaterialsMaterial struct {
	MaterialType string
	Value        string
}

type SubmitMaterialsData struct {
	VerifyStatus SubmitMaterialsVerifyStatus
}

type SubmitMaterialsVerifyStatus struct {
	StatusCode      int
	TrustedScore    float32
	SimilarityScore float32
}
type SubmitMaterialsArgs struct {
	ResourceOwnerId int64
	Materials       SubmitMaterialsMaterialList
	VerifyToken     string
}
type SubmitMaterialsResponse struct {
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      SubmitMaterialsData
}

type SubmitMaterialsMaterialList []SubmitMaterialsMaterial

func (list *SubmitMaterialsMaterialList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitMaterialsMaterial)
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

func (c *CloudauthClient) GetStatus(req *GetStatusArgs) (resp *GetStatusResponse, err error) {
	resp = &GetStatusResponse{}
	err = c.Request("GetStatus", req, resp)
	return
}

type GetStatusData struct {
	StatusCode       int
	TrustedScore     float32
	SimilarityScore  float32
	AuditConclusions string
}
type GetStatusArgs struct {
	ResourceOwnerId int64
	Biz             string
	TicketId        string
}
type GetStatusResponse struct {
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      GetStatusData
}
