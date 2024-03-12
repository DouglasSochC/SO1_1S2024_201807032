'use client'

import React, { useState, useEffect } from 'react'
import Layout from '@/components/Layout/page';
import { DataSet, Network } from 'vis-network/standalone';
import './style.css'

export default function DiagramaEstados() {

  // Definir estados para nodos y bordes
  const [nodes, setNodes] = useState(new DataSet([
    { id: 1, label: 'Node 1' },
    { id: 2, label: 'Node 2' },
    { id: 3, label: 'Node 3' },
    { id: 4, label: 'Node 4' },
    { id: 5, label: 'Node 5' }
  ]));
  const [edges, setEdges] = useState(new DataSet([
    { from: 1, to: 3 },
    { from: 1, to: 2 },
    { from: 2, to: 4 },
    { from: 2, to: 5 }
  ]));

  const fetchNuevoProceso = async () => {
    try {
      console.log("Nuevo Proceso");
    } catch (error) {
      console.error('Error al obtener datos de la API:', error);
    }
  };

  const fetchPararProceso = async () => {
    try {
      console.log("Proceso en Pausa");
    } catch (error) {
      console.error('Error al obtener datos de la API:', error);
    }
  };

  const fetchProcesoListo = async () => {
    try {
      console.log("Proceso Listo");
    } catch (error) {
      console.error('Error al obtener datos de la API:', error);
    }
  };

  const fetchMatarProceso = async () => {
    try {
      console.log("Matar Proceso");
    } catch (error) {
      console.error('Error al obtener datos de la API:', error);
    }
  };

  useEffect(() => {
    const container = document.getElementById('mynetwork');
    const data = {
      nodes: nodes,
      edges: edges
    };
    const options = {};
    const network = new Network(container, data, options);
    return () => {
      network.destroy();
    };
  }, [nodes, edges]);

  return (<>
    <Layout
      pageTitle=""
    >
      <div className="min-h-screen flex flex-col justify-center items-center">
        <div className="m-auto text-center">
          <h1 className="main-title">DIAGRAMA DE ESTADOS</h1>
          <div className="button-container">
            <button onClick={fetchNuevoProceso} style={{
              padding: '15px',
              backgroundColor: '#4CAF50',
              color: 'white',
              border: 'none',
              borderRadius: '4px',
              cursor: 'pointer',
            }}>New</button>
            <button onClick={fetchPararProceso} style={{
              padding: '15px',
              backgroundColor: '#ACAE30',
              color: 'white',
              border: 'none',
              borderRadius: '4px',
              cursor: 'pointer',
            }}>Stop</button>
            <button onClick={fetchProcesoListo} style={{
              padding: '15px',
              backgroundColor: '#356BDE',
              color: 'white',
              border: 'none',
              borderRadius: '4px',
              cursor: 'pointer',
            }}>Ready</button>
            <button onClick={fetchMatarProceso} style={{
              padding: '15px',
              backgroundColor: '#DE3535',
              color: 'white',
              border: 'none',
              borderRadius: '4px',
              cursor: 'pointer',
            }}>Kill</button>
          </div>
          <br />
          <div className="charts-container">
            <div style={{ width: '1000px', height: '500px' }}>
              <div id="mynetwork" style={{ width: '1000px', height: '500px', border: '1px solid lightgray' }}></div>
            </div>
          </div>
        </div>
      </div>
    </Layout>
  </>);
}