
document.getElementById("aceptar").addEventListener("click", function() {
    crearNuevoProducto();
});
function crearNuevoProducto() {
    const nombre = document.querySelector('input[placeholder="Nombre"]').value;
    const tipo = document.querySelector('input[placeholder="Tipo"]').value;
    const peso = document.querySelector('input[placeholder="Peso unitario"]').value;
    const precio = document.querySelector('input[placeholder="Precio"]').value;
    const stock = document.querySelector('input[placeholder="Stock"]').value;
    const stockMinimo = document.querySelector('input[placeholder="Stock minimo"]').value;

    const nuevoProducto = {
        id: "",
        nombre: nombre,
        TipoProducto: tipo,
        peso: peso,
        precio: precio,
        stock: stock,
        stockMinimo: stockMinimo,
        Tipo: "",
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
            window.location.href = "/productos";
        } else {
            // Manejar errores
            console.error("Error al crear un nuevo producto");
        }
    })
    .catch(error => {
        console.error("Error al crear un nuevo producto:", error);
    });
}