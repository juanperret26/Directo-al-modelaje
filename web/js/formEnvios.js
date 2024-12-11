document.getElementById("aceptar").addEventListener("click", function() {
    crear();
});
function crear() {
    const Patente = document.querySelector('input[placeholder="Patente"]').value;
    const CiudadDestino = document.querySelector('input[placeholder="Ciudad destino"]').value;
    const KM = parseInt(document.querySelector('input[placeholder="Km"]').value);
    const Id = obtenerDatosSeleccionados();

    const nuevo = {
        pedidos: [Id],
        destino: {
            ciudad: CiudadDestino,
            kilometros: KM
        },
        patente_camion: Patente,
    };

    console.log("Datos enviados:", JSON.stringify(nuevo, null, 2));


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
        window.location.href = "/html/envios.html";
        // Realiza otras acciones si es necesario
    }
  
    function errorSolicitud(status, response) {
        console.error("Error . Estado:", status, "Respuesta:", response);
        // Maneja el error de acuerdo a tus necesidades
    }
}
  
//cargar lista pedidos
document.addEventListener("DOMContentLoaded",function(){
    cargarDatos();  
  });    
  
  function cargarDatos(){

    const url = `http://localhost:8080/pedidos/`;
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
    console.log(datos);
    datos.forEach(function(element){
        var fila = document.createElement("tr");
        
        var celdaId = document.createElement("td");
        celdaId.textContent = element.id;
        celdaId.className = "nombreCelda";
        fila.appendChild(celdaId);
        /*
        var celdaProductos = document.createElement("td");
        celdaProductos.textContent = element.PedidoProductos;
        fila.appendChild(celdaProductos);*/

        var celdaProductos = document.createElement("td");
        if (element.PedidoProductos && Array.isArray(element.PedidoProductos)) {
                var nombresProductos = element.PedidoProductos.map(function(producto) {
                return producto.nombre;
        });
            celdaProductos.textContent = nombresProductos.join(", ");
        } else {
            celdaProductos.textContent = "N/A";
        }
        fila.appendChild(celdaProductos);

        var celdaCiudad = document.createElement("td");
        celdaCiudad.textContent = element.destino;
        fila.appendChild(celdaCiudad);

        var celdaEstado = document.createElement("td");
        celdaEstado.textContent = element.estado;
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
            datosSeleccionados = CodigoProducto.toString();
        }
    });

    return datosSeleccionados;
}

document.addEventListener("DOMContentLoaded", function () {
    const boton_crear = document.getElementById("aceptar");
    const boton_editar = document.getElementById("editar");

    // Inicializar el formulario en modo crear
    boton_crear.style.display = "block";
    boton_editar.style.display = "none";

    // Obtener los parámetros de la URL
    const queryString = window.location.search;
    const urlParams = new URLSearchParams(queryString);

    var envioID = "";

    if (urlParams.has("envioID")) {
        envioID = urlParams.get("envioID");
    }

    // Función para cargar los valores en el formulario
    function cargarValorEnCampo(placeholder, paramName, data) {
        const input = document.querySelector(`input[placeholder="${placeholder}"]`);
        const keys = paramName.split('.');  // Permite acceder a campos anidados como "destino.ciudad"
        let value = data;

        keys.forEach(key => {
            value = value[key];
        });

        input.value = value || ''; // Establecer el valor del campo
    }

    // Si existe envioID, realizar la carga de datos para edición
    if (envioID) {
        boton_crear.style.display = "none";
        boton_editar.style.display = "block";

        // Cargar los datos del envío en el formulario usando makeRequest (Petición GET)
        const url = `http://localhost:8080/envios/${envioID}`;
        const datos = null;

        makeRequest(
            url,
            Method.GET,  // Método GET
            datos,
            ContentType.JSON,  // Tipo de contenido JSON
            CallType.PRIVATE,  // Llamada privada
            exitoSolicitud,
            errorSolicitud
        );

        function exitoSolicitud(data) {
            // Llenar los campos del formulario con los datos obtenidos
            cargarValorEnCampo("Patente", "patente_camion", data);
            cargarValorEnCampo("Km", "destino.kilometros", data);
            cargarValorEnCampo("Ciudad destino", "destino.ciudad", data);
        }

        function errorSolicitud(status, response) {
            console.error("Error al cargar los datos del envío. Estado:", status, "Respuesta:", response);
        }
    }

    // Acción para el botón de editar
    document.getElementById("editar").addEventListener("click", function () {
        editar(envioID);  // Pasar envioID a la función editar
    });
});




function editar(envioID) {
    const Patente = document.querySelector('input[placeholder="Patente"]').value;
    const CiudadDestino = document.querySelector('input[placeholder="Ciudad destino"]').value;
    const KM = parseFloat(document.querySelector('input[placeholder="Km"]').value);
    const Id = obtenerDatosSeleccionados();
    const id = envioID;
    estado = null;

    if (envioID) {
        // Solicitar los detalles del envío por ID usando makeRequest
        const url = `http://localhost:8080/envios/${envioID}`;
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
            estado = data.estado;
        }
    }


    const objetoEditado = {
        estado: estado,
        id: id,
        pedidos: [Id],
        destino: {
            ciudad: CiudadDestino,
            kilometros: KM
        },
        patente_camion: Patente,
    };

    // Realizar la solicitud PUT para editar el envío
    const url = `http://localhost:8080/envios/`;
    makeRequest(
        url,
        Method.PUT,
        objetoEditado,
        ContentType.JSON,
        CallType.PRIVATE,
        exitoSolicitud,
        errorSolicitud
    );

    function exitoSolicitud(data) {
        console.log("Éxito al editar el envío.");
        window.location.href = "http://localhost:8081/html/envios.html"; // Redirigir a la página principal de envíos
    }

    function errorSolicitud(status, response) {
        console.error("Error al editar. Estado:", status, "Respuesta:", response);
        window.location.href = "http://localhost:8081/html/envios.html";
    }
}
