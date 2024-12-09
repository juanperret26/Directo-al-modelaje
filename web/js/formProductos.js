
document.getElementById("aceptar").addEventListener("click", function() {
    crear();
});
function crear() {
    const nombre = document.querySelector('input[placeholder="Nombre"]').value;
    const tipo = document.querySelector('input[placeholder="Tipo"]').value;
    const peso = parseFloat(document.querySelector('input[placeholder="Peso unit."]').value);
    const precio = parseFloat(document.querySelector('input[placeholder="Precio"]').value);
    const stock = parseFloat(document.querySelector('input[placeholder="Stock"]').value);
    const stockMinimo = parseInt(document.querySelector('input[placeholder="Stock minimo"]').value);

    const nuevoProducto = {
        id: " ",
        nombre: nombre,
        TipoProducto: tipo,
        peso_unitario: peso,
        precio: precio,
        stock: stock,
        stockMinimo: stockMinimo,
        Tipo: " "
    };

    const url = `http://localhost:8080/productos/`;
    const datos = nuevoProducto;
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
        window.location.href = "/html/productos.html"
        // Realiza otras acciones si es necesario
    }
  
    function errorSolicitud(status, response) {
        console.error("Error al crear producto . Estado:", status, "Respuesta:", response);
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
  
    var productoID = "";

    if (urlParams.has("productoID")) {
        productoID = urlParams.get("productoID");
    }

    function cargarValorEnCampo(placeholder, paramName) {
      if (urlParams.has(paramName)) {

        const boton_crear = document.getElementById("aceptar");
        boton_crear.style.display = "none";
        const boton_editar = document.getElementById("editar");
        boton_editar.style.display = "block";

        const input = document.querySelector(`input[placeholder="${placeholder}"]`);
        input.value = urlParams.get(paramName);
      }
    }
  
    cargarValorEnCampo("Nombre", "Nombre");
    cargarValorEnCampo("Tipo", "Tipo");
    cargarValorEnCampo("Peso unit.", "Peso unit.");
    cargarValorEnCampo("Precio", "Precio");
    cargarValorEnCampo("Stock", "Stock");
    cargarValorEnCampo("Stock minimo", "StockMinimo");
    cargarValorEnCampo("Tipo", "Tipo");

    document.getElementById("editar").addEventListener("click", function() {
        editar(productoID); // Pasar productoID a la función editar
    });
});


//funcion para el boton editar
//document.getElementById("editar").addEventListener("click", function() {
//    editar();
//});
function editar(productoID) {

    const codigo = productoID;
    const nombre = document.querySelector('input[placeholder="Nombre"]').value;
    const tipo = document.querySelector('input[placeholder="Tipo"]').value;
    const peso = parseFloat(document.querySelector('input[placeholder="Peso unit."]').value);
    const precio = parseFloat(document.querySelector('input[placeholder="Precio"]').value);
    const stock = parseFloat(document.querySelector('input[placeholder="Stock"]').value);

    const objetoEditado = {
        id: codigo,
        nombre: nombre,
        TipoProducto: tipo,
        peso_unitario: peso,
        precio: precio,
        stock: stock
    };

    
        const id = productoID;
        const url = `http://localhost:8080/productos/${id}`;
        const datos = objetoEditado;
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
            window.location.href = "http://localhost:8081/html/productos.html";
            // Realiza otras acciones si es necesario
        }
      
        function errorSolicitud(status, response) {
            console.error("Error al crear el producto . Estado:", status, "Respuesta:", response);
            // Maneja el error de acuerdo a tus necesidades
        }   

}
