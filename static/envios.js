document.addEventListener("DOMContentLoaded",function(){
    let arreglo = new Array;
    arreglo[0] = {id: "1321", estado: "Pendiente", paradas: "Rafaela, Santa Fe", destino: "Rosario", creacion: "24/10/2023", pedido: "Heladera, Silla, Mesa, Horno", actualizacion:"25/10/2023", costo: "234532"};
    arreglo[1] = {id: "5657", estado: "Pendiente", paradas: "Rafaela, Santa Fe", destino: "Rosario", creacion: "24/10/2023", pedido: "Heladera, Silla, Mesa, Horno", actualizacion:"25/10/2023", costo: "234532"};
    arreglo[2] = {id: "6432", estado: "Pendiente", paradas: "Rafaela, Santa Fe", destino: "Rosario", creacion: "24/10/2023", pedido: "Heladera, Silla, Mesa, Horno", actualizacion:"25/10/2023", costo: "234532"};
    arreglo[3] = {id: "8756", estado: "Pendiente", paradas: "Rafaela, Santa Fe", destino: "Rosario", creacion: "24/10/2023", pedido: "Heladera, Silla, Mesa, Horno", actualizacion:"25/10/2023", costo: "234532"};
    arreglo[4] = {id: "0912", estado: "Pendiente", paradas: "Rafaela, Santa Fe", destino: "Rosario", creacion: "24/10/2023", pedido: "Heladera, Silla, Mesa, Horno", actualizacion:"25/10/2023", costo: "234532"};

    var table = document.getElementById("TablaPrincipal");
    var tbody = document.getElementById("TableBody");

    arreglo.forEach(function(element){
        var fila = document.createElement("tr");
        
        var celdaId = document.createElement("td");
        celdaId.textContent = element.id;
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
})
document.addEventListener("keyup",e=>{
    if(e.target.matches("#barraBuscador")){
        document.querySelectorAll(".nombreCelda").forEach(id =>{
            id.textContent.toLowerCase().includes(e.target.value.toLowerCase())
            ? id.parentElement.classList.remove("filtro")
            : id.parentElement.classList.add("filtro");    
        });
    };
});