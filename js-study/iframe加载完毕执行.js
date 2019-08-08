var iframe = document.createElement("iframe");
iframe.src = "http://www.planabc.net";
if (iframe.attachEvent){
	iframe.attachEvent("onload", function(){
		alert("Local iframe is now loaded.");
		});
} else {
	iframe.onload = function(){
		alert("Local iframe is now loaded.");
		};
}

document.body.appendChild(iframe);

document.getElementById("iframe1").contentWindow.document.body.innerHTML;


RegExp("(^|\\?|&)keyId=([^&]*)(\\s|&|$)", "i")