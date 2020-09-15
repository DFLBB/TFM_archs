
$(document).ready(function () {
	
	var paramstr = window.location.search.substr(1);
	console.log(paramstr);

    const URL = "http://netcan-env-2.eba-dmmnb4wj.eu-west-3.elasticbeanstalk.com/api/v1/";

    $("#subir").click(function (e) {
        e.preventDefault();
        Bperro = {};
        /*Bafijo["nombre"] = $("#Nombre").val();*/
        Bperro["id"] = $("#idPerro").val();
        Bperro["fechaDefuncion"] = $("#FDefuncion").val();

        if ($("#idPerro").val() != "" && $("#FDefuncion").val() !="") {      
            console.log(JSON.stringify(Bperro));     
            $.ajax({
                type: 'POST',
				url: URL + "perros/baja",
				data: JSON.stringify(Bperro),
				contentType: 'application/json; charset=utf-8',
                success: function (data, textStatus, request) {
					console.log("OK");
					if(confirm("¿Volver al cuadro de mandos?")){
						window.location = "../index.html?" + paramstr;
					}else{
						window.location = "BajaPerro.html?" + paramstr;
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
