function foo()
{
	alert('dddddddddddddd');
}

function start_ajax_create_note( url, callback_type )
{

	var post_key_values = "";
	post_key_values = post_key_values + 'note_type=' + document.getElementById('id_note_type').value;
	post_key_values = post_key_values + '&note_detail=' + document.getElementById('id_note_detail').value;
	
	ajax_blocking( url,
		function(what)
		{
			alert("what"+what);
			if( callback_type == 11 )
			{
				
			}
		}, 
		'POST', 
		post_key_values
	);
}

function show_create_new_note_frame( which, action, ajax_callback )
{

	var which_html_obj = document.getElementById('id_' + which);
	var submit_button_function = '';
	if( typeof(ajax_callback) != 'undefined')
	{
		submit_button_function = "onclick='return start_ajax_create_note(\""+action+"\", "+ ajax_callback +");'";
	}

	
	which_html_obj.innerHTML = 
		"<form action="+action+" method='POST'>\
			<div class='one-line-container'><span>type:</span>\
				<select id='id_note_type' name='note_type' >\
					<option value=1 selected='selected' >TECH</option>\
					<option value=2 >Daily normal</option>\
					<option value=3 >Link</option>\
				</select>\
			</div>\
			<div class='one-line-container'><span>note:</span><textarea id='id_note_detail' name='note_detail' ></textarea></div>\
			<div class='one-line-container'><input type='submit' "+ submit_button_function +" value='Submit'/></div>\
		</form>";
		

			
}