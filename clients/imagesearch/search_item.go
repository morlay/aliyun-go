
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
RequestId string
Success bool
Message string
Code int
Auctions SearchItemAuctionList
Head SearchItemHead
PicInfo SearchItemPicInfo
}

type SearchItemAuction struct {
CustContent string
ProductId string
SortExprValues string
CatId string
PicName string
}

type SearchItemHead struct {
SearchTime int
DocsFound int
DocsReturn int
}

type SearchItemPicInfo struct {
Category string
Region string
AllCategory SearchItemCategoryList
}

type SearchItemCategory struct {
Name string
Id string
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
                    
