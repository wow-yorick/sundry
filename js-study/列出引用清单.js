// JavaScript Document
function displayCitations(){
	var quotes = document.getElementByTagName("blackquote");
	for(var i = 0;i<quotes.length;i++){
		if(!quotes[i].getAttribute("cite")){
			continue;
			var url = quotes[i].getAttribute("cite");
			var quoteChilden = quotes[i].getElementByTagName("*");
			if(quoteChilden.length<1) continue;
			var elem = quoteChilden[quoteChilden.length-1];
			var alink = document.createElement("a");
			var link_text = document.createTextNode("source");
			alink.appendChild(link_text);
			alink.setAttribute("herf",url);
			var superscript = document.createElement("sup");
			superscript.appendChild("alink");
			elem.appendChild(superscript);
			}
		}
	}
	
