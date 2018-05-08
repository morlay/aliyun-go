
package imagesearch

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type SearchItemRequest struct {
requests.RoaRequest
Start SearchItemint `name:"Start"`
Num SearchItemint `name:"Num"`
CatId string `name:"CatId"`
SearchPicture SearchItembyte[] `name:"SearchPicture"`
InstanceName string `position:"Query" name:"InstanceName"`
}

func (req *SearchItemRequest) Invoke(client *sdk.Client) (resp *SearchItemResponse, err error) {
	req.InitWithApiInfo("ImageSearch", "2018-01-20", "SearchItem", "/item/search","", "")
    req.Method = "POST"

    resp = &SearchItemResponse{}
	err = client.DoAction(req, resp)
	return
}

type SearchItemResponse struct {
responses.BaseResponse
RequestId common.String
Success bool
Message common.String
Code common.Integer
Auctions SearchItemAuctionList
Head SearchItemHead
PicInfo SearchItemPicInfo
}

type SearchItemAuction struct {
CustContent common.String
ProductId common.String
SortExprValues common.String
CatId common.String
PicName common.String
}

type SearchItemHead struct {
SearchTime common.Integer
DocsFound common.Integer
DocsReturn common.Integer
}

type SearchItemPicInfo struct {
Category common.String
Region common.String
AllCategory SearchItemCategoryList
}

type SearchItemCategory struct {
Name common.String
Id common.String
}

                    type SearchItemAuctionList []SearchItemAuction

                    func (list *SearchItemAuctionList) UnmarshalJSON(data []byte) error {
                        m := make(map[string][]SearchItemAuction)
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
                    
                    type SearchItemCategoryList []SearchItemCategory

                    func (list *SearchItemCategoryList) UnmarshalJSON(data []byte) error {
                        m := make(map[string][]SearchItemCategory)
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
                    
