
document.getElementById("aceptar").addEventListener("click", function() {
    crear();
});
function crear() {
    const nombre = document.querySelector('input[placeholder="Nombre"]').value;
    const tipo = document.querySelector('input[placeholder="Tipo"]').value;
    const peso = parseFloat(document.querySelector('input[placeholder="Peso unit."]').value);
    const precio = parseFloat(document.querySelector('input[placeholder="Precio"]').value);
    const stock = parseFloat(document.querySelector('input[placeholder="Stock"]').value);
    const stockMinimo = parseInt(document.querySelector('input[placeholder="Stock minimo"]').value);

    const nuevoProducto = {
        id: " ",
        nombre: nombre,
        TipoProducto: tipo,
        peso_unitario: peso,
        precio: precio,
        stock: stock,
        stockMinimo: stockMinimo,
        Tipo: " "
    };

    fetch("/productos", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(nuevoProducto)
    })
    .then(response => {
        if (response.ok) {
            // La solicitud se realizó con éxito, puedes redirigir a la página de productos u realizar alguna otra acción
            window.location.href = "/htmlproductos";
        } else {
            // Manejar errores
            response.json().then(data => {
                console.log(data);});
            console.error("Error al crear un nuevo producto");
        }
    })
    .catch(error => {
        console.error("Error al crear un nuevo producto:", error);
    });
}

document.addEventListener("DOMContentLoaded", function () {
    const boton_crear = document.getElementById("aceptar");
    boton_crear.style.display = "block";
    const boton_editar = document.getElementById("editar");
    boton_editar.style.display = "none";

    const queryString = window.location.search;
    const urlParams = new URLSearchParams(queryString);
  
    var productoID = "";

    if (urlParams.has("productoID")) {
        productoID = urlParams.get("productoID");
    }

    function cargarValorEnCampo(placeholder, paramName) {
      if (urlParams.has(paramName)) {

        const boton_crear = document.getElementById("aceptar");
        boton_crear.style.display = "none";
        const boton_editar = document.getElementById("editar");
        boton_editar.style.display = "block";

        const input = document.querySelector(`input[placeholder="${placeholder}"]`);
        input.value = urlParams.get(paramName);
      }
    }
  
    cargarValorEnCampo("Nombre", "Nombre");
    cargarValorEnCampo("Tipo", "Tipo");
    cargarValorEnCampo("Peso unit.", "Peso unit.");
    cargarValorEnCampo("Precio", "Precio");
    cargarValorEnCampo("Stock", "Stock");
    cargarValorEnCampo("Stock minimo", "StockMinimo");
    cargarValorEnCampo("Tipo", "Tipo");

    document.getElementById("editar").addEventListener("click", function() {
        editar(productoID); // Pasar productoID a la función editar
    });
});


//funcion para el boton editar
//document.getElementById("editar").addEventListener("click", function() {
//    editar();
//});
function editar(productoID) {

    const codigo = productoID;
    const nombre = document.querySelector('input[placeholder="Nombre"]').value;
    const tipo = document.querySelector('input[placeholder="Tipo"]').value;
    const peso = parseFloat(document.querySelector('input[placeholder="Peso unit."]').value);
    const precio = parseFloat(document.querySelector('input[placeholder="Precio"]').value);
    const stock = parseFloat(document.querySelector('input[placeholder="Stock"]').value);

    const objetoEditado = {
        Id: codigo,
        nombre: nombre,
        TipoProducto: tipo,
        peso_unitario: peso,
        precio: precio,
        stock: stock
    };

    fetch("/productos", {
        method: "PUT",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(objetoEditado)
    })
    .then(response => {
        if (response.ok) {
            // La solicitud se realizó con éxito, puedes redirigir a la página de productos u realizar alguna otra acción
            window.location.href = "/htmlproductos";
        } else {
            // Manejar errores
            response.json().then(data => {
                console.log(data);});
            console.error("Error al crear un nuevo producto");
        }
    })
    .catch(error => {
        console.error("Error al crear un nuevo producto:", error);
    });
}
  