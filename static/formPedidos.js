document.getElementById("aceptar").addEventListener("click", function() {
    crear();
});
function crear() {
    const destino = document.querySelector('input[placeholder="Destino"]').value;
    const datosSeleccionados = obtenerDatosSeleccionados();

    const nuevo = {
        PedidoProductos:datosSeleccionados,
        Destino: destino
    };

    fetch("/pedidos", {
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
    });
}


//traer tabla pedidos
document.addEventListener("DOMContentLoaded",function(){
    cargarDatos();  
  });    
  
  function cargarDatos(){
      fetch("/productos", { method: "GET" })
      .then(response => {
        if (!response.ok) {
          throw new Error("Error al obtener datos de productos.");
        }
        return response.json();
      })
      .then(data => {
        mostrarDatosTabla(data);
      })
      .catch(error => {
        console.error("Error al obtener datos de productos:", error);
      });
  };
  
  function mostrarDatosTabla(datos){
      var table = document.getElementById("TablaPrincipal");
      var tbody = document.getElementById("TableBody");
  
      datos.forEach(function(element,index){
          var fila = document.createElement("tr");
          
          var celdaId = document.createElement("td");
          celdaId.textContent = element.Id;
          celdaId.className = "nombreCelda";
          fila.appendChild(celdaId);
  
          var celdaNombre = document.createElement("td");
          celdaNombre.textContent = element.Nombre;
          fila.appendChild(celdaNombre);
  
          var celdaTipo = document.createElement("td");
          celdaTipo.textContent = element.TipoProducto;
          fila.appendChild(celdaTipo);
  
          var celdaPeso = document.createElement("td");
          celdaPeso.textContent = element.Peso_unitario;
          fila.appendChild(celdaPeso);
  
          var celdaPrecio = document.createElement("td");
          celdaPrecio.textContent = element.Precio;
          fila.appendChild(celdaPrecio);
  
          var celdaStock = document.createElement("td");
          celdaStock.textContent = element.Stock;
          fila.appendChild(celdaStock);
  
          var celdaActualizacion = document.createElement("td");
          celdaActualizacion.textContent = obtenerFechaDesdeCadena(element.Actualizacion);
          fila.appendChild(celdaActualizacion);

          var celdaCreacion = document.createElement("td");
          celdaCreacion.textContent = obtenerFechaDesdeCadena(element.Creacion);
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
    const partes = cadenaFechaHora.split("T");
    const fecha = partes[0];
    
    return fecha;
}

//obtener filas check
function obtenerDatosSeleccionados() {
    const filas = document.querySelectorAll("#TableBody tr");
    const datosSeleccionados = [];

    filas.forEach(function (fila) {
        const checkbox = fila.querySelector(".checkbox-fila");
        const inputCantidad = fila.querySelector(".input-cantidad");

        if (checkbox.checked) {
            const celdaCodigo = fila.querySelector(".nombreCelda");
            const celdaNombre = fila.querySelector("td:nth-child(3)"); // Cambia el índice según la posición de la columna Nombre
            const celdaPrecio = fila.querySelector("td:nth-child(5)"); // Cambia el índice según la posición de la columna Precio

            const CodigoProducto = celdaCodigo.textContent;
            const Nombre = celdaNombre.textContent;
            const Cantidad = parseInt(inputCantidad.value);
            const Precio_unitario = parseInt(celdaPrecio.textContent);            

            datosSeleccionados.push({ CodigoProducto, Nombre, Cantidad, Precio_unitario });
        }
    });

    return datosSeleccionados;
}