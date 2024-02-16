import React, { useState, useEffect } from 'react';
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';
import { Pie } from 'react-chartjs-2';
import { Greet } from "../wailsjs/go/main/App";

ChartJS.register(ArcElement, Tooltip, Legend);

export function App() {

    const [ramTotal, setRamTotal] = useState(0);
    const [ramEnUso, setRamEnUso] = useState(0);
    const [ramLibre, setRamLibre] = useState(0);
    const [chartData, setChartData] = useState({
        labels: ['En Uso', 'Libre'],
        datasets: [
            {
                label: '%',
                data: [2, 80],
                backgroundColor: [
                    'rgba(255, 99, 132, 0.2)', // Rojo
                    'rgba(75, 192, 192, 0.2)' // Verde
                ],
                borderColor: [
                    'rgba(255, 99, 132, 1)', // Rojo
                    'rgba(75, 192, 192, 1)', // Verde
                ],
                borderWidth: 1,
            },
        ],
    });

    function greet() {
        Greet('').then((result) => {
            const resultado = JSON.parse(result);
            setRamTotal(resultado.totalRam);
            setRamEnUso(resultado.memoriaEnUso);
            setRamLibre(resultado.libre);
            setChartData((prevChartData) => {
                const newChartData = { ...prevChartData };
                newChartData.datasets[0].data = [resultado.porcentaje, 100 - resultado.porcentaje];
                return newChartData;
            });
        });
    }

    useEffect(() => {
        greet();
        const intervalo = setInterval(greet, 500);
        return () => clearInterval(intervalo);
    }, []);


    return (
        <div style={{ width: '400px', margin: '0 auto', textAlign: 'center' }}>
            <p><b>Nombre: </b> Douglas Alexander Soch Catalán</p>
            <p><b>Carnet: </b> 201807032</p>
            <h2>Parametros</h2>
            <b>RAM Total:</b> {ramTotal}
            <br />
            <b>RAM En Uso:</b> {ramEnUso}
            <br />
            <b>RAM Libre:</b> {ramLibre}
            <br /><br /><br />
            <Pie key={JSON.stringify(chartData.datasets[0].data)} data={chartData} />
        </div>
    );
}
