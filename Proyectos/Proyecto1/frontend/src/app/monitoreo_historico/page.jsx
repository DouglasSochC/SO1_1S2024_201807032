'use client'

import React, { useState, useEffect } from 'react'
import Layout from '@/components/Layout/page';
import { Pie } from 'react-chartjs-2';
import 'chart.js/auto';
import './style.css'

export default function MonitoreoHistorico() {

  const [data, setData] = useState([0, 0]);

  useEffect(() => {
    const actualizarEstadoCadaSegundo = () => {
      // Generar un nuevo valor aleatorio entre 1 y 100
      const nuevoValor = Math.floor(Math.random() * 100) + 1;

      // Actualizar el estado con el nuevo valor
      setData([100 - nuevoValor, nuevoValor]);
    };

    const intervalId = setInterval(actualizarEstadoCadaSegundo, 1000);

    return () => clearInterval(intervalId);
  }, []);

  return (<>
    <Layout
      pageTitle=""
    >
      <div className="min-h-screen flex flex-col justify-center items-center">
        <div className="m-auto text-center">
          <h1 className="main-title">MONITOREO HISTORICO</h1>
          <div className="charts-container">
            <div className="chart">
              <h1>Memoria RAM</h1>

              <Pie data={{
                labels: ['En Uso', 'Libre'],
                datasets: [
                  {
                    data: data,
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
                    data: data,
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