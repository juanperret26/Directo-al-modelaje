document.addEventListener("DOMContentLoaded",function(){
    cargarDatos();  
});    
  
  function cargarDatos(){

    const url = `/camiones`;
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

      /*fetch("/camiones", { method: "GET" })
      .then(response => {
        if (!response.ok) {
          throw new Error("Error al obtener datos de camiones.");
        }
        return response.json();
      })
      .then(data => {
        mostrarDatosTabla(data);
      })
      .catch(error => {
        console.error("Error al obtener datos de camiones:", error);
      });*/  
    };



function mostrarDatosTabla(datos){

    var table = document.getElementById("TablaPrincipal");
    var tbody = document.getElementById("TableBody");
    
    datos.forEach(function(element){
        var fila = document.createElement("tr");
        
        var celdaId = document.createElement("td");
        celdaId.textContent = element.ID;
        celdaId.className = "nombreCelda";
        fila.appendChild(celdaId);

        var celdaPatente = document.createElement("td");
        celdaPatente.textContent = element.Patente;
        fila.appendChild(celdaPatente);

        var celdaPeso = document.createElement("td");
        celdaPeso.textContent = element.Peso_maximo;
        fila.appendChild(celdaPeso);

        var celdaCosto = document.createElement("td");
        celdaCosto.textContent = element.Costo_km;
        fila.appendChild(celdaCosto);

        var celdaEditar = document.createElement("td");
        var botonEditar = document.createElement("button");
        botonEditar.className = "boton-editar";
        botonEditar.innerHTML = `<i class="fa-solid fa-pen" style="color: #ffffff;"></i>`;
        celdaEditar.appendChild(botonEditar);
        fila.appendChild(celdaEditar);

        var celdaEliminar = document.createElement("td");
        var botonEliminar = document.createElement("button");
        botonEliminar.className = "boton-eliminar";
        botonEliminar.innerHTML = `<i class="fa-solid fa-trash" style="color: #ffffff;"></i>`;
        celdaEliminar.appendChild(botonEliminar);
        fila.appendChild(celdaEliminar);

        tbody.appendChild(fila);
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
}

document.addEventListener("keyup",e=>{
    if(e.target.matches("#barraBuscador")){
        document.querySelectorAll(".nombreCelda").forEach(id =>{
            id.textContent.toLowerCase().includes(e.target.value.toLowerCase())
            ? id.parentElement.classList.remove("filtro")
            : id.parentElement.classList.add("filtro");    
        });
    };
});

function eliminar(ID) {
  const id = ID;
  const url = `/camiones/${id}`;
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
        throw new Error("Error al eliminar.");
      }
      location.reload();
      console.log("eliminado con éxito.");
    })
    .catch(error => {
      console.error("Error al eliminar:", error);
    });*/
}