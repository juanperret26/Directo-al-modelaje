document.addEventListener("DOMContentLoaded",function(){
    cargarDatos();  
  });    
  
  function cargarDatos(){
      fetch("/pedidos", { method: "GET" })
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
      });
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
        if (element.productos && Array.isArray(element.productos)) {
                var nombresProductos = element.productos.map(function(producto) {
                return producto.Nombre;
        });
            celdaProductos.textContent = nombresProductos.join(", ");
        } else {
            celdaProductos.textContent = "N/A"; // O cualquier otro mensaje que desees mostrar
        }
        fila.appendChild(celdaProductos);

        var celdaCiudad = document.createElement("td");
        celdaCiudad.textContent = element.destino;
        fila.appendChild(celdaCiudad);

        var celdaEstado = document.createElement("td");
        celdaEstado.textContent = element.estado;
        fila.appendChild(celdaEstado);

        var celdaAceptar = document.createElement("td");
        var botonAceptar = document.createElement("button");
        botonAceptar.className = "boton-aceptar";
        botonAceptar.innerHTML = `<i class="fa-solid fa-check" style="color: #ffffff;"></i>`;
        celdaAceptar.appendChild(botonAceptar);
        fila.appendChild(celdaAceptar);

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
    //eliminar
    tbody.addEventListener("click", function (event) {
        if (event.target.classList.contains("boton-eliminar")) {
            const botonEditar = event.target;
            const fila = botonEditar.closest("tr");
            const primeraCelda = fila.querySelector("td:first-child");
            const textoPrimeraCelda = primeraCelda.textContent;
            eliminar(textoPrimeraCelda);
        }
    });
    tbody.addEventListener("click", function (event) {
      if (event.target.classList.contains("boton-aceptar")) {
          const botonEditar = event.target;
          const fila = botonEditar.closest("tr");
          const primeraCelda = fila.querySelector("td:first-child");
          const textoPrimeraCelda = primeraCelda.textContent;
          aceptar(textoPrimeraCelda);
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

};
document.addEventListener("keyup",e=>{
    if(e.target.matches("#barraBuscador")){
        document.querySelectorAll(".nombreCelda").forEach(nombre =>{
            nombre.textContent.toLowerCase().includes(e.target.value.toLowerCase())
            ? nombre.parentElement.classList.remove("filtro")
            : nombre.parentElement.classList.add("filtro");    
        });
    };
});

//eliminar
function eliminar(ID) {
    const id = ID;
    const url = `/pedidos/${id}`;
    fetch(url, {
      method: "DELETE"
    })
      .then(response => {
        if (!response.ok) {
          throw new Error("Error al eliminar el pedido.");
        }
        location.reload();
        console.log("Pedido eliminado con éxito.");
      })
      .catch(error => {
        console.error("Error al eliminar el pedido:", error);
      });
};

function aceptar(ID) {
  /*const id = ID;
  const url = `/pedidos/${id}`;
  const datos = {};
  makeRequest(
      url,
      Method.PUT, 
      datos,
      ContentType.JSON,
      CallType.PRIVATE,
      exitoSolicitud,
      errorSolicitud
  );
*/

const id = ID;
const url = `/pedidos/${id}`;
fetch(url, {
  method: "PUT",
  headers: {
    'Content-Type': 'application/json'
  },
  body: JSON.stringify({}),

})
  .then(response => {
    if (!response.ok) {
      throw new Error("Error al aceptar el pedido.");
    }
    location.reload();
    console.log("Pedido aceptado con éxito.");
  })
  .catch(error => {
    console.error("Error al aceptar el pedido:", error);
  });

  function exitoSolicitud(data) {
      console.log("éxito.");
      location.reload();
      // Realiza otras acciones si es necesario
  }

  function errorSolicitud(status, response) {
      console.error("Error . Estado:", status, "Respuesta:", response);
      // Maneja el error de acuerdo a tus necesidades
  }
}


//Boton editar
function manejarEdicion(fila) {
  // Obtener elementos de la fila
  const celdas = fila.querySelectorAll("td");

  // Crear un objeto para almacenar los valores de las celdas
  const valoresCeldas = {};
  var pedidoID = "";
  const encabezado = document.querySelector("table thead tr");

  celdas.forEach((celda, index) => {
    if (index >= 0 && index <= 7) {
      if(index == 0){
        var tituloCelda = encabezado.querySelectorAll("th")[index].textContent;
        valoresCeldas[tituloCelda] = celda.textContent.trim();    
        pedidoID = valoresCeldas[tituloCelda].toString();
      }
      else{
      var tituloCelda = encabezado.querySelectorAll("th")[index].textContent;
      valoresCeldas[tituloCelda] = celda.textContent.trim();
      }
    }
  });

  // Redirigir a la página "formPedidos" y enviar los valores como parámetros
  const queryString = new URLSearchParams(valoresCeldas).toString();
  window.location.href = `/formPedidos?pedidoID=${pedidoID}&${queryString}`;
}
