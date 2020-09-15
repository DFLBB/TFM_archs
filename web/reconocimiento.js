$(document).ready(function (e) {
	
	const URL = 'http://netcan-env-2.eba-dmmnb4wj.eu-west-3.elasticbeanstalk.com/api/v1/bigdata/image';
	$('#identificarRaza').hide();
	 $("#procesando").hide();
	$('#resultado').hide();
	var fileName = "";
	
	$('input[type="file"]').change(function(e){
		fileName = e.target.files[0];
	});
		
	$('#upload').on('click', function () {
		
		if(fileName != ""){
			console.log(fileName);
			var file_data = $('#file-input').prop('files')[0];
			var form_data = new FormData();
			form_data.append('file', file_data);
		$.ajax({
			url: URL, // point to server-side controller method
			dataType: 'text', // what to expect back from the server
			cache: false,
			contentType: false,
			processData: false,
			data: form_data,
			type: 'post',
			success: function (response) {
				$('#identificarRaza').fadaIn();
			},
			error: function (response) {
				alert("Error al subir la imagen al servidor");
			}
		});
		}else{
			alert("No se ha introducido ninguna imagen");
		}
	});
	
	$('#identificar').on('click', function () {
		var imagenURL = URL+"?path=http://netcan-env-2.eba-dmmnb4wj.eu-west-3.elasticbeanstalk.com/web/uploads/" + fileName.name;
		$("#procesando").fadeIn();
		$.ajax({
                type: 'GET',
                url: imagenURL,
                success: function (data, textStatus, request) {
					$("#imagenPerro").attr("src","http://netcan-env-2.eba-dmmnb4wj.eu-west-3.elasticbeanstalk.com/web/uploads/" + fileName.name);
					$("#imagenPerro").attr("height","200");
					$("#imagenPerro").attr("width","400");
					$("#raza").text("Unai comeme las pelotas");
                    console.log(data);
                },
                error: function (request, textStatus, errorThrown) {
					$("#procesando").hide();
					console.log(request.responseText);
					console.log(textStatus);
					console.log(errorThrown);
					$("#resultado").fadeIn();
					$("#imagenPerro").attr("src","http://netcan-env-2.eba-dmmnb4wj.eu-west-3.elasticbeanstalk.com/web/uploads/" + fileName.name);
					$("#imagenPerro").attr("height","200");
					$("#imagenPerro").attr("width","400");
					$("#raza").text(request.responseText);

                }
            });
	});
});
