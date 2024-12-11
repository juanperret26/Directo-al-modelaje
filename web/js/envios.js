document.addEventListener("DOMContentLoaded",function(){
    cargarDatos();  
  });    
  
  function cargarDatos(){

    const url = `http://localhost:8080/envios/`;
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

    datos.forEach(function (element) {
      var fila = document.createElement("tr");

      // Celda ID
      var celdaId = document.createElement("td");
      celdaId.textContent = element.id;
      celdaId.className = "nombreCelda";
      fila.appendChild(celdaId);

      // Celda Estado
      var celdaEstado = document.createElement("td");
      celdaEstado.textContent = element.estado;
      fila.appendChild(celdaEstado);

      // Celda Paradas
      var celdaParadas = document.createElement("td");
      if (element.paradas && element.paradas.length > 0) {
          // Si hay paradas, mostrar la ciudad de la primera parada
          celdaParadas.textContent = element.paradas[0].ciudad;
      } else {
          // Si no hay paradas, mostrar "Sin paradas asignadas"
          celdaParadas.textContent = "Sin paradas asignadas";
      }
      fila.appendChild(celdaParadas);

      // Celda Destino
      var celdaDestino = document.createElement("td");
      celdaDestino.textContent = element.destino.ciudad;
      fila.appendChild(celdaDestino);

      // Celda Creación
      var celdaCreacion = document.createElement("td");
      celdaCreacion.textContent = element.fecha_creacion;
      fila.appendChild(celdaCreacion);

      // Celda Pedido
      var celdaPedido = document.createElement("td");
      celdaPedido.textContent = element.pedidos;
      fila.appendChild(celdaPedido);

      // Celda Actualización
      var celdaActualizacion = document.createElement("td");
      celdaActualizacion.textContent = element.actualizacion;
      fila.appendChild(celdaActualizacion);

      // Celda Costo
      var celdaCosto = document.createElement("td");
      celdaCosto.textContent = element.costo_total;
      fila.appendChild(celdaCosto);

      // Celda Editar
      var celdaEditar = document.createElement("td");
      var botonEditar = document.createElement("button");
      botonEditar.className = "boton-editar";
      botonEditar.innerHTML = `<i class="fa-solid fa-pen" style="color: #ffffff;"></i>`;
      celdaEditar.appendChild(botonEditar);
      fila.appendChild(celdaEditar);

      // Celda Paradas (botón de eliminar)
      var celdaEliminar = document.createElement("td");
      var botonEliminar = document.createElement("button");
      botonEliminar.className = "boton-eliminar";
      botonEliminar.innerHTML = `<i class="fa-solid fa-trash" style="color: #ffffff;"></i>`;
      celdaEliminar.appendChild(botonEliminar);
      fila.appendChild(celdaEliminar);

      // Celda Paradas (botón de paradas)
      var celdaParadasBoton = document.createElement("td");
      var botonParadas = document.createElement("button");
      botonParadas.className = "boton-paradas";
      botonParadas.innerHTML = `<i class="fa-solid fa-hand" style="color: #ffffff;"></i>`;
      celdaParadasBoton.appendChild(botonParadas);
      fila.appendChild(celdaParadasBoton);

      tbody.appendChild(fila);
  });

    tbody.addEventListener("click", function (event) {
      if (event.target.classList.contains("boton-eliminar")) {
          event.preventDefault(); // Evita la recarga de la página
  
          const botonEliminar = event.target;
          const fila = botonEliminar.closest("tr"); // Encuentra la fila correspondiente
          const primeraCelda = fila.querySelector("td:first-child"); // Encuentra la primera celda
          const textoPrimeraCelda = primeraCelda.textContent.trim(); // Obtén y limpia el texto del ID
  
          console.log("ID del producto a eliminar:", textoPrimeraCelda);
  
          eliminar(textoPrimeraCelda);
  
          fila.remove();
      }
    });

      //evento boton editar  
  tbody.addEventListener("click", function (event) {
    if (event.target.classList.contains("boton-editar")) {
        const botonEditar = event.target;
        const fila = botonEditar.closest("tr");
        manejarEdicion(fila);
      }
  });
  tbody.addEventListener("click", function (event) {
    if (event.target.classList.contains("boton-paradas")) {
        const botonParada = event.target;
        const fila = botonParada.closest("tr");
        manejarParada(fila);
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

//eliminar
function eliminar(ID) {
  const id = ID;
  const url = `http://localhost:8080/envios/${id}`;
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
      console.log(id);
      console.log("éxito.");
      // Realiza otras acciones si es necesario
  }

  function errorSolicitud(status, response) {
      console.error("Error . Estado:", status, "Respuesta:", response);
      // Maneja el error de acuerdo a tus necesidades
  }
}

//Boton editar
function manejarEdicion(fila) {
  const celdas = fila.querySelectorAll("td");
  const encabezados = document.querySelector("table thead tr").querySelectorAll("th");
  const datosEnvios = {};
  let envioID;

  celdas.forEach((celda, index) => {
    const tituloCelda = encabezados[index].textContent;
    datosEnvios[tituloCelda] = celda.textContent.trim();
    if (index === 0) {
      envioID = datosEnvios[tituloCelda];
    }
  });

  const queryString = new URLSearchParams(datosEnvios).toString();
  window.location.href = `/html/formEnvios.html?envioID=${envioID}&${queryString}`;
}



function manejarParada(fila) {
  const celdas = fila.querySelectorAll("td");
  const encabezados = document.querySelector("table thead tr").querySelectorAll("th");
  const datosEnvios = {};
  let envioID;

  // Verificar que hay suficientes encabezados para las celdas
  if (encabezados.length !== celdas.length) {
    console.error("El número de celdas no coincide con el número de encabezados.");
    return; // Salir de la función si no coinciden
  }

  celdas.forEach((celda, index) => {
    const tituloCelda = encabezados[index]?.textContent;  // Usar el operador opcional de encadenamiento
    if (tituloCelda) {
      datosEnvios[tituloCelda] = celda.textContent.trim();
      if (index === 0) {
        envioID = datosEnvios[tituloCelda];
      }
    } else {
      console.error(`No se encontró encabezado para la celda en el índice ${index}`);
    }
  });

  const queryString = new URLSearchParams(datosEnvios).toString();
  window.location.href = `/html/formParadas.html?paradaID=${envioID}&${queryString}`;
}
