(function(){
	var obj = document.getElementById("single-page-tbar");
	var top = getTop(obj);
	var isIE6 = /msie 6/i.test(navigator.userAgent);
	window.onscroll = function(){
		var bodyScrollTop = document.documentElement.scrollTop || document.body.scrollTop;
		if (bodyScrollTop > top){
			obj.style.position = (isIE6) ? "absolute" : "fixed";
			obj.style.top = (isIE6) ? bodyScrollTop + "px" : "0px";
		} else {
			obj.style.position = "static";
		}
	}
	function getTop(e){
		var offset = e.offsetTop;
		if(e.offsetParent != null) offset += getTop(e.offsetParent);
		return offset;
	}
})()