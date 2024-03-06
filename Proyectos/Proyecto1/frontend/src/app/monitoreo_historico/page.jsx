'use client'

import React, { useState } from 'react'
import Layout from '@/components/Layout/page';
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
} from 'chart.js';
import { Line } from 'react-chartjs-2';
import 'chart.js/auto';
import './style.css'

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend
);

export default function MonitoreoHistorico() {

  const [dataRAM, setDataRAM] = useState([]);
  const [labelsRAM, setLabelRAM] = useState([]);
  const [dataCPU, setDataCPU] = useState([]);
  const [labelsCPU, setLabelCPU] = useState([]);

  const fetchData = async () => {
    try {
      const response = await fetch('http://localhost:8080/monitoreo-historico');
      const data = await response.json();
      setDataRAM(data.ram.data)
      setLabelRAM(data.ram.labels)
      setDataCPU(data.cpu.data)
      setLabelCPU(data.cpu.labels)
    } catch (error) {
      console.error('Error al obtener datos de la API:', error);
    }
  };

  return (<>
    <Layout
      pageTitle=""
    >
      <div className="min-h-screen flex flex-col justify-center items-center">
        <div className="m-auto text-center">
          <h1 className="main-title">MONITOREO HISTORICO</h1>
          <button onClick={fetchData} style={{
            padding: '10px',
            backgroundColor: '#4CAF50',
            color: 'white',
            border: 'none',
            borderRadius: '4px',
            cursor: 'pointer',
          }}>Actualizar</button>
          <br />
          <div className="charts-container">
            <div className="chart" style={{ width: '600px', height: '400px' }}>
              <h1>Memoria RAM</h1>
              <Line data={{
                labels: labelsRAM,
                datasets: [
                  {
                    label: '% de Utilizacion',
                    data: dataRAM,
                    borderColor: 'rgb(255, 99, 132)',
                    backgroundColor: 'rgba(255, 99, 132, 0.5)',
                  }
                ],
              }} />;
            </div>
            <div className="chart" style={{ width: '600px', height: '400px' }}>
              <h1>CPU</h1>
              <Line data={{
                labels: labelsCPU,
                datasets: [
                  {
                    label: '% de Utilizacion',
                    data: dataCPU,
                    borderColor: 'rgb(255, 99, 132)',
                    backgroundColor: 'rgba(255, 99, 132, 0.5)',
                  }
                ],
              }} />;
            </div>
          </div>
        </div>
      </div>
    </Layout>
  </>);
}