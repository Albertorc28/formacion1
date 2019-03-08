$(document).ready (function () {
    $("#btn-enviar").click(function(){
        var texto = $("#txt-texto").val();
        var fecha = new Date();
        var peticion= {
            Palabra: texto,
            Fecha: fecha.getDate() + '/' + fecha.getMonth()+1 + '/'+ fecha.getFullYear()
        }
        console.log(peticion);
        $.ajax({
            type: 'POST',
            url: 'http://localhost:8080/ejercicio',
            data: JSON.stringify(peticion),
            dataType: 'json',
            contentType:'application/json',
        })
        .done(function (data) {
            $("#response p").html(data.Palabra);
            console.log( "correctamente" );
        })
        .fail(function (data) {
            console.log( "no correctamente" );
        })
        .always(function (data) {
            console.log( "complete" );
        });
    });
});