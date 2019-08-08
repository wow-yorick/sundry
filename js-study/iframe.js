function SetCwinHeight(ifmid){ 
	var iframeid=document.getElementById(ifmid); //iframe id 
	alert(iframeid)
	if (document.getElementById){ 
		if (iframeid && !window.opera){ 
     		if (iframeid.contentDocument && iframeid.contentDocument.body.offsetHeight){ 
      			iframeid.height = iframeid.contentDocument.body.offsetHeight+50; 
			}else if(iframeid.Document && iframeid.Document.body.scrollHeight){ 
      			iframeid.height = iframeid.Document.body.scrollHeight+50; 
     		} 
		} 
	} 
} 