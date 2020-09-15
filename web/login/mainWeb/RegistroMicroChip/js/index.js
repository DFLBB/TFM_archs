
$(document).ready(function () {
	
	var paramstr = window.location.search.substr(1);
	console.log(paramstr);

    const URL = "http://netcan-env-2.eba-dmmnb4wj.eu-west-3.elasticbeanstalk.com/api/v1/";

    $("#subir").click(function (e) {
        e.preventDefault();        
        Microchip = {};
        Microchip["idVeterinario"] = $("#IdVeterinario").val();
        Microchip["codMicrochip"] = $("#CodMicrochip").val();
        Microchip["idPerro"] = $("#IdPerro").val();

        if ($("#IdVeterinario").val() != "" && $("#CodMicrochip").val() != "" && $("#IdPerro").val() != "") {
            console.log(JSON.stringify(Microchip));
            $.ajax({
                type: 'POST',
				url: URL + "chips",
				data: JSON.stringify(Microchip),
				contentType: 'application/json; charset=utf-8',
                success: function (data, textStatus, request) {
					console.log("OK");
					if(confirm("¿Volver al cuadro de mandos?")){
						window.location = "../index.html?" + paramstr;
					}else{
						 window.location = "registroMicroChip.html?" + paramstr;
					}
                       
                },
                error: function (request, textStatus, errorThrown) {
					console.log(errorThrown);
					console.log(textStatus);
					console.log(request);
                    alert("Error en los parametros introducidos");
                }
            });

        }else{
				alert("Debe introducir todos los parametros");
			}
    });
	
		$("#cuadroMandos").click(function (e) {
        e.preventDefault();
        window.location = "../index.html?" + paramstr;

    });

});
