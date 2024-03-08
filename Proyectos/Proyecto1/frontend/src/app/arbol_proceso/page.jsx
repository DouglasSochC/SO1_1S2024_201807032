'use client'

import React, { useEffect, useState } from 'react';
import Layout from '@/components/Layout/page';
import { Tree } from 'react-d3-tree';
import Select from 'react-select';

const data = [{"id":1,"name":"Abuelo","children":[{"id":2,"name":"Padre","children":[{"id":4,"name":"Hijo1","children":[{"id":8,"name":"Nieto1"}]},{"id":5,"name":"Hijo2","children":[{"id":9,"name":"Nieto2"}]}]},{"id":3,"name":"Tío","children":[{"id":6,"name":"Sobrino1","children":[{"id":10,"name":"NuevoHijo"}]}]}]}]

const treeConfig = {
    orientation: 'vertical'
};



const TreeComponent = () => {
    const [mounted, setMounted] = useState(false);
    const [processes, setProcesses] = useState([]);
    const [selectedProcess, setSelectedProcess] = useState('');

    const fetchData = () => {

        console.log("AQUI");
        // Reemplaza la URL del endpoint con tu propio endpoint
        fetch('http://localhost:8080/procesos-actuales')
            .then(response => response.json())
            .then(data => {
                // Asegúrate de que la respuesta tenga el formato adecuado
                if (Array.isArray(data)) {
                    setProcesses(data);
                } else {
                    console.error('La respuesta del endpoint no tiene el formato esperado.');
                }
            })
            .catch(error => console.error('Error al obtener procesos:', error));
    };

    useEffect(() => {
        setMounted(true);
    }, []);

    return (
        <Layout
            pageTitle=""
        >
            <div className="min-h-screen flex flex-col justify-center items-center">
                <div className="m-auto text-center">
                    <h1 className="main-title">ARBOL DE UN PROCESO</h1>

                    <div style={{ marginBottom: '20px' }}>
                        <button onClick={fetchData} style={{
                            padding: '10px',
                            backgroundColor: '#4CAF50',
                            color: 'white',
                            border: 'none',
                            borderRadius: '4px',
                            cursor: 'pointer',
                        }}>Obtener Procesos Actuales</button>
                    </div>

                    <div>
                        <label htmlFor="processSelect">Seleccionar Proceso:</label>
                        <Select
                            id="processSelect"
                            value={selectedProcess}
                            onChange={(selectedOption) => setSelectedProcess(selectedOption)}
                            options={processes}
                        />
                    </div>
                    <br />

                    <div className="charts-container">
                        <div style={{ width: '1000px', height: '500px' }}>
                            {mounted && <Tree data={data} orientation={treeConfig.orientation} />}
                        </div>

                    </div>
                </div>
            </div>
        </Layout>
    );
};

export default TreeComponent;