
$(document).ready(function () {
	
	var paramstr = window.location.search.substr(1);
	console.log(paramstr);

    const URL = "http://netcan-env-2.eba-dmmnb4wj.eu-west-3.elasticbeanstalk.com/api/v1/";

    $("#subir").click(function (e) {
        e.preventDefault(); 
        vacunas = {};
        vacunas["idPerro"] = $("#IdPerro").val();
        vacunas["idVeterinatio"] = $("#IdVeterinario").val();
        vacunas["codVacuna"] = $("#CodVacuna").val();
        vacunas["fechaBaja"] = $("#FAlta").val();
        vacunas["fechaBajaProteccion"] = $("#FBaja").val();
        vacunas["idProteccion"] = $("#Descripcion").val();

        if ($("#IdPerro").val() != "" && $("#IdVeterinario").val() != "" && $("#CodVacuna").val() != "" && $("#FAlta").val() != "" && $("#FBaja").val() != "", $("#Descripcion").val() != "") {
            console.log(JSON.stringify(vacunas));
            $.ajax({
                type: 'POST',
				url: URL + "vacunas",
				data: JSON.stringify(vacunas),
				contentType: 'application/json; charset=utf-8',
                success: function (data, textStatus, request) {
					console.log("OK");
					if(confirm("¿Volver al cuadro de mandos?")){
						window.location = "../index.html?" + paramstr;
					}else{
						window.location = "RegistroVacuna.html?" + paramstr;
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
