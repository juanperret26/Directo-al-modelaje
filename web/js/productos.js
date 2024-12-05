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

    /*fetch("/productos", { method: "GET" })
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
    });*/
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
        celdaNombre.textContent = element.nombre;
        fila.appendChild(celdaNombre);

        var celdaTipo = document.createElement("td");
        celdaTipo.textContent = element.tipo_producto;
        fila.appendChild(celdaTipo);

        var celdaPeso = document.createElement("td");
        celdaPeso.textContent = element.peso;
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

        var celdaEditar = document.createElement("td");
        var botonEditar = document.createElement("button");
        botonEditar.className = "boton-editar";
        botonEditar.id = "editar-" + index;
        botonEditar.innerHTML = `<i class="fa-solid fa-pen" style="color: #ffffff;"></i>`;
        celdaEditar.appendChild(botonEditar);
        fila.appendChild(celdaEditar);

        var celdaEliminar = document.createElement("td");
        var botonEliminar = document.createElement("button");
        botonEliminar.className = "boton-eliminar";
        botonEliminar.id = "eliminar-" + index;
        botonEliminar.innerHTML = `<i class="fa-solid fa-trash" style="color: #ffffff;"></i>`;
        celdaEliminar.appendChild(botonEliminar);
        fila.appendChild(celdaEliminar);

        tbody.appendChild(fila);
    });
  //evento boton editar  
    tbody.addEventListener("click", function (event) {
      if (event.target.classList.contains("boton-editar")) {
          const botonEditar = event.target;
          const fila = botonEditar.closest("tr");
          manejarEdicion(fila);
      }
  });
  //evento boton eliminar
  tbody.addEventListener("click", function (event) {
    if (event.target.classList.contains("boton-eliminar")) {
        const botonEditar = event.target;
        const fila = botonEditar.closest("tr");
        const primeraCelda = fila.querySelector("td:first-child");
        const textoPrimeraCelda = primeraCelda.textContent;
        eliminar(textoPrimeraCelda);
    }
});
};



document.addEventListener("keyup",e=>{
    if(e.target.matches("#barraBuscador")){
        document.querySelectorAll(".nombreCelda").forEach(id =>{
            id.textContent.toLowerCase().includes(e.target.value.toLowerCase())
            ? id.parentElement.classList.remove("filtro")
            : id.parentElement.classList.add("filtro");    
        });
    };
});

function obtenerFechaDesdeCadena(cadenaFechaHora) {
    const partes = cadenaFechaHora.split("T");
    const fecha = partes[0];
    
    return fecha;
}

//Boton editar
function manejarEdicion(fila) {
  // Obtener elementos de la fila
  const celdas = fila.querySelectorAll("td");

  // Crear un objeto para almacenar los valores de las celdas
  const valoresCeldas = {};
  var productoID = "";
  const encabezado = document.querySelector("table thead tr");

  celdas.forEach((celda, index) => {
    if (index >= 0 && index <= 7) {
      if(index == 0){
        var tituloCelda = encabezado.querySelectorAll("th")[index].textContent;
        valoresCeldas[tituloCelda] = celda.textContent.trim();    
        productoID = valoresCeldas[tituloCelda].toString();
      }
      else{
      var tituloCelda = encabezado.querySelectorAll("th")[index].textContent;
      valoresCeldas[tituloCelda] = celda.textContent.trim();
      }
    }
  });

  // Redirigir a la página "formProductos" y enviar los valores como parámetros
  const queryString = new URLSearchParams(valoresCeldas).toString();
  window.location.href = `/front/html/formProductos.html?productoID=${productoID}&${queryString}`;
}

//eliminar
function eliminar(ID) {
  const id = ID;
  const url = `/productos/${id}`;
  const datos = null;
  makeRequest(
      url,
      Method.DELETE, 
      datos,
      ContentType.JSON,
      CallType.PRIVATE,
      exitoSolicitud,
      errorSolicitud
  );
  function exitoSolicitud(data) {
      console.log("éxito.");
      location.reload();
      // Realiza otras acciones si es necesario
  }

  function errorSolicitud(status, response) {
      console.error("Error . Estado:", status, "Respuesta:", response);
      // Maneja el error de acuerdo a tus necesidades
  }
  
  /*fetch(url, {
    method: "DELETE"
  })
    .then(response => {
      if (!response.ok) {
        throw new Error("Error al eliminar el producto.");
      }
      location.reload();
      console.log("Producto eliminado con éxito.");
    })
    .catch(error => {
      console.error("Error al eliminar el producto:", error);
    });*/
}
