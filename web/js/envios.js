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

      /*fetch("/envios", { method: "GET" })
      .then(response => {
        if (!response.ok) {
          throw new Error("Error al obtener datos de envios.");
        }
        return response.json();
      })
      .then(data => {
        mostrarDatosTabla(data);
      })
      .catch(error => {
        console.error("Error al obtener datos de envios:", error);
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

        var celdaEstado = document.createElement("td");
        celdaEstado.textContent = element.estado;
        fila.appendChild(celdaEstado);

        var celdaParadas = document.createElement("td");
        celdaParadas.textContent = element.paradas;
        fila.appendChild(celdaParadas);

        var celdaDestino = document.createElement("td");
        celdaDestino.textContent = element.destino;
        fila.appendChild(celdaDestino);

        var celdaCreacion = document.createElement("td");
        celdaCreacion.textContent = element.creacion;
        fila.appendChild(celdaCreacion);

        var celdaPedido = document.createElement("td");
        celdaPedido.textContent = element.pedido;
        fila.appendChild(celdaPedido);

        var celdaActualizacion = document.createElement("td");
        celdaActualizacion.textContent = element.actualizacion;
        fila.appendChild(celdaActualizacion);

        var celdaCosto = document.createElement("td");
        celdaCosto.textContent = element.costo;
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