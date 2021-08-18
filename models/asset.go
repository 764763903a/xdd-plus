package models

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/httplib"
	"github.com/buger/jsonparser"
)

type Asset struct {
	Nickname string
	Bean     struct {
		Total       int
		TodayIn     int
		TodayOut    int
		YestodayIn  int
		YestodayOut int
		ToExpire    []int
	}
	RedPacket struct {
		Total      float64
		ToExpire   float64
		ToExpireJd float64
		ToExpireJx float64
		ToExpireJs float64
		ToExpireJk float64
		Jd         float64
		Jx         float64
		Js         float64
		Jk         float64
	}
	Other struct {
		JsCoin   float64
		NcStatus float64
		McStatus float64
	}
}

var Int = func(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

var Float64 = func(s string) float64 {
	i, _ := strconv.ParseFloat(s, 64)
	return i
}

func DailyAssetsPush() {
	for _, ck := range GetJdCookies() {
		if (ck.QQ != 0 && Config.QQID != 0 && SendQQ != nil) || ck.PushPlus != "" {
			msg := ck.Query()
			if ck.QQ != 0 && Config.QQID != 0 && SendQQ != nil {
				SendQQ(int64(ck.QQ), msg)
			}
			if ck.PushPlus != "" {
				pushPlus(ck.PushPlus, msg)
			}
		}
	}
}

func (ck *JdCookie) Query() string {
	msgs := []string{
		fmt.Sprintf("账号昵称：%s", ck.Nickname),
	}
	asset := Asset{}
	if CookieOK(ck) {
		cookie := fmt.Sprintf("pt_key=%s;pt_pin=%s;", ck.PtKey, ck.PtPin)
		var rpc = make(chan []RedList)
		var fruit = make(chan string)
		var pet = make(chan string)
		var gold = make(chan int64)
		go redPacket(cookie, rpc)
		go initFarm(cookie, fruit)
		go initPetTown(cookie, pet)
		go jsGold(cookie, gold)
		today := time.Now().Local().Format("2006-01-02")
		yestoday := time.Now().Local().Add(-time.Hour * 24).Format("2006-01-02")
		page := 1
		end := false
		for {
			if end {
				msgs = append(msgs, []string{
					fmt.Sprintf("昨日收入：%d京豆", asset.Bean.YestodayIn),
					fmt.Sprintf("昨日支出：%d京豆", asset.Bean.YestodayOut),
					fmt.Sprintf("今日收入：%d京豆", asset.Bean.TodayIn),
					fmt.Sprintf("今日支出：%d京豆", asset.Bean.TodayOut),
				}...)
				break
			}
			bds := getJingBeanBalanceDetail(page, cookie)
			if bds == nil {
				end = true
				msgs = append(msgs, "京豆数据异常")
				break
			}
			for _, bd := range bds {
				amount := Int(bd.Amount)
				if strings.Contains(bd.Date, today) {
					if amount > 0 {
						asset.Bean.TodayIn += amount
					} else {
						asset.Bean.TodayOut += -amount
					}
				} else if strings.Contains(bd.Date, yestoday) {
					if amount > 0 {
						asset.Bean.YestodayIn += amount
					} else {
						asset.Bean.YestodayOut += -amount
					}
				} else {
					end = true
					break
				}
			}
			page++
		}
		msgs = append(msgs, fmt.Sprintf("当前京豆：%v京豆", ck.BeanNum))
		ysd := int(time.Now().Add(24 * time.Hour).Unix())
		if rps := <-rpc; len(rps) != 0 {
			for _, rp := range rps {
				b := Float64(rp.Balance)
				asset.RedPacket.Total += b
				if strings.Contains(rp.ActivityName, "京喜") {
					asset.RedPacket.Jx += b
					if ysd >= rp.EndTime {
						asset.RedPacket.ToExpireJx += b
						asset.RedPacket.ToExpire += b
					}
				} else if strings.Contains(rp.ActivityName, "极速版") {
					asset.RedPacket.Js += b
					if ysd >= rp.EndTime {
						asset.RedPacket.ToExpireJs += b
						asset.RedPacket.ToExpire += b
					}

				} else if strings.Contains(rp.ActivityName, "京东健康") {
					asset.RedPacket.Jk += b
					if ysd >= rp.EndTime {
						asset.RedPacket.ToExpireJk += b
						asset.RedPacket.ToExpire += b
					}
				} else {
					asset.RedPacket.Jd += b
					if ysd >= rp.EndTime {
						asset.RedPacket.ToExpireJd += b
						asset.RedPacket.ToExpire += b
					}
				}
			}
			msgs = append(msgs, []string{
				fmt.Sprintf("所有红包：%.2f(今日总过期%.2f)元🧧", asset.RedPacket.Total, asset.RedPacket.ToExpire),
				fmt.Sprintf("京喜红包：%.2f(今日总过期%.2f)元🧧", asset.RedPacket.Jx, asset.RedPacket.ToExpireJx),
				fmt.Sprintf("极速红包：%.2f(今日总过期%.2f)元🧧", asset.RedPacket.Js, asset.RedPacket.ToExpireJs),
				fmt.Sprintf("健康红包：%.2f(今日总过期%.2f)元🧧", asset.RedPacket.Jk, asset.RedPacket.ToExpireJk),
				fmt.Sprintf("京东红包：%.2f(今日总过期%.2f)元🧧", asset.RedPacket.Jd, asset.RedPacket.ToExpireJd),
			}...)
		} else {
			msgs = append(msgs, "暂无红包数据🧧")
		}
		msgs = append(msgs, fmt.Sprintf("东东农场：%s", <-fruit))
		msgs = append(msgs, fmt.Sprintf("东东萌宠：%s", <-pet))
		gn := <-gold
		msgs = append(msgs, fmt.Sprintf("极速金币：%d(≈%.2f元)", gn, float64(gn)/1000))
	} else {
		msgs = append(msgs, []string{
			"提醒：该账号已过期，请重新登录",
		}...)
	}
	ck.PtPin, _ = url.QueryUnescape(ck.PtPin)
	return strings.Join(msgs, "\n")
}

type BeanDetail struct {
	Date         string `json:"date"`
	Amount       string `json:"amount"`
	EventMassage string `json:"eventMassage"`
}

func getJingBeanBalanceDetail(page int, cookie string) []BeanDetail {
	type AutoGenerated struct {
		Code       string       `json:"code"`
		DetailList []BeanDetail `json:"detailList"`
	}
	a := AutoGenerated{}
	req := httplib.Post(`https://api.m.jd.com/client.action?functionId=getJingBeanBalanceDetail`)
	req.Header("User-Agent", ua)
	req.Header("Host", "api.m.jd.com")
	req.Header("Content-Type", "application/x-www-form-urlencoded")
	req.Header("Cookie", cookie)
	req.Body(fmt.Sprintf(`body={"pageSize": "20", "page": "%d"}&appid=ld`, page))
	data, err := req.Bytes()
	if err != nil {
		return nil
	}
	json.Unmarshal(data, &a)
	return a.DetailList
}

type RedList struct {
	ActivityName string `json:"activityName"`
	Balance      string `json:"balance"`
	BeginTime    int    `json:"beginTime"`
	DelayRemark  string `json:"delayRemark"`
	Discount     string `json:"discount"`
	EndTime      int    `json:"endTime"`
	HbID         string `json:"hbId"`
	HbState      int    `json:"hbState"`
	IsDelay      bool   `json:"isDelay"`
	OrgLimitStr  string `json:"orgLimitStr"`
}

func redPacket(cookie string, rpc chan []RedList) {
	type UseRedInfo struct {
		Count   int       `json:"count"`
		RedList []RedList `json:"redList"`
	}
	type Data struct {
		AvaiCount      int        `json:"avaiCount"`
		Balance        string     `json:"balance"`
		CountdownTime  string     `json:"countdownTime"`
		ExpiredBalance string     `json:"expiredBalance"`
		ServerCurrTime int        `json:"serverCurrTime"`
		UseRedInfo     UseRedInfo `json:"useRedInfo"`
	}
	type AutoGenerated struct {
		Data    Data   `json:"data"`
		Errcode int    `json:"errcode"`
		Msg     string `json:"msg"`
	}
	a := AutoGenerated{}
	req := httplib.Get(`https://m.jingxi.com/user/info/QueryUserRedEnvelopesV2?type=1&orgFlag=JD_PinGou_New&page=1&cashRedType=1&redBalanceFlag=1&channel=1&_=` + fmt.Sprint(time.Now().Unix()) + `&sceneval=2&g_login_type=1&g_ty=ls`)
	req.Header("User-Agent", ua)
	req.Header("Host", "m.jingxi.com")
	req.Header("Accept", "*/*")
	req.Header("Connection", "keep-alive")
	req.Header("Accept-Language", "zh-cn")
	req.Header("Accept-Encoding", "gzip, deflate, br")
	req.Header("Referer", "https://st.jingxi.com/my/redpacket.shtml?newPg=App&jxsid=16156262265849285961")
	req.Header("Cookie", cookie)
	data, _ := req.Bytes()
	json.Unmarshal(data, &a)
	rpc <- a.Data.UseRedInfo.RedList
}

func initFarm(cookie string, state chan string) {
	type RightUpResouces struct {
		AdvertID string `json:"advertId"`
		Name     string `json:"name"`
		AppImage string `json:"appImage"`
		AppLink  string `json:"appLink"`
		CxyImage string `json:"cxyImage"`
		CxyLink  string `json:"cxyLink"`
		Type     string `json:"type"`
		OpenLink bool   `json:"openLink"`
	}
	type TurntableInit struct {
		TimeState int `json:"timeState"`
	}
	type MengchongResouce struct {
		AdvertID string `json:"advertId"`
		Name     string `json:"name"`
		AppImage string `json:"appImage"`
		AppLink  string `json:"appLink"`
		CxyImage string `json:"cxyImage"`
		CxyLink  string `json:"cxyLink"`
		Type     string `json:"type"`
		OpenLink bool   `json:"openLink"`
	}
	type GUIDPopupTask struct {
		GUIDPopupTask string `json:"guidPopupTask"`
	}
	type IosConfigResouces struct {
		AdvertID string `json:"advertId"`
		Name     string `json:"name"`
		AppImage string `json:"appImage"`
		AppLink  string `json:"appLink"`
		CxyImage string `json:"cxyImage"`
		CxyLink  string `json:"cxyLink"`
		Type     string `json:"type"`
		OpenLink bool   `json:"openLink"`
	}
	type TodayGotWaterGoalTask struct {
		CanPop bool `json:"canPop"`
	}
	type LeftUpResouces struct {
		AdvertID string `json:"advertId"`
		Name     string `json:"name"`
		AppImage string `json:"appImage"`
		AppLink  string `json:"appLink"`
		CxyImage string `json:"cxyImage"`
		CxyLink  string `json:"cxyLink"`
		Type     string `json:"type"`
		OpenLink bool   `json:"openLink"`
	}
	type RightDownResouces struct {
		AdvertID string `json:"advertId"`
		Name     string `json:"name"`
		AppImage string `json:"appImage"`
		AppLink  string `json:"appLink"`
		CxyImage string `json:"cxyImage"`
		CxyLink  string `json:"cxyLink"`
		Type     string `json:"type"`
		OpenLink bool   `json:"openLink"`
	}
	type FarmUserPro struct {
		TotalEnergy     int    `json:"totalEnergy"`
		TreeState       int    `json:"treeState"`
		CreateTime      int64  `json:"createTime"`
		TreeEnergy      int    `json:"treeEnergy"`
		TreeTotalEnergy int    `json:"treeTotalEnergy"`
		ShareCode       string `json:"shareCode"`
		WinTimes        int    `json:"winTimes"`
		NickName        string `json:"nickName"`
		CouponKey       string `json:"couponKey"`
		CouponID        string `json:"couponId"`
		CouponEndTime   int64  `json:"couponEndTime"`
		Type            string `json:"type"`
		SimpleName      string `json:"simpleName"`
		Name            string `json:"name"`
		GoodsImage      string `json:"goodsImage"`
		SkuID           string `json:"skuId"`
		LastLoginDate   int64  `json:"lastLoginDate"`
		NewOldState     int    `json:"newOldState"`
		OldMarkComplete int    `json:"oldMarkComplete"`
		CommonState     int    `json:"commonState"`
		PrizeLevel      int    `json:"prizeLevel"`
	}
	type LeftDownResouces struct {
		AdvertID string `json:"advertId"`
		Name     string `json:"name"`
		AppImage string `json:"appImage"`
		AppLink  string `json:"appLink"`
		CxyImage string `json:"cxyImage"`
		CxyLink  string `json:"cxyLink"`
		Type     string `json:"type"`
		OpenLink bool   `json:"openLink"`
	}
	type LoadFriend struct {
		Code            string      `json:"code"`
		StatisticsTimes interface{} `json:"statisticsTimes"`
		SysTime         int64       `json:"sysTime"`
		Message         interface{} `json:"message"`
		FirstAddUser    bool        `json:"firstAddUser"`
	}
	type AutoGenerated struct {
		Code                  string                `json:"code"`
		RightUpResouces       RightUpResouces       `json:"rightUpResouces"`
		TurntableInit         TurntableInit         `json:"turntableInit"`
		IosShieldConfig       interface{}           `json:"iosShieldConfig"`
		MengchongResouce      MengchongResouce      `json:"mengchongResouce"`
		ClockInGotWater       bool                  `json:"clockInGotWater"`
		GUIDPopupTask         GUIDPopupTask         `json:"guidPopupTask"`
		ToFruitEnergy         int                   `json:"toFruitEnergy"`
		StatisticsTimes       interface{}           `json:"statisticsTimes"`
		SysTime               int64                 `json:"sysTime"`
		CanHongbaoContineUse  bool                  `json:"canHongbaoContineUse"`
		ToFlowTimes           int                   `json:"toFlowTimes"`
		IosConfigResouces     IosConfigResouces     `json:"iosConfigResouces"`
		TodayGotWaterGoalTask TodayGotWaterGoalTask `json:"todayGotWaterGoalTask"`
		LeftUpResouces        LeftUpResouces        `json:"leftUpResouces"`
		MinSupportAPPVersion  string                `json:"minSupportAPPVersion"`
		LowFreqStatus         int                   `json:"lowFreqStatus"`
		FunCollectionHasLimit bool                  `json:"funCollectionHasLimit"`
		Message               interface{}           `json:"message"`
		TreeState             int                   `json:"treeState"`
		RightDownResouces     RightDownResouces     `json:"rightDownResouces"`
		IconFirstPurchaseInit bool                  `json:"iconFirstPurchaseInit"`
		ToFlowEnergy          int                   `json:"toFlowEnergy"`
		FarmUserPro           FarmUserPro           `json:"farmUserPro"`
		RetainPopupLimit      int                   `json:"retainPopupLimit"`
		ToBeginEnergy         int                   `json:"toBeginEnergy"`
		LeftDownResouces      LeftDownResouces      `json:"leftDownResouces"`
		EnableSign            bool                  `json:"enableSign"`
		LoadFriend            LoadFriend            `json:"loadFriend"`
		HadCompleteXgTask     bool                  `json:"hadCompleteXgTask"`
		OldUserIntervalTimes  []int                 `json:"oldUserIntervalTimes"`
		ToFruitTimes          int                   `json:"toFruitTimes"`
		OldUserSendWater      []string              `json:"oldUserSendWater"`
	}
	a := AutoGenerated{}
	req := httplib.Post(`https://api.m.jd.com/client.action?functionId=initForFarm`)
	req.Header("accept", "*/*")
	req.Header("accept-encoding", "gzip, deflate, br")
	req.Header("accept-language", "zh-CN,zh;q=0.9")
	req.Header("cache-control", "no-cache")
	req.Header("cookie", cookie)
	req.Header("origin", "https://home.m.jd.com")
	req.Header("pragma", "no-cache")
	req.Header("referer", "https://home.m.jd.com/myJd/newhome.action")
	req.Header("sec-fetch-dest", "empty")
	req.Header("sec-fetch-mode", "cors")
	req.Header("sec-fetch-site", "same-site")
	req.Header("User-Agent", ua)
	req.Header("Content-Type", "application/x-www-form-urlencoded")
	req.Body(`body={"version":4}&appid=wh5&clientVersion=9.1.0`)
	data, _ := req.Bytes()
	json.Unmarshal(data, &a)

	rt := a.FarmUserPro.Name
	if rt == "" {
		rt = "数据异常"
	} else {
		if a.TreeState == 2 || a.TreeState == 3 {
			rt += "已可领取⏰"
		} else if a.TreeState == 1 {
			rt += fmt.Sprintf("种植中，进度%.2f%%🍒", float64(a.FarmUserPro.TreeTotalEnergy)/float64(a.FarmUserPro.TreeEnergy))
		} else if a.TreeState == 0 {
			rt = "您忘了种植新的水果⏰"
		}
	}
	state <- rt
}

func initPetTown(cookie string, state chan string) {
	type ResourceList struct {
		AdvertID string `json:"advertId"`
		ImageURL string `json:"imageUrl"`
		Link     string `json:"link"`
		ShopID   string `json:"shopId"`
	}
	type PetPlaceInfoList struct {
		Place  int `json:"place"`
		Energy int `json:"energy"`
	}
	type PetInfo struct {
		AdvertID     string `json:"advertId"`
		NickName     string `json:"nickName"`
		IconURL      string `json:"iconUrl"`
		ClickIconURL string `json:"clickIconUrl"`
		FeedGifURL   string `json:"feedGifUrl"`
		HomePetImage string `json:"homePetImage"`
		CrossBallURL string `json:"crossBallUrl"`
		RunURL       string `json:"runUrl"`
		TickleURL    string `json:"tickleUrl"`
	}
	type GoodsInfo struct {
		GoodsName        string `json:"goodsName"`
		GoodsURL         string `json:"goodsUrl"`
		GoodsID          string `json:"goodsId"`
		ExchangeMedalNum int    `json:"exchangeMedalNum"`
		ActivityID       string `json:"activityId"`
		ActivityIds      string `json:"activityIds"`
	}
	type Result struct {
		ShareCode              string             `json:"shareCode"`
		HisHbFlag              bool               `json:"hisHbFlag"`
		MasterHelpPeoples      []interface{}      `json:"masterHelpPeoples"`
		HelpSwitchOn           bool               `json:"helpSwitchOn"`
		UserStatus             int                `json:"userStatus"`
		TotalEnergy            int                `json:"totalEnergy"`
		MasterInvitePeoples    []interface{}      `json:"masterInvitePeoples"`
		ShareTo                string             `json:"shareTo"`
		PetSportStatus         int                `json:"petSportStatus"`
		UserImage              string             `json:"userImage"`
		MasterHelpReward       int                `json:"masterHelpReward"`
		ShowHongBaoExchangePop bool               `json:"showHongBaoExchangePop"`
		ShowNeedCollectPop     bool               `json:"showNeedCollectPop"`
		PetSportReward         string             `json:"petSportReward"`
		NewhandBubble          bool               `json:"newhandBubble"`
		ResourceList           []ResourceList     `json:"resourceList"`
		ProjectBubble          bool               `json:"projectBubble"`
		MasterInvitePop        bool               `json:"masterInvitePop"`
		MasterInviteReward     int                `json:"masterInviteReward"`
		MedalNum               int                `json:"medalNum"`
		MasterHelpPop          bool               `json:"masterHelpPop"`
		MeetDays               int                `json:"meetDays"`
		PetPlaceInfoList       []PetPlaceInfoList `json:"petPlaceInfoList"`
		MedalPercent           float64            `json:"medalPercent"`
		CharitableSwitchOn     bool               `json:"charitableSwitchOn"`
		PetInfo                PetInfo            `json:"petInfo"`
		NeedCollectEnergy      int                `json:"needCollectEnergy"`
		FoodAmount             int                `json:"foodAmount"`
		InviteCode             string             `json:"inviteCode"`
		RulesURL               string             `json:"rulesUrl"`
		PetStatus              int                `json:"petStatus"`
		GoodsInfo              GoodsInfo          `json:"goodsInfo"`
	}
	type AutoGenerated struct {
		Code       string `json:"code"`
		ResultCode string `json:"resultCode"`
		Message    string `json:"message"`
		Result     Result `json:"result"`
	}
	a := AutoGenerated{}
	req := httplib.Post(`https://api.m.jd.com/client.action?functionId=initPetTown`)
	req.Header("Host", "api.m.jd.com")
	req.Header("User-Agent", ua)
	req.Header("cookie", cookie)
	req.Header("Content-Type", "application/x-www-form-urlencoded")
	req.Body(`body={}&appid=wh5&loginWQBiz=pet-town&clientVersion=9.0.4`)
	data, _ := req.Bytes()
	fmt.Println(data)
	json.Unmarshal(data, &a)
	rt := ""
	if a.Code == "0" && a.ResultCode == "0" && a.Message == "success" {
		if a.Result.UserStatus == 0 {
			rt = "请手动开启活动⏰"
		} else if a.Result.GoodsInfo.GoodsName == "" {
			rt = "你忘了选购新的商品⏰"
		} else if a.Result.PetStatus == 5 {
			rt = a.Result.GoodsInfo.GoodsName + "已可领取⏰"
		} else if a.Result.PetStatus == 6 {
			rt = a.Result.GoodsInfo.GoodsName + "未继续领养新的物品⏰"
		} else {
			rt = a.Result.GoodsInfo.GoodsName + fmt.Sprintf("领养中，进度%.2f%%，勋章%d/%d🐶", a.Result.MedalPercent, a.Result.MedalNum, a.Result.GoodsInfo.ExchangeMedalNum)
		}
	} else {
		rt = "数据异常"
	}
	state <- rt
}

func jsGold(cookie string, state chan int64) {
	req := httplib.Post(`https://api.m.jd.com/`)
	req.Header("Origin", "https://gold.jd.com")
	req.Header("Host", "api.m.jd.com")
	req.Header("Accept-Encoding", "gzip, deflate, br")
	req.Header("User-Agent", ua)
	req.Header("cookie", cookie)
	req.Header("Content-Type", "application/x-www-form-urlencoded")
	req.Body(`functionId=MyAssetsService.execute&body={"method":"goldShopPage","data":{"channel":1}}&_t=1629271472844&appid=market-task-h5;`)
	data, _ := req.Bytes()
	fmt.Println(string(data))
	gold, _ := jsonparser.GetInt(data, "data.balanceVO.goldBalance")
	state <- gold
}
