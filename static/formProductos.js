
document.getElementById("aceptar").addEventListener("click", function() {
    crearNuevoCamion();
});
function crearNuevoCamion() {
    const nombre = document.querySelector('input[placeholder="Nombre"]').value;
    const tipo = document.querySelector('input[placeholder="Tipo"]').value;
    const peso = parseFloat(document.querySelector('input[placeholder="Peso unitario"]').value);
    const precio = parseFloat(document.querySelector('input[placeholder="Precio"]').value);
    const stock = parseFloat(document.querySelector('input[placeholder="Stock"]').value);
    const stockMinimo = parseInt(document.querySelector('input[placeholder="Stock minimo"]').value);

    const nuevoProducto = {
        id: " ",
        nombre: nombre,
        TipoProducto: tipo,
        peso: peso,
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