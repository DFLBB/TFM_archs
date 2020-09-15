
$(document).ready(function () {
	
	var paramstr = window.location.search.substr(1);
	console.log(paramstr);

    const URL = "http://netcan-env-2.eba-dmmnb4wj.eu-west-3.elasticbeanstalk.com/api/v1/";

    $("#subir").click(function (e) {
        e.preventDefault();
        Bafijo = {};
        Bafijo["nombre"] = $("#Nombre").val();
        Bafijo["id"] = $("#idAfijo").val();

        if ($("#idAfijo").val() != "" && $("#Nombre").val() != "") {      
            console.log(JSON.stringify(Bafijo));     
            $.ajax({
                type: 'POST',
				url: URL + "afijos/baja",
				data: JSON.stringify(Bafijo),
				contentType: 'application/json; charset=utf-8',
                success: function (data, textStatus, request) {
					console.log("OK");
					if(confirm("¿Volver al cuadro de mandos?")){
						window.location = "../index.html?" + paramstr;
					}else{
						window.location = "BajaAfijo.html?" + paramstr;
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
