// JavaScript Document
addLoadEvent(displayAbbreviations);


function addLoadEvent(func){
	var oldonload = window.onload;
	if (typeof window.onload != "function"){
	window.onload = func
	}else {
		window.onload = function () {
			oldonload();
			func();
			}
	}

}


function displayAbbreviations(){
	if(!document.getElementByTagName || !document.createElement || !document.createTextNode) return false;
	//ȡ���������Դ�
	var abbreviations = document.getElementByTagName("abbr");
	if(abbreviations.length<1) return false;
	var defs = new Array();
	//������Щ���Դ�
	for(var i=0; i<abbreviations.length; i++) {
		var current_abbr = abbreviations[i]
		if(current_abbr.childNodes.length < 1) continue;
		var definition = current_abbr.getAttribute("title");
		var key = abbreviations[i].lastChild.nodeValue;
		defs[key] = definition;
		}
	//���������б�
	var dlist = document.createElement("dl");
	//��������
	for(key in defs){
		var definition = defs[key];
		//�����������
		var dtitle = document.createElement("dt");
		var dtitle_text = document.createTextNode("key");
		dtitle.appendChild(dtitle_text);
		//������������
		var ddesc=document.createElement("dd");
		var ddesc_text = document.createTextNode("definition");
		ddesc.appendChild(ddesc_text);
		//��������ӵ������б�
		dlist.appendChild(dtitle);
		dlist.appendChild(ddesc);
		}
		if(dlist.childNodes.length < 1) return false;
		//��������
	var header = document.createElement("h2");
	var header_text = document.createTextNode("Abberviations");
	header.appendChild(header_text);
	//�ѱ�����뵽ҳ������
	document.getElementByTagName("body")[0].appendChild(header);
	//�Ѷ����б���ӵ�ҳ������
    document.getElementByTagName("body")[0].appendChild(dlist);
	}



