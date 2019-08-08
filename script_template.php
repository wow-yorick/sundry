<?php
error_reporting(0);

initialization();//脚本环境初始化
$crontabDB = getCrontabDB();//获取数据库对象
$requestData = getRequest();//获取请求数据
safetyVerification($requestData);//接口安全验证
apiLogicBody($crontabDB, $requestData);//接口逻辑主体(需要改的只有这个方法)


/**
 * initialization 
 * 程序环境初始化
 * @access public
 * @return void
 */
function initialization()
{
    define('SOURCE_PC', 'PC-SIMPLE_DESIGNER');
    define('SOURCE_WAP', 'WAP-MULTI_DESIGNER');
}

function apiLogicBody($crontabDB, $requestData)
{
        logicForWAP($crontabDB, $requestData);  
}


function logicForWAP($crontabDB, $requestData)
{
    $validateRule = array(
        'user_name' => array(
            'rule' => 'notnull',
            'declare' => 'xx',
        ),
        'sex' => array(
            'rule' => 'notnull',
            'declare' => 'x',
        ),
        'age' => array(
            'rule' => 'notnull',
            'declare' => 'xx',
        ),
        'education' => array(
            'rule' => 'notnull',
            'declare' => 'xx',
        ),
        'mobile' => array(
            'rule' => 'mobile',
            'declare' => 'xx',
        ),

        'service_year' => array(
            'rule' => 'notnull',
            'declare' => 'xx',
        ),
        'service_company' => array(
            'rule' => 'notnull',
            'declare' => 'xx',
        ),
        'title' => array(
            'rule' => 'notnull',
            'declare' => 'xx',
        ),
        'good_for' => array(
            'rule' => 'notnull',
            'declare' => 'xx',
        ),
        'idcard' => array(
            'rule' => 'all',
            'declare' => 'x',
        ),
        'share_mobile' => array(
            'rule' => 'empty|mobile',
            'declare' => 'xxx',
        ),
        'display_mobile' => array(
            'rule' => 'empty|mobile',
            'declare' => 'xxx',
        ),

        'source' => array(
            'rule' => 'notnull',
            'declare' => 'xx',
        ),
        'TJJID2' => array(
            'rule' => 'all',
            'declare' => 'xxx',
        ),
        'referrer' => array(
            'rule' => 'all',
            'declare' => 'xxx',
        ),
        'referrer_mobile' => array(
            'rule' => 'empty|mobile',
            'declare' => 'xx',
        ),
    );

    fieldVerify($validateRule, $requestData);//验证请求

    $designerEnterTable = $crontabDB['db']->selectCollection('designer_enter');
    applySafeLimit($designerEnterTable, $requestData);

    $nextId = dataInsert($requestData, $validateRule, $designerEnterTable, $crontabDB);

    //报名成功响应
    $retArr = responseStruct(200);//报名成功
    $retArr['designer_enter_id'] = $nextId;
    responseData($retArr);
}

/**
 * dataInsert 
 * 插入数据进数据库
 * @param mixed $requestData 验证通过后的请求字段
 * @param mixed $validateRule 字段验证规则
 * @param mixed $tableCollectObj 表操作对象
 * @param mixed $crontabDB 数据库对象
 * @access public
 * @return void
 */
function dataInsert($requestData, $validateRule, $tableCollectObj, $crontabDB)
{
    $data = array(
        'addtime' => time(),
        'addtime_str' => date('Y-m-d G:i:s', time()),
        'add_ip' => get_ip(),
        'self_url' => $requestData['HTTP_REFERER'],
        'TJJR1' => isset($requestData['TJJR1']) ? $requestData['TJJR1'] : '',//获得统计值（bi）
        'designer_enter_id' => $crontabDB['crontab']->auto_increment($tableCollectObj),
        'is_pass' => 0,
    );

    foreach ($validateRule as $field => $rule) {
        switch($rule['rule']) {
            case 'mobile' :
                $data[$field] = authCode($requestData[$field], "ENCODE");
                $data[$field.'_encrypt'] = md5($requestData[$field]);
            break;

            case 'number' :
                $data[$field] = intval($requestData[$field]);
            break;

            case 'empty|mobile' :
                if(isset($requestData[$field]) && is_mobile($requestData[$field])) {
                    $data[$field] = authCode($requestData[$field], "ENCODE");
                    $data[$field.'_encrypt'] = md5($requestData[$field]);
                } else {
                    $defaultVal = isset($rule['defaultval']) ? $rule['defaultval'] : '';
                    $data[$field] = isset($requestData[$field]) ? $requestData[$field] : $defaultVal;
                }
            break;

            default:
                $defaultVal = isset($rule['defaultval']) ? $rule['defaultval'] : '';
                $data[$field] = isset($requestData[$field]) ? $requestData[$field] : $defaultVal;
        }
    }

    $tableCollectObj->insert($data);

    $nextId = $data['designer_enter_id'];
    if(is_object($data['designer_enter_id']))
    {
        $nextId = $data['designer_enter_id']->value;
    }
    $nextId = intval($nextId);
    return $nextId;
}

