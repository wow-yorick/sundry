/*
function InitAjax(){ 
	var ajax; 
	if(window.ActiveXObject){ 
		var versions = ['Microsoft.XMLHTTP', 'MSXML.XMLHTTP', 'Microsoft.XMLHTTP', 'Msxml2.XMLHTTP.7.0', 'Msxml2.XMLHTTP.6.0', 'Msxml2.XMLHTTP.5.0', 'Msxml2.XMLHTTP.4.0', 'MSXML2.XMLHTTP.3.0', 'MSXML2.XMLHTTP']; 
		for(var i=0; i <versions.length; i++) { 
			try { 
				ajax = new ActiveXObject(versions[i]); 
				if(ajax) { 
					return ajax; 
				} 
			} catch(e) {} 
		} 
	}else if(window.XMLHttpRequest) { 
		ajax = new XMLHttpRequest(); 
	} 
	return ajax; 
} 

//js自动下载文件到本地 
var xh; 
function getXML(geturl){ 
	xh = InitAjax(); 
	xh.onreadystatechange = getReady; 
	xh.open("GET", geturl, true); 
	xh.send(); 
} 

function getReady(){ 
	//alert(xh.readyState); 
	if (xh.readyState == 4) { 
		alert(xh.status); 
		if (xh.status == 200) { 
			saveFile("d:\test.flv"); 
			return true; 
		}else{ return false; } 
	}else 
		return false; 
} 

function saveFile(tofile) { 
	alert(tofile); 
	var objStream; 
	var imgs; 
	imgs = xh.responseBody; 
	objStream = new ActiveXObject("ADODB.Stream"); 
	objStream.Type = 1; 
	objStream.open(); 
	objStream.write(imgs); 
	objStream.SaveToFile(tofile) 
} 
//getXML("http://10.76.3.116/2.bmp"); 
//js自动下载文件到本地结束 
*/