var xmlHttp
var xmlHttp_callback

function GetXmlHttpObject()
{
	var xmlHttp=null;
	
	try
	{
		// Firefox, Opera 8.0+, Safari
		xmlHttp=new XMLHttpRequest();
		
		
	}catch (e){

		// Internet Explorer
		try
		{
			xmlHttp=new ActiveXObject("Msxml2.XMLHTTP");
			
		}catch (e){
			xmlHttp=new ActiveXObject("Microsoft.XMLHTTP");
		}
	}
	
	return xmlHttp;
}

function stateChanged() 
{ 

	//if (xmlHttp.readyState==4 || xmlHttp.readyState=="complete")
	if (xmlHttp.readyState==4 )
	{ 
		if( xmlHttp.status == 200 || xmlHttp.status == 0 )
		{
			if( typeof( xmlHttp_callback ) != "undefined" )
			{
				xmlHttp_callback(xmlHttp.responseText)
				return ;
			}
		}
	} 
}

function ajax( url, callback, _method)
{
	var method = "GET";
	if( typeof(_method) != "undefined" )
	{
		method = _method;
	}
	
	xmlHttp=GetXmlHttpObject()
	if (xmlHttp==null)
	{
		alert ("Browser does not support HTTP Request")
		return
	} 
	
	xmlHttp_callback = callback

	xmlHttp.onreadystatechange=stateChanged 
	xmlHttp.open( method,url,true)
	xmlHttp.send(null)
}

function ajax_blocking( url, callback, _method, post_key_values )
{
	var method = "GET";
	if( typeof(_method) != "undefined" )
	{
		method = _method;
	}
	
	xmlHttp=GetXmlHttpObject()
	if (xmlHttp==null)
	{
		alert ("Browser does not support HTTP Request")
		return
	} 
	
	xmlHttp_callback = callback

	xmlHttp.onreadystatechange=stateChanged 
	xmlHttp.open( method,url,true)
	
	if( method == "POST" || method == "post" )
	{	
		xmlHttp.setRequestHeader("Content-Type","application/x-www-form-urlencoded"); 
		xmlHttp.send(post_key_values);
	}else{
		xmlHttp.send(null);
	}
}
