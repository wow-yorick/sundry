Math.floor(Math.random()*10+1) // 随机数
idString = idString.substr(1); //截取字符串
$('#data_formDetail').find("input[name=deliverymethod][value="+jsonDataInfo['deliverymethod']+"]").attr("checked",true); //jquery设置单选按钮的值

objiframe = $(document.getElementById('iframe1').contentWindow.document.body); // 获取框架页面内容

$(document).ready(function(){
	var height = $(window).height();
	var width = $(window).width();
	width = (width <1366)? 1366:width;
	height = (height <760)? 760:height;
	$('#container').height(height);
	$('#container').width(width);
})

setInterval("sdf()",1000);

setTimeout(‘setHint()’,1000);

stringObject.substr(start,length);

stringObject.indexOf(searchvalue,fromindex)


query取得iframe中元素的几种方法

 

在iframe子页面获取父页面元素
代码如下:

$('#objId', parent.document);
// 搞定...


在父页面 获取iframe子页面的元素
代码如下:

$("#objid",document.frames('iframename').document) 

 

$(document.getElementById('iframeId').contentWindow.document.body).html()
 
 显示iframe中body元素的内容。

 
$("#testId", document.frames("iframename").document).html();

 根据iframename取得其中ID为"testId"元素


$(window.frames["iframeName"].document).find("#testId").html()

 jQuery("#select1  option:selected").text();



用JS或jQuery访问页面内的iframe,兼容IE/FF 
注意:框架内的页面是不能跨域的!

假设有两个页面,在相同域下.

index.html 文件内含有一个iframe:

XML/HTML代码
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"

"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">  
<html xmlns="http://www.w3.org/1999/xhtml">  
<head>  
<meta http-equiv="Content-Type" content="text/html; charset=gb2312" />  
<title>页面首页</title>  
</head>  
  
<body>  
<iframe src="iframe.html" id="koyoz" height="0" width="0"></iframe>  
</body>  
</html>   
iframe.html 内容:

XML/HTML代码
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"

"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">  
<html xmlns="http://www.w3.org/1999/xhtml">  
<head>  
<meta http-equiv="Content-Type" content="text/html; charset=gb2312" />  
<title>iframe.html</title>  
</head>  
  
<body>  
<div id="test">www.koyoz.com</div>  
</body>  
</html>  

1. 在index.html执行JS直接访问:

JavaScript代码
document.getElementById('koyoz').contentWindow.document.getElementById

('test').style.color='red'   
通过在index.html访问ID名为'koyoz'的iframe页面,并取得此iframe页面内的ID为'test'的

对象,并将其颜色设置为红色.

此代码已经测试通过,能支持IE/firefox .

2. 在index.html里面借助jQuery访问:

JavaScript代码
$("#koyoz").contents().find("#test").css('color','red');   
此代码的效果和JS直接访问是一样的,由于借助于jQuery框架,代码就更短了.
 

收集网上的一些示例：
用jQuery在IFRAME里取得父窗口的某个元素的值
只好用DOM方法与jquery方法结合的方式实现了

1. 在父窗口中操作 选中IFRAME中的所有单选钮
$(window.frames["iframe1"].document).find("input:radio").attr("checked","true");

2. 在IFRAME中操作 选中父窗口中的所有单选钮
$(window.parent.document).find("input:radio").attr("checked","true");

父窗口想获得IFrame中的Iframe，就再加一个frames子级就行了，如：
$(window.frames["iframe1"].frames["iframe2"].document).find("input:radio").attr("checked","true");

3.在子窗口中调用父窗口中的另一个子窗口的方法(FRAME)：

  parent.frames["Main"].Fun();
4.调用父窗口的父窗口
  $(window.parent.parent.document) 
