'use client'

import React, { useState, useEffect } from 'react'
import Layout from '@/components/Layout/page';
import { DataSet, Network } from 'vis-network/standalone';
import './style.css'
import Swal from 'sweetalert2';

const showToast = (icon, title) => {
  const Toast = Swal.mixin({
    toast: true,
    position: "top-end",
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: (toast) => {
      toast.onmouseenter = Swal.stopTimer;
      toast.onmouseleave = Swal.resumeTimer;
    }
  });

  Toast.fire({
    icon: icon,
    title: title
  });
};

export default function DiagramaEstados() {

  const [pid, setPID] = useState(-1);
  const [idActual, setIDActual] = useState(-1);
  const [nodes, setNodes] = useState([]);
  const [edges, setEdges] = useState([]);

  const fetchNuevoProceso = async () => {
    try {

      if (pid != -1) {
        showToast("error", "Existe un proceso activo");
      } else {

        const response = await fetch('http://localhost:8080/crear-proceso');
        const data = await response.json();
        setNodes([]);
        setEdges([]);

        const nodoNew = { id: 1, label: 'New', color: '#26BBD2' };
        setNodes(prevNodos => [...prevNodos, nodoNew]);
        const nodoRunning = { id: 3, label: 'Running', color: '#26D243' };
        setNodes(prevNodos => [...prevNodos, nodoRunning]);
        const newEdge = { from: 1, to: 3, arrows: 'to' };
        setEdges(prevEdges => [...prevEdges, newEdge]);
        showToast("success", "Proceso creado correctamente");
        setPID(data);
        setIDActual(3);

      }
    } catch (error) {
      console.error('Error al obtener datos de la API:', error);
    }
  };

  const fetchPararProceso = async () => {
    try {

      if (pid == -1) {
        showToast("error", "Debe crear un proceso para utilizar esta opcion");
      }
      else if (idActual == 2) {
        showToast("error", "El proceso ya esta detenido");
      } else {
        await fetch('http://localhost:8080/parar-proceso/' + pid);
        const nodoNew = { id: 2, label: 'Ready', color: '#265DD2' };
        const existe = nodes.some(nodo => nodo.id === 2);
        if (!existe) {
          setNodes(prevNodos => [...prevNodos, nodoNew]);
        }
        const newEdge = { from: idActual, to: 2, arrows: 'to' };
        setEdges(prevEdges => [...prevEdges, newEdge]);
        showToast("success", "Proceso detenido correctamente");
        setIDActual(2);
      }
    } catch (error) {
      console.error('Error al obtener datos de la API:', error);
    }
  };

  const fetchProcesoListo = async () => {
    try {
      if (pid == -1) {
        showToast("error", "Debe crear un proceso para utilizar esta opcion");
      }
      else if (idActual == 3) {
        showToast("error", "El proceso ya esta corriendo");
      } else {
        await fetch('http://localhost:8080/iniciar-proceso/' + pid);
        const newEdge = { from: idActual, to: 3, arrows: 'to' };
        setEdges(prevEdges => [...prevEdges, newEdge]);
        showToast("success", "Proceso corriendo correctamente");
        setIDActual(3);
      }
    } catch (error) {
      console.error('Error al obtener datos de la API:', error);
    }
  };

  const fetchMatarProceso = async () => {
    try {
      if (pid == -1) {
        showToast("error", "Debe crear un proceso para utilizar esta opcion");
      }
      else if (idActual == 4) {
        showToast("error", "El proceso ya finalizo");
      } else {
        await fetch('http://localhost:8080/matar-proceso/' + pid);
        const nodoNew = { id: 4, label: 'Terminated', color: '#D24326' };
        const existe = nodes.some(nodo => nodo.id === 4);
        if (!existe) {
          setNodes(prevNodos => [...prevNodos, nodoNew]);
        }
        const newEdge = { from: idActual, to: 4, arrows: 'to' };
        setEdges(prevEdges => [...prevEdges, newEdge]);
        showToast("success", "Proceso terminado correctamente");
        setIDActual(4);
        setPID(-1);
      }
    } catch (error) {
      console.error('Error al obtener datos de la API:', error);
    }
  };

  useEffect(() => {
    const container = document.getElementById('mynetwork');
    const data = {
      nodes: new DataSet(nodes),
      edges: new DataSet(edges)
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
          {pid !== -1 && <h1 className='pid-title'>PID: {pid}</h1>}
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