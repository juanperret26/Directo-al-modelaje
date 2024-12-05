
document.getElementById("aceptar").addEventListener("click", function() {
    crearNuevoCamion();
});
function crearNuevoCamion() {
    const patente = document.querySelector('input[placeholder="Patente"]').value;
    const pesoMax = parseFloat(document.querySelector('input[placeholder="Peso maximo"]').value);
    const costoKM = parseInt(document.querySelector('input[placeholder="Costo por Km"]').value);
    const nuevoCamion = {
        patente: patente,
        peso_maximo: pesoMax,
        costo_km: costoKM
    };

    const url = "http://localhost:8080/camiones/";
    const datos = nuevoCamion;
    makeRequest(
        url,
        Method.POST, 
        datos,
        ContentType.JSON,
        CallType.PRIVATE,
        exitoSolicitud,
        errorSolicitud
    );
    function exitoSolicitud(data) {
        console.log("éxito.");
        window.location.href = "/camiones.html";
        // Realiza otras acciones si es necesario
    }
  
    function errorSolicitud(status, response) {
        console.error("Error . Estado:", status, "Respuesta:", response);
        // Maneja el error de acuerdo a tus necesidades
    }

    /*fetch("/camiones", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(nuevoCamion)
    })
    .then(response => {
        if (response.ok) {
            window.location.href = "/htmlcamiones";
        } else {
            // Manejar errores
            response.json().then(data => {
                console.log(data);});
            console.error("Error al crear un nuevo camion");
        }
    })
    .catch(error => {
        console.error("Error al crear un nuevo camion:", error);
    });*/
}