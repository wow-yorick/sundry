function setword(zh,en){
    var name =document.getElementsByName('name');
    var targetwin =document.getElementsByName('targetwin');
    
    var title =document.getElementsByName('title');
    var webtitle =document.getElementsByName('webtitle');
    var keywords =document.getElementsByName('keywords');
    var description =document.getElementsByName('description');
    name[0].value=en;
    targetwin[0].value=en;
    title[0].value=zh;
    webtitle[0].value=zh;
    keywords[0].value=zh;
    description[0].value=zh;
    
}
setword('城市设计分院作品集','cssheji')


function setword(zh,en){
    var name =document.getElementsByName('name');
    var targetwin =document.getElementsByName('targetwin');
    
    var title =document.getElementsByName('title');
    var webtitle =document.getElementsByName('webtitle');
    var keywords =document.getElementsByName('keywords');
    var description =document.getElementsByName('description');
    name[0].value=en;
    targetwin[0].value=en;
    title[0].value=zh;
    webtitle[0].value=zh;
    keywords[0].value=zh;
    description[0].value=zh;
    
}
document.getElementById('btnAdd').click();
//setword('景观二所作品集','cssheji');
setTimeout("setword('景观二所作品集','cssheji')",4000)
document.getElementById('btnSubmit').click();

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