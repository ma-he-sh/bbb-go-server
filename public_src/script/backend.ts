import $ from "jquery";
import swal from "sweetalert";

import "../styles/backend.scss";

window.onload = () => {

	$('#admin_create_event').submit(function(e:any) {
		e.preventDefault();

		var eventName:any = $(this).find('input#str_event_name').val();
		if (eventName == "") {
			return;
		}

		var data = {
			action: $(this).find('input#form_action').val(),
			active: ( $(this).find('input#toggle_active').prop('checked') ? '_checked' : '_uncheck' ),
			eventName: $(this).find('input#str_event_name').val(),
			moderatorpw: $(this).find('input#str_moderator_pw').val(),
			attendeepw : $(this).find('input#str_attendee_pw').val(),
			record: ( $(this).find('input#toggle_record').prop('checked') ? '_checked' : '_uncheck' ),
			welcome: $(this).find('input#str_event_message').val(),
		}

		var that = $(this);

		$.post(
			'/admin/create/event',
			JSON.stringify(data),
			function( response:any ) {
				console.log( response );
				if( response.success ) {
					that.trigger("reset");
					swal("Event", "Event Created", "success");
				} else {
					swal("Error", "Something went wrong", "warning");
				}
			},
			'json'
		);
	});
}