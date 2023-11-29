
// Ejemplo de uso
const url = "http://localhost:8080/aulas";

document.addEventListener("DOMContentLoaded", function(eventDOM){  
    if (!isUserLogged()){
      window.location = 'login.html?reason=login_required';
    }

    obtenerAulas();
    
    document.getElementById("btnFiltrar").addEventListener("click",function(eventClick){
        eventClick.preventDefault();
        obtenerAulas();        
    })
})

function obtenerAulas(){
  const nombre = document.getElementById("FiltroNombre").value;
  
  makeRequest(`${url}?nombre=${nombre}`, Method.GET, null, ContentType.JSON, CallType.PRIVATE, successFn, errorFn);
}

function successFn(data) {
    // Llenar la tabla con los datos obtenidos
    data.forEach(elemento => {
      const row = document.createElement("tr"); //crear una fila


      row.innerHTML = ` 
          <td>${elemento.ID}</td>
          <td>${elemento.Nombre}</td>
          <td class="acciones"><a href="form.html?id=${elemento.ID}&tipo=EDITAR">Editar</a> | <a href="form.html?id=${elemento.ID}&tipo=ELIMINAR">Eliminar</a></td>
      `; //crear una celda por cada campo que quiera mostrar

      elementosTable.appendChild(row);
  });
  }
  
  function errorFn(status, response) {
    console.log("Falla:", response);
  }