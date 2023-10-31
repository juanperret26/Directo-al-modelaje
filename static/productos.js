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
    /*fetch("/productos",{method: "GET"})
    .then(response =>response.json())
    .then(data =>
        {mostrarDatosTabla(data);
    })
    .catch(error => {
        console.error("Error al obtener datos de pedidos:", error);
        console.log(response.textContent);
    });*/
};

function mostrarDatosTabla(datos){
    var table = document.getElementById("TablaPrincipal");
    var tbody = document.getElementById("TableBody");

    datos.forEach(function(element){
        var fila = document.createElement("tr");
        
        var celdaId = document.createElement("td");
        celdaId.textContent = element.CodigoProducto;
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

function obtenerFechaDesdeCadena(cadenaFechaHora) {
    const partes = cadenaFechaHora.split("T");
    const fecha = partes[0];
    
    return fecha;
}