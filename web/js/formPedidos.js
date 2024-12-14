document.getElementById("aceptar").addEventListener("click", function() {
    crear();
});
function crear() {
    const destino = document.querySelector('input[placeholder="Destino"]').value;
    const datosSeleccionados = obtenerDatosSeleccionados();

    console.log(datosSeleccionados);

    const nuevo = {
        PedidoProductos:datosSeleccionados,
        Destino: destino,
        Estado: "pendiente"
    };

    console.log(nuevo);

    const url = `http://localhost:8080/pedidos/`;
    const datos = nuevo;
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
        window.location.href = "/html/pedidos.html"
        // Realiza otras acciones si es necesario
    }
  
    function errorSolicitud(status, response) {
        console.error("Error . Estado:", status, "Respuesta:", response);
        // Maneja el error de acuerdo a tus necesidades
    }

    /*fetch("/pedidos", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(nuevo)
    })
    .then(response => {
        if (response.ok) {
            // La solicitud se realizó con éxito, puedes redirigir a la página de productos u realizar alguna otra acción
            window.location.href = "/htmlpedidos";
        } else {
            // Manejar errores
            response.json().then(data => {
                console.log(data);});
            console.error("Error al crear un nuevo pedido");
        }
    })
    .catch(error => {
        console.error("Error al crear un nuevo pedido:", error);
    });*/
}


//traer tabla productos
document.addEventListener("DOMContentLoaded",function(){
    cargarDatos();  
  });    
  
  function cargarDatos(){

    const url = `http://localhost:8080/productos/`;
    const datos = null;
    makeRequest(
        url,
        Method.GET, 
        datos,
        ContentType.JSON,
        CallType.PRIVATE,
        exitoSolicitud,
        errorSolicitud
    );
    function exitoSolicitud(data) {
        console.log("éxito.");
        mostrarDatosTabla(data);
        // Realiza otras acciones si es necesario
    }
  
    function errorSolicitud(status, response) {
        console.error("Error . Estado:", status, "Respuesta:", response);
        // Maneja el error de acuerdo a tus necesidades
    }
  };

  function mostrarDatosTabla(datos){
      var table = document.getElementById("TablaPrincipal");
      var tbody = document.getElementById("TableBody");
  
      datos.forEach(function(element,index){
          var fila = document.createElement("tr");
          
          var celdaId = document.createElement("td");
          celdaId.textContent = element.id;
          celdaId.className = "nombreCelda";
          fila.appendChild(celdaId);
  
          var celdaNombre = document.createElement("td");
          celdaNombre.textContent = element.nombre;
          fila.appendChild(celdaNombre);
  
          var celdaTipo = document.createElement("td");
          celdaTipo.textContent = element.tipoProducto;
          fila.appendChild(celdaTipo);
  
          var celdaPeso = document.createElement("td");
          celdaPeso.textContent = element.peso_unitario;
          fila.appendChild(celdaPeso);
  
          var celdaPrecio = document.createElement("td");
          celdaPrecio.textContent = element.precio;
          fila.appendChild(celdaPrecio);
  
          var celdaStock = document.createElement("td");
          celdaStock.textContent = element.stock;
          fila.appendChild(celdaStock);

          var celdaActualizacion = document.createElement("td");
          celdaActualizacion.textContent = obtenerFechaDesdeCadena(element.actualizacion);
          fila.appendChild(celdaActualizacion);

          var celdaCreacion = document.createElement("td");
          celdaCreacion.textContent = obtenerFechaDesdeCadena(element.creacion);
          fila.appendChild(celdaCreacion);  
  
          var celdaCheckbox = document.createElement("td");
          var checkbox = document.createElement("input");
          checkbox.type = "checkbox";
          checkbox.className = "checkbox-fila";
          celdaCheckbox.appendChild(checkbox);
          fila.appendChild(celdaCheckbox);

          var celdaCantidad = document.createElement("td");
          var inputCantidad = document.createElement("input");
          inputCantidad.type = "number";
          inputCantidad.className = "input-cantidad";
          celdaCantidad.appendChild(inputCantidad);
          fila.appendChild(celdaCantidad);
          tbody.appendChild(fila);
      });
};

function obtenerFechaDesdeCadena(cadenaFechaHora) {
    if (!cadenaFechaHora || typeof cadenaFechaHora !== "string") {
        return "Fecha no disponible"; // Devuelve un valor por defecto
    }

    const partes = cadenaFechaHora.split("T");
    const fecha = partes[0];
    return fecha;
}

//obtener filas check
function obtenerDatosSeleccionados() {
    const filas = document.querySelectorAll("#TableBody tr"); // Selecciona todas las filas
    const productosSeleccionados = [];

    filas.forEach(fila => {
        const checkbox = fila.querySelector(".checkbox-fila");
        const inputCantidad = fila.querySelector(".input-cantidad");

        if (checkbox.checked) {
            const codigoProducto = fila.cells[0].textContent; // Columna "Código"
            const nombre = fila.cells[1].textContent; // Columna "Nombre"
            const tipoProducto = fila.cells[2].textContent; // Columna "Tipo"
            const precio_unitario = parseInt(fila.cells[5].textContent, 10); // Columna "Precio"
            const cantidad = parseInt(inputCantidad.value, 10) || 1;

            // Crear objeto correctamente
            productosSeleccionados.push({
                codigoProducto: codigoProducto,
                nombre: nombre, // Asegura que esto sea el nombre (HP Victus 16, etc.)
                tipoProducto: tipoProducto,
                precio_unitario: precio_unitario,
                cantidad: cantidad
            });
        }
    });

    return productosSeleccionados;
}



//editar



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
        Id: codigo,
        nombre: nombre,
        TipoProducto: tipo,
        peso_unitario: peso,
        precio: precio,
        stock: stock
    };

    
    const url = `http://localhost:8080/productos/`;
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
        window.location.href = "/front/html/productos.html";
        // Realiza otras acciones si es necesario
    }
  
    function errorSolicitud(status, response) {
        console.error("Error . Estado:", status, "Respuesta:", response);
        // Maneja el error de acuerdo a tus necesidades
    }

    /*fetch("/productos", {
        method: "PUT",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(objetoEditado)
    })
    .then(response => {
        if (response.ok) {
            // La solicitud se realizó con éxito, puedes redirigir a la página de productos u realizar alguna otra acción
            window.location.href = "/htmlproductos";
        } else {
            // Manejar errores
            response.json().then(data => {
                console.log(data);});
            console.error("Error al crear un nuevo producto");
        }
    })
    .catch(error => {
        console.error("Error al crear un nuevo producto:", error);
    });*/
}
  