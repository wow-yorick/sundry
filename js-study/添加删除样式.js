addLoadEvent();

function addLoadEvent(func){
	var oldonload = window.onload;
	if(typeof window.onload !='function'){
		window.onload = func;
	} else {
		window.onload = function() {
			oldonload();
			func();
		}
	}
}

/*function isclick_(){
	$=function(obj){return document.getElementById(obj);}
	var list = $('nav').getElementsByTagName('a');
	for(var i=0;i<list.length;i++){
		deletClass(list[i],'curr');
	}	
	for(var i=0;i<list.length;i++){
		
		list[i].onclick=function(){
			addClass(this,'curr');
		}
	}

}*/

function isclick(obj){
	$=function(obj){return document.getElementById(obj);}
	var list = $('nav').getElementsByTagName('a');
		for(var i=0;i<list.length;i++){
			deletClass(list[i],'curr');
		}	
		obj.onclick=function(){
			addClass(this,'curr');
		}
		//释放click的焦点,解决ie67的兼容问题
	$('temp').focus();
}

function addClass(element,value){
	if(!element.className){
		element.className = value;
	}else{
		ishas = element.className.indexOf(value);
		if(element.className.indexOf(value) < 0){
			var newclass = element.className;
			newclass+=' ';
			newclass+= value;
			element.className = newclass;
		}		
	}
}

function deletClass(element,value){
	patt = new RegExp(value);
	if(patt.test(element.className)){
		element.className = element.className.replace(value,'');
	}
}