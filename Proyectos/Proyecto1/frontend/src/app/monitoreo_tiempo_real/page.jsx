'use client'

import React, { useState, useEffect } from 'react'
import Layout from '@/components/Layout/page';
import { Pie } from 'react-chartjs-2';
import 'chart.js/auto';
import './style.css'

export default function MonitoreoTiempoReal() {

  const [dataRAM, setDataRAM] = useState([0, 0]);
  const [dataCPU, setDataCPU] = useState([0, 0]);

  useEffect(() => {
    const actualizarEstadoCadaSegundo = async () => {
      try {
        // Se obtienen los datos para la RAM
        const responseRAM = await fetch('http://localhost:8080/monitoreo-tiempo-real');
        if (responseRAM.ok) {
          const jsonData = await responseRAM.json();
          setDataRAM([jsonData.ram.memoria_porcentaje_uso, 100 - jsonData.ram.memoria_porcentaje_uso]);
          setDataCPU([jsonData.cpu.cpu_porcentaje, 100 - jsonData.cpu.cpu_porcentaje]);
        } else {
          console.error('Error al obtener datos del endpoint');
        }
      } catch (error) {
        console.error('Error de red:', error);
      }
    };

    const intervalId = setInterval(actualizarEstadoCadaSegundo, 500);

    return () => clearInterval(intervalId);
  }, []);

  return (<>
    <Layout
      pageTitle=""
    >
      <div className="min-h-screen flex flex-col justify-center items-center">
        <div className="m-auto text-center">
          <h1 className="main-title">MONITOREO EN TIEMPO REAL</h1>
          <div className="charts-container">
            <div className="chart">
              <h1>Memoria RAM</h1>

              <Pie data={{
                labels: ['En Uso', 'Libre'],
                datasets: [
                  {
                    data: dataRAM,
                    backgroundColor: [
                      'rgba(255, 99, 132, 0.6)',
                      'rgba(54, 162, 235, 0.6)',
                    ],
                    borderColor: [
                      'rgba(255, 99, 132, 1)',
                      'rgba(54, 162, 235, 1)',
                    ],
                    borderWidth: 1,
                  },
                ],
              }} />
            </div>
            <div className="chart">
              <h1>CPU</h1>
              <Pie data={{
                labels: ['En Uso', 'Libre'],
                datasets: [
                  {
                    data: dataCPU,
                    backgroundColor: [
                      'rgba(255, 99, 132, 0.6)',
                      'rgba(54, 162, 235, 0.6)',
                    ],
                    borderColor: [
                      'rgba(255, 99, 132, 1)',
                      'rgba(54, 162, 235, 1)',
                    ],
                    borderWidth: 1,
                  },
                ],
              }} />
            </div>
          </div>
        </div>
      </div>
    </Layout>
  </>);
}