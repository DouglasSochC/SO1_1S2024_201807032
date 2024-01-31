import React, { useState } from 'react';
import axios from 'axios';
import './App.css';

function App() {
  const [datos, setDatos] = useState('');

  const mostrarDatos = async () => {
    try {
      const response = await axios.get('http://localhost:8080/data');
      setDatos(response.data);
    } catch (error) {
      console.error('Error al obtener datos:', error);
    }
  };

  return (
    <div className="App">
      <main className="App-header">
        <h1>Tarea 1 - SO1 - 1S2024</h1>
        <button onClick={mostrarDatos}>Mostrar Datos</button>
        <br/>
        {
          datos && (
            <>
              <b>Carnet:</b><p>{datos.carnet}</p>
              <b>Nombre:</b><p>{datos.nombre}</p>
            </>
          )
        }
      </main>
    </div>
  );
}

export default App;