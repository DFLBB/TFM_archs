$(document).ready(function () {

	$("#perrete").slideUp(3000);
	
	var paramstr = window.location.search.substr(1);
	console.log(paramstr);
    var paramarr = paramstr.split("?");
    var params = paramarr[0].split("&");
    var sep = params[1].split("=");
    var id = sep[1];
	var tipo = params[1];
	console.log(tipo);
	
	if(tipo == "tipo=P"){
		$("#propietario").fadeIn("slow");
        $("#veterinario").hide();
        $("#federacion").hide();
	}else if(tipo == "tipo=V"){
		$("#veterinario").fadeIn("slow");
        $("#propietario").hide();
        $("#federacion").hide();
	}else{
		$("#federacion").fadeIn("slow");
        $("#propietario").hide();
        $("#veterinario").hide();
	}
	
	$("#registraPerroPropietario").click(function (e) {
        e.preventDefault();
        window.location = "RegistroPerro/registroPerro.html?" + paramstr;
    });
	
	$("#registraMuertePropietario").click(function (e) {
        e.preventDefault();
        window.location = "BajaPerro/BajaPerro.html?" + paramstr;
    });
	
	$("#registrarAfijoFederacion").click(function (e) {
        e.preventDefault();
        window.location = "RegistroAfijo/registroAfijo.html?" + paramstr;
    });
	
	$("#bajaAfijoFederacion").click(function (e) {
        e.preventDefault();
        window.location = "BajaAfijo/BajaAfijo.html?" + paramstr;
    });
	
	$("#registroPerroVeterinario").click(function (e) {
        e.preventDefault();
        window.location = "RegistroPerro/registroPerro.html?" + paramstr;
    });
	
	$("#registrarDefuncionVeterinario").click(function (e) {
        e.preventDefault();
        window.location = "BajaPerro/BajaPerro.html?" + paramstr;
    });
	
	$("#registrarChipVeterinario").click(function (e) {
        e.preventDefault();
        window.location = "RegistroMicroChip/registroMicroChip.html?" + paramstr;
    });
	
	$("#registroVacunaVeterinario").click(function (e) {
        e.preventDefault();
        window.location = "RegistroVacuna/registroVacuna.html?" + paramstr;
    });
	
	
	
	

});