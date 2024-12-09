
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
        window.location.href = "http://localhost:8081/html/camiones.html";
        // Realiza otras acciones si es necesario
    }
  
    function errorSolicitud(status, response) {
        console.error("Error . Estado:", status, "Respuesta:", response);
        // Maneja el error de acuerdo a tus necesidades
    }
}


document.addEventListener("DOMContentLoaded", function () {
    const boton_crear = document.getElementById("aceptar");
    boton_crear.style.display = "block";
    const boton_editar = document.getElementById("editar");
    boton_editar.style.display = "none";

    const queryString = window.location.search;
    const urlParams = new URLSearchParams(queryString);
    
    console.log(urlParams.toString());

    var camionID = "";

    if (urlParams.has("camionID")) {
        camionID = urlParams.get("camionID");
    }

    function cargarValorEnCampo(placeholder, paramName) {
      if (urlParams.has(paramName)) {

        const boton_crear = document.getElementById("aceptar");
        boton_crear.style.display = "none";
        const boton_editar = document.getElementById("editar");
        boton_editar.style.display = "block";

        const input = document.querySelector(`input[placeholder="${placeholder}"]`);
        input.value = urlParams.get(paramName);
        console.log(input.value);
      }
    }
  
    cargarValorEnCampo("Patente", "Patente");
    cargarValorEnCampo("Peso maximo", "Peso max.");
    cargarValorEnCampo("Costo por Km", "Costo/KM")

    document.getElementById("editar").addEventListener("click", function() {
        editar(camionID); // Pasar productoID a la función editar
    });

    
});



function editar(camionID) {

    const codigo = camionID;
    const patente = document.querySelector('input[placeholder="Patente"]').value;
    const peso_maximo = parseFloat(document.querySelector('input[placeholder="Peso maximo"]').value);
    const costo_Km = parseFloat(document.querySelector('input[placeholder="Costo por Km"]').value);
    
    const objetoEditado = {
        id: codigo,
        patente: patente,
        peso_maximo: peso_maximo,
        costo_km: costo_Km,
    };

    
        const id = camionID;
        const url = `http://localhost:8080/camiones/`;
        const datos = objetoEditado;
        console.log(objetoEditado);

        makeRequest(
            url,
            Method.PUT, 
            datos,
            ContentType.JSON,
            CallType.PRIVATE,
            exitoSolicitud,
            errorSolicitud
        );
        function exitoSolicitud(data) {
            console.log("éxito.");
            window.location.href = "http://localhost:8081/html/camiones.html";
            // Realiza otras acciones si es necesario
        }
      
        function errorSolicitud(status, response) {
            console.error("Error al crear el producto . Estado:", status, "Respuesta:", response);
            // Maneja el error de acuerdo a tus necesidades
        }   

}