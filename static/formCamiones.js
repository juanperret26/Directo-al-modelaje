
document.getElementById("aceptar").addEventListener("click", function() {
    crearNuevoCamion();
});
function crearNuevoCamion() {
    const patente = document.querySelector('input[placeholder="Patente"]').value;
    const pesoMax = parseFloat(document.querySelector('input[placeholder="Peso maximo"]').value);
    const costoKM = parseInt(document.querySelector('input[placeholder="Costo por Km"]').value);
    const nuevoCamion = {
        patente: patente,
        pesoMax: pesoMax,
        costoKM: costoKM
    };

    fetch("/camiones", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(nuevoCamion)
    })
    .then(response => {
        if (response.ok) {
            window.location.href = "/htmlcamiones";
        } else {
            // Manejar errores
            response.json().then(data => {
                console.log(data);});
            console.error("Error al crear un nuevo camion");
        }
    })
    .catch(error => {
        console.error("Error al crear un nuevo camion:", error);
    });
}