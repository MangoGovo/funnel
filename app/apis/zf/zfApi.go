package zf

import (
	"funnel/app/apis"
	"funnel/config"
	"strconv"
	"strings"
	"time"
)

func ChooseURL(jf bool) string {
	if jf {
		return apis.ZF_JF_URL
	}
	if config.Redis.Exists("zf_url").Val() != 1 {
		config.Redis.Set("zf_url", "bk", 0)
	}
	if strings.Compare(config.Redis.Get("zf_url").String(), "new") == 0 {
		return apis.ZF_URL
	} else {
		return apis.ZF_BK_URL
	}
}

func ZfLoginGetPublickey() string {
	return ChooseURL(false) + "xtgl/login_getPublicKey.html?time=" + strconv.FormatInt(time.Now().Unix()*1000, 10)
}
func ZfLoginHome() string {
	return ChooseURL(false) + "xtgl/login_slogin.html?time=" + strconv.FormatInt(time.Now().Unix()*1000, 10)
}
func ZfExamInfo(jf bool) string {
	return ChooseURL(jf) + "kwgl/kscx_cxXsksxxIndex.html?doType=query"
}
func ZfClassTable(jf bool) string {
	return ChooseURL(jf) + "kbcx/xskbcx_cxXsgrkb.html?gnmkdm=N2151&su="
}
func ZfScore(jf bool) string {
	return ChooseURL(jf) + "cjcx/cjcx_cxDgXscj.html?doType=query"
}
func ZfScoreDetail(jf bool) string {
	return ChooseURL(jf) + "cjcx/cjcx_cxXsKccjList.html?doType=query"
}
func ZfMinTermScore(jf bool) string {
	return ChooseURL(jf) + "design/funcData_cxFuncDataList.html?func_widget_guid=5EF567BFD3CE243EE053A11310AC1252&gnmkdm=N305013"
}
func ZfEmptyClassRoom(jf bool) string {
	return ChooseURL(jf) + "cdjy/cdjy_cxKxcdlb.html?doType=query"
}
func ZfUserInfo(jf bool) string {
	return ChooseURL(jf) + "xsxxxggl/xsgrxxwh_cxXsgrxx.html?gnmkdm=N100801&layout=default"
}
func ZfPY(jf bool) string {
	return ChooseURL(jf) + "pyfagl/pyfaxxck_dyPyfaxx.html?id="
}
