'use client'

import React, { useEffect, useState } from 'react';
import Layout from '@/components/Layout/page';
import { Tree } from 'react-d3-tree';
import Select from 'react-select';
import withReactContent from 'sweetalert2-react-content';
import Swal from 'sweetalert2';

const treeConfig = {
    orientation: 'vertical'
};

const MySwal = withReactContent(Swal);

const TreeComponent = () => {
    const [mounted, setMounted] = useState(false);
    const [processes, setProcesses] = useState([]);
    const [arbol, setArbol] = useState([{}]);

    const fetchData = () => {

        fetch('http://localhost:8080/procesos-actuales')
            .then(response => response.json())
            .then(data => {
                if (Array.isArray(data)) {
                    setProcesses(data);
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
                        icon: "success",
                        title: "Procesos obtenidos correctamente"
                    });
                } else {
                    console.error('La respuesta del endpoint no tiene el formato esperado.');
                }
            })
            .catch(error => console.error('Error al obtener procesos:', error));
    };

    const handleProcessChange = (selectedOption) => {

        fetch('http://localhost:8080/arbol-proceso/' + selectedOption.value)
            .then(response => response.json())
            .then(data => {
                console.log(data);
                setArbol(data);
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
                            onChange={handleProcessChange}
                            options={processes}
                        />
                    </div>
                    <br />

                    <div className="charts-container">
                        <div style={{ width: '1000px', height: '500px' }}>
                            {mounted && <Tree data={arbol} orientation={treeConfig.orientation} />}
                        </div>

                    </div>
                </div>
            </div>
        </Layout>
    );
};

export default TreeComponent;