/**
 * applySafeLimit 
 * 报名次数安全限制
 * @access public
 * @return void
 */
function applySafeLimit($applyTableObj, $requestData)
{
    //查找60秒是否存在提交数据 按照IP和手机号码限制
    //$mobile = $requestData['mobile '];
    //if ($requestData['mobile '] == '13200000001') {
        //return true;
    //}
    $start_time = time() - 60;
    $count_minute = $applyTableObj->count(
        array(
            '$or' => array(
                array(
                    'addtime' => array('$gt' => $start_time),
                    'mobile_encrypt' => md5($requestData['mobile']),
                ),
                array(
                    'addtime' => array('$gt' => $start_time),
                    'add_ip' => get_ip(),
                )
            )
        )
    );

    //直接返回报名成功
    if ($count_minute > 0){
        $retArr = responseStruct(200);
        responseData($retArr);
    }

    //一个手机号只能报一次名
    $start_time = strtotime('today');
    $count_day = $applyTableObj->count(
        array(
            //'addtime' => array('$gt' => $start_time),
            'mobile_encrypt' => md5($requestData['mobile']),
        )
    );
    if ($count_day >= 1){
        $retArr = responseStruct(200);
        $retArr['msg'] = 'xxx!';
        responseData($retArr);
    }

    return true;
}


/**
 * fieldVerify 
 * 字段验证
 * @access public
 * @return void
 */
function fieldVerify(Array $validateRule, $requestData)
{
    if(empty($requestData)) {
        $requestData = getRequest();//获取请求数据
    }
    foreach ($validateRule as $field => $rule) {
        switch($rule['rule']) {
            case 'notnull' :
                if (empty($requestData[$field])) {
                    $retArr = responseStruct(404);
                    $retArr['msg'] = $rule['declare'].'不能为空!';
                    responseData($retArr);
                }
            break;
            case 'mobile' :
                if (!is_mobile($requestData[$field])) {
                    $retArr = responseStruct(404);
                    $retArr['msg'] = $rule['declare'].'格式错误!';
                    responseData($retArr);
                }
            case 'number' :
                if (!is_numeric($requestData[$field])) {
                    $retArr = responseStruct(404);
                    $retArr['msg'] = $rule['declare'].'需要是数字!';
                    responseData($retArr);
                }
            break;
        }
    }
}


/**
 * getCrontabDB 
 * 获取crontab
 * @access public
 * @return void
 */
function getCrontabDB()
{
    $retMap = array();

    global $mdb_name;
    $crontab = new Crontab();
    $obj_mongo = $crontab->get_mongo_object();
    $xz_db = $obj_mongo->selectDB($mdb_name['zx']);

    $retMap['db'] = $xz_db;
    $retMap['crontab'] = $crontab;
    return $retMap;
}



/**
 * safetyVerification 
 * 接口安全验证
 * @param mixed $requestData 
 * @access public
 * @return void
 */
function safetyVerification($requestData)
{

    //来源url限制
    if (empty($requestData['HTTP_REFERER']) || (false === strpos($requestData['HTTP_REFERER'], 'xxx.com') && false === strpos($requestData['HTTP_REFERER'], 'xxx.com.cn')))
    {
        $retArr = responseStruct(404);
        responseData($retArr);
    }

    //ip黑名单限制
    $result = api_post("http://10.10.21.164/api/sys/ip/check", 'POST', json_encode(array("ip" => get_ip())));//132.123.222.222
    $result_json = json_decode($result, true);
    //需要封掉的ip地址
    if ($result_json['result']['type'] == "black") {
        $retArr = responseStruct(404);
        $retArr['msg'] = "IP 地址限制";
        responseData($retArr);
    }

    return true;
}


/**
 * getRequest 
 * 获取请求数据
 * @access public
 * @return void
 */
