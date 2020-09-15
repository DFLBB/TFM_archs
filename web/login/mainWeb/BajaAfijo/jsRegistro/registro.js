
$(document).ready(function () {

    const URL = "http://netcan-env-1.eba-6e2hpndt.eu-west-3.elasticbeanstalk.com/api/v1/";

    $("#propietario").hide();
    $("#veterinario").hide();
    $("#federacion").hide();
    console.log($("#tipo").val());

    $("#continuar").click(function (e) {
        e.preventDefault();
        var email = $("#email").val();
        var nickname = $("#nickname").val();
        var password = $("#password").val();
        var telefono = $("#telefono").val();
        var tipo = $("#tipo").val();

        if (email != "" && nickname != "" && password != "" && telefono != "" && tipo != "") {
            if (tipo == "Propietario") {
                $("#propietario").fadeIn("slow");
                $("#veterinario").hide();
                $("#federacion").hide();
            } else if (tipo == "Veterinario") {
                $("#veterinario").fadeIn("slow");
                $("#propietario").hide();
                $("#federacion").hide();
            } else if (tipo == "Federacion"){
                $("#federacion").fadeIn("slow");
                $("#propietario").hide();
                $("#veterinario").hide();
            }else{
				alert("Es necesario elegir un tipo");
			}
        } else {
            alert("Es necesario rellenar todos los campos");
        }
    });

    $("#botonPropietario").click(function (e) {
        e.preventDefault();
        var vacio = false;
        login = {};
        login["nickname"] = $("#nickname").val();
        login["email"] = $("#email").val();
        login["password"] = $("#password").val();
        login["telefono"] = $("#telefono").val();
        login["tipo"] = $("#tipo").val();
        propietario = {};
        camposPropietario = ["nombre", "apellido1", "apellido2", "tipoDoc", "documento", "direccion", "ciudad", "pais"];
        for (var i = 1; i < 9; i++) {
            var campo = "#propietario" + i;
            if ($(campo).val() != "") {
                propietario[camposPropietario[i - 1]] = $(campo).val();
            } else {
                vacio = true;
            }
        }

        login["propietario"] = propietario;
		
        if (!vacio) {
			console.log(JSON.stringify(login));
            $.ajax({
                type: 'POST',
				url: URL + "login",
				data: JSON.stringify(login),
				contentType: 'application/json; charset=utf-8',
                success: function (data, textStatus, request) {
					console.log("OK");
                        window.location = "login.html";
                },
                error: function (request, textStatus, errorThrown) {
					console.log(errorThrown);
					console.log(textStatus);
					console.log(request);
                    alert("Error en los parametros introducidos");
                }
            });
        } else {
            alert("Se deben de rellenar todos los parametros");
        }

    });
	
	$("#botonVeterinario").click(function (e) {
        e.preventDefault();
        var vacio = false;
        login = {};
        login["nickname"] = $("#nickname").val();
        login["email"] = $("#email").val();
        login["password"] = $("#password").val();
        login["telefono"] = $("#telefono").val();
        login["tipo"] = "V";
        veterinario = {};
        camposVeterinario = ["nombre", "apellido1", "apellido2", "colegiado","pais" , "direccion", "ciudad",  "clinicaVeterinaria"];
        for (var i = 1; i < 9; i++) {
            var campo = "#veterinario" + i;
            if ($(campo).val() != "") {
                veterinario[camposVeterinario[i - 1]] = $(campo).val();
            } else {
                vacio = true;
            }
        }

        login["veterinario"] = veterinario;
		
        if (!vacio) {
			console.log(JSON.stringify(login));
            $.ajax({
                type: 'POST',
				url: URL + "login",
				data: JSON.stringify(login),
				contentType: 'application/json; charset=utf-8',
                success: function (data, textStatus, request) {
					console.log("OK");
                        window.location = "login.html";
                },
                error: function (request, textStatus, errorThrown) {
					console.log(errorThrown);
					console.log(textStatus);
					console.log(request);
                    alert("Error en los parametros introducidos");
                }
            });
        } else {
            alert("Se deben de rellenar todos los parametros");
        }

    });
	
	$("#botonFederacion").click(function (e) {
        e.preventDefault();
        var vacio = false;
        login = {};
        login["nickname"] = $("#nickname").val();
        login["email"] = $("#email").val();
        login["password"] = $("#password").val();
        login["telefono"] = $("#telefono").val();
        login["tipo"] = "F";
        federacion = {};
        camposFederacion = ["nombre", "pais"];
        for (var i = 1; i < 3; i++) {
            var campo = "#federacion" + i;
            if ($(campo).val() != "") {
                federacion[camposFederacion[i - 1]] = $(campo).val();
            } else {
                vacio = true;
            }
        }

        login["federacion"] = federacion;
		
        if (!vacio) {
			console.log(JSON.stringify(login));
            $.ajax({
                type: 'POST',
				url: URL + "login",
				data: JSON.stringify(login),
				contentType: 'application/json; charset=utf-8',
                success: function (data, textStatus, request) {
					console.log("OK");
                        window.location = "login.html";
                },
                error: function (request, textStatus, errorThrown) {
					console.log(errorThrown);
					console.log(textStatus);
					console.log(request);
                    alert("Error en los parametros introducidos");
                }
            });
        } else {
            alert("Se deben de rellenar todos los parametros");
        }

    });

});
