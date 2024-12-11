document.addEventListener('DOMContentLoaded', function() {
    // Asegurarse de que el código se ejecute después de que el DOM esté completamente cargado
    document.getElementById("aceptar").addEventListener("click", function() {
        crear();
    });

    function crear() {
        // Crear el objeto URLSearchParams para acceder a los parámetros de la URL
        const urlParams = new URLSearchParams(window.location.search);
        
        // Acceder a los valores de los inputs
        const ciudadInput = document.querySelector('input[placeholder="Ciudad"]');
        let kmInput = document.querySelector('input[placeholder="Km"]');

        // Verificar si los inputs existen
        if (!ciudadInput || !kmInput) {
            console.error('No se encuentran los inputs con los placeholders correctos');
            return; // Detener la ejecución si no se encuentran los elementos
        }

        const Ciudad = ciudadInput.value;
        // Convertir el valor de "Km" a entero y verificar si es válido
        kmInput = parseInt(kmInput.value, 10);

        // Verificar si los campos están vacíos o si "Km" no es un número válido
        if (!Ciudad || isNaN(kmInput) || kmInput <= 0) {
            console.error('Por favor, completa todos los campos correctamente. Asegúrate de que "Km" sea un número válido.');
            // Mostrar un mensaje de advertencia en la UI
            alert('Por favor, completa todos los campos correctamente. Asegúrate de que "Km" sea un número válido.');
            return; // Detener la ejecución si los campos están vacíos o "Km" no es un número válido
        }

        let productoID;
        
        // Verificar si el parámetro "productoID" existe en la URL
        if (urlParams.has("paradaID")) {
            productoID = urlParams.get("paradaID");
        }

        const nuevaParada = {
            ciudad: Ciudad,
            kilometros: kmInput // Ahora "kmInput" es un número entero
        };

        const url = `http://localhost:8080/envios/${productoID}/parada`;

        makeRequest(
            url,
            Method.PUT, 
            nuevaParada,
            ContentType.JSON,
            CallType.PRIVATE,
            exitoSolicitud,
            errorSolicitud
        );
        
        function exitoSolicitud(data) {
            console.log("éxito.");
            window.location.href = "/html/envios.html"; // Redirige después de éxito
        }
      
        function errorSolicitud(status, response) {
            console.error("Error al crear producto. Estado:", status, "Respuesta:", response);
            // Mostrar el error de manera clara
            alert(`Error al crear producto. Estado: ${status}. Respuesta: ${JSON.stringify(response)}`);
        }
    }
});
