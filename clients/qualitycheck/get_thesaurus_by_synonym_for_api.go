package qualitycheck

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetThesaurusBySynonymForApiRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *GetThesaurusBySynonymForApiRequest) Invoke(client *sdk.Client) (resp *GetThesaurusBySynonymForApiResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "GetThesaurusBySynonymForApi", "", "")
	resp = &GetThesaurusBySynonymForApiResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetThesaurusBySynonymForApiResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
	Data      GetThesaurusBySynonymForApiThesaurusPoList
}

type GetThesaurusBySynonymForApiThesaurusPo struct {
	Id          common.Long
	Business    common.String
	SynonymList GetThesaurusBySynonymForApiSynonymListList
}

type GetThesaurusBySynonymForApiThesaurusPoList []GetThesaurusBySynonymForApiThesaurusPo

func (list *GetThesaurusBySynonymForApiThesaurusPoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetThesaurusBySynonymForApiThesaurusPo)
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

type GetThesaurusBySynonymForApiSynonymListList []common.String

func (list *GetThesaurusBySynonymForApiSynonymListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
