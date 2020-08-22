import $ from "jquery";
import swal from "sweetalert";

window.onload = () => {

	function REST( url:string, data:any={}, callback:Function ) {
		$.post( url, JSON.stringify(data), function(response:any) {
			callback(response);
		} );
	}

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
			email: ( $(this).find('input#toggle_email').prop('checked') ? '_checked' : '_uncheck' ),
			welcome: $(this).find('input#str_event_message').val(),
		}

		var that = $(this);

		$.post(
			'/admin/create/event',
			JSON.stringify(data),
			function( response:any ) {
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

	$(document).on('click', '[data-action]', function() {
		var action = $(this).data('action');
		var eventID= $(this).data('event-id');
		var data:any = {
			eventid: eventID,
		}

		var requestURL = '';
		if( 'gen_join_link' === action ) {
			requestURL = '/admin/gen/joinlink';

			var username = $(`#user_name_${eventID}`).val();
			if(!username) return;
			data.username = username;

			REST(requestURL, data, function(resp:any) {
				if( resp.success ) {
					$(`#join_link_${eventID}`).val( resp.payload.url );
				} else {
					swal("Error", "Something went wrong", "warning");
				}
			});
		} else if ( 'delete' === action ) {
			requestURL = '/admin/delete/event';
			swal({
				title: "Are you sure you want to delete event?",
				text: "This action cannot be undone",
				icon: "warning",
				dangerMode: true,
			}).then((remove) => {
				if( remove ) {
					REST(requestURL, data, function(resp:any) {
						if( resp.success ) {
							swal("Event", "Event Deleted", "success");
							$(`#event_wrapper_${eventID}`).remove();
						} else {
							swal("Error", "Something went wrong", "warning");
						}
					});
				}
			})
			.catch((err)=>{
				throw new Error(err);
			})
		}
	});
}