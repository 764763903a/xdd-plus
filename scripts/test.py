import json, requests, re
import sys
requests.packages.urllib3.disable_warnings()
ws=sys.argv[1]
def getsign():
    try:
        sign = 'https://pan.smxy.xyz/sign'
        s = requests.get(url=sign).json()
        svv = s['sv']
        stt = s['st']
        suid = s['uuid']
        sign = s['sign']
        return svv, stt, suid, sign
    except:
        print('调用sign出错！')


def getToken(wskey):
    sv, st, uuid, sign = getsign()
    headers = {'cookie': wskey,
               'User-Agent': 'okhttp/3.12.1;jdmall;android;version/10.1.2;build/89743;screen/1440x3007;os/11;network/wifi;',
               'content-type': 'application/x-www-form-urlencoded; charset=UTF-8', 'charset': 'UTF-8',
               'accept-encoding': 'br,gzip,deflate'}
    params = {'functionId': 'genToken', 'clientVersion': '10.1.2', 'client': 'android', 'uuid': uuid, 'st': st,
              'sign': sign, 'sv': sv}
    url = 'https://api.m.jd.com/client.action'
    data = 'body=%7B%22action%22%3A%22to%22%2C%22to%22%3A%22https%253A%252F%252Fplogin.m.jd.com%252Fcgi-bin%252Fm%252Fthirdapp_auth_page%253Ftoken%253DAAEAIEijIw6wxF2s3bNKF0bmGsI8xfw6hkQT6Ui2QVP7z1Xg%2526client_type%253Dandroid%2526appid%253D879%2526appup_type%253D1%22%7D&'
    res = requests.post(url=url, params=params, headers=headers, data=data, verify=False)
    res_json = json.loads(res.text)
    tokenKey = res_json['tokenKey']
    return appjmp(wskey, tokenKey)

def appjmp(wskey, tokenKey):
    headers = {
        'User-Agent': 'okhttp/3.12.1;jdmall;android;version/10.1.2;build/89743;screen/1440x3007;os/11;network/wifi;',
        'accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3', }
    params = {'tokenKey': tokenKey,
              'to': 'https://plogin.m.jd.com/cgi-bin/m/thirdapp_auth_page?token=AAEAIEijIw6wxF2s3bNKF0bmGsI8xfw6hkQT6Ui2QVP7z1Xg',
              'client_type': 'android', 'appid': 879, 'appup_type': 1, }
    url = 'https://un.m.jd.com/cgi-bin/app/appjmp'
    res = requests.get(url=url, headers=headers, params=params, verify=False, allow_redirects=False)
    res_set = res.cookies.get_dict()
    pt_key = 'pt_key=' + res_set['pt_key']
    pt_pin = 'pt_pin=' + res_set['pt_pin']
    jd_ck = str(pt_key) + ';' + str(pt_pin) + ';'
    # print(jd_ck)
    wskey = wskey.split(";")[0]
    if 'fake' in pt_key:
        print("wskey格式错误！")
    else:
        print(jd_ck)
        
        
def checkwskey(wskey):
    flag = "wskey=" in wskey
    flag1 = "pin=" in wskey
    if flag == True & flag1 == True:
       return True
    else :
       return False
        
if __name__ == '__main__':
    getToken(ws)