function getRequest()
{
    $requestDataArr = array();

    $post_data = file_get_contents("php://input");
    if ($post_data) {
        $requestDataArr = json_decode($post_data, true);
    } else {
        $requestDataArr = $_REQUEST;
    }
    $requestDataArr['HTTP_REFERER'] = isset($_SERVER['HTTP_REFERER']) ? $_SERVER['HTTP_REFERER'] : '';
    return $requestDataArr;
}

/**
 * responseStructBase 
 * 响应数据结构
 * @param mixed $code 
 * @access public
 * @return void
 */
function responseStruct($code, $customArr = array())
{
    $codeMap = array(
        200 => '请求成功!',
        401 => '未知错误!',
        404 => '非法访问!',
    );

    $retMap = array(
        'code' => $code,
        'msg' => $codeMap[$code],
    );
    $retMap = array_merge($retMap, $customArr);
    return $retMap;
}

/**
 * responseData 
 * 接口响应
 * @param mixed $data 
 * @access public
 * @return void
 */
function responseData($data)
{
    $requestData = getRequest();
	if (isset($requestData['callback']) && $requestData['callback'])
		echo  $requestData['callback'] . '(' . json_encode($data, JSON_UNESCAPED_UNICODE) . ')';
	else
		echo json_encode($data, JSON_UNESCAPED_UNICODE);
    die;
}




/**
 * is_mobile 
 * 判断是否为手机号
 * @param mixed $str 
 * @access public
 * @return void
 */
function is_mobile($str)
{
	if (preg_match("/1[34578]{1}[0-9]{9}$/", $str))
		return true;
	else
		return false;
}

/**
 * authCode 
 * 手机号码加解密
 * @param mixed $string 
 * @param string $operation 
 * @param string $key 
 * @param int $expiry 
 * @access public
 * @return void
 */
function authCode($string, $operation = 'DECODE', $key = '', $expiry = 0)
{
	return $string;
}

/**
 * get_ip 
 * 获取IP
 * @access public
 * @return void
 */
function get_ip()
{
	if (getenv('HTTP_CLIENT_IP') && strcasecmp(getenv('HTTP_CLIENT_IP'), 'unknown')) {
		$onlineip = getenv('HTTP_CLIENT_IP');
	} elseif (getenv('HTTP_X_FORWARDED_FOR') && strcasecmp(getenv('HTTP_X_FORWARDED_FOR'), 'unknown')) {
		$onlineip = getenv('HTTP_X_FORWARDED_FOR');
	} elseif (getenv('REMOTE_ADDR') && strcasecmp(getenv('REMOTE_ADDR'), 'unknown')) {
		$onlineip = getenv('REMOTE_ADDR');
	} elseif (isset($_SERVER['REMOTE_ADDR']) && $_SERVER['REMOTE_ADDR'] && strcasecmp($_SERVER['REMOTE_ADDR'], 'unknown')) {
		$onlineip = $_SERVER['REMOTE_ADDR'];
	}

	preg_match("/[\d\.]{7,15}/", $onlineip, $onlineipmatches);
	$ip = $onlineipmatches[0] ? $onlineipmatches[0] : '';

	return $ip;
}

/**
 * api_post 
 * 接口请求方法
 * @param mixed $url 
 * @param mixed $method 
 * @param mixed $data 
 * @access public
 * @return void
 */
function api_post($url, $method, $data = NULL)
{
	$headers = array(
		'Content-Type: text/html; charset=UTF-8'
	);
	$handle = curl_init();
	curl_setopt($handle, CURLOPT_URL, $url);
	curl_setopt($handle, CURLOPT_HTTPHEADER, $headers);
	curl_setopt($handle, CURLOPT_TIMEOUT, 10); //设置cURL允许执行的最长秒数。
	curl_setopt($handle, CURLOPT_RETURNTRANSFER, true);
	curl_setopt($handle, CURLOPT_SSL_VERIFYHOST, false);
	curl_setopt($handle, CURLOPT_SSL_VERIFYPEER, false);
	switch ($method) {
		case 'GET':
			break;
		case 'POST':
			curl_setopt($handle, CURLOPT_POST, true);
			curl_setopt($handle, CURLOPT_POSTFIELDS, $data);
			break;
		case 'PUT':
			curl_setopt($handle, CURLOPT_CUSTOMREQUEST, 'PUT');
			curl_setopt($handle, CURLOPT_POSTFIELDS, http_build_query($data));
			break;
		case 'DELETE':
			curl_setopt($handle, CURLOPT_CUSTOMREQUEST, 'DELETE');
			break;
	}
	$response = curl_exec($handle);
	return $response;
}

