document.addEventListener("DOMContentLoaded",function(){
    let arreglo = new Array;
    arreglo[0] = {id: "1321", productos: "Bicicleta, Inflador, Rueda, Frenos...", ciudad: "Rafaela", estado: "Pendiente"};
    arreglo[1] = {id: "5657", productos: "Bicicleta, Inflador, Rueda, Frenos...", ciudad: "Rafaela", estado: "Pendiente"};
    arreglo[2] = {id: "6432", productos: "Bicicleta, Inflador, Rueda, Frenos...", ciudad: "Rafaela", estado: "Pendiente"};
    arreglo[3] = {id: "8756", productos: "Bicicleta, Inflador, Rueda, Frenos...", ciudad: "Rafaela", estado: "Pendiente"};
    arreglo[4] = {id: "0912", productos: "Bicicleta, Inflador, Rueda, Frenos...", ciudad: "Rafaela", estado: "Pendiente"};

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
})
document.addEventListener("keyup",e=>{
    if(e.target.matches("#barraBuscador")){
        document.querySelectorAll(".nombreCelda").forEach(nombre =>{
            nombre.textContent.toLowerCase().includes(e.target.value.toLowerCase())
            ? nombre.parentElement.classList.remove("filtro")
            : nombre.parentElement.classList.add("filtro");    
        });
    };
});