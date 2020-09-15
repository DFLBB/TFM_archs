
$(document).ready(function () {

    const URL = "http://netcan-env-2.eba-dmmnb4wj.eu-west-3.elasticbeanstalk.com/api/v1/";

    $("#subir").click(function (e) {
        e.preventDefault();
        var email = $("#email").val();
        var nickname = $("#nickname").val();
        var password = $("#password").val();

        if (email != "" && nickname != "" && password != "") {

            $.ajax({
                type: 'GET',
                url: URL + "login?nickname=" + nickname + "&email=" + email + "&password=" + password,
                success: function (data, textStatus, request) {
                    console.log(data);
                    window.location = "mainWeb/index.html?id=" + data.id + "&tipo=" + data.tipo;
                },
                error: function (request, textStatus, errorThrown) {
                    alert("Error en los parametros introducidos");
                }
            });

        }
    });

});
