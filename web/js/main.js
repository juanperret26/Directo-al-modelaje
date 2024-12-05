document.addEventListener("DOMContentLoaded",function(){
    let arreglo = new Array;
    arreglo[0] = {id: "1321", productos: "Rafaela", ciudad: "4 de enero", estado: "pendiente"};
    arreglo[1] = {id: "5657", productos: "Rafaela", ciudad: "4 de enero", estado: "pendiente"};
    arreglo[2] = {id: "6432", productos: "Rafaela", ciudad: "4 de enero", estado: "pendiente"};
    arreglo[3] = {id: "8756", productos: "Rafaela", ciudad: "4 de enero", estado: "pendiente"};
    arreglo[4] = {id: "0912", productos: "Rafaela", ciudad: "4 de enero", estado: "pendiente"};

    var table = document.getElementById("TablaPrincipal");
    var tbody = document.getElementById("TableBody");

    arreglo.forEach(function(element){
        var fila = document.createElement("tr");
        
        var celdaId = document.createElement("td");
        celdaId.textContent = element.id;
        celdaId.className = "nombreCelda";
        fila.appendChild(celdaId);

        var celdaProductos = document.createElement("td");
        celdaProductos.textContent = element.productos;
        fila.appendChild(celdaProductos);

        var celdaCiudad = document.createElement("td");
        celdaCiudad.textContent = element.ciudad;
        fila.appendChild(celdaCiudad);

        var celdaEstado = document.createElement("td");
        celdaEstado.textContent = element.estado;
        fila.appendChild(celdaEstado);

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