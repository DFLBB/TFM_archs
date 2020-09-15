
$(document).ready(function () {
	
	var paramstr = window.location.search.substr(1);
	console.log(paramstr);

    const URL = "http://netcan-env-2.eba-dmmnb4wj.eu-west-3.elasticbeanstalk.com/api/v1/";

    $("#subir").click(function (e) {
        e.preventDefault();
        afijo = {};
        afijo["nombre"] = $("#Nombre").val();
        afijo["idPersona1"] = $("#Propietario1").val();
		
		if($("#Propietario2").val() == ""){
			afijo["idPersona2"] = 0;
		}else{
			afijo["idPersona2"] = $("#Propietario2").val();
		}
		
		if($("#Propietario2").val() == ""){
			afijo["idPersona3"] = 0;
		}else{
			afijo["idPersona3"] = $("#Propietario3").val();
		}
       
        if ($("#Nombre").val() != "" && $("#Propietario1").val() != "") {   		
            console.log(JSON.stringify(afijo));     
            $.ajax({
                type: 'POST',
				url: URL + "afijos",
				data: JSON.stringify(afijo),
				contentType: 'application/json; charset=utf-8',
                success: function (data, textStatus, request) {
					console.log("OK");
					if(confirm("¿Volver al cuadro de mandos?")){
						window.location = "../index.html?" +paramstr;
					}else{
						window.location = "registroAfijo.html?" + paramstr;
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
