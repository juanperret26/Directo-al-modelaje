document.getElementById("aceptar").addEventListener("click", function() {
    crear();
});
function crear() {
    const Patente = document.querySelector('input[placeholder="Patente"]').value;
    const CiudadDestino = document.querySelector('input[placeholder="Ciudad destino"]').value;
    const KM = parseInt(document.querySelector('input[placeholder="Km"]').value);
    const Id = obtenerDatosSeleccionados();
    /*<input class="input" type="text" placeholder="Patente">
        <input class="input" type="text" placeholder="Tipo">
        <input class="input" type="text" placeholder="Peso unitario">
        <input class="input" type="text" placeholder="Precio">
        <input class="input" type="text" placeholder="Stock">
        <input class="input" type="text" placeholder="Stock minimo">
        <input class="input" type="text" placeholder="Tipo">
        
        {
        "patente_camion": "ABC123",
        "destino": {
             "Ciudad": "Santa Fe",
             "Kilometros": 100
        },

        "pedidos": ["654525ae2f0e529636076e85"]
}
        
        */

    const nuevo = {
        patente_camion: Patente,
        destino: {
            Ciudad: CiudadDestino,
            Kilometros: KM
        },
        pedidos: [Id]
    };

    const url = `http://localhost:8080/envios/`;
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
        window.location.href = "/front/html/envios.html";
        // Realiza otras acciones si es necesario
    }
  
    function errorSolicitud(status, response) {
        console.error("Error . Estado:", status, "Respuesta:", response);
        // Maneja el error de acuerdo a tus necesidades
    }

    /*fetch("/envios", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(nuevo)
    })
    .then(response => {
        if (response.ok) {
            // La solicitud se realizó con éxito, puedes redirigir a la página de productos u realizar alguna otra acción
            window.location.href = "/htmlenvios";
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

document.addEventListener("DOMContentLoaded", function () {
    //const boton_crear = document.getElementById("aceptar");
    //boton_crear.style.display = "block";
    //const boton_editar = document.getElementById("editar");
    //boton_editar.style.display = "none";

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

    /*document.getElementById("editar").addEventListener("click", function() {
        editar(productoID); // Pasar productoID a la función editar
    });*/
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

    const url = `/productos`;
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

  
//cargar lista pedidos
document.addEventListener("DOMContentLoaded",function(){
    cargarDatos();  
  });    
  
  function cargarDatos(){

    const url = `/pedidos`;
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

      /*fetch("/pedidos", { method: "GET" })
      .then(response => {
        if (!response.ok) {
          throw new Error("Error al obtener datos de pedidos.");
        }
        return response.json();
      })
      .then(data => {
        mostrarDatosTabla(data);
      })
      .catch(error => {
        console.error("Error al obtener datos de pedidos:", error);
      });*/
  };
  

function mostrarDatosTabla(datos){
    var table = document.getElementById("TablaPrincipal");
    var tbody = document.getElementById("TableBody");
    datos.forEach(function(element){
        var fila = document.createElement("tr");
        
        var celdaId = document.createElement("td");
        celdaId.textContent = element.Id;
        celdaId.className = "nombreCelda";
        fila.appendChild(celdaId);
        /*
        var celdaProductos = document.createElement("td");
        celdaProductos.textContent = element.PedidoProductos;
        fila.appendChild(celdaProductos);*/

        var celdaProductos = document.createElement("td");
        if (element.PedidoProductos && Array.isArray(element.PedidoProductos)) {
                var nombresProductos = element.PedidoProductos.map(function(producto) {
                return producto.Nombre;
        });
            celdaProductos.textContent = nombresProductos.join(", ");
        } else {
            celdaProductos.textContent = "N/A";
        }
        fila.appendChild(celdaProductos);

        var celdaCiudad = document.createElement("td");
        celdaCiudad.textContent = element.Destino;
        fila.appendChild(celdaCiudad);

        var celdaEstado = document.createElement("td");
        celdaEstado.textContent = element.Estado;
        fila.appendChild(celdaEstado);

        var celdaCheckbox = document.createElement("td");
        var checkbox = document.createElement("input");
        checkbox.type = "checkbox";
        checkbox.className = "checkbox-fila";
        celdaCheckbox.appendChild(checkbox);
        fila.appendChild(celdaCheckbox);

        tbody.appendChild(fila);
    });
};

function obtenerDatosSeleccionados() {
    const filas = document.querySelectorAll("#TableBody tr");
    var datosSeleccionados = "";

    filas.forEach(function (fila) {
        const checkbox = fila.querySelector(".checkbox-fila");
        const inputCantidad = fila.querySelector(".input-cantidad");

        if (checkbox.checked) {
            const celdaCodigo = fila.querySelector(".nombreCelda");

            const CodigoProducto = celdaCodigo.textContent;
            datosSeleccionados = CodigoProducto;
        }
    });

    return datosSeleccionados;
}