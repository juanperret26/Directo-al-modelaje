document.getElementById("btnDibujarGrafico").addEventListener("click", dibujarGraficoPedidos);
document.getElementById("btnDibujarGraficoEnvios").addEventListener("click", dibujarGraficoEnvios);
  
  function obtenerBeneficioEntreFechas() {
    var fechaDesde = document.getElementById("FechaDesde").value;
    var fechaHasta = document.getElementById("FechaHasta").value;
  
    var urlConFiltro = `http://localhost:8080/envios/`;
  
    //Si fechaDesde esta vacio, no se agrega al filtro
    if (fechaDesde != "") {
      urlConFiltro += `?fechaDesde=${fechaDesde}`;
    }
  
    //Si fechaHasta esta vacio, no se agrega al filtro
    if (fechaHasta != "") {
      if (fechaDesde != "") {
        urlConFiltro += `&fechaHasta=${fechaHasta}`;
      } else {
        urlConFiltro += `?fechaHasta=${fechaHasta}`;
      }
    }
  
    //urlConFiltro = `http://localhost:8080/envios/beneficioEntreFechas?fechaDesde=${fechaDesde}&fechaHasta=${fechaHasta}`;
  
    makeRequest(
      urlConFiltro,
      Method.GET,
      null,
      ContentType.JSON,
      CallType.PRIVATE,
      exitoObtenerBeneficioEntreFechas,
      errorGraficos
    );
  }
  
  function exitoObtenerBeneficioEntreFechas(data) {
    var montoFechasMes = [];
    var montoFechasAnio = [];
    var meses = [];
    var anios = [];
  
    Chart.defaults.font.size = 16;
  
    // Verificar si data.meses y data.anios existen y son arreglos
    if (!data.meses || !Array.isArray(data.meses) || data.meses.length === 0) {
        document.getElementById("mensajeSinBeneficio").innerHTML = "No hay beneficios mensuales cargados en esas fechas";
        return; // Si no hay datos en 'meses', salimos de la función
    }

    if (!data.anios || !Array.isArray(data.anios) || data.anios.length === 0) {
        document.getElementById("mensajeSinBeneficio").innerHTML = "No hay beneficios anuales cargados en esas fechas";
        return; // Si no hay datos en 'anios', salimos de la función
    }
  
    // Procesar datos por mes
    data.meses.forEach((element) => {
        montoFechasMes.push(element.Monto);
        meses.push(element.Nombre);
    });
  
    // Procesar datos por año
    data.anios.forEach((element) => {
        montoFechasAnio.push(element.Monto);
        anios.push(element.Nombre);
    });
  
    const configuracionBarras = {
        plugins: {
            legend: {
                position: 'top',
            },
            title: {
                display: true,
                text: 'Beneficio Mensual'
            }
        },
        responsive: true
    };
  
    const datosMeses = {
        labels: meses,
        datasets: [
            {
                data: montoFechasMes,
                backgroundColor: ["#FF5733", "#FFC300", "#33FF57", "#339CFF", "#FFA500"],
            },
        ],
    };
  
    // Obtener el contexto del lienzo de barras de meses
    const contextoBarrasMeses = document.getElementById('graficoBeneficioMes').getContext('2d');
  
    // Crear el gráfico de barras de meses
    const configBarrasMeses = {
        type: 'bar',
        data: datosMeses,
        options: configuracionBarras,
    };
  
    // Destroy existing chart if it exists
    if (window.myChartMeses) {
        window.myChartMeses.destroy();
    }
  
    // Create the new chart for months
    window.myChartMeses = new Chart(contextoBarrasMeses, configBarrasMeses);
  
    const datosAnios = {
        labels: anios,
        datasets: [
            {
                data: montoFechasAnio,
                backgroundColor: ["#FF5733", "#FFC300", "#33FF57", "#339CFF", "#FFA500"],
            },
        ],
    };
  
    // Obtener el contexto del lienzo de barras de años
    const contextoBarrasAnio = document.getElementById('graficoBeneficioAnio').getContext('2d');
  
    const configuracionBarrasAnio = {
        plugins: {
            legend: {
                position: 'top',
            },
            title: {
                display: true,
                text: 'Beneficio Anual'
            }
        },
        responsive: true
    };
  
    // Crear el gráfico de barras de años
    const configBarrasAnios = {
        type: 'bar',
        data: datosAnios,
        options: configuracionBarrasAnio,
    };
  
    // Destroy existing chart if it exists
    if (window.myChartAnios) {
        window.myChartAnios.destroy();
    }
  
    // Create the new chart for years
    window.myChartAnios = new Chart(contextoBarrasAnio, configBarrasAnios);
}

  
  
  function errorGraficos(status, body) {
    alert(body.error);
    console.log(body.json());
    throw new Error(status.Error);
  }
  
  function dibujarGraficoPedidos() {

    const estadoInput = document.getElementById("estado").value;

    if (!estadoInput) {
        alert("El estado no puede estar vacío.");
        return;
    }

    var urlConFiltro = `http://localhost:8080/pedidos/estado/${estadoInput}`;
  
    makeRequest(
      urlConFiltro,
      Method.GET,
      null,
      ContentType.JSON,
      CallType.PRIVATE,
      exitoObtenerGraficoPedidos,
      errorGraficos
    );
  }
  
  function exitoObtenerGraficoPedidos(data) {
    var cantidadPedidos = [];
    var estadoPedidos = [];
  
    Chart.defaults.font.size = 16;
  
    if (data.length == 0) {
      document.getElementById("mensajeSinPedidos").innerHTML = "No hay pedidos cargados";
    }
  
    for (let i = 0; i < data.length; i++) {
      const element = data[i];
      cantidadPedidos.push(element.Cantidad);
      estadoPedidos.push(element.Estado);
    }
  
    const datos = {
      labels: estadoPedidos,
      datasets: [
        {
          data: cantidadPedidos, // Cantidad de pedidos por estado
          backgroundColor: [
            "#FF5733",
            "#FFC300",
            "#33FF57",
            "#339CFF",
            "#FFA500",
          ], // Colores para cada sector del gráfico
        },
      ],
    };
  
    // Configuración del gráfico
    const config = {
      type: "pie",
      data: datos,
    };
  
    // Dibuja el gráfico en el elemento canvas 
    const ctx = document.getElementById("graficoPedidosTorta").getContext("2d");
    new Chart(ctx, config);
  
    // Configuración del gráfico de barras
    var configuracionBarras = {
      plugins: {
        legend: {
          position: 'top',
        },
        title: {
          display: true,
          text: 'Cantidad Pedidos por Estado'
        }
      },
      responsive: true,
    };
  
    // Obtener el contexto del lienzo de barras
    const contextoBarras = document.getElementById('graficoPedidosBarra').getContext('2d');
  
    // Crear el gráfico de barras
    const configBarras = {
        type: 'bar',
        data: datos,
        options: configuracionBarras
    };
  
    // Dibuja el gráfico de barras en el elemento canvas 
    new Chart(contextoBarras, configBarras);
  }
  
  function dibujarGraficoEnvios() {
    const estadoInput = document.getElementById("estadoEnvio").value;

    if (!estadoInput) {
        alert("El estado no puede estar vacío.");
        return;
    }

    var urlConFiltro = `http://localhost:8080/envios/estado/${estadoInput}`;
  
    makeRequest(
      urlConFiltro,
      Method.GET,
      null,
      ContentType.JSON,
      CallType.PRIVATE,
      exitoObtenerGraficoEnvios,
      errorGraficos
    );
  }
  
  function exitoObtenerGraficoEnvios(data) {
    console.log("Respuesta de la API:", data);  // Inspeccionamos la respuesta completa

    var cantidadEnvios = [];
    var estadoEnvios = [];

    // Verificar si 'cantidad' está presente en la respuesta
    if (data && 'cantidad' in data) {
        console.log("Cantidad de envíos:", data.cantidad);  // Verificar la cantidad
        cantidadEnvios.push(data.cantidad);  // Acceder a 'cantidad' de la respuesta
        estadoEnvios.push('A despachar'); // Puedes usar el estado que te corresponda, ya que solo tienes la cantidad
    } else {
        document.getElementById("mensajeSinEnvios").innerHTML = "No hay envíos cargados";
        return;
    }

    // Si 'cantidad' es cero, mostrar el mensaje correspondiente
    if (data.cantidad === 0) {
        document.getElementById("mensajeSinEnvios").innerHTML = "No hay envíos disponibles para mostrar";
        return;
    }

    // Configuración del gráfico
    const datos = {
        labels: estadoEnvios,  // Estado de los envíos
        datasets: [
            {
                data: cantidadEnvios,  // Cantidad de envíos
                backgroundColor: [
                    "#FF5733", "#FFC300", "#33FF57", "#339CFF", "#FFA500",
                ],  // Colores para cada sector del gráfico
            },
        ],
    };

    const config = {
        type: "pie",  // Tipo de gráfico: torta
        data: datos,
    };

    // Dibuja el gráfico de torta en el elemento canvas
    const ctx = document.getElementById("graficoEnviosTorta").getContext("2d");
    new Chart(ctx, config);

    // Configuración del gráfico de barras
    var configuracionBarras = {
        plugins: {
            legend: {
                position: 'top',
            },
            title: {
                display: true,
                text: 'Cantidad Envíos por Estado',
            },
        },
        responsive: true,
    };

    const configBarras = {
        type: 'bar',  // Tipo de gráfico: barras
        data: datos,
        options: configuracionBarras,
    };

    // Dibuja el gráfico de barras en el elemento canvas
    const ctxBarras = document.getElementById("graficoEnviosBarra").getContext("2d");
    new Chart(ctxBarras, configBarras);
}


  