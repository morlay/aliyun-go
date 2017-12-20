package teslamaxcompute

import (
	"encoding/json"
)

func (c *TeslamaxcomputeClient) GetProjectInfo(req *GetProjectInfoArgs) (resp *GetProjectInfoResponse, err error) {
	resp = &GetProjectInfoResponse{}
	err = c.Request("GetProjectInfo", req, resp)
	return
}

type GetProjectInfoData struct {
	Total  int
	Detail GetProjectInfoInstanceList
}

type GetProjectInfoInstance struct {
	Project     string
	InstanceId  string
	Status      string
	UserAccount string
	NickName    string
	Cluster     string
	RunTime     string
}
type GetProjectInfoArgs struct {
	PageSize int
	Project  string
	PageNum  int
	Status   string
}
type GetProjectInfoResponse struct {
	Code      int
	Message   string
	RequestId string
	Data      GetProjectInfoData
}

type GetProjectInfoInstanceList []GetProjectInfoInstance

func (list *GetProjectInfoInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetProjectInfoInstance)
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

func (c *TeslamaxcomputeClient) GetClusterInfo(req *GetClusterInfoArgs) (resp *GetClusterInfoResponse, err error) {
	resp = &GetClusterInfoResponse{}
	err = c.Request("GetClusterInfo", req, resp)
	return
}

type GetClusterInfoData struct {
	Total  int
	Detail GetClusterInfoInstanceList
}

type GetClusterInfoInstance struct {
	Project     string
	InstanceId  string
	Status      string
	UserAccount string
	NickName    string
	Cluster     string
	RunTime     string
}
type GetClusterInfoArgs struct {
	Cluster  string
	PageSize int
	PageNum  int
	Status   string
}
type GetClusterInfoResponse struct {
	Code      int
	Message   string
	RequestId string
	Data      GetClusterInfoData
}

type GetClusterInfoInstanceList []GetClusterInfoInstance

func (list *GetClusterInfoInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetClusterInfoInstance)
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

func (c *TeslamaxcomputeClient) GetQuotaInfo(req *GetQuotaInfoArgs) (resp *GetQuotaInfoResponse, err error) {
	resp = &GetQuotaInfoResponse{}
	err = c.Request("GetQuotaInfo", req, resp)
	return
}

type GetQuotaInfoData struct {
	Total  int
	Detail GetQuotaInfoInstanceList
}

type GetQuotaInfoInstance struct {
	Project     string
	InstanceId  string
	Status      string
	UserAccount string
	NickName    string
	Cluster     string
	RunTime     string
}
type GetQuotaInfoArgs struct {
	Cluster  string
	PageSize int
	QuotaId  string
	PageNum  int
	Status   string
}
type GetQuotaInfoResponse struct {
	Code      int
	Message   string
	RequestId string
	Data      GetQuotaInfoData
}

type GetQuotaInfoInstanceList []GetQuotaInfoInstance

func (list *GetQuotaInfoInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetQuotaInfoInstance)
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

func (c *TeslamaxcomputeClient) GetUserInfo(req *GetUserInfoArgs) (resp *GetUserInfoResponse, err error) {
	resp = &GetUserInfoResponse{}
	err = c.Request("GetUserInfo", req, resp)
	return
}

type GetUserInfoData struct {
	Total  int
	Detail GetUserInfoInstanceList
}

type GetUserInfoInstance struct {
	Project     string
	InstanceId  string
	Status      string
	UserAccount string
	NickName    string
	Cluster     string
	RunTime     string
}
type GetUserInfoArgs struct {
	PageSize int
	PageNum  int
	User     string
	Status   string
}
type GetUserInfoResponse struct {
	Code      int
	Message   string
	RequestId string
	Data      GetUserInfoData
}

type GetUserInfoInstanceList []GetUserInfoInstance

func (list *GetUserInfoInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetUserInfoInstance)
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
