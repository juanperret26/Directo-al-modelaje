document.addEventListener("DOMContentLoaded",function(){
    let arreglo = new Array;
    arreglo[0] = {id: "1321", patente: "324UEQ", peso: "500", costo: "1000"};
    arreglo[1] = {id: "5657", patente: "574HJK", peso: "500", costo: "1000"};
    arreglo[2] = {id: "6432", patente: "928BNM", peso: "500", costo: "1000"};
    arreglo[3] = {id: "8756", patente: "431DFG", peso: "500", costo: "1000"};
    arreglo[4] = {id: "0912", patente: "390MNF", peso: "500", costo: "1000"};

    var table = document.getElementById("TablaPrincipal");
    var tbody = document.getElementById("TableBody");

    arreglo.forEach(function(element){
        var fila = document.createElement("tr");
        
        var celdaId = document.createElement("td");
        celdaId.textContent = element.id;
        celdaId.className = "nombreCelda";
        fila.appendChild(celdaId);

        var celdaPatente = document.createElement("td");
        celdaPatente.textContent = element.patente;
        fila.appendChild(celdaPatente);

        var celdaPeso = document.createElement("td");
        celdaPeso.textContent = element.peso;
        fila.appendChild(celdaPeso);

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